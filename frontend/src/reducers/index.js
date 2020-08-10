import { combineReducers } from "redux";
import raceReducer from "./raceReducer";
import userReducer from "./userReducer";
import errorReducer from "./errorReducer";
import routerReducer from "./routerReducer";

export default combineReducers({
  race: raceReducer,
  user: userReducer,
  error: errorReducer,
  router: routerReducer
});
