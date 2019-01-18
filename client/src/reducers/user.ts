import { actionTypes } from '../constants/action-types'
import { CurrentUserAction } from '../actions/users'

export function currentUser(state: any = null, action: CurrentUserAction) {
  switch (action.type) {
    case actionTypes.USER_API_REQUEST:
      return [
        ...state,
        {
          isLoading: true,
          items: [],
        },
      ]

    case actionTypes.USER_API_SUCCESS:
      return [
        ...state,
        {
          isLoading: false,
          items: action.payload.currentUser,
        },
      ]

    case actionTypes.USER_API_FAILURE:
      return [
        ...state,
        {
          isLoading: false,
          items: action.payload.error,
        },
      ]

    case actionTypes.SESSION_API_REQUEST:
      return [
        ...state,
        {
          isLoading: true,
          items: [],
        },
      ]

    case actionTypes.SESSION_API_SUCCESS:
      return [
        ...state,
        {
          isLoading: false,
          items: action.payload.currentUser,
        },
      ]

    case actionTypes.SESSION_DELETE_SUCCESS:
      return [
        ...state,
        {
          isLoading: false,
          items: null,
        },
      ]

    case actionTypes.SESSION_API_FAILURE:
      return [
        ...state,
        {
          isLoading: false,
          items: action.payload.error,
        },
      ]

    default:
      return state
  }
}
