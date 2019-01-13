import React from 'react'
import { Field, getIn } from 'formik'
import './style.css'

interface ErrorMessageProps {
  name: string
}

const ErrorMessage: React.SFC<ErrorMessageProps> = ({ name }: any) => (
  <Field
    name={name}
    render={({ form }: any) => {
      const error = getIn(form.errors, name)
      const touch = getIn(form.touched, name)
      return touch && error ? error : null
    }}
  />
)

export default ErrorMessage
