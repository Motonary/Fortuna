import React from 'react'
import { connect } from 'react-redux'

// import { SignInForm, SignUpForm } from '../../organisms/Form'
import { SignUpForm } from '../../organisms/Form'

import { createUser, createSession } from '../../../actions/users'

interface TopPageTemplateProps {
  createUser: Function
  createSession: Function
}

const TopPageTemplate: React.FC<TopPageTemplateProps> = ({ createSession, createUser }) => (
  <div>
    <SignInForm actionFunc={createSession} email="" password="" />
    <SignUpForm actionFunc={createUser} name="" email="" password="" confirmation="" />
  </div>
)

export default connect(
  null,
  { createUser, createSession }
)(TopPageTemplate)
