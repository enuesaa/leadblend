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

const getCometQuery = `query ($id: ID!) {
  getComet(id: $id) {
    id
    data
  }
}`

export const getComet = (id: string) =>
	createQuery<Comet>({
		queryKey: [getCometQuery, id],
		queryFn: async () => {
			const res = await runQuery(getCometQuery, { id })
			return res.data.getComet
		}
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
