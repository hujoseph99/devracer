import React, { Component } from "react";

// import LoginModal from "./LoginModal";
// import RegisterModal from "./RegisterModal";
// import GuestModal from "./GuestModal";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGithub } from '@fortawesome/free-brands-svg-icons'

import "../../css/verticalCenter.css";
import "../../css/MainMenu.css";

export default class MainMenu extends Component {
  render() {
    return (
      <div className="fullscreen">
        <h1 className="header">Typers</h1>
        <div className="horizontal debug-border">
        </div>
        <a target="_blank" href="https://github.com/hujoseph99/typing">
          <FontAwesomeIcon className="icon-color" icon={faGithub} size="lg"/>
        </a>
      </div>
    );
  }
}
