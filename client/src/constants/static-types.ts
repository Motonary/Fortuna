export interface BaseAction {
  type: string
  payload?: any
}

export interface User {
  id: number
  name: string
  email: string
  password: string
  createdAt: any // string?
  updatedAt: any // string?
}
