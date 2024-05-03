import { queryStore, gql } from '@urql/svelte'
import { useClient } from '$lib/graphql/client'
import { derived } from 'svelte/store'
import type { Space } from './types'

const getCurrentSpace = queryStore<{getCurrentSpace: Space}>({
  client: useClient(),
  query: gql`
    query {
      getCurrentSpace {
        name
      }
    }
  `
})

export const currentSpace = derived(getCurrentSpace, ($getCurrentSpace) => {
  if ($getCurrentSpace.data === undefined) {
    return {name: ''}
  }
  return $getCurrentSpace.data.getCurrentSpace
})
