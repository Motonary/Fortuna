import { actionTypes } from '../constants'
import { User } from '../constants/static-types'
import { CurrentUserAction } from '../actions/users'
import { ReduxAPIStruct, defaultSet } from './static-types'

export const currentUser = (
  state: ReduxAPIStruct<User> = defaultSet(),
  action: CurrentUserAction
): ReduxAPIStruct<User> => {
  switch (action.type) {
    case actionTypes.USER_API_REQUEST:
      return { ...state, status: 'fetching' }

    case actionTypes.USER_API_FAILURE:
      return { ...state, status: 'failure', error: action.payload.error }

    case actionTypes.SET_CURRENT_USER:
      return { ...state, status: 'success', data: action.payload.currentUser }

    case actionTypes.UPDATE_CURRENT_USER:
      return { ...state, status: 'success', data: action.payload.updatedUser }

    case actionTypes.DELETE_SESSION:
      return { ...state, status: 'success', data: null }
  }
  return state
}
