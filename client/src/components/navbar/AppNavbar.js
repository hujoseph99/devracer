import React, { Component } from "react";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";

class AppNavbar extends Component {
  render() {
    return (
      <div>
        <Navbar bg="dark" variant="dark">
          <Navbar.Brand href="#home">Typers.io</Navbar.Brand>
          <Nav className="mr-auto"></Nav>
        </Navbar>
      </div>
    );
  }
}

export default AppNavbar;
