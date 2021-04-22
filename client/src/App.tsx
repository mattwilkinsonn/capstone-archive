import React from 'react'
import { QueryClient, QueryClientProvider } from 'react-query'
import { ReactQueryDevtools } from 'react-query/devtools'
import { BrowserRouter, Route } from 'react-router-dom'
import './App.css'
import { CreateCapstone } from './components/pages/CreateCapstonePage'
import HomePage from './components/pages/HomePage'
import { LoginPage } from './components/pages/LoginPage'
import SearchPage from './components/pages/SearchPage'
import MoreviewPage from './components/pages/ViewPage'

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
