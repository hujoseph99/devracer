import React, { Component } from "react";
import TextField from "./TextField";

class Game extends Component {
  render() {
    return (
      <div
        className="d-flex justify-content-center"
        style={{ marginTop: "3rem" }}
      >
        <TextField />
      </div>
    );
  }
}

export default Game;
