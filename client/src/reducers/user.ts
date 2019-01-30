import { actionTypes } from '../constants/action-types'
// import { CurrentUserAction } from '../actions/users'

export function currentUser(state: any = null, action: any) {
  switch (action.type) {
    case actionTypes.SET_CURRENT_USER:
    case actionTypes.UPDATE_CURRENT_USER:
      return action.payload.currentUser

    default:
      return state
  }
}

export function IsLoadingCurrentUser(state: boolean = false, action: any) {
  switch (action.type) {
    case actionTypes.CURRENT_USER_SET_IS_LOADING:
      return action.payload.isLoading

    default:
      return state
  }
}
