import * as React from 'react'
import ReactDOM from 'react-dom'
// import { Provider } from 'react-redux'
import { createStore, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
// import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom'
// import promise from 'redux-promise'
import reducers from './reducers'

import Hello from './components/Hello'

const store = createStore(reducers, applyMiddleware(thunk))

ReactDOM.render(<Hello />, document.getElementById('app'))
