import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import { AuthState, LoginBody, LoginResponse } from "./types"

import axios from 'axios';
import { RootState } from "../../app/store";

const AUTH_SLICE_NAME = "auth"

export const login = createAsyncThunk<LoginResponse, LoginBody>(
	`${AUTH_SLICE_NAME}/login`,
	async body => {
		const response = await axios.post<LoginResponse>(
			'http://localhost:8080/auth/login',
			body
		)
		return response.data;
	}
)

const initialState: AuthState = {
	accessToken: "",
	refreshToken: "",
	status: 'idle'
};

export const authSlice = createSlice({
	name: AUTH_SLICE_NAME,
	initialState,
	reducers: {},
	extraReducers: builder => {
		builder.addCase(login.fulfilled, (state, action) => {
			state.status = 'succeeded';
			state.accessToken = action.payload.accessToken;
			state.refreshToken = action.payload.refreshToken;
		});
		builder.addCase(login.pending, state => {
			state.status = 'loading';
		});
		builder.addCase(login.rejected, state => {
			state.status = 'failed';
		});
	},
});

export default authSlice.reducer;

export const selectAccessToken = (state: RootState) => state.auth.accessToken;
export const selectRefreshToken = (state: RootState) => state.auth.refreshToken;
export const selectStatus = (state: RootState) => state.auth.status;
