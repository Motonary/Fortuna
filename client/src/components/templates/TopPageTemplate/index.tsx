import React from 'react'

import Form from '../../organisms/Form'

// interface TopPageTemplateProps {
// }

const TopPageTemplate: React.SFC = () => (
  <div>
    <Form formType="signIn" email="" password="" />
    <Form formType="signUp" name="" email="" password="" confirmation="" />
  </div>
)

export default TopPageTemplate
