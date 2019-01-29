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
