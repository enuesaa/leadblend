import { describe, expect, it } from 'vitest'
import { convertCometData } from './data'

describe('convertCometData', () => {
	it('empty object', () => {
		expect(convertCometData('{}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [],
		})
	})

	it('key value', () => {
		expect(convertCometData('{"a":"b", "c":"d"}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [
				{ type: 'string', key: 'a', value: 'b' },
				{ type: 'string', key: 'c', value: 'd' },
			],
		})
	})

	it('string value', () => {
		expect(convertCometData('{"a":"b"}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [{ type: 'string', key: 'a', value: 'b' }],
		})
	})

	it('number value', () => {
		expect(convertCometData('{"a":8}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [{ type: 'number', key: 'a', value: 8 }],
		})
	})

	it('boolean value', () => {
		expect(convertCometData('{"a":true}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [{ type: 'boolean', key: 'a', value: true }],
		})
	})

	it('null value', () => {
		expect(convertCometData('{"a":null}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [{ type: 'null', key: 'a', value: null }],
		})
	})

	it('object values', () => {
		expect(convertCometData('{"a":{"b":"c"}}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [{ type: 'object', key: 'a', values: [{ type: 'string', key: 'b', value: 'c' }] }],
		})
	})

	it('array values', () => {
		expect(convertCometData('{"a":["b", "c"]}')[0]).toStrictEqual({
			type: 'object',
			key: '',
			values: [
				{
					type: 'array',
					key: 'a',
					values: [
						{ type: 'string', key: '0', value: 'b' },
						{ type: 'string', key: '1', value: 'c' },
					],
				},
			],
		})
	})
})
