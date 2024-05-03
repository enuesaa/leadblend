import { mutationStore, gql } from '@urql/svelte'
import { client, executeQuery } from '$lib/graphql/client'
import type { Planet, MutationCreatePlanetArgs } from './types'
import { createQuery } from '@tanstack/svelte-query'

export const listPlanets = () => createQuery<Planet[]>({
  queryKey: ['listPlanets'],
  queryFn: async () => {
    const query = gql`
      query {
        listPlanets {
          id
          name
          comment
        }
      }
    `
    const res = await executeQuery(query)
    return res.data.listPlanets
  },
  initialData: [],
})

export const createPlanet = ({ name, comment }: MutationCreatePlanetArgs) => {
  const result = mutationStore({
    client,
    query: gql`
      mutation ($name: String!, $comment: String!) {
        createPlanet(name: $name, comment: $comment)
      }
    `,
    variables: { name, comment },
  });
  console.log(result)
}
