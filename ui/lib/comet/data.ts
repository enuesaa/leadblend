export const convertCometData = (jsondata: string): CometObject => {
  const data: CometObject = {
    type: 'object',
    values: [],
  }

  try {
    const parsed = JSON.parse(jsondata)
    for (const [key, value] of Object.entries(parsed)) {
      if (typeof value === 'string') {
        data.values.push({
          type: 'string',
          value,
        })
        continue
      }
      if (typeof value === 'number') {
        data.values.push({
          type: 'number',
          value,
        })
        continue
      }
      if (typeof value === 'boolean') {
        data.values.push({
          type: 'boolean',
          value,
        })
        continue
      }
      if (value === null) {
        data.values.push({
          type: 'null',
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
  } catch (err) {}

  return data
}

export type CometObject = {
  type: 'object'
  values: (CometValue|CometObject|CometArray)[]
}

type CometArray = {
  type: 'array'
  values: (CometValue|CometObject|CometArray)[]
}

type CometValue = CometNumber|CometBoolean|CometNull|CometString

type CometNumber = {
  type: 'number'
  value: number
}
type CometBoolean = {
  type: 'boolean'
  value: boolean
}
type CometNull = {
  type: 'null'
  value: null
}
type CometString = {
  type: 'string'
  value: string
}
