type status = 'idle' | 'loading' | 'succeeded' | 'failed';

export interface Profile {
  totalWordsTyped: number;
  racesCompleted: number;
  racesWon: number;
  maxTPM: number;
  averageTPMAllTime: number;
  averageTPMLast10: number;
}

export interface Preferences {
	displayName: string;
}

export interface UserState {
	profile: Profile;
	preferences: Preferences;
	status: status;
}

export interface UserResponse {
	profile: Profile;
	preferences: Preferences;
}

export interface UserRequest {
	userid: string;
}

export interface UserStateGQLResponse { 
	data: {
		user: UserResponse;
	}
}
