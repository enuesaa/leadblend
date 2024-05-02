import { queryStore, gql } from '@urql/svelte'
import { clientStore } from '$lib/graphql/client'
import { get } from 'svelte/store'
import type { Planet } from './types'

export const listPlanets = queryStore<{listPlanets: Array<Planet>}>({
  client: get(clientStore),
  query: gql`
    query {
      listPlanets {
        id
        name
        comment
      }
    }
  `,
})
