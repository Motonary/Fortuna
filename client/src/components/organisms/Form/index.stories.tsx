import * as React from 'react'
import { storiesOf } from '@storybook/react'
// import { SignInForm, SignUpForm } from './index'
import { SignUpForm } from './index'

storiesOf('Form', module)
  // .add('Sign-In', () => <SignInForm email="" password="" />)
  .add('Sign-Up', () => <SignUpForm name="" email="" password="" confirmation="" />)
