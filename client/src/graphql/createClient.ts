import { GraphQLClient } from 'graphql-request'

export const createClient = (): GraphQLClient => {
  const client = new GraphQLClient('http://localhost:4000/graphql')
  return client
}
