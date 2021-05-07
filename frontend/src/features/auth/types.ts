type status = 'idle' | 'loading' | 'succeeded' | 'failed';

export interface AuthState {
	accessToken: string;
	refreshToken: string;
	isLoggedIn: boolean;
	status: status;
};

export interface LoginResponse {
	accessToken: string;
	refreshToken: string;
}

export interface LoginBody {
	username: string;
	password: string;
	rememberMe: boolean;
}

export interface RegisterResponse {
	accessToken: string;
	refreshToken: string;
}

export interface RegisterBody {
	username: string;
	password: string;
	email: string;
	nickname: string;
}

export interface RefreshResponse {
	accessToken: string;
	refreshToken: string;
}

export interface RefreshBody {
	refreshToken: string;
}

export interface LogoutBody {
	refreshToken: string;
}

export interface JWTPayload {
	exp: number;
	userid: string;
}
