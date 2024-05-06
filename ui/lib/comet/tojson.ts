import { type CometArray, type CometObject, type CometValue } from './types'

export const convertCometDataToJson = (cometdata: CometObject): string => {
  const jsondata = convertObject(cometdata)
  return JSON.stringify(jsondata)
}

const convertArray = (data: CometArray): any[] => {
  const ret: any[] = []

  for (const value of data.values) {
    if (value.type === 'array') {
      ret.push(convertArray(value))
    } else if (value.type === 'object') {
      ret.push(convertObject(value))
    } else {
      ret.push(value.value)
    }
  }
  return ret
}

const convertObject = (data: CometObject): any => {
  const ret: any = {}

  for (const value of data.values) {
    if (value.type === 'array') {
      ret[value.key] = convertArray(value)
    } else if (value.type === 'object') {
      ret[value.key] = convertObject(value)
    } else {
      ret[value.key] = value.value
    }
  }
  return ret
}
