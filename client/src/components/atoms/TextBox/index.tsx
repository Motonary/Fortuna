import * as React from 'react'
import './styles.css'

interface TextBoxProps {
  className: string
}

const TextBox: React.SFC<TextBoxProps> = ({ className, ...props }: any) => (
  <input type="text" className={['textbox', className].join(' ')} {...props} />
)

export default TextBox
