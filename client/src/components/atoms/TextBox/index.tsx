import * as React from 'react'
import * as styles from './style.css'

interface TextBoxProps {
  style: string
}

const TextBox: React.SFC<TextBoxProps> = ({ style }) => (
  <input type="text" className={[styles.textbox, style].join(' ')} />
)

export default TextBox
