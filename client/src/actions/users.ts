import axios from 'axios'
import { actionTypes } from '../constants/action-types'
import { BaseAction } from '../constants/static-types'
import { ROOT_URL } from '../constants/url'

// -------------------------------------------------------------------------------------
// CurrentUser
// -------------------------------------------------------------------------------------
interface CreateUserAction extends BaseAction {
  type: string
  payload: { currentUser: Object }
}

interface UpdateUserAction extends BaseAction {
  type: string
  payload: { updatedUser: Object }
}

interface CreateSessionAction extends BaseAction {
  type: string
  payload: { currentUser: Object }
}

interface DeleteSessionAction extends BaseAction {
  type: string
}

export type CurrentUserAction =
  | CreateUserAction
  | UpdateUserAction
  | CreateSessionAction
  | DeleteSessionAction

export function createUser(
  name: string,
  email: string,
  password: string,
  password_confirmation: string
): Promise<CreateUserAction | void> {
  return axios
    .post(`${ROOT_URL}/api/v1/users`, {
      user: { name, email, password, password_confirmation },
    })
    .then(res => {
      return createSession(email, password).then(() => {
        return {
          type: actionTypes.SET_USER_SESSION,
          payload: { currentUser: res.data },
        }
      })
    })
    .catch(() => alert('Sorry, something went wrong. Please reload.'))
}

export function updateUser(
  name: any = null,
  email: any = null,
  password: any,
  password_confirmation: any
): Promise<UpdateUserAction | void> {
  return axios({
    method: 'put',
    // TODO:RESTfulなURLを考慮
    url: `${ROOT_URL}/api/v1/users/`,
    data: { user: { name, email, password, password_confirmation } },
    headers: { Authorization: `Bearer ${sessionStorage.getItem('jwt')}` },
  })
    .then(res => {
      return {
        type: actionTypes.UPDATE_USER,
        payload: { updatedUser: res.data },
      }
    })
    .catch(() => alert('Sorry, something went wrong. Please reload.'))
}

export function createSession(
  email: string,
  password: string
): Promise<CreateSessionAction | void> {
  return axios
    .post(`${ROOT_URL}/api/v1/session`, {
      auth: { email: email, password: password },
    })
    .then(res => {
      sessionStorage.setItem('jwt', res.data.jwt.token)
      setTimeout(() => alert('Successfully signed in!'), 100)
      const user: Object = res.data.signinUser
      return {
        type: actionTypes.SET_USER_SESSION,
        payload: { currentUser: user },
      }
    })
    .catch(() => alert('Sorry, something went wrong. Please reload.'))
}

export function deleteSession(callback: Function): DeleteSessionAction {
  sessionStorage.removeItem('jwt')
  callback()
  setTimeout(() => alert('Successfully signed out.'), 100)
  return { type: actionTypes.DELETE_USER_SESSION }
}
