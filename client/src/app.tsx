import * as React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { createStore, applyMiddleware } from 'redux'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import reducers from './reducers'
import thunk from 'redux-thunk'
import axios from 'axios'

import TopPage from './components/pages/TopPage'
import Auth from './components/utils/auth'

import { ROOT_URL } from './constants'

const store = createStore(reducers, applyMiddleware(thunk))

ReactDOM.render(
  <Provider store={store}>
    <Router>
      <Switch>
        <Route exact path="/" component={TopPage} />
        <Auth>
          <Switch>
            <Route render={() => <h2>404 Not Found</h2>} />
          </Switch>
        </Auth>
      </Switch>
    </Router>
  </Provider>,
  document.getElementById('app')
)

if (process.env.NODE_ENV === 'production') {
  window.addEventListener('error', e => {
    const { message, filename, lineno, colno } = e
    // TODO: Backendでデータを整形しSlackへ投稿
    // NOTE: 'Script error.'しか帰ってこない場合は以下参照 -> https://qiita.com/sue71/items/885caeedb02ae6dc48c4
    axios.post(`${ROOT_URL}/errors`, { message, filename, lineno, colno })
  })
}
