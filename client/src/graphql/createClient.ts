import {request, GraphQLClient} from 'graphql-request'

const createClient = () => {
    const client = new GraphQLClient('http://localhost:4000/graphql')
    return client
}
