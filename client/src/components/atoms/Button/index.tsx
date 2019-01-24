import * as React from 'react'
import * as styles from './style.css'

interface ButtonProps {
  type: 'normal' | 'primary' | 'warning'
  label: string
}

const Button: React.SFC<ButtonProps> = ({ type, label }) => (
  <button className={[styles.button, styles[type]].join(' ')}>{label}</button>
)

export default Button
