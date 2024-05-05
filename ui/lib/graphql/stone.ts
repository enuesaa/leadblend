import { get } from '$lib/graphql/client'
import type { Stone } from './types'

const listQuery = `query ($islandId: ID!) {
  listStones (islandId: $islandId) {
    id
		data
		islandId
    patternId
  }
}`
export const listStones = (islandId: string) => get<Stone[]>(listQuery, {
	vars: {islandId},
	usekey: 'listStones',
	initialData: [],
})
