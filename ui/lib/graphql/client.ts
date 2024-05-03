import { writable } from 'svelte/store'
import { Client, cacheExchange, fetchExchange } from '@urql/svelte'
import { PUBLIC_GRAPHQL_ENDPOINT } from '$env/static/public'

const client = new Client({
  url: PUBLIC_GRAPHQL_ENDPOINT,
  exchanges: [cacheExchange, fetchExchange],
})
export const clientStore = writable<Client>(client)
