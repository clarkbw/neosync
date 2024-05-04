'use client';
import {
  SINGLE_TABLE_SCHEMA_FORM_SCHEMA,
  SingleTableSchemaFormValues,
} from '@/app/(mgmt)/[account]/new/job/schema';
import {
  SchemaTable,
  extractAllFormErrors,
} from '@/components/jobs/SchemaTable/SchemaTable';
import { getSchemaConstraintHandler } from '@/components/jobs/SchemaTable/schema-constraint-handler';
import { useAccount } from '@/components/providers/account-provider';
import { Alert, AlertTitle } from '@/components/ui/alert';
import { Button } from '@/components/ui/button';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { useToast } from '@/components/ui/use-toast';
import { useGetConnectionForeignConstraints } from '@/libs/hooks/useGetConnectionForeignConstraints';
import { useGetConnectionPrimaryConstraints } from '@/libs/hooks/useGetConnectionPrimaryConstraints';
import {
  ConnectionSchemaMap,
  useGetConnectionSchemaMap,
} from '@/libs/hooks/useGetConnectionSchemaMap';
import { useGetConnectionUniqueConstraints } from '@/libs/hooks/useGetConnectionUniqueConstraints';
import { useGetJob } from '@/libs/hooks/useGetJob';
import { getErrorMessage } from '@/util/util';
import {
  convertJobMappingTransformerFormToJobMappingTransformer,
  convertJobMappingTransformerToForm,
} from '@/yup-validations/jobs';
import { yupResolver } from '@hookform/resolvers/yup';
import {
  GenerateSourceOptions,
  GenerateSourceSchemaOption,
  GenerateSourceTableOption,
  Job,
  JobMapping,
  JobMappingTransformer,
  JobSource,
  JobSourceOptions,
  UpdateJobSourceConnectionRequest,
  UpdateJobSourceConnectionResponse,
} from '@neosync/sdk';
import { ExclamationTriangleIcon } from '@radix-ui/react-icons';
import { ReactElement, useEffect, useMemo, useState } from 'react';
import { useFieldArray, useForm } from 'react-hook-form';
import SchemaPageSkeleton from './SchemaPageSkeleton';
import { getFkIdFromGenerateSource, getOnSelectedTableToggle } from './util';

interface Props {
  jobId: string;
}

export default function DataGenConnectionCard({ jobId }: Props): ReactElement {
  const { toast } = useToast();
  const { account } = useAccount();

  const {
    data,
    mutate,
    isLoading: isJobLoading,
  } = useGetJob(account?.id ?? '', jobId);
  const fkSourceConnectionId = getFkIdFromGenerateSource(data?.job?.source);

  const {
    data: connectionSchemaDataMap,
    isLoading: isSchemaDataMapLoading,
    isValidating: isSchemaMapValidating,
  } = useGetConnectionSchemaMap(account?.id ?? '', fkSourceConnectionId ?? '');

  const { data: primaryConstraints, isValidating: isPkValidating } =
    useGetConnectionPrimaryConstraints(
      account?.id ?? '',
      fkSourceConnectionId ?? ''
    );

  const { data: foreignConstraints, isValidating: isFkValidating } =
    useGetConnectionForeignConstraints(
      account?.id ?? '',
      fkSourceConnectionId ?? ''
    );

  const { data: uniqueConstraints, isValidating: isUCValidating } =
    useGetConnectionUniqueConstraints(
      account?.id ?? '',
      fkSourceConnectionId ?? ''
    );

  const form = useForm<SingleTableSchemaFormValues>({
    resolver: yupResolver(SINGLE_TABLE_SCHEMA_FORM_SCHEMA),
    values: getJobSource(data?.job, connectionSchemaDataMap?.schemaMap),
    context: { accountId: account?.id },
  });

  const schemaConstraintHandler = useMemo(
    () =>
      getSchemaConstraintHandler(
        connectionSchemaDataMap?.schemaMap ?? {},
        primaryConstraints?.tableConstraints ?? {},
        foreignConstraints?.tableConstraints ?? {},
        uniqueConstraints?.tableConstraints ?? {}
      ),
    [isSchemaMapValidating, isPkValidating, isFkValidating, isUCValidating]
  );
  const [selectedTables, setSelectedTables] = useState<Set<string>>(new Set());
  const { append, remove, fields } = useFieldArray<SingleTableSchemaFormValues>(
    {
      control: form.control,
      name: 'mappings',
    }
  );

  useEffect(() => {
    if (isJobLoading || isSchemaDataMapLoading || selectedTables.size > 0) {
      return;
    }
    const js = getJobSource(data?.job, connectionSchemaDataMap?.schemaMap);
    setSelectedTables(
      new Set(
        js.mappings.map((mapping) => `${mapping.schema}.${mapping.table}`)
      )
    );
  }, [isJobLoading, isSchemaDataMapLoading]);

  if (isJobLoading || isSchemaDataMapLoading) {
    return <SchemaPageSkeleton />;
  }

  async function onSubmit(values: SingleTableSchemaFormValues) {
    const job = data?.job;
    if (!job) {
      return;
    }
    try {
      await updateJobConnection(account?.id ?? '', job, values);
      toast({
        title: 'Successfully updated job source connection!',
        variant: 'success',
      });
      mutate();
    } catch (err) {
      console.error(err);
      toast({
        title: 'Unable to update job source connection',
        description: getErrorMessage(err),
        variant: 'destructive',
      });
    }
  }

  const onSelectedTableToggle = getOnSelectedTableToggle(
    connectionSchemaDataMap?.schemaMap ?? {},
    selectedTables,
    setSelectedTables,
    fields,
    remove,
    append
  );
  const formMappings = form.watch('mappings');

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="numRows"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Number of Rows</FormLabel>
              <FormDescription>The number of rows to generate.</FormDescription>
              <FormControl>
                <Input
                  {...field}
                  type="number"
                  onChange={(e) => {
                    const numberValue = e.target.valueAsNumber;
                    if (!isNaN(numberValue)) {
                      field.onChange(numberValue);
                    }
                  }}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <SchemaTable
          data={formMappings}
          jobType="generate"
          constraintHandler={schemaConstraintHandler}
          schema={connectionSchemaDataMap?.schemaMap ?? {}}
          isSchemaDataReloading={isSchemaMapValidating}
          selectedTables={selectedTables}
          onSelectedTableToggle={onSelectedTableToggle}
          formErrors={extractAllFormErrors(form.formState.errors, formMappings)}
        />

        {form.formState.errors.mappings && (
          <Alert variant="destructive">
            <AlertTitle className="flex flex-row space-x-2 justify-center">
              <ExclamationTriangleIcon />
              <p>Please fix form errors and try again.</p>
            </AlertTitle>
          </Alert>
        )}
        <div className="flex flex-row gap-1 justify-end">
          <Button key="submit" type="submit">
            Update
          </Button>
        </div>
      </form>
    </Form>
  );
}

