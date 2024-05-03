import { queryStore, mutationStore, gql } from '@urql/svelte'
import { useClient } from '$lib/graphql/client'
import { derived } from 'svelte/store'
import type { Planet, MutationCreatePlanetArgs } from './types'

const listPlanets = queryStore<{listPlanets: Array<Planet>}>({
  client: useClient(),
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

export const createPlanet = ({ name, comment }: MutationCreatePlanetArgs) => {
  const result = mutationStore({
    client: useClient(),
    query: gql`
      mutation ($name: String!, $comment: String!) {
        createPlanet(name: $name, comment: $comment)
      }
    `,
    variables: { name, comment },
  });
  console.log(result)
}
