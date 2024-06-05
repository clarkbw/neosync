package genbenthosconfigs_activity

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/doug-martin/goqu/v9"
	sql_manager "github.com/nucleuscloud/neosync/backend/pkg/sqlmanager"
	tabledependency "github.com/nucleuscloud/neosync/backend/pkg/table-dependency"
	pg_query "github.com/pganalyze/pg_query_go/v5"
	pgquery "github.com/wasilibs/go-pgquery"
	"github.com/xwb1989/sqlparser"

	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/exp"
	neosync_benthos "github.com/nucleuscloud/neosync/worker/internal/benthos"
	"github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/activities/shared"
)

const (
	innerJoin joinType = "INNER"
)

type joinType string

type sqlJoin struct {
	JoinType       joinType
	JoinTable      string
	BaseTable      string
	Alias          *string
	JoinColumnsMap map[string]string // map of joinColumn to baseColumn
}

type SubsetReferenceKey struct {
	Table         string
	Columns       []string
	OriginalTable *string
}
type SubsetColumnConstraint struct {
	Columns     []string
	NotNullable []bool
	ForeignKey  *SubsetReferenceKey
}

func buildSelectQuery(
	driver, table string,
	columns []string,
	whereClause *string,
) (string, error) {
	builder := goqu.Dialect(driver)
	sqltable := goqu.I(table)

	selectColumns := make([]any, len(columns))
	for i, col := range columns {
		selectColumns[i] = col
	}
	query := builder.From(sqltable).Select(selectColumns...)

	if whereClause != nil && *whereClause != "" {
		query = query.Where(goqu.L(*whereClause))
	}
	sql, _, err := query.ToSQL()
	if err != nil {
		return "", err
	}

	return formatSqlQuery(sql), nil
}

func buildSelectJoinQuery(
	driver, table string,
	columns []string,
	joins []*sqlJoin,
	whereClauses []string,
) (string, error) {
	builder := goqu.Dialect(driver)
	sqltable := goqu.I(table)

	selectColumns := make([]any, len(columns))
	for i, col := range columns {
		selectColumns[i] = buildSqlIdentifier(table, col)
	}
	query := builder.From(sqltable).Select(selectColumns...)
	// joins
	for _, j := range joins {
		if j == nil {
			continue
		}
		joinConditionTable := j.JoinTable
		if j.Alias != nil && *j.Alias != "" {
			joinConditionTable = *j.Alias
		}
		joinCondition := goqu.Ex{}
		for joinCol, baseCol := range j.JoinColumnsMap {
			joinCondition[buildSqlIdentifier(joinConditionTable, joinCol)] = goqu.I(buildSqlIdentifier(j.BaseTable, baseCol))
		}
		if j.JoinType == innerJoin {
			var joinTable exp.Expression
			joinTable = goqu.I(j.JoinTable)
			if j.Alias != nil && *j.Alias != "" {
				joinTable = goqu.I(j.JoinTable).As(*j.Alias)
			}
			query = query.InnerJoin(
				joinTable,
				goqu.On(joinCondition),
			)
		}
	}
	// where
	goquWhere := []exp.Expression{}
	for _, w := range whereClauses {
		goquWhere = append(goquWhere, goqu.L(w))
	}
	query = query.Where(goqu.And(goquWhere...))

	sql, _, err := query.ToSQL()
	if err != nil {
		return "", err
	}
	return formatSqlQuery(sql), nil
}

