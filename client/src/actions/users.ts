import { AnyAction } from 'redux'
import { ThunkAction, ThunkDispatch } from 'redux-thunk'
import axios from 'axios'

import { actionTypes, ROOT_URL } from '../constants'
import { BaseAction } from './static-types'
import { User } from '../constants/static-types'
import { ReduxAPIError } from '../reducers/static-types'
import { CreateUserValues, CreateSessionValues } from '../components/organisms/Form/types'
import { toHash } from '../constants/functions'

export type CurrentUserThunkActionType = ThunkAction<Promise<void>, {}, {}, AnyAction> // TODO: Rename

interface UserAPIRequest extends BaseAction {
  type: string
}

interface UserAPIFailure extends BaseAction {
  type: string
  payload: { error: ReduxAPIError }
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

const userAPIFailure = (error: ReduxAPIError) => ({
  type: actionTypes.USER_API_FAILURE,
  payload: { error },
})

export const createUser = (values: CreateUserValues): CurrentUserThunkActionType => {
  const { confirmation, ...sentValues } = values
  sentValues.password = toHash(values.password)

  // TODO: ThunkDispatchの型微妙
  return (dispatch: ThunkDispatch<{}, {}, UserAPIRequest | CreateUserAction>) => {
    dispatch({ type: actionTypes.USER_API_REQUEST })
    return axios
      .post(`${ROOT_URL}/api/v1/users`, {
        user: sentValues,
      })
      .then(res => {
        dispatch({
          type: actionTypes.SET_CURRENT_USER,
          payload: { currentUser: res.data },
        })
      })
      .catch((err: ReduxAPIError) => {
        if ('statusCode' in err && 'message' in err) {
          dispatch(userAPIFailure(err))
        } else {
          dispatch(userAPIFailure({ statusCode: 500, message: 'Unexpected error' }))
        }
      })
  }
}

export const createSession = (values: CreateSessionValues): CurrentUserThunkActionType => {
  const sentValues = Object.assign({}, values)
  sentValues.password = toHash(values.password)

  // TODO: ThunkDispatchの型微妙
  return (dispatch: ThunkDispatch<{}, {}, UserAPIRequest | CreateSessionAction>) => {
    dispatch({ type: actionTypes.USER_API_REQUEST })
    return axios
      .post(`${ROOT_URL}/api/v1/session`, {
        auth: sentValues,
      })
      .then(res => {
        sessionStorage.setItem('jwt', res.data.jwt.token)
        dispatch({
          type: actionTypes.SET_CURRENT_USER,
          payload: { currentUser: res.data },
        })
      })
      .catch((err: ReduxAPIError) => {
        if ('statusCode' in err && 'message' in err) {
          dispatch(userAPIFailure(err))
        } else {
          dispatch(userAPIFailure({ statusCode: 500, message: 'Unexpected error' }))
        }
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
      .catch((err: ReduxAPIError) => {
        if ('statusCode' in err && 'message' in err) {
          dispatch(userAPIFailure(err))
        } else {
          dispatch(userAPIFailure({ statusCode: 500, message: 'Unexpected error' }))
        }
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
