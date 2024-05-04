import { createMutation } from '@tanstack/svelte-query'
import { que, runMutation } from './client'
import type { Comet, MutationCreateCometArgs, MutationDeleteCometArgs } from './types'

const listQuery = `query {
  listComets {
    id
    data
  }
}`

export const listComets = () => que<Comet[]>(listQuery, {
	usekey: 'listComets',
	initialData: [],
})

const getQuery = `query ($id: ID!) {
  getComet(id: $id) {
    id
    data
  }
}`

export const getComet = (id: string) => que<Comet>(getQuery, {
	vars: { id },
	usekey: 'getComet',
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

const deleteCometQuery = `mutation ($id: ID!) {
  deleteComet(id: $id)
}`

export const useDeleteComet = () =>
  createMutation({
    mutationFn: async (data: MutationDeleteCometArgs) => {
      const res = await runMutation(deleteCometQuery, data)
      return res.data?.deleteComet
    }
  })
  