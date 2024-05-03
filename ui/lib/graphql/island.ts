import { runQuery, runMutation } from '$lib/graphql/client'
import type { Island, MutationCreateIslandArgs } from './types'
import { createMutation, createQuery } from '@tanstack/svelte-query'

const createIslandQuery = `mutation ($title: String!, $planetId: String!, $content: String!, $comment: String!) {
  createIsland(title: $title, planetId: $planetId, content: $content, comment: $comment)
}`

export const useCreateIsland = () =>
	createMutation({
		mutationFn: async (data: MutationCreateIslandArgs) => {
			const res = await runMutation(createIslandQuery, data)
			return res.data?.createIsland
		}
	})
