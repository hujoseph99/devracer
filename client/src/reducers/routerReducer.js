import { JOIN_GAME, RETURN_MENU, ENTER_PRACTICE } from "../actions/types";

const initialState = {
  inGame: false,
  inMenu: true,
  inPractice: false,
};

export default function(state = initialState, action) {
  switch (action.type) {
    case JOIN_GAME:
      return {
        ...state,
        inGame: true,
        inMenu: false,
        inPractice: false
      };
    case ENTER_PRACTICE:
      return {
        ...state,
        inGame: false,
        inMenu: false,
        inPractice: true
      };
    case RETURN_MENU:
      return {
        ...state,
        inGame: false,
        inMenu: true,
        inPractice: false
      };
    default:
      return state;
  }
}
