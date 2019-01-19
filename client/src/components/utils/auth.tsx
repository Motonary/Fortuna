import React from 'react'
import { Redirect } from 'react-router-dom'

const Auth = ({currentUser, children}) => (currentUser ? children : <Redirect to={'/'} />)
export default Auth
