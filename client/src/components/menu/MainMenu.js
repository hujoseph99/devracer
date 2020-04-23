import React, { Component } from "react";

import LoginModal from "./LoginModal";
import RegisterModal from "./RegisterModal";
import GuestModal from "./GuestModal";

import "../../css/verticalCenter.css";
import "../../css/MainMenu.css";

export default class MainMenu extends Component {
  render() {
    return (
      <div className="fullscreen">
        <h1 className="header">Typers</h1>
        <div className="horizontal debug-border">
        </div>
        <i className="fab fa-github"></i>
      </div>
    );
  }
}
