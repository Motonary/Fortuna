import { FormikActions } from 'formik'

// Form values
interface CreateUserValues {
  username: string
  email: string
  password: string
  confirmation: string
}

interface CreateSessionValues {
  email: string
  password: string
}

export type FormValues = CreateSessionValues | CreateUserValues

export interface FormErrors {
  username?: 'Username required' | 'Too long username'
  email?: 'Required' | 'Invalid email address'
  password?: 'Required'
  confirmation?: 'Password confirmation required' | 'Not match password'
}

// Functions
export type SubmitHofType = ((values: FormValues, actions: FormikActions<FormValues>) => void)

export type FormGeneratorType = (type: string) => JSX.Element

export type FormType = ({ actionFunc, ...rest }: { actionFunc: Function } & any) => JSX.Element
