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
  username: "",
  nickname: "",
  wpm: 0,
  token: "",
  id: "",
  isGuest: true
};

export default function(state = initialState, action) {
  switch (action.type) {
    case LOGIN_SUCCESS:
      return {
        ...state,
        username: action.payload.user.username,
        nickname: action.payload.user.nickname,
        wpm: action.payload.user.wpm,
        token: action.payload.token,
        id: action.payload.user.id,
        isGuest: false
      };
    default:
      return state;
  }
}
