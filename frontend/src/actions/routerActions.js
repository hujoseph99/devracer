import { 
  JOIN_GAME, 
  RETURN_MENU, 
  ENTER_PRACTICE, 
  GO_TO_REGISTER_MENU, 
  GO_TO_LOGIN_MENU 
} from "./types";

export const joinGame = () => dispatch => {
  dispatch({
    type: JOIN_GAME
  });
};

export const returnMenu = () => dispatch => {
  dispatch({
    type: RETURN_MENU
  });
};

export const enterPractice = () => dispatch => {
  dispatch({
    type: ENTER_PRACTICE
  });
};

export const goToLoginMenu = () => dispatch => {
  dispatch({
    type: GO_TO_LOGIN_MENU
  });
}

export const goToRegisterMenu = () => dispatch => {
  dispatch({
    type: GO_TO_REGISTER_MENU
  });
};
