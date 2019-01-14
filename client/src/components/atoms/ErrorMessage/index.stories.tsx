import * as React from 'react'
import { storiesOf } from '@storybook/react'
import ErrorMessage from './index'

storiesOf('ErrorMessage', module).add('デフォルト', () => <ErrorMessage name="email" />)
