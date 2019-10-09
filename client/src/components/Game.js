import React, { Component } from "react";
import TextField from "./TextField";
import InputField from "./InputField";

class Game extends Component {
  render() {
    return (
      <div className="container-fluid">
        <div
          className="d-flex flex-column align-items-center"
          style={{ marginTop: "3rem" }}
        >
          <TextField />
          <InputField />
        </div>
      </div>
    );
  }
}

export default Game;
