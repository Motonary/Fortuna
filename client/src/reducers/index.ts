import { combineReducers } from 'redux'
import { currentUser, isLoadingCurrentUser } from './users'

const rootReducer = combineReducers({
  // Users
  currentUser,

  // Loading
  isLoadingCurrentUser,
})

export default rootReducer
