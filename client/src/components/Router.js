import React, { Component } from "react";
import { connect } from "react-redux";
import PropTypes from "prop-types";

import AppNavbar from "./navbar/AppNavbar";
import Game from "./game/Game";
import MainMenu from "./menu/MainMenu";

import Login from "./menu/Login";

class Router extends Component {
  static propTypes = {
    inGame: PropTypes.bool.isRequired,
    inMenu: PropTypes.bool.isRequired
  };

  getRenderPage() {
    if (this.props.inGame) {
      return (
        <div>
          <AppNavbar />
          <Game />
        </div>
      );
    } else if (this.props.inMenu) {
      // return <MainMenu />;
      return <Login />
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
