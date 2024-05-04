import { runMutation, que } from '$lib/graphql/client'
import type { Planet, MutationCreatePlanetArgs } from './types'
import { createMutation } from '@tanstack/svelte-query'

const listQuery = `query {
  listPlanets {
    id
    name
    comment
  }
}`

export const listPlanets = () => que<Planet[]>(listQuery, {
	usekey: 'listPlanets',
	initialData: [],
})

const getQuery = `query ($name: String!) {
	getPlanet(name: $name) {
    id
    name
    comment
	}
}`

export const getPlanet = (name: string) => que<Planet>(getQuery, {
	vars: {name},
	usekey: 'getPlanet',
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
