import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";

import axios from 'axios'

import { newPracticeRaceGQL } from "./graphql";
import { Race, NewPracticeRaceGQLResponse, RaceFieldState } from './types'
import { mapGQLPracticeRaceToNewPracticeRace } from "./utils";
import { BACKEND_HOSTNAME } from "../../config";

// redux prefix for this slice
const RACE_FIELD = 'raceField';

// fetches a new practice race from the backend using graphql
export const fetchNewPracticeRace = createAsyncThunk<Race>(
	`${RACE_FIELD}/fetchNewPracticeRace`, 
	async () => {
		const response = await axios.post<NewPracticeRaceGQLResponse>(
			`${BACKEND_HOSTNAME}/graphql`, 
			{ query: newPracticeRaceGQL }
		);
		return mapGQLPracticeRaceToNewPracticeRace(response.data);
	}
);

const initialState: RaceFieldState = {
	race: {
		snippet: {
			id: '0',
			raceContent: '',
			tokenCount: 0,
			language: 'plain_text',
		},
		typedSoFar: '',
		timeLimit: 0
	},
	status: 'idle'
}

const raceFieldSlice = createSlice({
	name: RACE_FIELD,
	initialState,
	reducers: {
		// action payload should just contain the string that the user has typed
		// so far
		typedSoFarChanged: (state, action: PayloadAction<string>) => {
			state.race.typedSoFar = action.payload;
		}
	},
	extraReducers: builder => {
		builder.addCase(fetchNewPracticeRace.fulfilled, (state, action) => {
			state.status = 'succeeded';
			state.race = action.payload;
		});
		// TODO: Add some error handling here, have to figure out what to do in the
		// case of an error
		builder.addCase(fetchNewPracticeRace.rejected, (state, action) => {
			console.log(action);
		})
	}
})

export default raceFieldSlice.reducer;

export const { typedSoFarChanged } = raceFieldSlice.actions;

export const selectSnippet = (state: RootState) => state.raceField.race.snippet;
export const selectTypedSoFar = (state: RootState) => state.raceField.race.typedSoFar;
