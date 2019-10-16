import React, { Component } from "react";
import Card from "react-bootstrap/Card";
import "./css/textField.css";

import { getRace } from "../actions/raceAction";
import { connect } from "react-redux";
import PropTypes from "prop-types";

class TextField extends Component {
  static propTypes = {
    snippet: PropTypes.string.isRequired,
    correctEnd: PropTypes.number.isRequired,
    incorrectStart: PropTypes.number.isRequired,
    incorrectEnd: PropTypes.number.isRequired,
    currWordStart: PropTypes.number.isRequired,
    getRace: PropTypes.func.isRequired
  };

  componentDidMount() {
    this.props.getRace();
  }

  render() {
    const successText = this.props.snippet.slice(0, this.props.correctEnd);
    const failText = this.props.snippet.slice(
      this.props.incorrectStart,
      this.props.incorrectEnd
    );
    const restText = this.props.snippet.slice(this.props.incorrectEnd);

    return (
      <div className="mb-3">
        <Card style={{ width: "40rem" }}>
          <Card.Body className="gameField">
            <Card.Text>
              <span className="text-success">{successText}</span>
              <span className="bg-danger">{failText}</span>
              {restText}
            </Card.Text>
          </Card.Body>
        </Card>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  ...state.race
});

export default connect(
  mapStateToProps,
  { getRace }
)(TextField);
