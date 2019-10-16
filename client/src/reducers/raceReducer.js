import {
  GET_RACE,
  INPUT_CORRECT,
  INPUT_INCORRECT,
  INPUT_FINISHED_WORD
} from "../actions/types";

const initialState = {
  snippet: "",
  correctEnd: 0,
  incorrectStart: 0,
  incorrectEnd: 0,
  currWordStart: 0
};

export default function(state = initialState, action) {
  switch (action.type) {
    case GET_RACE:
      return {
        ...state,
        snippet: action.payload.snippet
      };
    case INPUT_CORRECT:
      return {
        ...state,
        correctEnd: action.payload.end,
        incorrectStart: action.payload.end,
        incorrectEnd: action.payload.end
      };
    case INPUT_INCORRECT:
      return {
        ...state,
        incorrectStart: action.payload.start,
        incorrectEnd: action.payload.end
      };
    case INPUT_FINISHED_WORD:
      return {
        ...state,
        currWordStart: action.payload.newStart,
        correctEnd: action.payload.newStart,
        incorrectStart: action.payload.newStart,
        incorrectEnd: action.payload.newStart
      };
    default:
      return state;
  }
}
