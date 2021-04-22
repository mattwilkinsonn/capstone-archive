import React from 'react'
import { QueryClient, QueryClientProvider } from 'react-query'
import { ReactQueryDevtools } from 'react-query/devtools'
import { Router, Route, BrowserRouter } from 'react-router-dom'
import './App.css'
import { CreateCapstone } from './components/pages/CreateCapstonePage'
import HomePage from './components/pages/HomePage'
import { LoginPage } from './components/pages/LoginPage'
import MoreviewPage from './components/pages/ViewPage'
import SearchPage from './components/pages/SearchPage'

function App(): JSX.Element {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <div>
        <BrowserRouter>
          <Route path="/" exact component={HomePage} />
          <Route path="/view/:slug" exact component={MoreviewPage} />
          <Route path="/add" exact component={CreateCapstone} />
          <Route path="/search" exact component={SearchPage} />
          <Route path="/login" exact component={LoginPage} />
        </BrowserRouter>
      </div>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}

export default App
