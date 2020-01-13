import React, { Component } from "react";

import { login } from "../../actions/userActions";

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

  handleSubmit = () => {
    this.props.login(this.state.username, this.state.password);
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
              <Button variant="primary" onClick={this.handleSubmit}>
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

export default connect(mapStateToProps, { login })(LoginModal);
