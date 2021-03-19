import React from 'react'
import './App.css'
import { Route } from 'react-router-dom'
import HomePage from './components/pages/HomePage'
import MoreviewPage from './components/pages/MoreviewPage'
import { QueryClientProvider, QueryClient } from 'react-query'

function App() {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <div>
        <Route path="/" exact component={HomePage} />
        <Route path="/View" exact component={MoreviewPage} />
      </div>
    </QueryClientProvider>
  )
}

export default App
