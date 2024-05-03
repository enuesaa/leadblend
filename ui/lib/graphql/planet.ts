import { queryStore, gql } from '@urql/svelte'
import { clientStore } from '$lib/graphql/client'
import { derived, get } from 'svelte/store'
import type { Planet } from './types'

const listPlanets = queryStore<{listPlanets: Array<Planet>}>({
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

export const planets = derived(listPlanets, ($listPlanets) => {
  if ($listPlanets.data === undefined || $listPlanets.data === null) {
    return []
  }
  return $listPlanets.data.listPlanets
})
