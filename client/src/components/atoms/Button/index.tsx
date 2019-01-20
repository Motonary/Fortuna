import * as React from 'react'
import * as styles from './style.css'

function buttonFactory(type: 'normal' | 'primary' | 'warning') {
  return ({ children, className, ...props }: any) => (
    <button className={[styles.button, styles[type], className].join(' ')}>{children}</button>
  )
}

export const Button = buttonFactory('normal')
export const PrimaryButton = buttonFactory('primary')

export default Button

// interface ButtonProps {
//   type: 'normal' | 'primary' | 'warning'
//   label: string
// }

// const Button: React.SFC<ButtonProps> = ({ type, label }) => (
//   <button className={[styles.button, styles[type]].join(' ')}>{label}</button>
// )

// export default Button
