import { combineReducers } from 'redux'
import { currentUser } from './users'

const rootReducer = combineReducers({
  // Users
  currentUser,
})

export default rootReducer
