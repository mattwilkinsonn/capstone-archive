import { GraphQLClient } from 'graphql-request'

export const createClient = (): GraphQLClient => {
  const api = process.env.REACT_APP_API_URL
  if (typeof api != 'string') throw new Error('API variable missing')

  const client = new GraphQLClient(api, {
    credentials: 'include',
  })
  return client
}
