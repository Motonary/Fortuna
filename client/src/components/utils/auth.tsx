import React from 'react'
import { Redirect } from 'react-router-dom'

const Auth = (props: any) => (props.currentUser.isLoggedIn ? props.children : <Redirect to={'/'} />)
const Auth = (props: any) => (props.currentUser ? props.children : <Redirect to={'/'} />)
export default Auth
