import React, { Component } from "react";
import Card from "react-bootstrap/Card";
import "./css/textField.css";
import axios from "axios";

class TextField extends Component {
  state = {
    text: ""
  };

  componentDidMount() {
    axios
      .get("/api/race")
      .then(res => this.setState({ text: res.data.snippet }))
      .catch(err => console.log(err));
  }

  render() {
    return (
      <div>
        <Card style={{ width: "40rem" }}>
          <Card.Body className="gameField">
            <Card.Text>{this.state.text}</Card.Text>
          </Card.Body>
        </Card>
      </div>
    );
  }
}

export default TextField;
