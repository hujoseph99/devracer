import React, { Component } from "react";
import TextField from "./TextField";
import InputField from "./InputField";

import { getRace } from "../actions/raceAction";
import { connect } from "react-redux";
import PropTypes from "prop-types";

import io from "socket.io-client";

class Game extends Component {
  static propTypes = {
    roomNum: PropTypes.string.isRequired,
    getRace: PropTypes.func.isRequired
  };

  componentDidMount() {
    this.props.getRace();
  }

  // TODO: This will be called every time props updated.  Need to change it so
  //  that it won't keep connecting every time props are updated
  componentDidUpdate() {
    if (this.props.roomNum) {
      var room = io.connect("localhost:5000/" + this.props.roomNum, {
        query: {
          roomNum: this.props.roomNum,
          username: Math.floor(Math.random() * 20) > 10 ? "John" : "Doe",
          wpm: Math.floor(Math.random() * 20) > 5 ? 100 : 120
        }
      });
      room.on("updateWPM", room => {
        console.log(room);
      });
    }
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

const mapStateToProps = state => ({
  roomNum: state.race.roomNum
});

export default connect(
  mapStateToProps,
  { getRace }
)(Game);
