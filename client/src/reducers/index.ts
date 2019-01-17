import { combineReducers } from 'redux'
import { currentUser } from './user'

const rootReducer = combineReducers({
  // Users
  currentUser,
})

export default rootReducer
