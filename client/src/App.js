import React from "react";
import "./App.css";

import AppNavbar from "./components/AppNavbar";
import Game from "./components/Game";
import { Provider } from "react-redux";

import store from "./store";

function App() {
  return (
    <Provider store={store}>
      <div className="App">
        <AppNavbar />
        <Game />
      </div>
    </Provider>
  );
}

export default App;
