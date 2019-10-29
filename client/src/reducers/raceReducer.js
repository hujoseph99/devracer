import {
  GET_RACE,
  INPUT_CORRECT,
  INPUT_INCORRECT,
  INPUT_FINISHED_WORD,
  UPDATE_WPM,
  SET_START_TIME,
  SET_END_TIME
} from "../actions/types";

const initialState = {
  snippet: "",
  roomNum: 0,
  correctEnd: 0,
  incorrectStart: 0,
  incorrectEnd: 0,
  currWordStart: 0,
  wpm: 0,
  startTime: 0,
  endTime: 0,
  wordsTyped: 0
};

export default function(state = initialState, action) {
  switch (action.type) {
    case GET_RACE:
      return {
        ...state,
        roomNum: action.payload.roomNum,
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
        incorrectEnd: action.payload.newStart,
        wordsTyped: state.wordsTyped + 1
      };
    case UPDATE_WPM:
      return {
        ...state,
        wpm: action.payload.newWPM
      };
    case SET_START_TIME:
      return {
        ...state,
        startTime: action.payload.startTime
      };
    case SET_END_TIME:
      return {
        ...state,
        endTime: action.payload.endTime
      };
    default:
      return state;
  }
}
