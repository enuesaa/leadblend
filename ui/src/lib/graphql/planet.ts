import { queryStore, gql } from '@urql/svelte'
import { clientStore } from '$lib/graphql/client'
import { get } from 'svelte/store'

export const listPlanets = queryStore({
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
