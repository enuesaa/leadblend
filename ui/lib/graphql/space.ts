import { gql, executeQuery } from '$lib/graphql/client'
import type { Space } from './types'
import { createQuery } from '@tanstack/svelte-query'

export const getCurrentSpace = () => createQuery<Space>({
  queryKey: ['getCurrentSpace'],
  queryFn: async () => {
    const query = gql`
      query {
        getCurrentSpace {
          name
        }
      }
    `
    const res = await executeQuery(query)
    return res.data.getCurrentSpace
  },
})
