import { actionTypes } from '../constants/action-types'
import { CurrentUserAction } from '../actions/users'

export function currentUser(state: any = null, action: CurrentUserAction) {
  switch (action.type) {
    case actionTypes.SET_USER_SESSION:
      return action.payload.currentUser

    case actionTypes.DELETE_USER_SESSION:
      return null

    case actionTypes.UPDATE_USER:
      return action.payload.updatedUser

    default:
      return state
  }
}
