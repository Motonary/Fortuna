import { AnyAction } from 'redux'
import { ThunkAction, ThunkDispatch } from 'redux-thunk'
import axios from 'axios'

import { actionTypes } from '../constants/action-types'
import { BaseAction } from '../constants/static-types'
import { ROOT_URL } from '../constants/url'

// -------------------------------------------------------------------------------------
// CurrentUser
// -------------------------------------------------------------------------------------

interface UserApiRequest extends BaseAction {
  type: string
}

interface UserApiSuccess extends BaseAction {
  type: string
  payload: { currentUser: Object }
}

interface UserApiFailure extends BaseAction {
  type: string
  payload: { error: any }
}

export type CurrentUserAction = UserApiRequest | UserApiSuccess | UserApiFailure | SessionAction

export type CurrentUserActionType = ThunkAction<
  Promise<CurrentUserAction | void>,
  {},
  {},
  AnyAction
>

const userApiRequest = () => {
  return {
    type: actionTypes.USER_API_REQUEST,
  }
}

const userApiSuccess = (json: any) => {
  return {
    type: actionTypes.USER_API_SUCCESS,
    payload: { currentUser: json },
  }
}

const userApiFailure = (error: any) => {
  return {
    type: actionTypes.USER_API_FAILURE,
    payload: { error },
  }
}

export const createUser = (
  name: string,
  email: string,
  password: string,
  password_confirmation: string
): CurrentUserActionType => {
  return (dispatch: ThunkDispatch<{}, {}, CurrentUserAction>) => {
    dispatch(userApiRequest())
    return axios
      .post(`${ROOT_URL}/api/v1/users`, {
        user: { name, email, password, password_confirmation },
      })
      .then(res => dispatch(userApiSuccess(res.data)))
      .catch(err => dispatch(userApiFailure(err)))
  }
}

export const updateUser = (
  name: string,
  email: string,
  password: string,
  password_confirmation: string
): CurrentUserActionType => {
  return (dispatch: ThunkDispatch<{}, {}, CurrentUserAction>) => {
    dispatch(userApiRequest())
    return axios({
      method: 'put',
      // TODO:RESTfulなURLを考慮
      url: `${ROOT_URL}/api/v1/users/`,
      data: { user: { name, email, password, password_confirmation } },
      headers: { Authorization: `Bearer ${sessionStorage.getItem('jwt')}` },
    })
      .then(res => dispatch(userApiSuccess(res.data)))
      .catch(err => dispatch(userApiFailure(err)))
  }
}

// -------------------------------------------------------------------------------------
// UserSession (CurrentUser)
// -------------------------------------------------------------------------------------

// interface
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
