import axios from "axios";
import {
  GET_RACE,
  INPUT_CORRECT,
  INPUT_INCORRECT,
  INPUT_FINISHED_WORD
} from "./types";

// getRace gets a single race from backend
export const getRace = () => dispatch => {
  axios
    .get("/api/race")
    .then(res =>
      dispatch({
        type: GET_RACE,
        payload: res.data
      })
    )
    .catch(err => console.log(err));
};

// updates state to reflect the new ending index of the correct portion of the snippet
export const inputCorrect = end => dispatch => {
  dispatch({
    type: INPUT_CORRECT,
    payload: { end }
  });
};

// updates the state to reflect the new starting and ending indices of the incorrect portion of the snippet
export const inputIncorrect = (start, end) => dispatch => {
  dispatch({
    type: INPUT_INCORRECT,
    payload: { start, end }
  });
};

// updates the state to reflect the new current word (i.e., called when a word is completed)
export const inputFinishedWord = newStart => dispatch => {
  dispatch({
    type: INPUT_FINISHED_WORD,
    payload: { newStart }
  });
};
