import axios from "axios";
import { GET_RACE } from "./types";

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
