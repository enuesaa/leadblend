import { get } from '$lib/graphql/client'
import type { Space } from './types'

const getCurrentSpaceQuery = `query {
  getCurrentSpace {
    name
  }
}`

export const getCurrentSpace = () =>
	get<Space>(getCurrentSpaceQuery, {
		usekey: 'getCurrentSpace',
	})
