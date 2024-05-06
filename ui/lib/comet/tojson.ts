import { type CometObject } from './types'

export const convertCometDataToJson = (cometdata: CometObject): string => {
  const jsondata: any = {}

  for (const child of cometdata.values) {
    if (child.type === 'array') {
      continue;
    }
    if (child.type === 'object') {
      continue;
    }
    jsondata[child.key] = child.value
  }

  return JSON.stringify(jsondata)
}
