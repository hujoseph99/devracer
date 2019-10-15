import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { InputGroup, FormControl } from "react-bootstrap";

import {
  inputCorrect,
  inputIncorrect,
  inputFinishedWord
} from "../actions/raceAction";

class InputField extends Component {
  static propTypes = {
    snippet: PropTypes.string.isRequired,
    correctEnd: PropTypes.number.isRequired,
    incorrectStart: PropTypes.number.isRequired,
    incorrectEnd: PropTypes.number.isRequired,
    currWordStart: PropTypes.number.isRequired,
    inputCorrect: PropTypes.func.isRequired,
    inputIncorrect: PropTypes.func.isRequired,
    inputFinishedWord: PropTypes.func.isRequired
  };

  componentDidMount() {
    this.props.inputCorrect(10);
    this.props.inputIncorrect(10, 15);
  }

  render() {
    return (
      <div>
        <InputGroup className="mb-3" style={{ width: "40rem" }}>
          <FormControl id="formCurrWord" placeholder="Enter text" type="text" />
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
  { inputCorrect, inputIncorrect, inputFinishedWord }
)(InputField);