function getJobSource(
  job?: Job,
  connSchemaMap?: ConnectionSchemaMap
): SingleTableSchemaFormValues {
  if (!job || !connSchemaMap) {
    return {
      mappings: [],
      numRows: 0,
    };
  }
  let numRows = 0;
  if (job.source?.options?.config.case === 'generate') {
    const srcSchemas = job.source.options.config.value.schemas;
    if (srcSchemas.length > 0) {
      const tables = srcSchemas[0].tables;
      if (tables.length > 0) {
        numRows = Number(tables[0].rowCount); // this will be an issue if the number is bigger than what js allows
      }
    }
  }

  const mapData: Record<string, Set<string>> = {};

  const mappings: SingleTableSchemaFormValues['mappings'] = (
    job.mappings ?? []
  ).map((mapping) => {
    const tkey = `${mapping.schema}.${mapping.table}`;
    const uniqcols = mapData[tkey];
    if (uniqcols) {
      uniqcols.add(mapping.column);
    } else {
      mapData[tkey] = new Set([mapping.column]);
    }

    return {
      schema: mapping.schema,
      table: mapping.table,
      column: mapping.column,
      transformer: mapping.transformer
        ? convertJobMappingTransformerToForm(mapping.transformer)
        : convertJobMappingTransformerToForm(new JobMappingTransformer()),
    };
  });

  Object.entries(mapData).forEach(([key, currcols]) => {
    const dbcols = connSchemaMap[key];
    if (!dbcols) {
      return;
    }
    dbcols.forEach((dbcol) => {
      if (!currcols.has(dbcol.column)) {
        mappings.push({
          schema: dbcol.schema,
          table: dbcol.table,
          column: dbcol.column,
          transformer: convertJobMappingTransformerToForm(
            new JobMappingTransformer()
          ),
        });
      }
    });
  });

  return {
    mappings: mappings,
    numRows: numRows,
  };
}

async function updateJobConnection(
  accountId: string,
  job: Job,
  values: SingleTableSchemaFormValues
): Promise<UpdateJobSourceConnectionResponse> {
  const schema = values.mappings.length > 0 ? values.mappings[0].schema : null;
  const table = values.mappings.length > 0 ? values.mappings[0].table : null;
  const res = await fetch(
    `/api/accounts/${accountId}/jobs/${job.id}/source-connection`,
    {
      method: 'PUT',
      headers: {
        'content-type': 'application/json',
      },
      body: JSON.stringify(
        new UpdateJobSourceConnectionRequest({
          id: job.id,
          mappings: values.mappings.map((m) => {
            return new JobMapping({
              schema: m.schema,
              table: m.table,
              column: m.column,
              transformer:
                convertJobMappingTransformerFormToJobMappingTransformer(
                  m.transformer
                ),
            });
          }),
          source: new JobSource({
            options: new JobSourceOptions({
              config: {
                case: 'generate',
                value: new GenerateSourceOptions({
                  fkSourceConnectionId: getFkIdFromGenerateSource(job.source),
                  schemas:
                    schema && table
                      ? [
                          new GenerateSourceSchemaOption({
                            schema: schema,
                            tables: [
                              new GenerateSourceTableOption({
                                table: table,
                                rowCount: BigInt(values.numRows),
                              }),
                            ],
                          }),
                        ]
                      : [],
                }),
              },
            }),
          }),
        })
      ),
    }
  );
  if (!res.ok) {
    const body = await res.json();
    throw new Error(body.message);
  }
  return UpdateJobSourceConnectionResponse.fromJson(await res.json());
}
