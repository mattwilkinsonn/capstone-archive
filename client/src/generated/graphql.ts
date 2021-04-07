import { GraphQLClient } from 'graphql-request';
import { useQuery, UseQueryOptions, useMutation, UseMutationOptions } from 'react-query';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };

function fetcher<TData, TVariables>(client: GraphQLClient, query: string, variables?: TVariables) {
  return async (): Promise<TData> => client.request<TData, TVariables>(query, variables);
}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Todo = {
  __typename?: 'Todo';
  id: Scalars['ID'];
  text: Scalars['String'];
  done: Scalars['Boolean'];
  user: User;
};

export type Capstone = {
  __typename?: 'Capstone';
  id: Scalars['ID'];
  title: Scalars['String'];
  description: Scalars['String'];
  author: Scalars['String'];
  createdAt: Scalars['Int'];
  updatedAt: Scalars['Int'];
  semester: Scalars['String'];
};

export type PaginatedCapstones = {
  __typename?: 'PaginatedCapstones';
  capstones: Array<Maybe<Capstone>>;
  hasMore: Scalars['Boolean'];
};

export type NewCapstone = {
  title: Scalars['String'];
  description: Scalars['String'];
  author: Scalars['String'];
};

export type UserError = {
  __typename?: 'UserError';
  field: Scalars['String'];
  message: Scalars['String'];
};

export type PublicUser = {
  __typename?: 'PublicUser';
  username: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createCapstone: Capstone;
  register: UserResponse;
  login: UserResponse;
  logout: Scalars['Boolean'];
};


export type MutationCreateCapstoneArgs = {
  input: NewCapstone;
};


export type MutationRegisterArgs = {
  input: Register;
};


export type MutationLoginArgs = {
  input: Login;
};

export type Login = {
  usernameOrEmail: Scalars['String'];
  password: Scalars['String'];
};

export type UserResponse = {
  __typename?: 'UserResponse';
  user?: Maybe<User>;
  error?: Maybe<UserError>;
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID'];
  username: Scalars['String'];
  email: Scalars['String'];
  createdAt: Scalars['Int'];
  updatedAt: Scalars['Int'];
};

export type Register = {
  username: Scalars['String'];
  email: Scalars['String'];
  password: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  searchCapstones: PaginatedCapstones;
  capstones: PaginatedCapstones;
  capstone?: Maybe<Capstone>;
  users: Array<PublicUser>;
  me?: Maybe<User>;
};


export type QuerySearchCapstonesArgs = {
  query: Scalars['String'];
  limit: Scalars['Int'];
  offset?: Maybe<Scalars['Int']>;
};


export type QueryCapstonesArgs = {
  limit: Scalars['Int'];
  cursor?: Maybe<Scalars['Int']>;
};


export type QueryCapstoneArgs = {
  id: Scalars['Int'];
};

export type CapstoneQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type CapstoneQuery = (
  { __typename?: 'Query' }
  & { capstone?: Maybe<(
    { __typename?: 'Capstone' }
    & Pick<Capstone, 'id' | 'title' | 'description' | 'author' | 'createdAt' | 'updatedAt' | 'semester'>
  )> }
);

export type CapstonesQueryVariables = Exact<{
  limit: Scalars['Int'];
  cursor?: Maybe<Scalars['Int']>;
}>;


export type CapstonesQuery = (
  { __typename?: 'Query' }
  & { capstones: (
    { __typename?: 'PaginatedCapstones' }
    & Pick<PaginatedCapstones, 'hasMore'>
    & { capstones: Array<Maybe<(
      { __typename?: 'Capstone' }
      & Pick<Capstone, 'id' | 'title' | 'description' | 'createdAt' | 'semester'>
    )>> }
  ) }
);

export type CreateCapstoneMutationVariables = Exact<{
  input: NewCapstone;
}>;


export type CreateCapstoneMutation = (
  { __typename?: 'Mutation' }
  & { createCapstone: (
    { __typename?: 'Capstone' }
    & Pick<Capstone, 'id' | 'title' | 'description' | 'createdAt' | 'author' | 'updatedAt' | 'semester'>
  ) }
);

export type SearchCapstonesQueryVariables = Exact<{
  query: Scalars['String'];
  limit: Scalars['Int'];
  offset?: Maybe<Scalars['Int']>;
}>;


export type SearchCapstonesQuery = (
  { __typename?: 'Query' }
  & { searchCapstones: (
    { __typename?: 'PaginatedCapstones' }
    & Pick<PaginatedCapstones, 'hasMore'>
    & { capstones: Array<Maybe<(
      { __typename?: 'Capstone' }
      & Pick<Capstone, 'id' | 'title' | 'description' | 'createdAt' | 'updatedAt' | 'author' | 'semester'>
    )>> }
  ) }
);

export type LoginMutationVariables = Exact<{
  input: Login;
}>;


export type LoginMutation = (
  { __typename?: 'Mutation' }
  & { login: (
    { __typename?: 'UserResponse' }
    & { user?: Maybe<(
      { __typename?: 'User' }
      & Pick<User, 'id' | 'username' | 'email' | 'createdAt' | 'updatedAt'>
    )>, error?: Maybe<(
      { __typename?: 'UserError' }
      & Pick<UserError, 'field' | 'message'>
    )> }
  ) }
);

export type LogoutMutationVariables = Exact<{ [key: string]: never; }>;


export type LogoutMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'logout'>
);

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = (
  { __typename?: 'Query' }
  & { me?: Maybe<(
    { __typename?: 'User' }
    & Pick<User, 'id' | 'username' | 'email' | 'createdAt' | 'updatedAt'>
  )> }
);

export type RegisterMutationVariables = Exact<{
  input: Register;
}>;


