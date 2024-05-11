import { get, mutate } from './client'
import type { MutationCreatePatternArgs, Pattern, Trait } from './types'

const listQuery = `query {
  listPatterns {
    id
    title
    color
    traits {
      defaultValue
      path
      required
      type
    }
  }
}`
export const listPatterns = () =>
	get<Pattern[]>(listQuery, {
		usekey: 'listPatterns',
		initialData: [],
	})

const createQuery = `mutation ($title: String!, $color: String, $traits: [TraitInput!]!) {
  createPattern(title: $title, color: $color, traits: $traits)
}`
export const useCreatePattern = () =>
	mutate<MutationCreatePatternArgs>(createQuery, {
		usekey: 'createPattern',
	})
