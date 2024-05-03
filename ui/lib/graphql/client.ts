import { Client, createRequest, fetchExchange, type OperationResult, type TypedDocumentNode } from '@urql/svelte'
import { PUBLIC_GRAPHQL_ENDPOINT } from '$env/static/public'

export { gql } from '@urql/svelte'

export const client = new Client({
  url: PUBLIC_GRAPHQL_ENDPOINT,
  exchanges: [fetchExchange],
})

export const executeQuery = async (query: TypedDocumentNode): Promise<OperationResult> => {
  const request = createRequest(query, {})
  const res = await client.executeQuery(request, {})
  return res
}

export const executeMutation = async (query: TypedDocumentNode, variables: Record<string, any>) => {
  const request = createRequest(query, variables)
  const res = await client.executeMutation(request, {})
  return res
}
