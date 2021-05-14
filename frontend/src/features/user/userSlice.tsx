import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { RootState } from "../../app/store";
import { UserRequest, UserState, UserResponse, UserStateGQLResponse } from "./types";
import { fetchUserDataGQL } from './graphql';
import { BACKEND_HOSTNAME } from "../../config";

const USER_SLICE_NAME = "user";

export const fetchUserData = createAsyncThunk<UserResponse, UserRequest>(
	`${USER_SLICE_NAME}/fetchUserData`,
	async ({ userid }) => {
		const response = await axios.post<UserStateGQLResponse>(
			`${BACKEND_HOSTNAME}/graphql`,
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
	reducers: {
		resetUser: state => {
			state.profile = initialState.profile;
			state.preferences = initialState.preferences;
			state.status = initialState.status;
		}
	},
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

export const { resetUser } = userSlice.actions;

export const selectDisplayName = (state: RootState) => {
	if (state.auth.isLoggedIn) {
		return state.user.preferences.displayName;
	}
	return "Guest"
}
