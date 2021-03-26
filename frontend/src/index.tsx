import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { store } from "./app/store";
import { Provider } from "react-redux";
import axios from 'axios'


const query = `
  query getPracticeRace {
    practiceRace {
      snippet {
        id
        raceContent
        tokenCount
      }
      timeLimit
    }
  }
`;

axios
  .post('http://localhost:8080/graphql', { query })
  .then(res => console.log(res));

ReactDOM.render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>,
  document.getElementById("root")
);
