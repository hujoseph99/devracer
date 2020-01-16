import { JOIN_GAME, RETURN_MENU } from "../actions/types";

const initialState = {
  inGame: false,
  inMenu: true
};

export default function(state = initialState, action) {
  switch (action.type) {
    case JOIN_GAME:
      return {
        ...state,
        inGame: true,
        inMenu: false
      };
    case RETURN_MENU:
      return {
        ...state,
        inGame: false,
        inMenu: true
      };
    default:
      return state;
  }
}
