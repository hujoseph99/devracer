import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { RootState } from "../../app/store";
import { UserRequest, UserState, UserResponse, UserStateGQLResponse } from "./types";
import { fetchUserDataGQL } from './graphql';

const USER_SLICE_NAME = "user";

export const fetchUserData = createAsyncThunk<UserResponse, UserRequest>(
	`${USER_SLICE_NAME}/fetchUserData`,
	async ({ userid }) => {
		const response = await axios.post<UserStateGQLResponse>(
			'http://localhost:8080/graphql',
			{ query: fetchUserDataGQL, variables: { userid } }
		)

		return response.data.data.user;
	}
)

const initialState: UserState = {
	profile: {
		totalWordsTyped: 0,
		racesCompleted: 0,
		racesWon: 0,
		maxTPM: 0,
		averageTPMAllTime: 0,
		averageTPMLast10: 0,
	},
	preferences: {
		displayName: "",
	},
	status: 'idle',
}

const userSlice = createSlice({
	name: USER_SLICE_NAME,
	initialState,
	reducers: {},
	extraReducers: builder => {
		builder.addCase(fetchUserData.fulfilled, (state, action) => {
			state.preferences = action.payload.preferences;
			state.profile = action.payload.profile;
			state.status = 'succeeded';
		});
		builder.addCase(fetchUserData.pending, state => {
			state.status = 'loading';
		});
		builder.addCase(fetchUserData.rejected, state => {
			state.preferences = initialState.preferences;
			state.profile = initialState.profile;
			state.status = 'failed';
		});
	},
});

export default userSlice.reducer;

export const selectDisplayName = (state: RootState) => state.user.preferences.displayName;
