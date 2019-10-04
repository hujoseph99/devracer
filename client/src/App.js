import React from "react";
import "./App.css";

import AppNavbar from "./components/AppNavbar";
import Game from "./components/Game";
import { Provider } from "react-redux";

import rootReducer from "./reducers";
import thunk from "redux-thunk";
import { createStore, applyMiddleware, compose } from "redux";

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(
  rootReducer,
  {},
  composeEnhancers(applyMiddleware(thunk))
);

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
