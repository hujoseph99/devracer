import React, { Component } from "react";
import TextField from "./TextField";
import InputField from "./InputField";
import io from "socket.io-client";

class Game extends Component {
  componentDidMount() {
    var socket = io.connect("http://localhost:5000");
    socket.on("test", function(data) {
      console.log(data);
    });
  }

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
