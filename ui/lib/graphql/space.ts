import { queryStore, gql } from '@urql/svelte'
import { clientStore } from '$lib/graphql/client'
import { derived, get } from 'svelte/store'
import type { Space } from './types'

const getCurrentSpace = queryStore<{getCurrentSpace: Space}>({
  client: get(clientStore),
  query: gql`
    query {
      getCurrentSpace {
        name
      }
    }
  `,
})

export const currentSpace = derived(getCurrentSpace, ($getCurrentSpace) => {
  if ($getCurrentSpace.data === undefined) {
    return {name: ''}
  }
  return $getCurrentSpace.data.getCurrentSpace
})
