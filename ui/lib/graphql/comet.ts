import { createMutation, createQuery } from '@tanstack/svelte-query'
import { runMutation, runQuery } from './client'
import type { Comet, MutationCreateCometArgs } from './types'

const listCometsQuery = `query {
  listComets {
    id
    data
  }
}`

export const listComets = () =>
	createQuery<Comet[]>({
		queryKey: [listComets],
		queryFn: async () => {
			const res = await runQuery(listCometsQuery, {})
			return res.data.listComets
		},
		initialData: []
	})

const createCometQuery = `mutation ($data: String!) {
	createComet(data: $data)
}`

export const useCreateComet = () =>
	createMutation({
		mutationFn: async (data: MutationCreateCometArgs) => {
			const res = await runMutation(createCometQuery, data)
			return res.data?.createIsland
		}
	})
