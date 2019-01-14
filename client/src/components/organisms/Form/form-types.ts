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

export type FormValues = CreateUserValues // | CreateSessionValues
