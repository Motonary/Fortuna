import keyMirror from 'keymirror'

export const actionTypes = keyMirror({
  // User
  USER_API_REQUEST: null,
  USER_API_FAILURE: null,

  // CurrentUser
  SET_CURRENT_USER: null,
  UPDATE_CURRENT_USER: null,
  DELETE_SESSION: null,
})
