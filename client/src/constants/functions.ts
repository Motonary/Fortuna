import * as crypto from 'crypto'

export const toHash = (string: string) => {
  const hash = crypto.createHash('sha256')
  hash.update(string)
  return hash.digest('hex')
}
