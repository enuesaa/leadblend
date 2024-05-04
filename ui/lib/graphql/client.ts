import { Client, createRequest, fetchExchange, type OperationResult, gql } from '@urql/svelte'
import { PUBLIC_GRAPHQL_ENDPOINT } from '$env/static/public'
import { createQuery } from '@tanstack/svelte-query'

export const client = new Client({
	url: PUBLIC_GRAPHQL_ENDPOINT,
	exchanges: [fetchExchange]
})

export const runQuery = async (query: string, variables: Record<string, any>): Promise<OperationResult> => {
	const request = createRequest(gql(query), variables)
	const res = await client.executeQuery(request, {})
	return res
}

export const runMutation = async (query: string, variables: Record<string, any>) => {
	const request = createRequest(gql(query), variables)
	const res = await client.executeMutation(request, {})
	return res
}

const calcCacheKey = (query: string, variables: Record<string, any>): string[] => {
  return [`${query}-${JSON.stringify(variables)}`]
}

type QueOptions = {
	vars: Record<string, any>;
	usekey: string;
	initialData: any;
}
export const que = <T>(query: string, options: Partial<QueOptions> = {}) => {
	const vars = options.vars ?? {}
	const usekey = options.usekey ?? ''
	const initialData = options.initialData ?? undefined

	const creation = createQuery<T>({
		queryKey: calcCacheKey(query, vars),
		queryFn: async () => {
			const request = createRequest(gql(query), vars)
			const res = await client.executeQuery(request, {})
			if (res.data === undefined || res.data === null) {
				return res.data
			}
			if (res.data.hasOwnProperty(usekey)) {
				return res.data[usekey]
			}
			return {}
		},
		initialData: initialData,
	})

	return creation
}
