export interface ReduxAPIError {
  readonly statusCode: string
  readonly message: string
}

export interface ReduxAPIStruct<T> {
  readonly status: 'default' | 'fetching' | 'success' | 'failure'
  readonly data: T | null
  readonly error: ReduxAPIError
}

export const defaultSet = <T>(defaultValue?: T): ReduxAPIStruct<T> => ({
  status: 'default',
  data: defaultValue || null,
  error: errorDefault(),
})

export const errorDefault = (): ReduxAPIError => ({
  statusCode: 'default',
  message: '',
})
