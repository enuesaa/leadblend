import { get, mutate } from '$lib/graphql/client'
import type { MutationDeleteStoneArgs, Stone } from './types'

const listQuery = `query ($islandId: ID!) {
  listStones (islandId: $islandId) {
    id
		data
		islandId
    patternId
  }
}`
export const listStones = (islandId: string) =>
	get<Stone[]>(listQuery, {
		vars: { islandId },
		usekey: 'listStones',
		initialData: [],
	})

const deleteQuery = `mutation ($id: ID!) {
  deleteStone(id: $id)
}`
export const useDeleteStone = () =>
	mutate<MutationDeleteStoneArgs>(deleteQuery, {
		usekey: 'deleteStone',
	})
