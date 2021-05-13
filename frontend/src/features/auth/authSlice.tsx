import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import { AuthState, JWTPayload, LoginBody, LoginResponse, LogoutBody, RefreshBody, RefreshResponse, RegisterBody, RegisterResponse } from "./types"
import jwtDecode from 'jwt-decode';

import axios from 'axios';
import { RootState } from "../../app/store";
import { resetUser } from "../user/userSlice";
import { BACKEND_HOSTNAME } from "../../config";

const AUTH_SLICE_NAME = "auth"

export const login = createAsyncThunk<LoginResponse, LoginBody>(
	`${AUTH_SLICE_NAME}/login`,
	async body => {
		const response = await axios.post<LoginResponse>(
			`${BACKEND_HOSTNAME}/auth/login`,
			body
		)
		return response.data;
	}
);


export const githubCallback = createAsyncThunk<LoginResponse, URLSearchParams>(
	`${AUTH_SLICE_NAME}/githubCallback`,
	async body => {
		const response = await axios.post<LoginResponse>(
			`${BACKEND_HOSTNAME}/auth/githubCallback?${body.toString()}`,
			body, { withCredentials: true }
		)
		return response.data;
	}
);

export const register = createAsyncThunk<RegisterResponse, RegisterBody>(
	`${AUTH_SLICE_NAME}/register`,
	async body => {
		const response = await axios.post<RegisterResponse>(
			`${BACKEND_HOSTNAME}/auth/register`,
			body
		);
		return response.data;
	}
);

export const logout = createAsyncThunk<{}, LogoutBody>(
	`${AUTH_SLICE_NAME}/logout`,
	async (body, thunkAPI)=> {
		await axios.post<{}>(
			`${BACKEND_HOSTNAME}/auth/logout`,
			body
		);
		thunkAPI.dispatch(resetUser());
		return {};
	}
)

export const refresh = createAsyncThunk<RefreshResponse, RefreshBody>(
	`${AUTH_SLICE_NAME}/refresh`,
	async body => {
		const response = await axios.post<RefreshResponse>(
			`${BACKEND_HOSTNAME}/auth/refresh`,
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

		builder.addCase(githubCallback.fulfilled, (state, action) => {
			state.status = 'succeeded';
			state.accessToken = action.payload.accessToken;
			state.refreshToken = action.payload.refreshToken;
			state.isLoggedIn = true;
		});
		builder.addCase(githubCallback.pending, state => {
			state.status = 'loading';
		});
		builder.addCase(githubCallback.rejected, state => {
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

		builder.addCase(logout.fulfilled, state => {
			state.accessToken = "";
			state.refreshToken = "";
			state.status = 'succeeded';
			state.isLoggedIn = false;
		})
		builder.addCase(logout.pending, state => {
			state.status = 'loading';
		});
		builder.addCase(logout.rejected, state => {
			state.status = 'failed';
		})
	},
});

export default authSlice.reducer;

export const { resetStatus } = authSlice.actions;

export const selectAccessToken = (state: RootState) => state.auth.accessToken;
export const selectRefreshToken = (state: RootState) => state.auth.refreshToken;
export const selectStatus = (state: RootState) => state.auth.status;
export const selectIsLoggedIn = (state: RootState) => state.auth.isLoggedIn;

export const selectUserID = (state:RootState) => {
	if (state.auth.accessToken) {
		const accessToken = state.auth.accessToken;
		const decoded = jwtDecode<JWTPayload>(accessToken);
		return decoded.userid;
	}
	return "";
}
