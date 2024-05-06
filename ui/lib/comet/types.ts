export type CometObject = {
  type: 'object'
  key: string
  values: (CometValue|CometObject|CometArray)[]
}

export type CometArray = {
  type: 'array'
  key: string
  values: (CometValue|CometObject|CometArray)[]
}

export type CometValue = CometNumber|CometBoolean|CometNull|CometString

export type CometNumber = {
  type: 'number'
  key: string
  value: number
}

export type CometBoolean = {
  type: 'boolean'
  key: string
  value: boolean
}

export type CometNull = {
  type: 'null'
  key: string
  value: null
}

export type CometString = {
  type: 'string'
  key: string
  value: string
}
