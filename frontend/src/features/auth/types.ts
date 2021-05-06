type status = 'idle' | 'loading' | 'succeeded' | 'failed';

export interface AuthState {
	accessToken: string;
	refreshToken: string;
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
