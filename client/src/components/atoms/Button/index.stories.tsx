import * as React from 'react'
import { storiesOf } from '@storybook/react'
import Button, { PrimaryButton } from './index'

storiesOf('Button', module)
  .add('デフォルト', () => <Button>デフォルト</Button>)
  .add('プライマリ', () => <PrimaryButton>プライマリ</PrimaryButton>)
