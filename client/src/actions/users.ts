import { AnyAction } from 'redux'
import { ThunkAction, ThunkDispatch } from 'redux-thunk'
import axios from 'axios'

import { actionTypes } from '../constants/action-types'
import { BaseAction } from '../constants/static-types'
import { ROOT_URL } from '../constants/url'

// -------------------------------------------------------------------------------------
// CurrentUser
// -------------------------------------------------------------------------------------
export type CurrentUserActionType = ThunkAction<Promise<void>, {}, {}, AnyAction>

const switchIsLoading = (isLoading: boolean) => {
  return {
    type: actionTypes.CURRENT_USER_SET_IS_LOADING,
    isLoading,
  }
}

export const createUser = (
  name: string,
  email: string,
  password: string
): CurrentUserActionType => {
  return (dispatch: ThunkDispatch<{}, {}, any>) => {
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

export const updateUser = (
  name: string,
  email: string,
  password: string
): CurrentUserActionType => {
  return (dispatch: ThunkDispatch<{}, {}, any>) => {
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
          payload: { currentUser: res.data },
        })
        dispatch(switchIsLoading(false))
      })
      .catch(err => {
        alert(err)
        dispatch(switchIsLoading(false))
      })
  }
}

// -------------------------------------------------------------------------------------
// UserSession (CurrentUser)
// -------------------------------------------------------------------------------------

// TODO: CreateUserと同様に修正(@https://github.com/Motonary/Fortuna/pull/72)
interface SessionApiRequest extends BaseAction {
  type: string
}

interface SessionApiSuccess extends BaseAction {
  type: string
  payload: { currentUser: Object }
}

interface SessionDeleteSuccess extends BaseAction {
  type: string
}

interface SessionApiFailure extends BaseAction {
  type: string
  payload: { error: any }
}

export type SessionAction =
  | SessionApiRequest
  | SessionApiSuccess
  | SessionApiFailure
  | SessionDeleteSuccess

// rename
export type SessionActionType = ThunkAction<Promise<SessionAction | void>, {}, {}, AnyAction>

const sessionApiRequest = () => {
  return {
    type: actionTypes.SESSION_API_REQUEST,
  }
}

const sessionApiSuccess = (json: any) => {
  return {
    type: actionTypes.SESSION_API_SUCCESS,
    payload: { currentUser: json },
  }
}

const sessionDeleteSuccess = () => {
  return {
    type: actionTypes.SESSION_DELETE_SUCCESS,
  }
}

const sessionApiFailure = (error: any) => {
  return {
    type: actionTypes.SESSION_API_FAILURE,
    payload: { error },
  }
}
export const createSession = (email: string, password: string): SessionActionType => {
  return (dispatch: ThunkDispatch<{}, {}, SessionAction>) => {
    dispatch(sessionApiRequest())
    return axios
      .post(`${ROOT_URL}/api/v1/session`, {
        auth: { email: email, password: password },
      })
      .then(res => {
        sessionStorage.setItem('jwt', res.data.jwt.token)
        setTimeout(() => alert('Successfully signed in!'), 100)
        dispatch(sessionApiSuccess(res.data.signinUser))
      })
      .catch(err => dispatch(sessionApiFailure(err)))
  }
}

export const deleteSession = (): SessionActionType => {
  return (dispatch: ThunkDispatch<{}, {}, SessionAction>) => {
    dispatch(sessionApiRequest())
    return axios
      .delete(`${ROOT_URL}/api/v1/session`)
      .then(res => {
        sessionStorage.removeItem('jwt')
        setTimeout(() => alert('Successfully signed out.'), 100)
        dispatch(sessionDeleteSuccess())
      })
      .catch(err => dispatch(sessionApiFailure(err)))
  }
}
