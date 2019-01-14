import * as React from 'react'
import './styles.css'

function buttonFactory(type: string) {
  return ({ children, className, ...props }: any) => (
    <button className={['button', type, className].join(' ')} {...props}>
      {children}
    </button>
  )
}

export const Button = buttonFactory('default')
export const PrimaryButton = buttonFactory('primary')

export default Button
