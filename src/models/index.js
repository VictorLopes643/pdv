// @ts-check
import { initSchema } from '@aws-amplify/datastore';
import { schema } from './schema';



const { Categorias, Produtos } = initSchema(schema);

export {
  Categorias,
  Produtos
};