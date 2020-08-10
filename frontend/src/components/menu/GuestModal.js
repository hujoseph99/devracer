import React, { Component } from "react";

import { guestLogin } from "../../actions/userActions";

import Modal from "react-bootstrap/Modal";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { connect } from "react-redux";
import PropTypes from "prop-types";

class GuestModal extends Component {
  state = {
    show: false,
    nickname: ""
  };

  static propTypes = {
    login: PropTypes.func.isRequired,
    isAuthenticated: PropTypes.bool.isRequired
  };

  componentDidUpdate(prevProps) {
    // If authenticated, close modal
    if (this.state.show) {
      if (this.props.isAuthenticated) {
        this.toggle();
      }
    }
  }

  toggle = () => {
    // Clear errors
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
    this.props.guestLogin(this.state.nickname);
  };

  render() {
    return (
      <div>
        <Button variant="primary" onClick={this.handleShow}>
          Guest
        </Button>

        <Modal show={this.state.show}>
          <Modal.Body>
            <Form>
              <Form.Group controlId="nickname">
                <Form.Label>Temporary Nickname</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Enter nickname"
                  onChange={this.handleChange}
                />
              </Form.Group>
              <Button variant="primary" onClick={this.handleSubmit}>
                Continue
              </Button>
            </Form>
          </Modal.Body>
        </Modal>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  isAuthenticated: state.user.isAuthenticated
});

export default connect(mapStateToProps, { guestLogin })(GuestModal);
