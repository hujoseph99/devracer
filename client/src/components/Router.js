import React, { Component } from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";

import AppNavbar from "./navbar/AppNavbar";
import Game from "./game/Game";

import MainMenu from "./menu/MainMenu";
import Login from "./menu/Login";
import Register from "./menu/Register";

import {
  MAIN_MENU,
  LOGIN_PAGE,
  REGISTER_PAGE,
  RACE_PAGE,
  PRACTICE_PAGE
} from "../types/pageTypes";

class Router extends Component {
  static propTypes = {
    inGame: PropTypes.bool.isRequired,
    inMenu: PropTypes.bool.isRequired,
    pageType: PropTypes.number.isRequired
  };

  getRenderPage = () => {
    switch(this.props.pageType) {
      case RACE_PAGE:
        return (
          <div>
            <AppNavbar />
            <Game />
          </div>
        );
      case MAIN_MENU:
        return <MainMenu />;
      case LOGIN_PAGE:
        return <Login />;
      case REGISTER_PAGE:
        return <Register />;
      default:
        return <MainMenu />;
    }
  }

  render() {
    return (
      <div>
        {this.getRenderPage()}
      </div>
    );
  }
}

const mapStateToProps = state => ({
  ...state.router
});

export default connect(mapStateToProps, {})(Router);
