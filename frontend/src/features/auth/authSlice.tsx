import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import { AuthState, JWTPayload, LoginBody, LoginResponse, RefreshBody, RefreshResponse, RegisterBody, RegisterResponse } from "./types"
import jwtDecode from 'jwt-decode';

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
);

export const register = createAsyncThunk<RegisterResponse, RegisterBody>(
	`${AUTH_SLICE_NAME}/register`,
	async body => {
		const response = await axios.post<RegisterResponse>(
			'http://localhost:8080/auth/register',
			body
		);
		return response.data;
	}
);

export const refresh = createAsyncThunk<RefreshResponse, RefreshBody>(
	`${AUTH_SLICE_NAME}/refresh`,
	async body => {
		const response = await axios.post<RefreshResponse>(
			'http://localhost:8080/auth/refresh',
			body
		);
		return response.data;
	}
)

const initialState: AuthState = {
	accessToken: "",
	refreshToken: "",
	status: 'idle',
	isLoggedIn: false,
};

export const authSlice = createSlice({
	name: AUTH_SLICE_NAME,
	initialState,
	reducers: {
		resetStatus: state => {
			state.status = 'idle'
		}
	},
	extraReducers: builder => {
		builder.addCase(login.fulfilled, (state, action) => {
			state.status = 'succeeded';
			state.accessToken = action.payload.accessToken;
			state.refreshToken = action.payload.refreshToken;
			state.isLoggedIn = true;
		});
		builder.addCase(login.pending, state => {
			state.status = 'loading';
		});
		builder.addCase(login.rejected, state => {
			state.status = 'failed';
		});

		builder.addCase(register.fulfilled, (state, action) => {
			state.status = 'succeeded';
			state.accessToken = action.payload.accessToken;
			state.refreshToken = action.payload.refreshToken;
			state.isLoggedIn = true;
		});
		builder.addCase(register.pending, state => {
			state.status = 'loading';
		});
		builder.addCase(register.rejected, state => {
			state.status = 'failed';
		});
		builder.addCase(refresh.fulfilled, (state, action) => {
			state.status = 'idle';
			state.accessToken = action.payload.accessToken;
			state.refreshToken = action.payload.refreshToken;
			state.isLoggedIn = true;
		});
		builder.addCase(refresh.pending, state => {
			state.status = 'loading';
		});
		builder.addCase(refresh.rejected, state => {
			state.status = 'failed';
			state.accessToken = "";
			state.refreshToken = "";
			state.isLoggedIn = false;
		});
	},
});

export default authSlice.reducer;

export const { resetStatus } = authSlice.actions;

export const selectAccessToken = (state: RootState) => state.auth.accessToken;
export const selectRefreshToken = (state: RootState) => state.auth.refreshToken;
export const selectStatus = (state: RootState) => state.auth.status;
export const selectIsLoggedIn = (state: RootState) => state.auth.isLoggedIn;

export const selectUserID = (state:RootState) => {
	const accessToken = state.auth.accessToken;
	const decoded = jwtDecode<JWTPayload>(accessToken);
	return decoded.userid;
}
