import { Client, createRequest, fetchExchange, type OperationResult, gql } from '@urql/svelte'
import { PUBLIC_GRAPHQL_ENDPOINT } from '$env/static/public'

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

export const cachekey = (query: string, variables: Record<string, any> = {}): string[] => {
	return [`${query}-${JSON.stringify(variables)}`]
}
