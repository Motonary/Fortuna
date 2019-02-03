import { actionTypes } from '../constants/action-types'
import { User } from '../constants/static-types'
import { SwitchIsLoadingAction, CurrentUserAction } from '../actions/users'
import { ReduxAPIStruct, defaultSet } from './reducer-type'

export function CurrentUser(state: ReduxAPIStruct<User> = defaultSet(), action: CurrentUserAction) {
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
      return null

    default:
      return state
  }
}
