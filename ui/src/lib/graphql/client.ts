import { writable } from 'svelte/store'
import { Client, cacheExchange, fetchExchange } from '@urql/svelte'

const client = new Client({
  url: 'http://localhost:3000/graphql',
  exchanges: [cacheExchange, fetchExchange],
})
export const clientStore = writable<Client>(client)
