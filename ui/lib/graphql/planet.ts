import { gql, executeQuery, executeMutation } from '$lib/graphql/client'
import type { Planet, MutationCreatePlanetArgs } from './types'
import { createMutation, createQuery } from '@tanstack/svelte-query'

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

export const useCreatePlanet = () => createMutation({
  mutationFn: async (data: MutationCreatePlanetArgs) => {
    const query = gql`
      mutation ($name: String!, $comment: String!) {
        createPlanet(name: $name, comment: $comment)
      }
    `
    const res = await executeMutation(query, data)
    return res.data?.createPlanet
  }
})
