import React, { Component } from "react";
import TextField from "./TextField";
import InputField from "./InputField";
import io from "socket.io-client";

class Game extends Component {
  componentDidMount() {
    let rand = Math.floor(Math.random() * 10);
    var socket = io.connect("http://localhost:5000", {
      query: {
        room: rand >= 3 ? "244" : "122",
        username: rand >= 5 ? "John" : "Dave",
        wpm: rand >= 7 ? "120" : "100"
      }
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
