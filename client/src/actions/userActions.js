import {
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  REGISTER_SUCCESS,
  REGISTER_FAIL,
  GUEST_LOGIN,
  ERROR_LOGIN_FAIL,
  ERROR_REGISTER_FAIL
} from "./types";
import { returnErrors } from "./errorActions";
import { returnMenu } from "./routerActions";
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
    .then(res => {
      dispatch({
        type: LOGIN_SUCCESS,
        payload: res.data
      });
      dispatch(returnMenu());
    })
    .catch(err => {
      dispatch(
        returnErrors(err.response.data, err.response.status, ERROR_LOGIN_FAIL)
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
        ERROR_REGISTER_FAIL
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
    .then(res => {
      dispatch({
        type: REGISTER_SUCCESS,
        payload: res.data
      });
      dispatch(returnMenu());
    })
    .catch(err => {
      console.log(err);
      dispatch(
        returnErrors(err.response.data, err.response.status, ERROR_REGISTER_FAIL)
      );
      dispatch({
        type: REGISTER_FAIL
      });
    });
};

export const guestLogin = nickname => dispatch => {
  dispatch({
    type: GUEST_LOGIN,
    payload: { nickname }
  });
  dispatch(returnMenu());
};
