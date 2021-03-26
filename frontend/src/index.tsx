import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { store } from "./app/store";
import { Provider } from "react-redux";

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

fetch('http://localhost:8080/graphql', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
  body: JSON.stringify({
    query
  })
})
.then(r => r.json())
.then(data => console.log('data returned:', data));

ReactDOM.render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>,
  document.getElementById("root")
);
