import { Client, createRequest, fetchExchange, type OperationResult, gql } from '@urql/svelte'
import { PUBLIC_GRAPHQL_ENDPOINT } from '$env/static/public'

export const client = new Client({
  url: PUBLIC_GRAPHQL_ENDPOINT,
  exchanges: [fetchExchange],
})

export const executeQuery = async (query: string): Promise<OperationResult> => {
  const request = createRequest(gql(query), {})
  const res = await client.executeQuery(request, {})
  return res
}

export const executeMutation = async (query: string, variables: Record<string, any>) => {
  const request = createRequest(gql(query), variables)
  const res = await client.executeMutation(request, {})
  return res
}
