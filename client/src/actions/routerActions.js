import { JOIN_GAME, RETURN_MENU, ENTER_PRACTICE } from "./types";

export const joinGame = () => dispatch => {
  dispatch({
    type: JOIN_GAME
  });
};

export const enterPractice = () => dispatch => {
  dispatch({
    type: ENTER_PRACTICE
  });
};

export const returnMenu = () => dispatch => {
  dispatch({
    type: RETURN_MENU
  });
};



