import React from "react";
import "./App.css";

import AppNavbar from "./components/AppNavbar";
import Game from "./components/Game";

function App() {
  return (
    <div className="App">
      <AppNavbar />
      <Game />
    </div>
  );
}

export default App;
