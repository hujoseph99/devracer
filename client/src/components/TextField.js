import React, { Component } from "react";
import Card from "react-bootstrap/Card";
import "./css/textField.css";

import { getRace } from "../actions/raceAction";
import { connect } from "react-redux";
import PropTypes from "prop-types";

class TextField extends Component {
  static propTypes = {
    snippet: PropTypes.string.isRequired,
    getRace: PropTypes.func.isRequired
  };

  componentDidMount() {
    this.props.getRace();
  }

  render() {
    return (
      <div className="mb-3">
        <Card style={{ width: "40rem" }}>
          <Card.Body className="gameField">
            <Card.Text>{this.props.snippet}</Card.Text>
          </Card.Body>
        </Card>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  snippet: state.race.snippet
});

export default connect(
  mapStateToProps,
  { getRace }
)(TextField);