func buildSelectRecursiveQuery(
	driver, table string,
	columns []string,
	columnInfoMap map[string]*sql_manager.ColumnInfo,
	dependencies []*selfReferencingCircularDependency,
	joins []*sqlJoin,
	whereClauses []string,
) (string, error) {
	recursiveCteAlias := "related"
	var builder goqu.DialectWrapper
	if driver == sql_manager.MysqlDriver {
		opts := goqu.DefaultDialectOptions()
		opts.QuoteRune = '`'
		opts.SupportsWithCTERecursive = true
		opts.SupportsWithCTE = true
		dialectName := "custom-mysql-dialect"
		goqu.RegisterDialect(dialectName, opts)
		builder = goqu.Dialect(dialectName)
	} else {
		builder = goqu.Dialect(driver)
	}

	sqltable := goqu.I(table)

	selectColumns := make([]any, len(columns))
	for i, col := range columns {
		colInfo := columnInfoMap[col]
		if driver == sql_manager.PostgresDriver && colInfo != nil && colInfo.DataType == "json" {
			selectColumns[i] = goqu.L("to_jsonb(?)", goqu.I(buildSqlIdentifier(table, col))).As(col)
		} else {
			selectColumns[i] = buildSqlIdentifier(table, col)
		}
	}
	selectQuery := builder.From(sqltable).Select(selectColumns...)

	initialSelect := selectQuery
	// joins
	for _, j := range joins {
		if j == nil {
			continue
		}
		joinCondition := goqu.Ex{}
		for joinCol, baseCol := range j.JoinColumnsMap {
			joinCondition[buildSqlIdentifier(j.JoinTable, joinCol)] = goqu.I(buildSqlIdentifier(j.BaseTable, baseCol))
		}
		if j.JoinType == innerJoin {
			table := goqu.I(j.JoinTable)
			initialSelect = initialSelect.InnerJoin(
				table,
				goqu.On(joinCondition),
			)
		}
	}

	// where
	goquWhere := []exp.Expression{}
	for _, w := range whereClauses {
		goquWhere = append(goquWhere, goqu.L(w))
	}
	initialSelect = initialSelect.Where(goqu.And(goquWhere...))

	// inner join on foreign keys
	goquOnOrEx := []exp.Expression{}
	for _, d := range dependencies {
		goquOnAndEx := []exp.Expression{}
		for _, fk := range d.ForeignKeyColumns {
			if len(fk) > 1 {
				for idx, col := range fk {
					goquOnAndEx = append(goquOnAndEx, goqu.Ex{buildSqlIdentifier(table, d.PrimaryKeyColumns[idx]): goqu.I(buildSqlIdentifier(recursiveCteAlias, col))})
				}
				goquOnOrEx = append(goquOnOrEx, goqu.And(goquOnAndEx...))
			} else {
				goquOnOrEx = append(goquOnOrEx, goqu.Ex{buildSqlIdentifier(table, d.PrimaryKeyColumns[0]): goqu.I(buildSqlIdentifier(recursiveCteAlias, fk[0]))})
			}
		}
	}

	recursiveSelect := selectQuery
	recursiveSelect = recursiveSelect.InnerJoin(goqu.I(recursiveCteAlias), goqu.On(goqu.Or(goquOnOrEx...)))

	// union
	unionQuery := initialSelect.Union(recursiveSelect)

	selectCols := make([]any, len(columns))
	for i, col := range columns {
		selectCols[i] = col
	}
	recursiveQuery := builder.From(goqu.T(recursiveCteAlias)).WithRecursive(recursiveCteAlias, unionQuery).SelectDistinct(selectCols...)
	sql, _, err := recursiveQuery.ToSQL()
	if err != nil {
		return "", err
	}
	return formatSqlQuery(sql), nil
}

// returns map of schema.table -> select query
func buildSelectQueryMap(
	driver string,
	tableDependencies map[string][]*sql_manager.ForeignConstraint,
	runConfigs []*tabledependency.RunConfig,
	subsetByForeignKeyConstraints bool,
	groupedColumnInfo map[string]map[string]*sql_manager.ColumnInfo,
) (map[string]map[tabledependency.RunType]string, error) {
	insertRunConfigMap := map[string]*tabledependency.RunConfig{}
	for _, cfg := range runConfigs {
		if cfg.RunType == tabledependency.RunTypeInsert {
			insertRunConfigMap[cfg.Table] = cfg
		}
	}

	// map of table -> where clause
	tableWhereMap := map[string]string{}
	for table, config := range insertRunConfigMap {
		if config.WhereClause != nil && *config.WhereClause != "" {
			qualifiedWhere, err := qualifyWhereColumnNames(driver, *config.WhereClause, table)
			if err != nil {
				return nil, err
			}
			tableWhereMap[table] = qualifiedWhere
		}
	}

	if !subsetByForeignKeyConstraints || len(tableDependencies) == 0 || len(tableWhereMap) == 0 {
		queryRunTypeMap, err := buildQueryMapNoSubsetConstraints(driver, runConfigs)
		if err != nil {
			return nil, err
		}
		return queryRunTypeMap, nil
	}

	subsetConfigs, err := buildTableSubsetQueryConfigs(driver, tableDependencies, tableWhereMap, insertRunConfigMap)
	if err != nil {
		return nil, err
	}

	queryRunTypeMap := map[string]map[tabledependency.RunType]string{}
	for _, runConfig := range runConfigs {
		if _, ok := queryRunTypeMap[runConfig.Table]; !ok {
			queryRunTypeMap[runConfig.Table] = map[tabledependency.RunType]string{}
		}
		columns := runConfig.SelectColumns
		subsetConfig := subsetConfigs[runConfig.Table]
		columnInfoMap := groupedColumnInfo[runConfig.Table]
		sql, err := buildTableQuery(driver, runConfig.Table, columns, subsetConfig, columnInfoMap)
		if err != nil {
			return nil, err
		}
		queryRunTypeMap[runConfig.Table][runConfig.RunType] = sql
	}
	return queryRunTypeMap, nil
}

