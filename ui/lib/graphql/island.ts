import { get, mutate } from '$lib/graphql/client'
import type { Island, MutationCreateIslandArgs, MutationDeleteIslandArgs } from './types'

const listQuery = `query ($planetId: ID!) {
  listIslands (planetId: $planetId) {
    id
		planetId
		title
		content
    comment
  }
}`
export const listIslands = (planetId: string) =>
	get<Island[]>(listQuery, {
		vars: { planetId },
		usekey: 'listIslands',
		initialData: [],
	})

const createQuery = `mutation ($title: String!, $planetId: String!, $content: String!, $comment: String!) {
	createIsland(title: $title, planetId: $planetId, content: $content, comment: $comment)
}`
export const useCreateIsland = () =>
	mutate<MutationCreateIslandArgs>(createQuery, {
		usekey: 'createIsland',
	})

const deleteQuery = `mutation ($id: ID!) {
	deleteIsland(id: $id)
}`
export const useDeleteIsland = () =>
	mutate<MutationDeleteIslandArgs>(deleteQuery, {
		usekey: 'deleteIsland',
	})
