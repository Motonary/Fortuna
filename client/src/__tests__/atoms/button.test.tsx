import * as React from 'react'
import { shallow, configure } from 'enzyme'
import Adapter from 'enzyme-adapter-react-16'

import Button from '../../components/atoms/Button'

configure({ adapter: new Adapter() })

describe('ButtonComponent', () => {
  it('<Button />', () => {
    const defaultButton = shallow(<Button>Default</Button>)
    expect(defaultButton.find('button').text()).toBe('Default')
  })
})