func buildTableQuery(
	driver, table string,
	columns []string,
	config *subsetQueryConfig,
	columnInfoMap map[string]*sql_manager.ColumnInfo,
) (string, error) {
	if len(config.SelfReferencingCircularDependency) != 0 {
		sql, err := buildSelectRecursiveQuery(
			driver,
			table,
			columns,
			columnInfoMap,
			config.SelfReferencingCircularDependency,
			config.Joins,
			config.WhereClauses,
		)
		if err != nil {
			return "", fmt.Errorf("unable to build recursive select query: %w", err)
		}
		return sql, err
	} else if len(config.Joins) == 0 {
		where := strings.Join(config.WhereClauses, " AND ")
		sql, err := buildSelectQuery(driver, table, columns, &where)
		if err != nil {
			return "", fmt.Errorf("unable to build select query: %w", err)
		}
		return sql, nil
	} else {
		sql, err := buildSelectJoinQuery(driver, table, columns, config.Joins, config.WhereClauses)
		if err != nil {
			return "", fmt.Errorf("unable to build select query with joins: %w", err)
		}
		return sql, nil
	}
}

func buildQueryMapNoSubsetConstraints(
	driver string,
	runConfigs []*tabledependency.RunConfig,
) (map[string]map[tabledependency.RunType]string, error) {
	queryRunTypeMap := map[string]map[tabledependency.RunType]string{}
	for _, config := range runConfigs {
		if _, ok := queryRunTypeMap[config.Table]; !ok {
			queryRunTypeMap[config.Table] = map[tabledependency.RunType]string{}
		}
		query, err := buildSelectQuery(driver, config.Table, config.SelectColumns, config.WhereClause)
		if err != nil {
			return nil, fmt.Errorf("unable to build select query: %w", err)
		}
		queryRunTypeMap[config.Table][config.RunType] = query
	}
	return queryRunTypeMap, nil
}

type tableSubset struct {
	Joins        []*sqlJoin
	WhereClauses []string
}

type tableAlias struct {
	Alias string
	Count int
}

