import * as React from 'react'
import { connect } from 'react-redux'
import { Field, Formik, ErrorMessage, FormikActions } from 'formik'
import _ from 'lodash'

import Button from '../../atoms/Button'

import styles from './style.css'

import { CreateUserValues, CreateSessionValues } from './form-types'

import { createUser } from '../../../actions/users'

const validate = (values: any) => {
  let errors: any

  if (!values.username) {
    errors.username = 'Username required'
  } else if (values.username && values.username.length > 50) {
    errors.username = 'Too long username'
  }

  if (!values.email) {
    errors.email = 'Required'
  } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)) {
    errors.email = 'Invalid email address'
  }
  if (!values.password) {
    errors.password = 'Required'
  }

  if (!values.confirmation) {
    errors.confirmation = 'Password confirmation required'
  } else if (values.password !== values.confirmation) {
    errors.confirmation = 'Not match password'
  }

  return errors
}

const submitSignUp = (values: CreateUserValues, actions: FormikActions<CreateUserValues>) => {
  createUser(values.email, values.password, '', '')
    .then(() => actions.setSubmitting(false))
    .catch(() => actions.setSubmitting(false))
}

const submitSignIn = (values: CreateSessionValues, actions: FormikActions<CreateSessionValues>) => {
  // this.props
  //   .createSession(values.email, values.password)
  //   .then(() => actions.setSubmitting(false))
  //   .catch(() => actions.setSubmitting(false))
}

const formGenerator = (type: 'signin' | 'signup') => {
  // const submitSignUp = (values: CreateUserValues, actions: FormikActions<CreateUserValues>) => {
  //   createUser(values.email, values.password, '', '')
  //     .then(() => actions.setSubmitting(false))
  //     .catch(() => actions.setSubmitting(false))
  // }

  return ({ createUser, ...rest }: any) => (
    <div id={styles.formContainer}>
      <h1 id={styles.formName}>{type.toUpperCase()}</h1>
      <Formik
        initialValues={{ ...rest }}
        validate={validate}
        onSubmit={type === 'signup' ? submitSignUp : submitSignIn}
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

// export const SignInForm = connect(
//   null,
//   { createSession }
// )(formGenerator('signin'))
export const SignUpForm = connect(
  null,
  { createUser }
)(formGenerator('signup'))
