import React from 'react'

// import { SignInForm, SignUpForm } from '../../organisms/Form'
import { SignUpForm } from '../../organisms/Form'

// interface TopPageTemplateProps {
// }

const TopPageTemplate: React.SFC = () => (
  <div>
    {/* <SignInForm email="" password="" /> */}
    <SignUpForm name="" email="" password="" confirmation="" />

    {/* TODO: あとで消す暫定的な新規ユーザー表示 */}
    <div style={{ textAlign: 'center' }}>
      <div>Successfully created.</div>
      <div>User: hoge</div>
    </div>
  </div>
)

export default TopPageTemplate
