import { GraphQLClient } from 'graphql-request'

export const createClient = () => {
  const client = new GraphQLClient('http://localhost:4000/graphql')
  return client
}
