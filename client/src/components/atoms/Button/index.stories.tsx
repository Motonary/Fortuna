import * as React from 'react'
import { storiesOf } from '@storybook/react'
import Button from './index'

storiesOf('Button', module)
  .add(
    'normal',
    () => (
      <Button type="submit" style="normal">
        normal
      </Button>
    ),
    { notes: 'This is a first note for Fortuna(test).' }
  )
  .add('primary', () => (
    <Button type="submit" style="primary">
      primary
    </Button>
  ))
  .add('warning', () => (
    <Button type="submit" style="warning">
      warning
    </Button>
  ))
  .add('form', () => (
    <Button type="submit" style="form">
      form
    </Button>
  ))
