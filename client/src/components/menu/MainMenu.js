import React, { Component } from "react";

import Jumbotron from 'react-bootstrap/Jumbotron';
import Button from 'react-bootstrap/Button'

import '../../css/verticalCenter.css'


export default class MainMenu extends Component {
  render() {
    return (
      <div className="vertical-center justify-content-center">
          <Jumbotron style={{ width: "50rem", height: "30rem"}}>
            <h1>INSERT GAME NAME HERE</h1>
            <h3>Join the race as...</h3>
            <Button variant="outline-primary">User</Button>
            <Button variant="outline-secondary">Guest</Button>

          </Jumbotron>
        </div>
    );
  }
}