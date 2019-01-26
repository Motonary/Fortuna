import * as React from 'react'
import { Redirect } from 'react-router-dom'

interface AuthProps {
  // currentUser: Object
  children: JSX.Element
}

// const Auth: React.SFC<AuthProps> = ({ currentUser, children }) => (
//   currentUser ? children : <Redirect to={'/'} />
// )

// export default Auth

const Auth: React.SFC<AuthProps> = ({ children }) => children
export default Auth
