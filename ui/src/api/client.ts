import { GraphQLClient } from "graphql-request";

const endpoint = process.env.REACT_APP_API_ENDPOINT || 'http://localhost:8080/graphql'

export const client = new GraphQLClient(endpoint)