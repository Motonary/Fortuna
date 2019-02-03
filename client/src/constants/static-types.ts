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

export type StatusCodeSuccess = 200 | 204

export type StatusCodeFailure = 400 | 401 | 403 | 404 | 422 | 500 | 503

export type StatusCode = StatusCodeSuccess | StatusCodeFailure
