import * as React from 'react'
import * as styles from './style.css'

interface ButtonProps {
  type: 'submit' | 'reset' | 'button'
  style: 'normal' | 'primary' | 'warning' | 'form'
  label: string
  disabled?: boolean
}

const Button: React.SFC<ButtonProps> = ({ type, style, label }) => (
  <button type={type} className={[styles.button, styles[style]].join(' ')}>
    {label}
  </button>
)

export default Button
