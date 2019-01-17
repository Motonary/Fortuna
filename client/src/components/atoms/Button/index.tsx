import * as React from 'react'
import * as styles from './style.css'

function buttonFactory(type: string) {
  return ({ children, className, ...props }: any) => (
    <button className={[styles.button, styles.normal, className].join(' ')} {...props}>
      {children}
    </button>
  )
}

export const Button = buttonFactory('normal')
export const PrimaryButton = buttonFactory('primary')

export default Button
