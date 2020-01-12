import {
  UPDATE_RACE,
  INPUT_CORRECT,
  INPUT_INCORRECT,
  INPUT_FINISHED_WORD,
  UPDATE_WPM,
  SET_START_TIME,
  SET_END_TIME
} from "./types";

export const updateRace = room => dispatch => {
  dispatch({
    type: UPDATE_RACE,
    payload: room
  })
}

// getRace gets a single race from backend
export const getRace = () => dispatch => {
  // axios
  //   .get("/api/race")
  //   .then(res => {
  //     dispatch({
  //       type: GET_RACE,
  //       payload: res.data
  //     });
  //   })
  //   .catch(err => console.log(err));
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

// updates the state to reflect the new current word (e.g., called when a word is completed)
export const inputFinishedWord = newStart => dispatch => {
  dispatch({
    type: INPUT_FINISHED_WORD,
    payload: { newStart }
  });
};

// updates the state to reflect the new current wpm
export const updateWPM = newWPM => dispatch => {
  dispatch({
    type: UPDATE_WPM,
    payload: { newWPM }
  });
};

// updates start time
export const setStartTime = startTime => dispatch => {
  dispatch({
    type: SET_START_TIME,
    payload: { startTime }
  });
};

// updates end time
export const setEndTime = endTime => dispatch => {
  dispatch({
    type: SET_END_TIME,
    payload: { endTime }
  });
};
