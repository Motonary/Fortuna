import * as React from 'react'
import './styles.css'

interface TextBoxProps {
  className: string
}

const TextBox: React.SFC<TextBoxProps> = ({ className }) => (
  <input type="text" className={['textbox', className].join(' ')} />
)

export default TextBox
