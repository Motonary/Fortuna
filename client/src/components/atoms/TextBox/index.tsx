import * as React from 'react'
import * as styles from './style.css'

interface TextBoxProps {
  style: 'textbox'
}

const TextBox: React.SFC<TextBoxProps> = ({ style }) => (
  <input type="text" className={[styles.textbox, styles[style]].join(' ')} />
)

export default TextBox
