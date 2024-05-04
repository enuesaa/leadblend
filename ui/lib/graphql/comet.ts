import { createMutation } from '@tanstack/svelte-query'
import { runMutation } from './client'
import type { MutationCreateCometArgs } from './types'

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
