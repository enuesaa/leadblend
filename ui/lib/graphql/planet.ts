import { executeQuery, executeMutation, states, runQuery } from '$lib/graphql/client'
import { derived } from 'svelte/store'
import type { Planet, MutationCreatePlanetArgs } from './types'
import { createMutation, createQuery } from '@tanstack/svelte-query'

const listPlanetsQuery = `query {
  listPlanets {
    id
    name
    comment
  }
}`
export const listPlanets = () => {
  const cacheKey = runQuery(listPlanetsQuery, {})
  return derived(states, (val): Planet[] => {
    return val.hasOwnProperty(cacheKey) ? val[cacheKey]?.data?.listPlanets ?? [] : []
  })
}


// export const listPlanets = () => createQuery<Planet[]>({
//   queryKey: [listPlanetsQuery],
//   queryFn: async () => {
//     const res = await executeQuery(listPlanetsQuery)
//     return res.data.listPlanets
//   },
//   initialData: [],
// })


const createPlanetQuery = `mutation ($name: String!, $comment: String!) {
  createPlanet(name: $name, comment: $comment)
}`

export const useCreatePlanet = () => createMutation({
  mutationFn: async (data: MutationCreatePlanetArgs) => {
    const res = await executeMutation(createPlanetQuery, data)
    return res.data?.createPlanet
  }
})
