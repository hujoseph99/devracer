import { createSelector, createSlice } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";

interface RaceTextFieldState {
	snippet: string
}

const initialState: RaceTextFieldState = {
	snippet: ''
}

const raceTextFieldSlice = createSlice({
	name: 'raceTextField',
	initialState,
	reducers: {},
	extraReducers: builder => {

	}
})

export default raceTextFieldSlice.reducer;

export const selectSnippet = (state: RootState) => state.RaceTextField.snippet;