// recusively builds join for subset table
func buildSubsetJoins(
	driver, table string,
	// data map[string][]*SubsetColumnConstraint,
	foreignKeyMap map[string][]*sql_manager.ForeignConstraint,
	whereClauses map[string]string,
	visited map[string]bool,
	aliasReference map[string]*tableAlias,
) (*tableSubset, error) {
	fmt.Println()
	fmt.Printf("table: %s \n", table)
	jsonF, _ := json.MarshalIndent(aliasReference, "", " ")
	fmt.Printf("aliasReference: %s \n", string(jsonF))

	joins := []*sqlJoin{}
	wheres := []string{}

	if seen := visited[table]; seen {
		return &tableSubset{
			Joins:        joins,
			WhereClauses: wheres,
		}, nil
	}

	condition, whereExists := whereClauses[table]
	if whereExists && len(visited) == 0 {
		fmt.Printf("adding where: %s \n", condition)
		wheres = append(wheres, condition)
	}

	// visited[table] = true

	fks, exists := foreignKeyMap[table]
	if !exists {
		return &tableSubset{
			Joins:        joins,
			WhereClauses: wheres,
		}, nil
	}
	for _, fk := range fks {
		if fk.ForeignKey.Table == "" && fk.ForeignKey.Columns == nil {
			continue
		}
		if !visited[fk.ForeignKey.Table] {
			// handle aliased table
			var alias *string
			joinTable := fk.ForeignKey.Table
			aliasRef, okAliasRef := aliasReference[fk.ForeignKey.Table]

			if okAliasRef {
				newTable := fmt.Sprintf("%s_%s_%d", strings.ReplaceAll(fk.ForeignKey.Table, ".", "_"), strings.Join(fk.Columns, "_"), aliasRef.Count)
				aliasTable := aliasHash(newTable)
				fmt.Printf("table: %s  alias: %s \n", newTable, aliasTable)
				aliasReference[fk.ForeignKey.Table] = &tableAlias{
					Alias: aliasTable,
					Count: aliasRef.Count + 1,
				}
				visited[aliasTable] = true
				alias = &aliasTable
			} else {
				newTable := fmt.Sprintf("%s_%s_%d", strings.ReplaceAll(fk.ForeignKey.Table, ".", "_"), strings.Join(fk.Columns, "_"), 0)
				aliasTable := aliasHash(newTable)
				fmt.Printf("table: %s  alias: %s \n", newTable, aliasTable)
				aliasReference[fk.ForeignKey.Table] = &tableAlias{
					Alias: aliasTable,
					Count: 1,
				}
				visited[aliasTable] = true
				alias = &aliasTable
			}
			// if ok := visited[aliasTable]; !ok {

			// }
			//  else {
			// 	joinTable = aliasTable
			// }

			if condition, whereOk := whereClauses[joinTable]; whereOk {
				aliasedWhere, err := qualifyWhereWithTableAlias(driver, condition, *alias)
				if err != nil {
					return nil, err
				}
				fmt.Printf("adding aliased where: %s \n", aliasedWhere)

				wheres = append(wheres, aliasedWhere)
			}

			baseTable := table
			if a, ok := aliasReference[baseTable]; ok {
				baseTable = a.Alias
			}

			joinColMap := map[string]string{}
			for idx, c := range fk.ForeignKey.Columns {
				joinColMap[c] = fk.Columns[idx]
			}
			joins = append(joins, &sqlJoin{
				JoinType:       innerJoin,
				JoinTable:      joinTable,
				BaseTable:      baseTable,
				Alias:          alias,
				JoinColumnsMap: joinColMap,
			})
			jsonF, _ := json.MarshalIndent(&sqlJoin{
				JoinType:       innerJoin,
				JoinTable:      joinTable,
				BaseTable:      baseTable,
				Alias:          alias,
				JoinColumnsMap: joinColMap,
			}, "", " ")
			fmt.Printf("join: %s \n", string(jsonF))

			sub, err := buildSubsetJoins(driver, fk.ForeignKey.Table, foreignKeyMap, whereClauses, visited, aliasReference)
			if err != nil {
				return nil, err
			}
			joins = append(joins, sub.Joins...)
			wheres = append(wheres, sub.WhereClauses...)
		}
	}
	return &tableSubset{
		Joins:        joins,
		WhereClauses: wheres,
	}, nil
}

type subsetQueryConfig struct {
	Joins                             []*sqlJoin
	WhereClauses                      []string
	SelfReferencingCircularDependency []*selfReferencingCircularDependency
}

