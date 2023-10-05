'use client';
import { useAccount } from '@/components/providers/account-provider';
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
import { Switch } from '@/components/ui/switch';
import { Transformer } from '@/neosync-api-client/mgmt/v1alpha1/transformer_pb';
import { yupResolver } from '@hookform/resolvers/yup';
import { ReactElement } from 'react';
import { useForm } from 'react-hook-form';
import * as Yup from 'yup';

const FORM_SCHEMA = Yup.object({
  name: Yup.string().required(),
  preserve_length: Yup.bool().required(),
  preserve_domain: Yup.bool().required(),
});

type FormValues = Yup.InferType<typeof FORM_SCHEMA>;

interface Props {
  transformer: Transformer;
}

export default function EmailTransformerForm(props: Props): ReactElement {
  const account = useAccount();

  const { transformer } = props;
  const form = useForm<FormValues>({
    resolver: yupResolver(FORM_SCHEMA),
    defaultValues: {
      name: transformer.name ?? '',
      preserve_length: true, //replace with transformer values
      preserve_domain: true,
    },
  });
  //   const router = useRouter();
  //   const [checkResp, setCheckResp] = useState<
  //     CheckConnectionConfigResponse | undefined
  //   >();

  async function onSubmit(values: FormValues) {
    if (!account) {
      return;
    }
    try {
      console.log('values', values);
      //   const connection = await createPostgresConnection(
      //     values.db,
      //     values.connectionName,
      //     account.id
      //   );
      //   if (connection.connection?.id) {
      //     router.push(`/connections/${connection.connection.id}`);
      //   } else {
      //     router.push(`/connections`);
      //   }
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormControl>
                  <Input placeholder="Transformer Name" {...field} />
                </FormControl>
                <FormDescription>
                  The unique name of the transformer
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="preserve_length"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                <div className="space-y-0.5">
                  <FormLabel>Preserve Length</FormLabel>
                  <FormDescription>
                    Set the length of the output email to be the same as the
                    input
                  </FormDescription>
                </div>
                <FormControl>
                  <Switch
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                </FormControl>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="preserve_domain"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                <div className="space-y-0.5">
                  <FormLabel>Preserve Domain</FormLabel>
                  <FormDescription>
                    Set the domain of the output email to be the same as the
                    input
                  </FormDescription>
                </div>
                <FormControl>
                  <Switch
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                </FormControl>
              </FormItem>
            )}
          />
          <div className="flex justify-end">
            <Button type="submit">Submit</Button>
          </div>
        </form>
      </Form>
    </div>
  );
}

// async function createPostgresConnection(
//   db: FormValues['db'],
//   name: string,
//   accountId: string
// ): Promise<CreateConnectionResponse> {
//   const res = await fetch(`/api/connections`, {
//     method: 'POST',
//     headers: {
//       'content-type': 'application/json',
//     },
//     body: JSON.stringify(
//       new CreateConnectionRequest({
//         accountId,
//         name: name,
//         connectionConfig: new ConnectionConfig({
//           config: {
//             case: 'pgConfig',
//             value: new PostgresConnectionConfig({
//               connectionConfig: {
//                 case: 'connection',
//                 value: new PostgresConnection({
//                   host: db.host,
//                   name: db.name,
//                   user: db.user,
//                   pass: db.pass,
//                   port: db.port,
//                   sslMode: db.sslMode,
//                 }),
//               },
//             }),
//           },
//         }),
//       })
//     ),
//   });
//   if (!res.ok) {
//     const body = await res.json();
//     throw new Error(body.message);
//   }
//   return CreateConnectionResponse.fromJson(await res.json());
// }

// async function checkPostgresConnection(
//   db: FormValues['db']
// ): Promise<CheckConnectionConfigResponse> {
//   const res = await fetch(`/api/connections/postgres/check`, {
//     method: 'POST',
//     headers: {
//       'content-type': 'application/json',
//     },
//     body: JSON.stringify(db),
//   });
//   if (!res.ok) {
//     const body = await res.json();
//     throw new Error(body.message);
//   }
//   return CheckConnectionConfigResponse.fromJson(await res.json());
// }
