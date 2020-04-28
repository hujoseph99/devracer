import { 
  JOIN_GAME, 
  RETURN_MENU, 
  ENTER_PRACTICE, 
  GO_TO_LOGIN_MENU, 
  GO_TO_REGISTER_MENU 
} from "../actions/types";

import {
  MAIN_MENU,
  LOGIN_PAGE,
  REGISTER_PAGE,
  RACE_PAGE,
  PRACTICE_PAGE
} from "../types/pageTypes";

const initialState = {
  inGame: false,
  inMenu: true,
  inPractice: false,
  pageType: LOGIN_PAGE
};

export default function(state = initialState, action) {
  switch (action.type) {
    case RETURN_MENU:
      return {
        ...state,
        pageType: MAIN_MENU
      };
    case GO_TO_LOGIN_MENU:
      return {
        ...state,
        pageType: LOGIN_PAGE
      };
    case GO_TO_REGISTER_MENU:
      return {
        ...state,
        pageType: REGISTER_PAGE
      };
    case JOIN_GAME:
      return {
        ...state,
        pageType: RACE_PAGE
      };
    case ENTER_PRACTICE:
      return {
        ...state,
        pageType: PRACTICE_PAGE
      };
    default:
      return state;
  }
}
