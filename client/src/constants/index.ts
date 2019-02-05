import keyMirror from 'keymirror'

// actionでもreducerでも使われるのでactions配下ではなくここに置く
export const actionTypes = keyMirror({
  // User
  USER_API_REQUEST: null,
  USER_API_FAILURE: null,

  // CurrentUser
  SET_CURRENT_USER: null,
  UPDATE_CURRENT_USER: null,
  DELETE_SESSION: null,
})

export const ROOT_URL = 'http://localhost:3000'
