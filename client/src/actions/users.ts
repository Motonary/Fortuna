import { AnyAction } from 'redux'
import { ThunkAction, ThunkDispatch } from 'redux-thunk'
import axios from 'axios'

import { actionTypes } from '../constants/action-types'
import { BaseAction, User } from '../constants/static-types'
import { ROOT_URL } from '../constants/url'

// -------------------------------------------------------------------------------------
// Loading
// -------------------------------------------------------------------------------------
export type SwitchIsLoadingAction = {
  type: 'SET_IS_LOADING__CURRENT_USER'
  payload: { isLoading: boolean }
}

const switchIsLoading = (isLoading: boolean) => {
  return {
    type: actionTypes.SET_IS_LOADING__CURRENT_USER,
    payload: { isLoading },
  }
}
// -------------------------------------------------------------------------------------
// CurrentUser
// -------------------------------------------------------------------------------------
export type CurrentUserThunkActionType = ThunkAction<Promise<void>, {}, {}, AnyAction> // TODO: Rename

interface CreateUserAction extends BaseAction {
  type: string
  payload: { currentUser: User } // TODO: 厳格に
}

interface CreateSessionAction extends BaseAction {
  type: string
  payload: { currentUser: User } // TODO: 厳格に
}

// interface FetchCurrentUserAction extends BaseAction {
//   type: string
//   payload: { currentUser: Object }
// }

interface DeleteSessionAction extends BaseAction {
  type: string
}

// interface RemoveFirstVisitFlagAction extends BaseAction {
//   type: string
// }

// interface UpdateUserImgAction extends BaseAction {
//   type: string
//   payload: { newAvatarUrl: string }
// }

interface UpdateProfileAction extends BaseAction {
  type: string
  payload: { updatedUser: User }
}

export type CurrentUserAction =
  | CreateUserAction
  | CreateSessionAction
  // | FetchCurrentUserAction
  | DeleteSessionAction
  // | RemoveFirstVisitFlagAction
  // | UpdateUserImgAction
  | UpdateProfileAction

export const createUser = (
  name: string,
  email: string,
  password: string
): CurrentUserThunkActionType => {
  return (dispatch: ThunkDispatch<{}, {}, CreateUserAction | SwitchIsLoadingAction>) => {
    dispatch(switchIsLoading(true))
    return axios
      .post(`${ROOT_URL}/api/v1/users`, {
        user: { name, email, password },
      })
      .then(res => {
        dispatch({
          type: actionTypes.SET_CURRENT_USER,
          payload: { currentUser: res.data },
        })
        dispatch(switchIsLoading(false))
      })
      .catch(err => {
        alert(err) // 暫定処理
        dispatch(switchIsLoading(false))
      })
  }
}

export const createSession = (email: string, password: string): CurrentUserThunkActionType => {
  return (dispatch: ThunkDispatch<{}, {}, CreateSessionAction | SwitchIsLoadingAction>) => {
    dispatch(switchIsLoading(true))
    return axios
      .post(`${ROOT_URL}/api/v1/session`, {
        auth: { email: email, password: password },
      })
      .then(res => {
        sessionStorage.setItem('jwt', res.data.jwt.token)
        dispatch({
          type: actionTypes.SET_CURRENT_USER,
          payload: { currentUser: res.data },
        })
      })
      .catch(err => {
        alert(err) // 暫定処理
        dispatch(switchIsLoading(false))
      })
  }
}

export const updateProfile = (
  name: string,
  email: string,
  password: string
): CurrentUserThunkActionType => {
  return (dispatch: ThunkDispatch<{}, {}, UpdateProfileAction | SwitchIsLoadingAction>) => {
    dispatch(switchIsLoading(true))
    return axios({
      method: 'put',
      // TODO:RESTfulなURLを考慮
      url: `${ROOT_URL}/api/v1/users/`,
      data: { user: { name, email, password } },
      headers: { Authorization: `Bearer ${sessionStorage.getItem('jwt')}` },
    })
      .then(res => {
        dispatch({
          type: actionTypes.UPDATE_CURRENT_USER,
          payload: { updatedUser: res.data },
        })
        dispatch(switchIsLoading(false))
      })
      .catch(err => {
        alert(err)
        dispatch(switchIsLoading(false))
      })
  }
}

export const deleteSession = (): ThunkAction<void, {}, {}, AnyAction> => {
  return (dispatch: ThunkDispatch<{}, {}, DeleteSessionAction>) => {
    sessionStorage.removeItem('jwt')
    // TODO: Redirectなどのcallbackを走らせる
    dispatch({ type: actionTypes.DELETE_SESSION })
  }
}
