import * as React from 'react'
import { storiesOf } from '@storybook/react'
import { SignInForm, SignUpForm } from './index'

// TODO: Actionをconnectしてpropsで渡す
storiesOf('Form', module)
  .add('Sign-In', () => <SignInForm email="" password="" />)
  .add('Sign-Up', () => <SignUpForm name="" email="" password="" confirmation="" />)