func buildTableSubsetQueryConfigs(driver string, tableConstraints map[string][]*sql_manager.ForeignConstraint, whereClauses map[string]string, runConfigMap map[string]*tabledependency.RunConfig) (map[string]*subsetQueryConfig, error) {
	configs := map[string]*subsetQueryConfig{}

	filteredConstraints := filterForeignKeysWithSubset(runConfigMap, tableConstraints, whereClauses)
	jsonF, _ := json.MarshalIndent(whereClauses, "", " ")
	fmt.Printf("\n whereClauses: %s \n", string(jsonF))
	jsonF, _ = json.MarshalIndent(tableConstraints, "", " ")
	fmt.Printf("\n constraints: %s \n", string(jsonF))
	jsonF, _ = json.MarshalIndent(runConfigMap, "", " ")
	fmt.Printf("\n runConfigMap: %s \n", string(jsonF))
	jsonF, _ = json.MarshalIndent(filteredConstraints, "", " ")
	fmt.Printf("\n filteredConstraints: %s \n", string(jsonF))
	for table, cfg := range runConfigMap {
		if cfg.RunType == tabledependency.RunTypeUpdate {
			continue
		}

		fmt.Println()
		fmt.Printf("------ begin build joins:  %s \n", table)
		visited := map[string]bool{}
		tableSubset, err := buildSubsetJoins(driver, table, filteredConstraints, whereClauses, visited, map[string]*tableAlias{})
		if err != nil {
			return nil, err
		}

		constraints := tableConstraints[table]
		selfRefCd := getSelfReferencingColumns(table, constraints)
		jsonF, _ = json.MarshalIndent(selfRefCd, "", " ")
		fmt.Printf("\n selfRefCd: %s \n", string(jsonF))
		configs[table] = &subsetQueryConfig{
			Joins:                             tableSubset.Joins,
			WhereClauses:                      tableSubset.WhereClauses,
			SelfReferencingCircularDependency: selfRefCd,
		}
	}
	return configs, nil
}

// check if table or any parent table has a where clause
func shouldSubsetTable(table string, data map[string][]*sql_manager.ForeignConstraint, whereClauses map[string]string, visited map[string]bool) bool {
	if _, exists := whereClauses[table]; exists {
		return true
	}
	if visited[table] {
		return false
	}
	visited[table] = true

	if columns, exists := data[table]; exists {
		for _, col := range columns {
			if col.ForeignKey.Table != "" {
				if shouldSubsetTable(col.ForeignKey.Table, data, whereClauses, visited) {
					return true
				}
			}
		}
	}
	return false
}

// filters out any foreign keys that are not involved in the subset (where clauses) and tables not in the map
// handles circular dependencies
// func filterForeignKeysWithSubset(runConfigMap map[string]*tabledependency.RunConfig, constraints map[string][]*sql_manager.ForeignConstraint, whereClauses map[string]string) map[string][]*sql_manager.ForeignConstraint {
// 	tableSubsetMap := map[string]bool{} // map of which tables to subset
// 	for table := range runConfigMap {
// 		visited := map[string]bool{}
// 		tableSubsetMap[table] = shouldSubsetTable(table, constraints, whereClauses, visited)
// 	}

// 	filteredConstraints := map[string][]*sql_manager.ForeignConstraint{}
// 	for table, cfg := range runConfigMap {
// 		if len(cfg.DependsOn) == 0 {
// 			// filter out dependencies that aren't in run configs
// 			// this handles circular dependencies
// 			continue
// 		}
// 		filteredConstraints[table] = []*sql_manager.ForeignConstraint{}
// 		if tableConstraints, ok := constraints[table]; ok {
// 			for _, colDef := range tableConstraints {
// 				if exists := tableSubsetMap[colDef.ForeignKey.Table]; exists {
// 					filteredConstraints[table] = append(filteredConstraints[table], colDef)
// 				}
// 			}
// 		}
// 	}
// 	return filteredConstraints
// }

// filters out any foreign keys that are not involved in the subset (where clauses) and tables not in the map
// handles circular dependencies
func filterForeignKeysWithSubset(runConfigMap map[string]*tabledependency.RunConfig, constraints map[string][]*sql_manager.ForeignConstraint, whereClauses map[string]string) map[string][]*sql_manager.ForeignConstraint {
	tableSubsetMap := map[string]bool{} // map of which tables to subset
	for table := range runConfigMap {
		visited := map[string]bool{}
		tableSubsetMap[table] = shouldSubsetTable(table, constraints, whereClauses, visited)
	}

	filteredConstraints := map[string][]*sql_manager.ForeignConstraint{}
	for table, cfg := range runConfigMap {
		if len(cfg.DependsOn) == 0 {
			// filter out dependencies that aren't in run configs
			// this handles circular dependencies
			continue
		}
		filteredConstraints[table] = []*sql_manager.ForeignConstraint{}
		if tableConstraints, ok := constraints[table]; ok {
			// for _, colDef := range tableConstraints {
			// 	if exists := tableSubsetMap[colDef.ForeignKey.Table]; exists {
			// 		filteredConstraints[table] = append(filteredConstraints[table], colDef)
			// 	}
			// }
			for _, d := range cfg.DependsOn {
				fk := getForeignKeyConstraint(tableConstraints, d)
				if fk != nil {
					if exists := tableSubsetMap[fk.ForeignKey.Table]; exists {
						filteredConstraints[table] = append(filteredConstraints[table], fk)
					}
				}

			}
		}
	}
	return filteredConstraints
}

