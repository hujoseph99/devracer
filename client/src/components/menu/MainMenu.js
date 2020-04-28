import React, { Component } from "react";
import { connect } from "react-redux";

import LoginModal from "./LoginModal";
// import RegisterModal from "./RegisterModal";
// import GuestModal from "./GuestModal";

import Button from "react-bootstrap/Button";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGithub } from '@fortawesome/free-brands-svg-icons';

import { joinGame, enterPractice } from '../../actions/routerActions';

import "../../css/verticalCenter.css";
import "../../css/MainMenu.css";

class MainMenu extends Component {
  render() {
    //TODO: Add login/register button to the top right
    return (
      <div className="fullscreen">
        <div className="horizontal debug-border">
          <span className="header">Typers</span>
          <LoginModal />
        </div>
        <div className="horizontal">
          <h1 className="subtitle">Choose your mode</h1>
          <Button size="lg" onClick={this.props.joinGame} className="btn-race">Race against others</Button>
          <Button size="lg" onClick={this.props.enterPractice} className="btn-practice">Solo practice</Button>
        </div>
        <a target="_blank" href="https://github.com/hujoseph99/typing">
          <FontAwesomeIcon className="icon-color" icon={faGithub} size="lg"/>
        </a>
      </div>
    );
  }
}

export default connect(null, { joinGame, enterPractice })(MainMenu);
