import React, { Component } from "react";
import Modal from "react-bootstrap/Modal";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { connect } from "react-redux";

class LoginModal extends Component {
  state = {
    show: true,
    username: "",
    password: ""
  };

  handleShow = () => {
    this.setState({
      show: true
    });
  };

  handleChange = e => {
    this.setState({
      [e.target.id]: e.target.value
    });
  };

  render() {
    return (
      <div>
        <Button variant="primary" onClick={this.handleShow}>
          User
        </Button>

        <Modal show={this.state.show}>
          <Modal.Body>
            <Form>
              <Form.Group controlId="username">
                <Form.Label>Username</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Enter username"
                  onChange={this.handleChange}
                />
              </Form.Group>

              <Form.Group controlId="password">
                <Form.Label>Password</Form.Label>
                <Form.Control
                  type="password"
                  placeholder="Password"
                  onChange={this.handleChange}
                />
              </Form.Group>
              <Button variant="primary" type="submit">
                Login
              </Button>
            </Form>
          </Modal.Body>
        </Modal>
      </div>
    );
  }
}

const mapStateToProps = state => ({});

export default connect(mapStateToProps, {})(LoginModal);
