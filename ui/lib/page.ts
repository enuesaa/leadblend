import { page } from '$app/stores'
import { get } from 'svelte/store'

export const useParams = () => get(page).params
