import { writable } from 'svelte/store'

export const cometVisible = writable<boolean>(true)
export const showComet = () => cometVisible.set(true)
export const hideComet = () => cometVisible.set(false)
