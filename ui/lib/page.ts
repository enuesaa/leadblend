import { page } from '$app/stores'
import { get } from 'svelte/store'

/**
 * @todo fix
 * 
 * now working after route change.
 */
export const useParams = () => get(page).params