export type RegisterMutation = (
  { __typename?: 'Mutation' }
  & { register: (
    { __typename?: 'UserResponse' }
    & { user?: Maybe<(
      { __typename?: 'User' }
      & Pick<User, 'id' | 'username' | 'email' | 'createdAt' | 'updatedAt'>
    )>, error?: Maybe<(
      { __typename?: 'UserError' }
      & Pick<UserError, 'field' | 'message'>
    )> }
  ) }
);


export const CapstoneDocument = `
    query Capstone($id: Int!) {
  capstone(id: $id) {
    id
    title
    description
    author
    createdAt
    updatedAt
    semester
  }
}
    `;
export const useCapstoneQuery = <
      TData = CapstoneQuery,
      TError = unknown
    >(
      client: GraphQLClient, 
      variables: CapstoneQueryVariables, 
      options?: UseQueryOptions<CapstoneQuery, TError, TData>
    ) => 
    useQuery<CapstoneQuery, TError, TData>(
      ['Capstone', variables],
      fetcher<CapstoneQuery, CapstoneQueryVariables>(client, CapstoneDocument, variables),
      options
    );
export const CapstonesDocument = `
    query Capstones($limit: Int!, $cursor: Int) {
  capstones(limit: $limit, cursor: $cursor) {
    capstones {
      id
      title
      description
      createdAt
      semester
    }
    hasMore
  }
}
    `;
export const useCapstonesQuery = <
      TData = CapstonesQuery,
      TError = unknown
    >(
      client: GraphQLClient, 
      variables: CapstonesQueryVariables, 
      options?: UseQueryOptions<CapstonesQuery, TError, TData>
    ) => 
    useQuery<CapstonesQuery, TError, TData>(
      ['Capstones', variables],
      fetcher<CapstonesQuery, CapstonesQueryVariables>(client, CapstonesDocument, variables),
      options
    );
export const CreateCapstoneDocument = `
    mutation createCapstone($input: NewCapstone!) {
  createCapstone(input: $input) {
    id
    title
    description
    createdAt
    author
    updatedAt
    semester
  }
}
    `;
export const useCreateCapstoneMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient, 
      options?: UseMutationOptions<CreateCapstoneMutation, TError, CreateCapstoneMutationVariables, TContext>
    ) => 
    useMutation<CreateCapstoneMutation, TError, CreateCapstoneMutationVariables, TContext>(
      (variables?: CreateCapstoneMutationVariables) => fetcher<CreateCapstoneMutation, CreateCapstoneMutationVariables>(client, CreateCapstoneDocument, variables)(),
      options
    );
export const SearchCapstonesDocument = `
    query searchCapstones($query: String!, $limit: Int!, $offset: Int) {
  searchCapstones(query: $query, limit: $limit, offset: $offset) {
    capstones {
      id
      title
      description
      createdAt
      updatedAt
      author
      semester
    }
    hasMore
  }
}
    `;
export const useSearchCapstonesQuery = <
      TData = SearchCapstonesQuery,
      TError = unknown
    >(
      client: GraphQLClient, 
      variables: SearchCapstonesQueryVariables, 
      options?: UseQueryOptions<SearchCapstonesQuery, TError, TData>
    ) => 
    useQuery<SearchCapstonesQuery, TError, TData>(
      ['searchCapstones', variables],
      fetcher<SearchCapstonesQuery, SearchCapstonesQueryVariables>(client, SearchCapstonesDocument, variables),
      options
    );
export const LoginDocument = `
    mutation Login($input: Login!) {
  login(input: $input) {
    user {
      id
      username
      email
      createdAt
      updatedAt
    }
    error {
      field
      message
    }
  }
}
    `;
export const useLoginMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient, 
      options?: UseMutationOptions<LoginMutation, TError, LoginMutationVariables, TContext>
    ) => 
    useMutation<LoginMutation, TError, LoginMutationVariables, TContext>(
      (variables?: LoginMutationVariables) => fetcher<LoginMutation, LoginMutationVariables>(client, LoginDocument, variables)(),
      options
    );
export const LogoutDocument = `
    mutation Logout {
  logout
}
    `;
export const useLogoutMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient, 
      options?: UseMutationOptions<LogoutMutation, TError, LogoutMutationVariables, TContext>
    ) => 
    useMutation<LogoutMutation, TError, LogoutMutationVariables, TContext>(
      (variables?: LogoutMutationVariables) => fetcher<LogoutMutation, LogoutMutationVariables>(client, LogoutDocument, variables)(),
      options
    );
export const MeDocument = `
    query Me {
  me {
    id
    username
    email
    createdAt
    updatedAt
  }
}
    `;
export const useMeQuery = <
      TData = MeQuery,
      TError = unknown
    >(
      client: GraphQLClient, 
      variables?: MeQueryVariables, 
      options?: UseQueryOptions<MeQuery, TError, TData>
    ) => 
    useQuery<MeQuery, TError, TData>(
      ['Me', variables],
      fetcher<MeQuery, MeQueryVariables>(client, MeDocument, variables),
      options
    );
export const RegisterDocument = `
    mutation Register($input: Register!) {
  register(input: $input) {
    user {
      id
      username
      email
      createdAt
      updatedAt
    }
    error {
      field
      message
    }
  }
}
    `;
export const useRegisterMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient, 
      options?: UseMutationOptions<RegisterMutation, TError, RegisterMutationVariables, TContext>
    ) => 
    useMutation<RegisterMutation, TError, RegisterMutationVariables, TContext>(
      (variables?: RegisterMutationVariables) => fetcher<RegisterMutation, RegisterMutationVariables>(client, RegisterDocument, variables)(),
      options
    );