import { AnyAction } from 'redux'
import { ThunkAction, ThunkDispatch } from 'redux-thunk'
import axios from 'axios'

import { actionTypes, ROOT_URL } from '../constants'
import { BaseAction } from './static-types'
import { User } from '../constants/static-types'

export type CurrentUserThunkActionType = ThunkAction<Promise<void>, {}, {}, AnyAction> // TODO: Rename

interface UserAPIRequest extends BaseAction {
  type: string
}

interface UserAPIFailure extends BaseAction {
  type: string
  payload: { error: any } // TODO: 厳格に
}

interface CreateUserAction extends BaseAction {
  type: string
  payload: { currentUser: User }
}

interface CreateSessionAction extends BaseAction {
  type: string
  payload: { currentUser: User }
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
  | UserAPIRequest
  | UserAPIFailure
  | CreateUserAction
  | CreateSessionAction
  // | FetchCurrentUserAction
  | DeleteSessionAction
  // | RemoveFirstVisitFlagAction
  // | UpdateUserImgAction
  | UpdateProfileAction

const userAPIFailure = (error: any) => ({
  type: actionTypes.USER_API_FAILURE,
  payload: { error },
})

export const createUser = (
  name: string,
  email: string,
  password: string
): CurrentUserThunkActionType => {
  // TODO: ThunkDispatchの型微妙
  return (dispatch: ThunkDispatch<{}, {}, UserAPIRequest | CreateUserAction>) => {
    dispatch({ type: actionTypes.USER_API_REQUEST })
    return axios
      .post(`${ROOT_URL}/api/v1/users`, {
        user: { name, email, password },
      })
      .then(res => {
        dispatch({
          type: actionTypes.SET_CURRENT_USER,
          payload: { currentUser: res.data },
        })
      })
      .catch(err => {
        // TODO: errの型が{status: string, message: string}でない場合(想定していないエラーの場合)、APIにエラーログを投げる
        dispatch(userAPIFailure(err))
      })
  }
}

export const createSession = (email: string, password: string): CurrentUserThunkActionType => {
  // TODO: ThunkDispatchの型微妙
  return (dispatch: ThunkDispatch<{}, {}, UserAPIRequest | CreateSessionAction>) => {
    dispatch({ type: actionTypes.USER_API_REQUEST })
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
        // TODO: errの型が{status: string, message: string}でない場合(想定していないエラーの場合)、APIにエラーログを投げる
        dispatch(userAPIFailure(err))
      })
  }
}

export const updateProfile = (
  name: string,
  email: string,
  password: string
): CurrentUserThunkActionType => {
  // TODO: ThunkDispatchの型微妙
  return (dispatch: ThunkDispatch<{}, {}, UserAPIRequest | UpdateProfileAction>) => {
    dispatch({ type: actionTypes.USER_API_REQUEST })
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
      })
      .catch(err => {
        // TODO: errの型が{status: string, message: string}でない場合(想定していないエラーの場合)、APIにエラーログを投げる
        dispatch(userAPIFailure(err))
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
