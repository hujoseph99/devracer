import React, { Component } from "react";

import { register } from "../../actions/userActions";
import { clearErrors } from "../../actions/errorActions";

import Alert from "react-bootstrap/Alert";
import Modal from "react-bootstrap/Modal";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { connect } from "react-redux";
import PropTypes from "prop-types";

class RegisterModal extends Component {
  state = {
    show: false,
    username: null,
    nickname: null,
    password: null,
    confirmPassword: null,
    msg: null
  };

  static propTypes = {
    isAuthenticated: PropTypes.bool.isRequired,
    error: PropTypes.object.isRequired,
    register: PropTypes.func.isRequired,
    clearErrors: PropTypes.func.isRequired
  };

  componentDidUpdate(prevProps) {
    const { isAuthenticated, error } = this.props;
    if (error !== prevProps.error) {
      // Check for register error
      if (error.id === "REGISTER_FAIL") {
        this.setState({ msg: error.msg.msg });
      } else {
        this.setState({ msg: null });
      }
    }

    // If authenticated, close modal
    if (this.state.show) {
      if (isAuthenticated) {
        this.toggle();
      }
    }
  }

  toggle = () => {
    // Clear errors
    this.props.clearErrors();
    this.setState({
      show: !this.state.show
    });
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

  handleSubmit = e => {
    e.preventDefault();

    this.props.register(
      this.state.username,
      this.state.nickname,
      this.state.password,
      this.state.confirmPassword
    );
  };

  render() {
    return (
      <div>
        <Button variant="primary" onClick={this.handleShow}>
          Register
        </Button>

        <Modal show={this.state.show}>
          <Modal.Body>
            {this.state.msg ? (
              <Alert variant="danger">{this.state.msg}</Alert>
            ) : null}
            <Form onSubmit={this.handleSubmit}>
              <Form.Group controlId="username">
                <Form.Label>Username</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Username"
                  onChange={this.handleChange}
                />
              </Form.Group>

              <Form.Group controlId="nickname">
                <Form.Label>Nickname</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Nickname"
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

              <Form.Group controlId="confirmPassword">
                <Form.Label>Confirm Password</Form.Label>
                <Form.Control
                  type="password"
                  placeholder="Password"
                  onChange={this.handleChange}
                />
              </Form.Group>
              <Button variant="primary" type="submit">
                Register
              </Button>
            </Form>
          </Modal.Body>
        </Modal>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  error: state.error,
  isAuthenticated: state.user.isAuthenticated
});

export default connect(mapStateToProps, { register, clearErrors })(
  RegisterModal
);
