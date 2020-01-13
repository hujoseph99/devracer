import { combineReducers } from "redux";
import race from "./raceReducer";
import user from "./userReducer";

export default combineReducers({ race, user });
