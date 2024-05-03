import { Client, createRequest, fetchExchange, type OperationResult, gql } from '@urql/svelte'
import { PUBLIC_GRAPHQL_ENDPOINT } from '$env/static/public'
import { writable } from 'svelte/store'

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

export const states = writable<Record<string, any>>({})

const calcCacheKey = (query: string, variables: Record<string, any>): string => {
  return `${query}-${JSON.stringify(variables)}`
}

export const runQuery = (query: string, variables: Record<string, any>): string => {
  const cacheKey = calcCacheKey(query, variables)

  const request = createRequest(gql(query), variables)
  client.executeQuery(request, {}).then(res => {
    states.update(val => ({...val, [cacheKey]: res}))
  })

  return cacheKey
}
