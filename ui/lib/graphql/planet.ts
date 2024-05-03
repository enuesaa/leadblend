import { runQuery, runMutation, cachekey } from '$lib/graphql/client'
import type { Planet, MutationCreatePlanetArgs } from './types'
import { createMutation, createQuery } from '@tanstack/svelte-query'

const listPlanetsQuery = `query {
  listPlanets {
    id
    name
    comment
  }
}`

export const listPlanets = () =>
	createQuery<Planet[]>({
		queryKey: cachekey(listPlanetsQuery),
		queryFn: async () => {
			const res = await runQuery(listPlanetsQuery, {})
			return res.data.listPlanets
		},
		initialData: []
	})

const createPlanetQuery = `mutation ($name: String!, $comment: String!) {
  createPlanet(name: $name, comment: $comment)
}`

export const useCreatePlanet = () =>
	createMutation({
		mutationFn: async (data: MutationCreatePlanetArgs) => {
			const res = await runMutation(createPlanetQuery, data)
			return res.data?.createPlanet
		}
	})
