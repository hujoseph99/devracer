import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";
import axios from 'axios'
import { NewPracticeRace, NewPracticeRaceGQLResponse, RaceTextFieldState } from './types'
import { newPracticeRaceGQL } from "./graphql";

// redux prefix for this slice
const RACE_TEXT_FIELD = 'raceTextField';

// fetches a new practice race from the backend using graphql
export const fetchNewPracticeRace = createAsyncThunk<NewPracticeRace>(
	`${RACE_TEXT_FIELD}/fetchNewPracticeRace`, 
	async () => {
		const response = await axios.post<NewPracticeRaceGQLResponse>(
			'http://localhost:8080/graphql', 
			{ query: newPracticeRaceGQL }
		);
		return response.data.data.practiceRace;
	}
);

const initialState: RaceTextFieldState = {
	practiceRace: {
		snippet: {
			id: '0',
			raceContent: '',
			tokenCount: 0
		},
		timeLimit: 0
	},
	status: 'idle'
}

const raceTextFieldSlice = createSlice({
	name: 'raceTextField',
	initialState,
	reducers: {},
	extraReducers: builder => {
		builder.addCase(fetchNewPracticeRace.fulfilled, (state, action) => {
			state.status = 'succeeded';
			state.practiceRace = action.payload;
		});
		builder.addCase(fetchNewPracticeRace.rejected, (state, action) => {
			console.log(action);
		})
	}
})

export default raceTextFieldSlice.reducer;

export const selectSnippet = (state: RootState) => state.raceTextField.practiceRace.snippet;
