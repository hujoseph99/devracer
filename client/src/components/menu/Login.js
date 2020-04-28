import React, { Component } from "react";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBolt } from "@fortawesome/free-solid-svg-icons";

import "../../css/login.css";

export default class Login extends Component {
	handleChange = () => {

	}
	
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
			// <div>
			// 	<form class="form-signin">
			// 		<img class="mb-4" src="https://getbootstrap.com/docs/4.0/assets/brand/bootstrap-solid.svg" alt="" width="72" height="72" />
			// 		<h1 class="h3 mb-3 font-weight-normal">Please sign in</h1>
			// 		<label for="inputEmail" class="sr-only">Email address</label>
			// 		<input type="email" id="inputEmail" class="form-control" placeholder="Email address" required="" autofocus="">
			// 		<label for="inputPassword" class="sr-only">Password</label>
			// 		<input type="password" id="inputPassword" class="form-control" placeholder="Password" required="">
			// 		<div class="checkbox mb-3">
			// 			<label>
			// 				<input type="checkbox" value="remember-me">
			// 			</label>
			// 		</div>
			// 		<button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
			// 		<p class="mt-5 mb-3 text-muted">© 2017-2018</p>
			// 	</form>
			// </div>
		);
	}
	
}
