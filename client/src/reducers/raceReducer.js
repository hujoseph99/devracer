import { GET_RACE } from "../actions/types";

const initialState = {
  snippet: ""
};

export default function(state = initialState, action) {
  switch (action.type) {
    case GET_RACE:
      return {
        ...state,
        snippet: action.payload.snippet
      };
    default:
      return state;
  }
}
