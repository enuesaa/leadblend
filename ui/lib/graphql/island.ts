import { runQuery, runMutation } from '$lib/graphql/client'
import type { Island, MutationCreateIslandArgs, MutationDeleteIslandArgs } from './types'
import { createMutation, createQuery } from '@tanstack/svelte-query'

const listIslandsQuery = `query ($planetId: ID!) {
  listIslands (planetId: $planetId) {
    id
		planetId
		title
		content
    comment
  }
}`

export const listIslands = (planetId: string) =>
	createQuery<Island[]>({
		queryKey: [listIslandsQuery],
		queryFn: async () => {
			const res = await runQuery(listIslandsQuery, {planetId})
			return res.data.listIslands
		},
		initialData: []
	})


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

const deleteIslandQuery = `mutation ($id: ID!) {
	deleteIsland(id: $id)
}`

export const useDeleteIsland = () =>
	createMutation({
		mutationFn: async (data: MutationDeleteIslandArgs) => {
			const res = await runMutation(deleteIslandQuery, data)
			return res.data?.deleteIsland
		}
	})
	