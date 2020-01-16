import {
  USER_LOADED,
  USER_LOADING,
  AUTH_ERROR,
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  LOGOUT_SUCCESS,
  REGISTER_SUCCESS,
  REGISTER_FAIL
} from "../actions/types";

const initialState = {
  token: localStorage.getItem("token"),
  username: null,
  nickname: null,
  wpm: null,
  id: null,
  isGuest: true,
  isAuthenticated: false
};

export default function(state = initialState, action) {
  switch (action.type) {
    case LOGIN_SUCCESS:
    case REGISTER_SUCCESS:
      localStorage.setItem("token", action.payload.token);
      return {
        ...state,
        username: action.payload.user.username,
        nickname: action.payload.user.nickname,
        wpm: action.payload.user.wpm,
        token: action.payload.token,
        id: action.payload.user._id,
        isGuest: false,
        isAuthenticated: true
      };
    case AUTH_ERROR:
    case LOGIN_FAIL:
    case LOGOUT_SUCCESS:
    case REGISTER_FAIL:
      localStorage.removeItem("token");
      return {
        ...state,
        username: null,
        nickname: null,
        wpm: null,
        token: null,
        id: null,
        isGuest: true,
        isAuthenticated: false
      };
    default:
      return state;
  }
}
