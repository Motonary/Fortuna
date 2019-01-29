import * as React from 'react'
import { Field, Formik, ErrorMessage, FormikActions } from 'formik'
import _ from 'lodash'

import Button from '../../atoms/Button'

import styles from './style.css'

import { FormValues, FormErrors } from './form-types'

const validate = (values: FormValues): FormErrors => {
  let errors: any

  if (!values.email) {
    errors.email = 'Required'
  } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)) {
    errors.email = 'Invalid email address'
  }

  if (!values.password) {
    errors.password = 'Required'
  }

  if ('username' in values && 'confirmation' in values) {
    if (!values.confirmation) {
      errors.confirmation = 'Password confirmation required'
    } else if (values.password !== values.confirmation) {
      errors.confirmation = 'Not match password'
    }

    if (!values.username) {
      errors.username = 'Username required'
    } else if (values.username && values.username.length > 50) {
      errors.username = 'Too long username'
    }
  }

  return errors
}

const submitHof = (actionFunc: Function) => {
  return function submit(values: FormValues, actions: FormikActions<FormValues>) {
    actionFunc({ ...values })
      .then(() => actions.setSubmitting(false))
      .catch(() => actions.setSubmitting(false))
  }
}

const formGenerator = (type: string): Function => {
  return ({ actionFunc, ...rest }: any) => (
    <div id={styles.formContainer}>
      <h1 id={styles.formName}>{type.toUpperCase()}</h1>
      <Formik
        initialValues={{ ...rest }}
        validate={validate}
        onSubmit={submitHof(actionFunc)}
        render={({ handleSubmit, isSubmitting }) => (
          <form className={styles.signForm} onSubmit={handleSubmit}>
            {_.map(Object.keys(rest), key => {
              return (
                <div className={styles.formFields} key={key}>
                  <Field
                    type={`${key}`}
                    name={`${key}`}
                    placeholder={`${key}`}
                    className={styles.formField}
                  />
                  <ErrorMessage name={`${key}`}>{msg => <div>{msg}</div>}</ErrorMessage>
                </div>
              )
            })}
            <Button type="submit" style="form" disabled={isSubmitting}>
              Submit
            </Button>
          </form>
        )}
      />
    </div>
  )
}

export const SignInForm = formGenerator('signin')
export const SignUpForm = formGenerator('signup')
