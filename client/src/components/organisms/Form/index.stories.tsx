import * as React from 'react'
import { storiesOf } from '@storybook/react'
import Form from './index'

storiesOf('Form', module)
  .add('Sign-In', () => <Form formType="signIn" email="" password="" />)
  .add('Sign-Up', () => <Form formType="signUp" name="" email="" password="" confirmation="" />)
