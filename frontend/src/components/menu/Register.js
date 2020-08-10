import React, { Component } from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";

import { register } from "../../actions/userActions";
import { clearErrors } from "../../actions/errorActions";
import { goToLoginMenu } from "../../actions/routerActions";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBolt } from "@fortawesome/free-solid-svg-icons";

import { ERROR_REGISTER_FAIL } from '../../types/errorTypes';

import "../../css/register.css";
import "../../css/colors.css"
import "../../css/forms.css";

class Register extends Component {
	state = {
		msg: "",
		email: "",
		username: "",
		nickname: "",
		password: "",
		confirmPassword: ""
  };

  static propTypes = {
    error: PropTypes.object.isRequired,
    register: PropTypes.func.isRequired,
		clearErrors: PropTypes.func.isRequired,
		goToLoginMenu: PropTypes.func.isRequired
	};
	
	componentDidUpdate(prevProps) {
    const error= this.props.error;
    if (error !== prevProps.error) {
      // Check for register error
      if (error.id === ERROR_REGISTER_FAIL) {
        this.setState({ msg: error.msg.msg });
      } else {
        this.setState({ msg: null });
      }
    }
  }
	
	getErrorMessage = () => {
		if (this.state.msg) {
			return (
				<div className="alert alert-danger mb-4 alert-fixed" role="alert">
					{this.state.msg}
				</div>
			);
		} else {
			return null;
		}
	}

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
	
	handleLoginClick = e => {
		e.preventDefault();
		this.props.clearErrors();
		this.props.goToLoginMenu();
	}

	render() {
		return (
			<div className="fullscreen">
				{this.getErrorMessage()}
				<div className="card form-background register-form-background-dimensions">
					<div className="register-form-dimensions">
						<form autocomplete="off">
							<FontAwesomeIcon icon={faBolt} size='5x' className="yellow mb-4"/>

							<h1 className="h3 font-weight-normal pink mb-3">Register</h1>
							<input type="text" id="email" className="form-control topInput" placeholder="Email" onChange={this.handleChange} required autofocus="true" />
							<input type="text" id="username" className="form-control middleInput" placeholder="Username" onChange={this.handleChange} required />
							<input type="text" id="nickname" className="form-control middleInput" placeholder="Nickname" onChange={this.handleChange} required />
							<input type="password" id="password" className="form-control middleInput" placeholder="Password" onChange={this.handleChange} required />
							<input type="password" id="confirmPassword" className="form-control bottomInput mb-4" placeholder="Confirm Password" onChange={this.handleChange} required />

							<button class="btn btn-lg btn-primary btn-block mb-4" onClick={this.handleSubmit}>Register</button>
						</form>
						<button class="formFooterText" onClick={this.handleLoginClick}>Return to login</button>
					</div>
				</div>
			</div>
		);
	}
}

const mapStateToProps = state => ({
  error: state.error
});

export default connect(mapStateToProps, { register, clearErrors, goToLoginMenu })(Register);
