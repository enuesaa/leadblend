import type { CometArray, CometObject, CometValue } from './types'

export const convertCometData = (jsondata: string): [CometObject, string] => {
	let data: CometObject = {
		type: 'object',
		key: '',
		values: [],
	}

	try {
		const parsed = JSON.parse(jsondata)
		data = convertObject('', parsed)
	} catch (err) {
		console.log(err)
		return [data, 'invalid json format']
	}

	return [data, '']
}

const convertObject = (key: string, values: any): CometObject => {
	const data: CometObject = {
		type: 'object',
		key,
		values: [],
	}
	for (const [k, value] of Object.entries(values)) {
		if (typeof value === 'string' || typeof value === 'boolean' || typeof value === 'number' || value === null) {
			data.values.push(convertScalar(k, value))
		} else if (Array.isArray(value)) {
			data.values.push(convertArray(k, value))
		} else {
			data.values.push(convertObject(k, value))
		}
	}
	return data
}

const convertScalar = (key: string, value: string | number | boolean | null): CometValue => {
	if (typeof value === 'string') {
		return { type: 'string', key, value }
	}
	if (typeof value === 'number') {
		return { type: 'number', key, value }
	}
	if (typeof value === 'boolean') {
		return { type: 'boolean', key, value }
	}
	return { type: 'null', key, value }
}

const convertArray = (key: string, values: any[]): CometArray => {
	const data: CometArray = {
		type: 'array',
		key,
		values: [],
	}
	let k = 0
	for (const value of values) {
		if (typeof value === 'string' || typeof value === 'boolean' || typeof value === 'number' || value === null) {
			data.values.push(convertScalar(`${k}`, value))
		} else if (Array.isArray(value)) {
			data.values.push(convertArray(`${k}`, value))
		} else {
			data.values.push(convertObject(`${k}`, value))
		}
		k++
	}
	return data
}
