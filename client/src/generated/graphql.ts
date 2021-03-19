import { GraphQLClient } from 'graphql-request'
import {
  useQuery,
  UseQueryOptions,
  useMutation,
  UseMutationOptions,
} from 'react-query'
export type Maybe<T> = T | null
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K]
}
export type MakeOptional<T, K extends keyof T> = Omit<T, K> &
  { [SubKey in K]?: Maybe<T[SubKey]> }
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> &
  { [SubKey in K]: Maybe<T[SubKey]> }

function fetcher<TData, TVariables>(
  client: GraphQLClient,
  query: string,
  variables?: TVariables
) {
  return async (): Promise<TData> =>
    client.request<TData, TVariables>(query, variables)
}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string
  String: string
  Boolean: boolean
  Int: number
  Float: number
}

export type NewCapstone = {
  title: Scalars['String']
  description: Scalars['String']
  author: Scalars['String']
}

export type UserError = {
  __typename?: 'UserError'
  field: Scalars['String']
  message: Scalars['String']
}

export type User = {
  __typename?: 'User'
  id: Scalars['ID']
  username: Scalars['String']
  email: Scalars['String']
  createdAt: Scalars['Int']
  updatedAt: Scalars['Int']
}

export type PaginatedCapstones = {
  __typename?: 'PaginatedCapstones'
  capstones: Array<Maybe<Capstone>>
  hasMore: Scalars['Boolean']
}

export type Query = {
  __typename?: 'Query'
  searchCapstones: PaginatedCapstones
  capstones: PaginatedCapstones
  users: Array<PublicUser>
  me?: Maybe<User>
}

export type QuerySearchCapstonesArgs = {
  query: Scalars['String']
  limit: Scalars['Int']
  offset?: Maybe<Scalars['Int']>
}

export type QueryCapstonesArgs = {
  limit: Scalars['Int']
  cursor?: Maybe<Scalars['Int']>
}

export type Mutation = {
  __typename?: 'Mutation'
  createCapstone: Capstone
  register: UserResponse
  login: UserResponse
  logout: Scalars['Boolean']
}

export type MutationCreateCapstoneArgs = {
  input: NewCapstone
}

export type MutationRegisterArgs = {
  input: Register
}

export type MutationLoginArgs = {
  input: Login
}

export type Register = {
  username: Scalars['String']
  email: Scalars['String']
  password: Scalars['String']
}

export type Capstone = {
  __typename?: 'Capstone'
  id: Scalars['ID']
  title: Scalars['String']
  description: Scalars['String']
  author: Scalars['String']
  createdAt: Scalars['Int']
  updatedAt: Scalars['Int']
  semester: Scalars['String']
}

export type Login = {
  usernameOrEmail: Scalars['String']
  password: Scalars['String']
}

export type Todo = {
  __typename?: 'Todo'
  id: Scalars['ID']
  text: Scalars['String']
  done: Scalars['Boolean']
  user: User
}

export type UserResponse = {
  __typename?: 'UserResponse'
  user?: Maybe<User>
  error?: Maybe<UserError>
}

export type PublicUser = {
  __typename?: 'PublicUser'
  username: Scalars['String']
}

export type CapstonesQueryVariables = Exact<{
  limit: Scalars['Int']
  cursor?: Maybe<Scalars['Int']>
}>

export type CapstonesQuery = { __typename?: 'Query' } & {
  capstones: { __typename?: 'PaginatedCapstones' } & Pick<
    PaginatedCapstones,
    'hasMore'
  > & {
      capstones: Array<
        Maybe<
          { __typename?: 'Capstone' } & Pick<
            Capstone,
            'id' | 'title' | 'description' | 'createdAt' | 'semester'
          >
        >
      >
    }
}

export type LoginMutationVariables = Exact<{
  input: Login
}>

export type LoginMutation = { __typename?: 'Mutation' } & {
  login: { __typename?: 'UserResponse' } & {
    user?: Maybe<
      { __typename?: 'User' } & Pick<
        User,
        'id' | 'username' | 'email' | 'createdAt' | 'updatedAt'
      >
    >
    error?: Maybe<
      { __typename?: 'UserError' } & Pick<UserError, 'field' | 'message'>
    >
  }
}

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
    `
export const useCapstonesQuery = <TData = CapstonesQuery, TError = unknown>(
  client: GraphQLClient,
  variables: CapstonesQueryVariables,
  options?: UseQueryOptions<CapstonesQuery, TError, TData>
) =>
  useQuery<CapstonesQuery, TError, TData>(
    ['Capstones', variables],
    fetcher<CapstonesQuery, CapstonesQueryVariables>(
      client,
      CapstonesDocument,
      variables
    ),
    options
  )
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
    `
export const useLoginMutation = <TError = unknown, TContext = unknown>(
  client: GraphQLClient,
  options?: UseMutationOptions<
    LoginMutation,
    TError,
    LoginMutationVariables,
    TContext
  >
) =>
  useMutation<LoginMutation, TError, LoginMutationVariables, TContext>(
    (variables?: LoginMutationVariables) =>
      fetcher<LoginMutation, LoginMutationVariables>(
        client,
        LoginDocument,
        variables
      )(),
    options
  )
