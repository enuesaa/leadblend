import { get, mutate } from './client'
import type { Comet, MutationCreateCometArgs, MutationDeleteCometArgs, MutationLinkCometArgs } from './types'

const listQuery = `query {
  listComets {
    id
    data
  }
}`
export const listComets = () => get<Comet[]>(listQuery, {
	usekey: 'listComets',
	initialData: [],
})

const getQuery = `query ($id: ID!) {
  getComet(id: $id) {
    id
    data
  }
}`
export const getComet = (id: string) => get<Comet>(getQuery, {
	vars: { id },
	usekey: 'getComet',
})

const createQuery = `mutation ($data: String!) {
	createComet(data: $data)
}`
export const useCreateComet = () => mutate<MutationCreateCometArgs>(createQuery, {
	usekey: 'createComet',
})

const linkQuery = `mutation ($id: ID!, $islandId: String!) {
	linkComet(id: $id, islandId: $islandId)
}`
export const useLinkComet = () => mutate<MutationLinkCometArgs>(linkQuery, {
	usekey: 'linkComet',
})

const deleteQuery = `mutation ($id: ID!) {
  deleteComet(id: $id)
}`
export const useDeleteComet = () => mutate<MutationDeleteCometArgs>(deleteQuery, {
	usekey: 'deleteComet',
})
