package genbenthosconfigs_querybuilder

import (
	"fmt"

	sqlmanager_shared "github.com/nucleuscloud/neosync/backend/pkg/sqlmanager/shared"
	tabledependency "github.com/nucleuscloud/neosync/backend/pkg/table-dependency"
	"github.com/stretchr/testify/require"
)

func (s *IntegrationTestSuite) Test_BuildQueryMap_DoubleReference() {
	whereId := "id = 1"
	tableDependencies := map[string][]*sqlmanager_shared.ForeignConstraint{
		"genbenthosconfigs_querybuilder.department": {
			{Columns: []string{"company_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.company", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.transaction": {
			{Columns: []string{"department_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.department", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.expense_report": {
			{Columns: []string{"department_source_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.department", Columns: []string{"id"}}},
			{Columns: []string{"department_destination_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.department", Columns: []string{"id"}}},
			{Columns: []string{"transaction_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.transaction", Columns: []string{"id"}}},
		},
	}
	dependencyConfigs := []*tabledependency.RunConfig{
		{Table: "genbenthosconfigs_querybuilder.company", RunType: tabledependency.RunTypeInsert, SelectColumns: []string{"id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}, WhereClause: &whereId},
		{Table: "genbenthosconfigs_querybuilder.department", RunType: tabledependency.RunTypeInsert, SelectColumns: []string{"id", "company_id"}, InsertColumns: []string{"id", "company_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.company", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.transaction", RunType: tabledependency.RunTypeInsert, SelectColumns: []string{"id", "department_id"}, InsertColumns: []string{"id", "department_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.department", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.expense_report", RunType: tabledependency.RunTypeInsert, SelectColumns: []string{"id", "department_source_id", "department_destination_id", "transaction_id"}, InsertColumns: []string{"id", "department_source_id", "department_destination_id", "transaction_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.department", Columns: []string{"id"}}, {Table: "genbenthosconfigs_querybuilder.transaction", Columns: []string{"id"}}}},
	}

	expectedValues := map[string]map[string][]int64{
		"genbenthosconfigs_querybuilder.company": {
			"id": {1},
		},
		"genbenthosconfigs_querybuilder.department": {
			"company_id": {1},
		},
		"genbenthosconfigs_querybuilder.expense_report": {
			"department_source_id":      {1},
			"department_destination_id": {2},
		},
		"genbenthosconfigs_querybuilder.transaction": {
			"department_id": {1, 2},
		},
	}

	expectedCount := map[string]int{
		"genbenthosconfigs_querybuilder.company":        1,
		"genbenthosconfigs_querybuilder.department":     2,
		"genbenthosconfigs_querybuilder.expense_report": 1,
		"genbenthosconfigs_querybuilder.transaction":    2,
	}

	sqlMap, err := BuildSelectQueryMap(sqlmanager_shared.PostgresDriver, tableDependencies, dependencyConfigs, true, map[string]map[string]*sqlmanager_shared.ColumnInfo{})
	require.NoError(s.T(), err)
	for table, selectQueryRunType := range sqlMap {
		sql := selectQueryRunType[tabledependency.RunTypeInsert]
		require.NotEmpty(s.T(), sql)
		rows, err := s.pgpool.Query(s.ctx, sql)
		require.NoError(s.T(), err)

		columnDescriptions := rows.FieldDescriptions()

		tableExpectedValues, ok := expectedValues[table]
		require.True(s.T(), ok)

		rowCount := 0
		for rows.Next() {
			rowCount++
			values, err := rows.Values()
			require.NoError(s.T(), err)

			for i, col := range values {
				colName := columnDescriptions[i].Name
				allowedValues, ok := tableExpectedValues[colName]
				if ok {
					value := col.(int64)
					require.Containsf(s.T(), allowedValues, value, fmt.Sprintf("table: %s column: %s ", table, colName))
				}
			}
		}
		rows.Close()
		require.Equalf(s.T(), rowCount, expectedCount[table], fmt.Sprintf("table: %s ", table))
	}
}

func (s *IntegrationTestSuite) Test_BuildQueryMap_DoubleRootSubset() {
	whereCreated := "created > '2023-06-03'"
	tableDependencies := map[string][]*sqlmanager_shared.ForeignConstraint{
		"genbenthosconfigs_querybuilder.test_2_c": {
			{Columns: []string{"a_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_2_a", Columns: []string{"id"}}},
			{Columns: []string{"b_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_2_b", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_2_d": {
			{Columns: []string{"c_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_2_c", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_2_e": {
			{Columns: []string{"c_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_2_c", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_2_a": {
			{Columns: []string{"x_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_2_x", Columns: []string{"id"}}},
		},
	}
	dependencyConfigs := []*tabledependency.RunConfig{
		{Table: "genbenthosconfigs_querybuilder.test_2_x", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}, WhereClause: &whereCreated},
		{Table: "genbenthosconfigs_querybuilder.test_2_b", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}, WhereClause: &whereCreated},
		{Table: "genbenthosconfigs_querybuilder.test_2_a", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "x_id"}, InsertColumns: []string{"id", "x_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_2_x", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_2_c", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "a_id", "b_id"}, InsertColumns: []string{"id", "a_id", "b_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_2_a", Columns: []string{"id"}}, {Table: "genbenthosconfigs_querybuilder.test_2_b", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_2_d", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "c_id"}, InsertColumns: []string{"id", "c_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_2_c", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_2_e", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "c_id"}, InsertColumns: []string{"id", "c_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_2_c", Columns: []string{"id"}}}},
	}

	expectedValues := map[string]map[string][]int64{
		"genbenthosconfigs_querybuilder.test_2_x": {
			"id": {3, 4, 5},
		},
		"genbenthosconfigs_querybuilder.test_2_b": {
			"id": {3, 4, 5},
		},
		"genbenthosconfigs_querybuilder.test_2_a": {
			"x_id": {3, 4, 5},
		},
		"genbenthosconfigs_querybuilder.test_2_c": {
			"a_id": {3, 4},
			"x_id": {3, 4},
		},
		"genbenthosconfigs_querybuilder.test_2_d": {
			"c_id": {3, 4},
		},
		"genbenthosconfigs_querybuilder.test_2_e": {
			"c_id": {3, 4},
		},
	}

	expectedCount := map[string]int{
		"genbenthosconfigs_querybuilder.test_2_x": 3,
		"genbenthosconfigs_querybuilder.test_2_b": 3,
		"genbenthosconfigs_querybuilder.test_2_a": 4,
		"genbenthosconfigs_querybuilder.test_2_c": 2,
		"genbenthosconfigs_querybuilder.test_2_d": 2,
		"genbenthosconfigs_querybuilder.test_2_e": 2,
	}

	sqlMap, err := BuildSelectQueryMap(sqlmanager_shared.PostgresDriver, tableDependencies, dependencyConfigs, true, map[string]map[string]*sqlmanager_shared.ColumnInfo{})

	require.NoError(s.T(), err)
	for table, selectQueryRunType := range sqlMap {
		sql := selectQueryRunType[tabledependency.RunTypeInsert]
		require.NotEmpty(s.T(), sql)

		rows, err := s.pgpool.Query(s.ctx, sql)
		require.NoError(s.T(), err)

		columnDescriptions := rows.FieldDescriptions()

		tableExpectedValues, ok := expectedValues[table]
		require.Truef(s.T(), ok, fmt.Sprintf("table: %s missing expected values", table))

		rowCount := 0
		for rows.Next() {
			rowCount++
			values, err := rows.Values()
			require.NoError(s.T(), err)

			for i, col := range values {
				colName := columnDescriptions[i].Name
				allowedValues, ok := tableExpectedValues[colName]
				if ok {
					value := col.(int64)
					require.Containsf(s.T(), allowedValues, value, fmt.Sprintf("table: %s column: %s ", table, colName))
				}
			}
		}
		rows.Close()
		require.Equalf(s.T(), rowCount, expectedCount[table], fmt.Sprintf("table: %s ", table))
	}
}

func (s *IntegrationTestSuite) Test_BuildQueryMap_MultipleRoots() {
	whereId := "id = 1"
	whereId4 := "id in (4,5)"
	tableDependencies := map[string][]*sqlmanager_shared.ForeignConstraint{
		"genbenthosconfigs_querybuilder.test_3_b": {
			{Columns: []string{"a_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_a", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_c": {
			{Columns: []string{"b_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_b", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_d": {
			{Columns: []string{"c_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_c", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_e": {
			{Columns: []string{"d_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_d", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_g": {
			{Columns: []string{"f_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_f", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_h": {
			{Columns: []string{"g_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_g", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_i": {
			{Columns: []string{"h_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_h", Columns: []string{"id"}}},
		},
	}
	dependencyConfigs := []*tabledependency.RunConfig{
		{Table: "genbenthosconfigs_querybuilder.test_3_a", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}},
		{Table: "genbenthosconfigs_querybuilder.test_3_b", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "a_id"}, InsertColumns: []string{"id", "a_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_a", Columns: []string{"id"}}}, WhereClause: &whereId},
		{Table: "genbenthosconfigs_querybuilder.test_3_c", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "b_id"}, InsertColumns: []string{"id", "b_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_b", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_d", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "c_id"}, InsertColumns: []string{"id", "c_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_c", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_e", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "d_id"}, InsertColumns: []string{"id", "d_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_d", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_f", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}, WhereClause: &whereId4},
		{Table: "genbenthosconfigs_querybuilder.test_3_g", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "f_id"}, InsertColumns: []string{"id", "f_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_f", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_h", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "g_id"}, InsertColumns: []string{"id", "g_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_g", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_i", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "h_id"}, InsertColumns: []string{"id", "h_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_h", Columns: []string{"id"}}}},
	}

	expectedValues := map[string]map[string][]int64{
		"genbenthosconfigs_querybuilder.test_3_a": {
			"id": {1, 2, 3, 4, 5},
		},
		"genbenthosconfigs_querybuilder.test_3_b": {
			"a_id": {3},
		},
		"genbenthosconfigs_querybuilder.test_3_c": {
			"b_id": {1},
		},
		"genbenthosconfigs_querybuilder.test_3_d": {
			"c_id": {3},
		},
		"genbenthosconfigs_querybuilder.test_3_e": {
			"d_id": {5},
		},
		"genbenthosconfigs_querybuilder.test_3_f": {
			"id": {4, 5},
		},
		"genbenthosconfigs_querybuilder.test_3_g": {
			"f_id": {4, 5},
		},
		"genbenthosconfigs_querybuilder.test_3_h": {
			"g_id": {1, 3},
		},
		"genbenthosconfigs_querybuilder.test_3_i": {
			"h_id": {4},
		},
	}

	expectedCount := map[string]int{
		"genbenthosconfigs_querybuilder.test_3_a": 5,
		"genbenthosconfigs_querybuilder.test_3_b": 1,
		"genbenthosconfigs_querybuilder.test_3_c": 1,
		"genbenthosconfigs_querybuilder.test_3_d": 1,
		"genbenthosconfigs_querybuilder.test_3_e": 1,
		"genbenthosconfigs_querybuilder.test_3_f": 2,
		"genbenthosconfigs_querybuilder.test_3_g": 2,
		"genbenthosconfigs_querybuilder.test_3_h": 2,
		"genbenthosconfigs_querybuilder.test_3_i": 1,
	}

	sqlMap, err := BuildSelectQueryMap(sqlmanager_shared.PostgresDriver, tableDependencies, dependencyConfigs, true, map[string]map[string]*sqlmanager_shared.ColumnInfo{})

	require.NoError(s.T(), err)
	for table, selectQueryRunType := range sqlMap {
		sql := selectQueryRunType[tabledependency.RunTypeInsert]
		require.NotEmpty(s.T(), sql)

		rows, err := s.pgpool.Query(s.ctx, sql)
		require.NoError(s.T(), err)

		columnDescriptions := rows.FieldDescriptions()
		tableExpectedValues, ok := expectedValues[table]
		require.Truef(s.T(), ok, fmt.Sprintf("table: %s missing expected values", table))

		rowCount := 0
		for rows.Next() {
			rowCount++
			values, err := rows.Values()
			require.NoError(s.T(), err)

			for i, col := range values {
				colName := columnDescriptions[i].Name
				allowedValues, ok := tableExpectedValues[colName]
				if ok {
					value := col.(int64)
					require.Containsf(s.T(), allowedValues, value, fmt.Sprintf("table: %s column: %s ", table, colName))
				}
			}
		}
		rows.Close()
		require.Equalf(s.T(), rowCount, expectedCount[table], fmt.Sprintf("table: %s ", table))
	}
}

func (s *IntegrationTestSuite) Test_BuildQueryMap_MultipleSubsets() {
	whereId := "id in (3,4,5)"
	whereId4 := "id = 4"
	tableDependencies := map[string][]*sqlmanager_shared.ForeignConstraint{
		"genbenthosconfigs_querybuilder.test_3_b": {
			{Columns: []string{"a_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_a", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_c": {
			{Columns: []string{"b_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_b", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_d": {
			{Columns: []string{"c_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_c", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_e": {
			{Columns: []string{"d_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_d", Columns: []string{"id"}}},
		},
	}
	dependencyConfigs := []*tabledependency.RunConfig{
		{Table: "genbenthosconfigs_querybuilder.test_3_a", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}, WhereClause: &whereId},
		{Table: "genbenthosconfigs_querybuilder.test_3_b", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "a_id"}, InsertColumns: []string{"id", "a_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_a", Columns: []string{"id"}}}, WhereClause: &whereId4},
		{Table: "genbenthosconfigs_querybuilder.test_3_c", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "b_id"}, InsertColumns: []string{"id", "b_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_b", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_d", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "c_id"}, InsertColumns: []string{"id", "c_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_c", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_e", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "d_id"}, InsertColumns: []string{"id", "d_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_d", Columns: []string{"id"}}}},
	}

	expectedValues := map[string]map[string][]int64{
		"genbenthosconfigs_querybuilder.test_3_a": {
			"id": {3, 4, 5},
		},
		"genbenthosconfigs_querybuilder.test_3_b": {
			"a_id": {4},
		},
		"genbenthosconfigs_querybuilder.test_3_c": {
			"b_id": {4},
		},
		"genbenthosconfigs_querybuilder.test_3_d": {
			"c_id": {2},
		},
		"genbenthosconfigs_querybuilder.test_3_e": {
			"d_id": {4},
		},
	}

	expectedCount := map[string]int{
		"genbenthosconfigs_querybuilder.test_3_a": 3,
		"genbenthosconfigs_querybuilder.test_3_b": 1,
		"genbenthosconfigs_querybuilder.test_3_c": 1,
		"genbenthosconfigs_querybuilder.test_3_d": 1,
		"genbenthosconfigs_querybuilder.test_3_e": 1,
	}

	sqlMap, err := BuildSelectQueryMap(sqlmanager_shared.PostgresDriver, tableDependencies, dependencyConfigs, true, map[string]map[string]*sqlmanager_shared.ColumnInfo{})

	require.NoError(s.T(), err)
	for table, selectQueryRunType := range sqlMap {
		sql := selectQueryRunType[tabledependency.RunTypeInsert]
		require.NotEmpty(s.T(), sql)

		rows, err := s.pgpool.Query(s.ctx, sql)
		require.NoError(s.T(), err)

		columnDescriptions := rows.FieldDescriptions()
		tableExpectedValues, ok := expectedValues[table]
		require.Truef(s.T(), ok, fmt.Sprintf("table: %s missing expected values", table))

		rowCount := 0
		for rows.Next() {
			rowCount++
			values, err := rows.Values()
			require.NoError(s.T(), err)

			for i, col := range values {
				colName := columnDescriptions[i].Name
				allowedValues, ok := tableExpectedValues[colName]
				if ok {
					value := col.(int64)
					require.Containsf(s.T(), allowedValues, value, fmt.Sprintf("table: %s column: %s ", table, colName))
				}
			}
		}
		rows.Close()
		require.Equalf(s.T(), rowCount, expectedCount[table], fmt.Sprintf("table: %s ", table))
	}
}

func (s *IntegrationTestSuite) Test_BuildQueryMap_MultipleSubsets_SubsetsByForeignKeysOff() {
	whereId := "id in (4,5)"
	whereId4 := "id = 4"
	tableDependencies := map[string][]*sqlmanager_shared.ForeignConstraint{
		"genbenthosconfigs_querybuilder.test_3_b": {
			{Columns: []string{"a_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_a", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_c": {
			{Columns: []string{"b_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_b", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_d": {
			{Columns: []string{"c_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_c", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.test_3_e": {
			{Columns: []string{"d_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.test_3_d", Columns: []string{"id"}}},
		},
	}
	dependencyConfigs := []*tabledependency.RunConfig{
		{Table: "genbenthosconfigs_querybuilder.test_3_a", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}, WhereClause: &whereId},
		{Table: "genbenthosconfigs_querybuilder.test_3_b", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "a_id"}, InsertColumns: []string{"id", "a_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_a", Columns: []string{"id"}}}, WhereClause: &whereId4},
		{Table: "genbenthosconfigs_querybuilder.test_3_c", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "b_id"}, InsertColumns: []string{"id", "b_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_b", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_d", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "c_id"}, InsertColumns: []string{"id", "c_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_c", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.test_3_e", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "d_id"}, InsertColumns: []string{"id", "d_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.test_3_d", Columns: []string{"id"}}}},
	}

	expectedValues := map[string]map[string][]int64{
		"genbenthosconfigs_querybuilder.test_3_a": {
			"id": {4, 5},
		},
		"genbenthosconfigs_querybuilder.test_3_b": {
			"a_id": {4},
		},
		"genbenthosconfigs_querybuilder.test_3_c": {
			"b_id": {1, 2, 3, 4, 5},
		},
		"genbenthosconfigs_querybuilder.test_3_d": {
			"c_id": {1, 2, 3, 4, 5},
		},
		"genbenthosconfigs_querybuilder.test_3_e": {
			"d_id": {1, 2, 3, 4, 5},
		},
	}

	expectedCount := map[string]int{
		"genbenthosconfigs_querybuilder.test_3_a": 2,
		"genbenthosconfigs_querybuilder.test_3_b": 1,
		"genbenthosconfigs_querybuilder.test_3_c": 5,
		"genbenthosconfigs_querybuilder.test_3_d": 5,
		"genbenthosconfigs_querybuilder.test_3_e": 5,
	}

	sqlMap, err := BuildSelectQueryMap(sqlmanager_shared.PostgresDriver, tableDependencies, dependencyConfigs, false, map[string]map[string]*sqlmanager_shared.ColumnInfo{})

	require.NoError(s.T(), err)
	for table, selectQueryRunType := range sqlMap {
		sql := selectQueryRunType[tabledependency.RunTypeInsert]
		require.NotEmpty(s.T(), sql)

		rows, err := s.pgpool.Query(s.ctx, sql)
		require.NoError(s.T(), err)

		columnDescriptions := rows.FieldDescriptions()
		tableExpectedValues, ok := expectedValues[table]
		require.Truef(s.T(), ok, fmt.Sprintf("table: %s missing expected values", table))

		rowCount := 0
		for rows.Next() {
			rowCount++
			values, err := rows.Values()
			require.NoError(s.T(), err)

			for i, col := range values {
				colName := columnDescriptions[i].Name
				allowedValues, ok := tableExpectedValues[colName]
				if ok {
					value := col.(int64)
					require.Containsf(s.T(), allowedValues, value, fmt.Sprintf("table: %s column: %s ", table, colName))
				}
			}
		}
		rows.Close()
		require.Equalf(s.T(), rowCount, expectedCount[table], fmt.Sprintf("table: %s ", table))
	}
}

func (s *IntegrationTestSuite) Test_BuildQueryMap_CircularDependency() {
	whereId := "id in (1,5)"
	tableDependencies := map[string][]*sqlmanager_shared.ForeignConstraint{
		"genbenthosconfigs_querybuilder.addresses": {
			{Columns: []string{"order_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.orders", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.customers": {
			{Columns: []string{"address_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.addresses", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.orders": {
			{Columns: []string{"customer_id"}, NotNullable: []bool{false}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.customers", Columns: []string{"id"}}},
		},
		"genbenthosconfigs_querybuilder.payments": {
			{Columns: []string{"customer_id"}, NotNullable: []bool{true}, ForeignKey: &sqlmanager_shared.ForeignKey{Table: "genbenthosconfigs_querybuilder.customers", Columns: []string{"id"}}},
		},
	}
	dependencyConfigs := []*tabledependency.RunConfig{
		{Table: "genbenthosconfigs_querybuilder.orders", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "customer_id"}, InsertColumns: []string{"id"}, DependsOn: []*tabledependency.DependsOn{}},
		{Table: "genbenthosconfigs_querybuilder.addresses", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "order_id"}, InsertColumns: []string{"id", "order_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.orders", Columns: []string{"id"}}}, WhereClause: &whereId},
		{Table: "genbenthosconfigs_querybuilder.customers", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "address_id"}, InsertColumns: []string{"id", "address_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.addresses", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.payments", RunType: tabledependency.RunTypeInsert, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "customer_id"}, InsertColumns: []string{"id", "customer_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.customers", Columns: []string{"id"}}}},
		{Table: "genbenthosconfigs_querybuilder.orders", RunType: tabledependency.RunTypeUpdate, PrimaryKeys: []string{"id"}, SelectColumns: []string{"id", "customer_id"}, InsertColumns: []string{"customer_id"}, DependsOn: []*tabledependency.DependsOn{{Table: "genbenthosconfigs_querybuilder.orders", Columns: []string{"id"}}, {Table: "genbenthosconfigs_querybuilder.customers", Columns: []string{"id"}}}},
	}

	// expectedValues := map[string]map[string][]int64{
	// 	"genbenthosconfigs_querybuilder.orders": {
	// 		"id": {4, 5},
	// 	},
	// 	"genbenthosconfigs_querybuilder.addresses": {
	// 		"a_id": {4},
	// 	},
	// 	"genbenthosconfigs_querybuilder.customers": {
	// 		"b_id": {1, 2, 3, 4, 5},
	// 	},
	// 	"genbenthosconfigs_querybuilder.payments": {
	// 		"c_id": {1, 2, 3, 4, 5},
	// 	},
	// }

	// expectedCount := map[string]int{
	// 	"genbenthosconfigs_querybuilder.orders":    2,
	// 	"genbenthosconfigs_querybuilder.addresses": 1,
	// 	"genbenthosconfigs_querybuilder.customers": 5,
	// 	"genbenthosconfigs_querybuilder.payments":  5,
	// }

	sqlMap, err := BuildSelectQueryMap(sqlmanager_shared.PostgresDriver, tableDependencies, dependencyConfigs, true, map[string]map[string]*sqlmanager_shared.ColumnInfo{})

	require.NoError(s.T(), err)
	for table, selectQueryRunType := range sqlMap {
		fmt.Printf("\n---------- %s \n", table)
		sql := selectQueryRunType[tabledependency.RunTypeInsert]
		require.NotEmpty(s.T(), sql)
		fmt.Println(sql)

		rows, err := s.pgpool.Query(s.ctx, sql)
		require.NoError(s.T(), err)

		columnDescriptions := rows.FieldDescriptions()

		// tableExpectedValues, ok := expectedValues[table]
		// require.Truef(s.T(), ok, fmt.Sprintf("table: %s missing expected values", table))

		rowCount := 0
		for rows.Next() {
			rowCount++
			values, err := rows.Values()
			require.NoError(s.T(), err)

			fmt.Print("(")
			for i, col := range values {
				colName := columnDescriptions[i].Name
				fmt.Printf(" %s: %d ", colName, col.(int64))
				// allowedValues, ok := tableExpectedValues[colName]
				// if ok {
				// 	value := col.(int64)
				// 	require.Containsf(s.T(), allowedValues, value, fmt.Sprintf("table: %s column: %s ", table, colName))
				// }
			}
			fmt.Print(") ")
		}
		fmt.Println()
		fmt.Println()
		rows.Close()
		// require.Equalf(s.T(), rowCount, expectedCount[table], fmt.Sprintf("table: %s ", table))
	}
}
