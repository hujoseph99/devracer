import React, { Component } from "react";
import { InputGroup, FormControl } from "react-bootstrap";

class InputField extends Component {
  render() {
    return (
      <div>
        <InputGroup className="mb-3" style={{ width: "40rem" }}>
          <FormControl id="basic-url" aria-describedby="basic-addon3" />
        </InputGroup>
      </div>
    );
  }
}

export default InputField;