func getForeignKeyConstraint(fks []*sql_manager.ForeignConstraint, dependsOn *tabledependency.DependsOn) *sql_manager.ForeignConstraint {
	for _, fk := range fks {
		if fk.ForeignKey.Table != dependsOn.Table {
			continue
		}
		if slices.Equal(fk.ForeignKey.Columns, dependsOn.Columns) {
			return fk
		}

	}
	return nil

}

func aliasHash(input string) string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
	if len(hash) > 14 {
		hash = hash[:14]
	}
	return hash
}

func qualifyWhereWithTableAlias(driver, where, alias string) (string, error) {
	query := goqu.Dialect(driver).From(goqu.T(alias)).Select("*").Where(goqu.L(where))
	sql, _, err := query.ToSQL()
	if err != nil {
		return "", err
	}
	var updatedSql string
	switch driver {
	case sql_manager.MysqlDriver:
		sql, err := qualifyMysqlWhereColumnNames(sql, nil, alias)
		if err != nil {
			return "", err
		}
		updatedSql = sql
	case sql_manager.PostgresDriver:
		sql, err := qualifyPostgresWhereColumnNames(sql, nil, alias)
		if err != nil {
			return "", err
		}
		updatedSql = sql
	default:
		return "", errors.New("unsupported sql driver type")
	}
	index := strings.Index(strings.ToLower(updatedSql), "where")
	if index == -1 {
		// "where" not found
		return "", fmt.Errorf("unable to qualify where column names")
	}
	startIndex := index + len("where")
	return strings.TrimSpace(updatedSql[startIndex:]), nil
}

func qualifyWhereColumnNames(driver, where, table string) (string, error) {
	sql, err := buildSelectQuery(driver, table, []string{"*"}, &where)
	if err != nil {
		return "", err
	}
	schema, tableName := shared.SplitTableKey(table)
	var updatedSql string
	switch driver {
	case sql_manager.MysqlDriver:
		sql, err := qualifyMysqlWhereColumnNames(sql, &schema, tableName)
		if err != nil {
			return "", err
		}
		updatedSql = sql
	case sql_manager.PostgresDriver:
		sql, err := qualifyPostgresWhereColumnNames(sql, &schema, tableName)
		if err != nil {
			return "", err
		}
		updatedSql = sql
	default:
		return "", errors.New("unsupported sql driver type")
	}
	index := strings.Index(strings.ToLower(updatedSql), "where")
	if index == -1 {
		// "where" not found
		return "", fmt.Errorf("unable to qualify where column names")
	}
	startIndex := index + len("where")

	return strings.TrimSpace(updatedSql[startIndex:]), nil
}

func getPrimaryToForeignTableMapFromRunConfigs(runConfigs []*tabledependency.RunConfig) map[string][]string {
	dpMap := map[string]map[string]struct{}{}

	for _, cfg := range runConfigs {
		if _, exists := dpMap[cfg.Table]; !exists {
			dpMap[cfg.Table] = map[string]struct{}{}
		}
		for _, dep := range cfg.DependsOn {
			if _, exists := dpMap[dep.Table]; !exists {
				dpMap[dep.Table] = map[string]struct{}{}
			}
			dpMap[dep.Table][cfg.Table] = struct{}{}
		}
	}
	tableDependencyMap := map[string][]string{}
	for table, fkTables := range dpMap {
		for t := range fkTables {
			tableDependencyMap[table] = append(tableDependencyMap[table], t)
		}
	}

	return tableDependencyMap
}

