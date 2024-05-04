import { get } from './client'
import type { Pattern, Trait } from './types'

const listQuery = `query {
  listPatterns {
    id
    title
    traits {
      defaultValue
      path
      required
      type
    }
  }
}`
export const listPatterns = () => get<Pattern[]>(listQuery, {
	usekey: 'listPatterns',
	initialData: [],
})
