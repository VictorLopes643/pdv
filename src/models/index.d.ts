import { ModelInit, MutableModel, __modelMeta__, ManagedIdentifier } from "@aws-amplify/datastore";
// @ts-ignore
import { LazyLoading, LazyLoadingDisabled, AsyncCollection } from "@aws-amplify/datastore";





type EagerCategorias = {
  readonly [__modelMeta__]: {
    identifier: ManagedIdentifier<Categorias, 'id'>;
    readOnlyFields: 'createdAt' | 'updatedAt';
  };
  readonly id: string;
  readonly name?: string | null;
  readonly Produtos?: (Produtos | null)[] | null;
  readonly createdAt?: string | null;
  readonly updatedAt?: string | null;
}

type LazyCategorias = {
  readonly [__modelMeta__]: {
    identifier: ManagedIdentifier<Categorias, 'id'>;
    readOnlyFields: 'createdAt' | 'updatedAt';
  };
  readonly id: string;
  readonly name?: string | null;
  readonly Produtos: AsyncCollection<Produtos>;
  readonly createdAt?: string | null;
  readonly updatedAt?: string | null;
}

export declare type Categorias = LazyLoading extends LazyLoadingDisabled ? EagerCategorias : LazyCategorias

export declare const Categorias: (new (init: ModelInit<Categorias>) => Categorias) & {
  copyOf(source: Categorias, mutator: (draft: MutableModel<Categorias>) => MutableModel<Categorias> | void): Categorias;
}

type EagerProdutos = {
  readonly [__modelMeta__]: {
    identifier: ManagedIdentifier<Produtos, 'id'>;
    readOnlyFields: 'createdAt' | 'updatedAt';
  };
  readonly id: string;
  readonly name?: string | null;
  readonly price?: number | null;
  readonly categoriasID: string;
  readonly createdAt?: string | null;
  readonly updatedAt?: string | null;
}

type LazyProdutos = {
  readonly [__modelMeta__]: {
    identifier: ManagedIdentifier<Produtos, 'id'>;
    readOnlyFields: 'createdAt' | 'updatedAt';
  };
  readonly id: string;
  readonly name?: string | null;
  readonly price?: number | null;
  readonly categoriasID: string;
  readonly createdAt?: string | null;
  readonly updatedAt?: string | null;
}

export declare type Produtos = LazyLoading extends LazyLoadingDisabled ? EagerProdutos : LazyProdutos

export declare const Produtos: (new (init: ModelInit<Produtos>) => Produtos) & {
  copyOf(source: Produtos, mutator: (draft: MutableModel<Produtos>) => MutableModel<Produtos> | void): Produtos;
}