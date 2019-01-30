import { actionTypes } from '../constants/action-types'
import { SwitchIsLoadingAction, CurrentUserAction } from '../actions/users'

// TODO: stateの型付け厳格に
export function currentUser(state: any = {}, action: CurrentUserAction) {
  switch (action.type) {
    case actionTypes.SET_CURRENT_USER:
      if ('currentUser' in action.payload) {
        return action.payload.currentUser
      }
      break

    case actionTypes.UPDATE_CURRENT_USER:
      if ('updatedUser' in action.payload) {
        return action.payload.updatedUser
      }
      break

    case actionTypes.DELETE_SESSION:
      if ('currentUser' in action.payload) {
        return null
      }
      break

    default:
      return state
  }
}

export function isLoadingCurrentUser(state: boolean = false, action: SwitchIsLoadingAction) {
  switch (action.type) {
    case actionTypes.CURRENT_USER_SET_IS_LOADING:
      return action.payload.isLoading

    default:
      return state
  }
}
