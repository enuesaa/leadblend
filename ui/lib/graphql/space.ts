import { runQuery } from '$lib/graphql/client'
import type { Space } from './types'
import { createQuery } from '@tanstack/svelte-query'

const getCurrentSpaceQuery = `query {
  getCurrentSpace {
    name
  }
}`

export const getCurrentSpace = () => createQuery<Space>({
  queryKey: [getCurrentSpaceQuery],
  queryFn: async () => {
    const res = await runQuery(getCurrentSpaceQuery, {})
    return res.data.getCurrentSpace
  },
})
