import {
  USER_LOADED,
  USER_LOADING,
  AUTH_ERROR,
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  LOGOUT_SUCCESS,
  REGISTER_SUCCESS,
  REGISTER_FAIL
} from "./types";
import { returnErrors } from "./errorActions";
import axios from "axios";

// Login User
export const login = (username, password) => dispatch => {
  // Headers
  const config = {
    headers: {
      "Content-Type": "application/json"
    }
  };

  // Request body
  const body = JSON.stringify({ username, password });

  axios
    .post("http://localhost:5000/api/auth/login", body, config)
    .then(res =>
      dispatch({
        type: LOGIN_SUCCESS,
        payload: res.data
      })
    )
    .catch(err => {
      dispatch(
        returnErrors(err.response.data, err.response.status, "LOGIN_FAIL")
      );
      dispatch({
        type: LOGIN_FAIL
      });
    });
};

// Register User
export const register = (
  username,
  nickname,
  password,
  confirmPassword
) => dispatch => {
  // if password is not equal to the confirmed password, then alert the user
  if (password !== confirmPassword) {
    dispatch(
      returnErrors(
        { msg: "The passwords do not match.  Please try again." },
        400,
        "REGISTER_FAIL"
      )
    );
    return;
  }

  // Headers
  const config = {
    headers: {
      "Content-Type": "application/json"
    }
  };

  // Request body
  const body = JSON.stringify({ username, nickname, password });

  axios
    .post("http://localhost:5000/api/auth/register", body, config)
    .then(res =>
      dispatch({
        type: REGISTER_SUCCESS,
        payload: res.data
      })
    )
    .catch(err => {
      console.log(err);
      dispatch(
        returnErrors(err.response.data, err.response.status, "REGISTER_FAIL")
      );
      dispatch({
        type: REGISTER_FAIL
      });
    });
};
