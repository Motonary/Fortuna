import keyMirror from 'keymirror'

export const actionTypes = keyMirror({
  // CurrentUser
  SESSION_API_REQUEST: null,
  SESSION_API_SUCCESS: null,
  SESSION_DELETE_SUCCESS: null,
  SESSION_API_FAILURE: null,
  SET_CURRENT_USER: null,
  UPDATE_CURRENT_USER: null,
  CURRENT_USER_SET_IS_LOADING: null,
})
