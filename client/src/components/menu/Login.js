import React, { Component } from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";

import { login } from "../../actions/userActions";
import { clearErrors } from "../../actions/errorActions";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBolt } from "@fortawesome/free-solid-svg-icons";

import "../../css/login.css";

class Login extends Component {
	state = {
    show: false,
    username: "",
    password: ""
  };

  static propTypes = {
    error: PropTypes.object.isRequired,
    login: PropTypes.func.isRequired,
    clearErrors: PropTypes.func.isRequired
  };

	render() {
		return (
			<div className="fullscreen">
				<div className="card login-form-background">
					<form className="login-form" autocomplete="off">

						<FontAwesomeIcon icon={faBolt} size='5x' className="logo mb-5"/>

						<h1 className="h3 font-weight-normal pink mb-3">Please sign in</h1>
						<input type="text" id="username" class="form-control username" placeholder="Username" required autofocus />
						<input type="password" id="password" class="form-control password mb-4" placeholder="Password" requireds />

						<button class="btn btn-lg btn-primary btn-block" type="submit">Sign In</button>
					</form>
				</div>
			</div>
		);
	}
}

const mapStateToProps = state => ({
  error: state.error,
  isAuthenticated: state.user.isAuthenticated
});

export default connect(mapStateToProps, { login, clearErrors })(Login);
