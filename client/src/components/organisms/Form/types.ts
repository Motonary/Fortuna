import { FormikActions } from 'formik'
import { CurrentUserThunkActionType } from '../../../actions/users'

// Form values
interface CreateUserValues {
  name: string
  email: string
  password: string
  confirmation: string
}

interface CreateSessionValues {
  email: string
  password: string
}

export type FormValues = CreateSessionValues | CreateUserValues

export interface FormErrorMsgs {
  name?: 'Username required' | 'Too long name'
  email?: 'Required' | 'Invalid email address'
  password?: 'Required'
  confirmation?: 'Password confirmation required' | 'Not match password'
}

// Functions
export type SubmitHofType = (values: FormValues, actions: FormikActions<FormValues>) => void

export type FormGeneratorType = (type: string) => JSX.Element

export type FormType = ({
  actionFunc,
  ...rest
}: { actionFunc: CurrentUserThunkActionType } & any) => JSX.Element