type selfReferencingCircularDependency struct {
	PrimaryKeyColumns []string
	ForeignKeyColumns [][]string
}

func getSelfReferencingColumns(table string, constraints []*sql_manager.ForeignConstraint) []*selfReferencingCircularDependency {
	result := []*selfReferencingCircularDependency{}
	resultMap := map[string]*selfReferencingCircularDependency{}

	for _, c := range constraints {
		if c.ForeignKey.Table != table {
			continue
		}
		pkHash := neosync_benthos.ToSha256(fmt.Sprintf("%s%s", c.ForeignKey.Table, strings.Join(c.ForeignKey.Columns, "")))
		if _, exists := resultMap[pkHash]; !exists {
			resultMap[pkHash] = &selfReferencingCircularDependency{
				PrimaryKeyColumns: c.ForeignKey.Columns,
				ForeignKeyColumns: [][]string{},
			}
		}
		resultMap[pkHash].ForeignKeyColumns = append(resultMap[pkHash].ForeignKeyColumns, c.Columns)
	}

	for _, r := range resultMap {
		result = append(result, r)
	}

	return result
}

func formatSqlQuery(sql string) string {
	return fmt.Sprintf("%s;", sql)
}

func buildSqlIdentifier(identifiers ...string) string {
	return strings.Join(identifiers, ".")
}

func qualifyPostgresWhereColumnNames(sql string, schema *string, table string) (string, error) {
	tree, err := pgquery.Parse(sql)
	if err != nil {
		return "", err
	}

	for _, stmt := range tree.GetStmts() {
		selectStmt := stmt.GetStmt().GetSelectStmt()

		if selectStmt.WhereClause != nil {
			updatePostgresExpr(schema, table, selectStmt.WhereClause)
		}
	}
	updatedSql, err := pgquery.Deparse(tree)
	if err != nil {
		return "", err
	}
	return updatedSql, nil
}

func updatePostgresExpr(schema *string, table string, node *pg_query.Node) {
	switch expr := node.Node.(type) {
	case *pg_query.Node_SubLink:
		updatePostgresExpr(schema, table, node.GetSubLink().GetTestexpr())
	case *pg_query.Node_BoolExpr:
		for _, arg := range expr.BoolExpr.GetArgs() {
			updatePostgresExpr(schema, table, arg)
		}
	case *pg_query.Node_AExpr:
		updatePostgresExpr(schema, table, expr.AExpr.GetLexpr())
		updatePostgresExpr(schema, table, expr.AExpr.Rexpr)
	case *pg_query.Node_ColumnDef:
	case *pg_query.Node_ColumnRef:
		col := node.GetColumnRef()
		isQualified := false
		var colName *string
		// find col name and check if already has schema + table name
		for _, f := range col.Fields {
			val := f.GetString_().GetSval()
			if schema != nil && val == *schema {
				continue
			}
			if val == table {
				isQualified = true
				break
			}
			colName = &val
		}
		if !isQualified && colName != nil && *colName != "" {
			fields := []*pg_query.Node{}
			if schema != nil && *schema != "" {
				fields = append(fields, pg_query.MakeStrNode(*schema))
			}
			fields = append(fields, []*pg_query.Node{
				pg_query.MakeStrNode(table),
				pg_query.MakeStrNode(*colName),
			}...)
			col.Fields = fields
		}
	}
}

func qualifyMysqlWhereColumnNames(sql string, schema *string, table string) (string, error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return "", err
	}

	switch stmt := stmt.(type) { //nolint:gocritic
	case *sqlparser.Select:
		err = sqlparser.Walk(func(node sqlparser.SQLNode) (kontinue bool, err error) {
			switch node := node.(type) { //nolint:gocritic
			case *sqlparser.ComparisonExpr:
				if col, ok := node.Left.(*sqlparser.ColName); ok {
					s := ""
					if schema != nil && *schema != "" {
						s = *schema
					}
					col.Qualifier.Qualifier = sqlparser.NewTableIdent(s)
					col.Qualifier.Name = sqlparser.NewTableIdent(table)
				}
				return false, nil
			}
			return true, nil
		}, stmt)
		if err != nil {
			return "", err
		}
	}

	return sqlparser.String(stmt), nil
}
