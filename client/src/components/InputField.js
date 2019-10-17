import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { InputGroup, FormControl } from "react-bootstrap";

import {
  inputCorrect,
  inputIncorrect,
  inputFinishedWord,
  updateWPM,
  setStartTime,
  setEndTime
} from "../actions/raceAction";

class InputField extends Component {
  static propTypes = {
    snippet: PropTypes.string.isRequired,
    correctEnd: PropTypes.number.isRequired,
    incorrectStart: PropTypes.number.isRequired,
    incorrectEnd: PropTypes.number.isRequired,
    currWordStart: PropTypes.number.isRequired,
    wpm: PropTypes.number.isRequired,
    startTime: PropTypes.number.isRequired,
    endTime: PropTypes.number.isRequired,
    wordsTyped: PropTypes.number.isRequired,
    inputCorrect: PropTypes.func.isRequired,
    inputIncorrect: PropTypes.func.isRequired,
    inputFinishedWord: PropTypes.func.isRequired,
    updateWPM: PropTypes.func.isRequired,
    setStartTime: PropTypes.func.isRequired,
    setEndTime: PropTypes.func.isRequired
  };

  firstDifference = (str1, str2) => {
    var shorterLength = Math.min(str1.length, str2.length);

    for (var i = 0; i < shorterLength; i++) {
      if (str1[i] !== str2[i]) return i;
    }

    if (str1.length !== str2.length) return shorterLength;

    return -1;
  };

  // TODO: change this implementation so maybe this is called every second??
  updateWPM = hasEnded => {
    var currDate = new Date().getTime();
    if (!this.props.startTime) {
      // sets the start time when the user starts typing in milliseconds from EPOCH
      this.props.setStartTime(currDate);
    } else if (hasEnded) {
      this.props.setEndTime(currDate);
    } else {
      var secondsElapsed = (currDate - this.props.startTime) / 1000;
      var wpm = Math.round((this.props.wordsTyped / secondsElapsed) * 60);
      this.props.updateWPM(wpm);
    }
  };

  handleChange = () => {
    var currWordStart = this.props.currWordStart;
    var currInput = document.getElementById("formCurrWord").value;
    var targetInput = this.props.snippet.slice(
      currWordStart,
      currWordStart + currInput.length
    );

    var difference = this.firstDifference(currInput, targetInput);

    this.updateWPM(false);

    // No difference
    if (difference === -1) {
      // end of word
      if (currInput.slice(-1) === " ") {
        this.props.inputFinishedWord(currWordStart + currInput.length);
        document.getElementById("formCurrWord").value = "";
      } else {
        this.props.inputCorrect(currWordStart + currInput.length);
      }
    } else {
      this.props.inputCorrect(currWordStart + difference);
      this.props.inputIncorrect(
        currWordStart + difference,
        currWordStart + currInput.length
      );
    }
  };

  render() {
    return (
      <div>
        <InputGroup className="mb-3" style={{ width: "40rem" }}>
          <FormControl
            id="formCurrWord"
            placeholder="Enter text"
            type="text"
            onChange={this.handleChange}
          />
        </InputGroup>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  ...state.race
});

export default connect(
  mapStateToProps,
  {
    inputCorrect,
    inputIncorrect,
    inputFinishedWord,
    updateWPM,
    setStartTime,
    setEndTime
  }
)(InputField);
