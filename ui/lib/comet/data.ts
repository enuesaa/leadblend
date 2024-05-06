export const convertCometData = (jsondata: string): [CometObject, string] => {
  const data: CometObject = {
    type: 'object',
    key: '',
    values: [],
  }

  try {
    const parsed = JSON.parse(jsondata)
    for (const [key, value] of Object.entries(parsed)) {
      if (typeof value === 'string') {
        data.values.push({
          type: 'string',
          key,
          value,
        })
        continue
      }
      if (typeof value === 'number') {
        data.values.push({
          type: 'number',
          key,
          value,
        })
        continue
      }
      if (typeof value === 'boolean') {
        data.values.push({
          type: 'boolean',
          key,
          value,
        })
        continue
      }
      if (value === null) {
        data.values.push({
          type: 'null',
          key,
          value,
        })
        continue
      }
      if (Array.isArray(value)) {
        // todo
        continue
      }
      // todo
      continue
    }
  } catch (err) {
    return [data, 'invalid json format']
  }

  return [data, '']
}

export type CometObject = {
  type: 'object'
  key: string
  values: (CometValue|CometObject|CometArray)[]
}

type CometArray = {
  type: 'array'
  key: string
  values: (CometValue|CometObject|CometArray)[]
}

type CometValue = CometNumber|CometBoolean|CometNull|CometString

type CometNumber = {
  type: 'number'
  key: string
  value: number
}
type CometBoolean = {
  type: 'boolean'
  key: string
  value: boolean
}
type CometNull = {
  type: 'null'
  key: string
  value: null
}
type CometString = {
  type: 'string'
  key: string
  value: string
}
