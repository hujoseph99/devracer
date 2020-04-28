import React, { Component } from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";

import { joinGame, enterPractice, goToLoginMenu } from '../../actions/routerActions';

import Button from "react-bootstrap/Button";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGithub } from '@fortawesome/free-brands-svg-icons';

import "../../css/MainMenu.css";

class MainMenu extends Component {
  static propTypes = {
    joinGame: PropTypes.func.isRequired,
		enterPractice: PropTypes.func.isRequired,
		goToLoginMenu: PropTypes.func.isRequired
  };
  
  handleLoginClick = () => {
      this.props.goToLoginMenu();
  }

  render() {
    //TODO: Add login/register button to the top right
    return (
      <div className="fullscreenMainMenu">
        <div className="horizontal">
          <h1 className="header mt-2">Typers</h1>
          <button className="loginText" onClick={this.handleLoginClick}>Login</button>
        </div>
        <div className="horizontal">
          <h1 className="subtitle">Choose your mode</h1>
          <Button size="lg" onClick={this.props.joinGame} className="btn-race">Race against others</Button>
          <Button size="lg" onClick={this.props.enterPractice} className="btn-practice">Solo practice</Button>
        </div>
        <a target="_blank" href="https://github.com/hujoseph99/typing">
          <FontAwesomeIcon className="icon-color mb-2" icon={faGithub} size="lg"/>
        </a>
      </div>
    );
  }
}

export default connect(null, { joinGame, enterPractice, goToLoginMenu })(MainMenu);
