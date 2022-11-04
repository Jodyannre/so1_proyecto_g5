import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css'
import { BrowserRouter as Router, Route } from 'react-router-dom'
//import { redirect } from "react-router-dom";

import Navigation from './components/Navigation'
//import GraficaLive from './components/GraficaLive'
import SelectorPartido from './components/SelectorPartido'
import MongoDash from './components/MongoDash'

/*
import NotesList from './components/NotesList'
import CreateNote from './components/CreateNote'

*/
import './App.css';

function App() {
  return (
    <Router>
      <Navigation />
      <div className="container p-4">
        {<Route path="/" exact component={SelectorPartido} />}
        {<Route path="/live" exact component={SelectorPartido} />}
        {/*<Route path="/" exact component={GraficaLive} />*/}
        {<Route path="/logs" exact component={MongoDash} />}
        {/*<Route path="/create" exact component={CreateCar} />*/}
        {/*<Route path="/edit/:id" component={CreateNote} />
        <Route path="/create" component={CreateNote} />
  */}
      </div>
    </Router>
  );
}

export default App;