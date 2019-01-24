import * as React from 'react'
import { storiesOf } from '@storybook/react'
import Button from './index'

storiesOf('Button', module)
  .add('normal', () => <Button type="submit" style="normal" label="normal" />)
  .add('primary', () => <Button type="submit" style="primary" label="primary" />)
  .add('warning', () => <Button type="submit" style="warning" label="warning" />)
  .add('form', () => <Button type="submit" style="form" label="form" />)
