import * as React from 'react'
import * as styles from './style.css'

interface ButtonProps {
  type: 'submit' | 'reset' | 'button'
  style: 'normal' | 'primary' | 'warning' | 'form'
  disabled?: boolean
}

const Button: React.SFC<ButtonProps> = ({ type, style, children }) => (
  <button type={type} className={[styles.button, styles[style]].join(' ')}>
    {children}
  </button>
)

export default Button
