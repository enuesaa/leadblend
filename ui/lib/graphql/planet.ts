import { runQuery, runMutation } from '$lib/graphql/client'
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
		queryKey: [listPlanetsQuery],
		queryFn: async () => {
			const res = await runQuery(listPlanetsQuery, {})
			return res.data.listPlanets
		},
		initialData: []
	})

const getPlanetQuery = `query ($name: String!) {
	getPlanet(name: $name) {
    id
    name
    comment
	}
}`

export const getPlanet = (name: string) =>
	createQuery<Planet>({
		queryKey: [getPlanetQuery, name],
		queryFn: async () => {
			const res = await runQuery(getPlanetQuery, { name })
			return res.data.getPlanet
		}
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
