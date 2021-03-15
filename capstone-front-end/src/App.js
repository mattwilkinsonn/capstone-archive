import React from 'react';
import './App.css';
import { Route } from 'react-router-dom';
import HomePage from './components/pages/HomePage';
import MoreviewPage from './components/pages/MoreviewPage';


function App() {
  return (
    <div>
      <Route path="/" exact component={HomePage}/>
      <Route path="/View" exact component={MoreviewPage}/>
    </div>
  );
}

export default App;
