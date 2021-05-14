export type language = 'plain_text' | 'c_cpp' | 'golang' | 'javascript' | 'python';
export type lobbyState = 'waiting' | 'countdown' | 'inProgress' | 'finished';

export interface SnippetState {
	id: string;
	snippet: string;
	tokenCount: number;
	language: language;
	dateCreated: Date;
}

export interface GameState {
	playerId: string;
	lobbyId: string;
	isHost: boolean;
	state: lobbyState;
	snippet: SnippetState;
	isQueued: boolean;
	gameProgress: GameProgress[];
	queuedPlayers: GameProgress[];
	placements: string[];
	countdown: number;
}

export interface GameProgress {
	playerId: string;
	displayName: string;
	percentCompleted: number;
	wpm: number;
}

export interface ErrorResponse {
	message: string;
}

export interface SnippetResponse {
	id: string;
	snippet: string;
	tokenCount: number;
	language: number;
	dateCreated: Date;
}

export interface CreateGameResponse {
	playerId: string;
	lobbyId: string;
	snippet: SnippetResponse;
}

export interface JoinGameResponse {
	playerId: string;
	snippet: SnippetResponse;
	gameProgress: GameProgress[];
	queuedPlayers: GameProgress[];
	placements: string[];
	wasQueued: boolean;
}

export interface NewPlayerResponse {
	playerId: string;
	displayName: string;
	percentCompleted: number;
	wasQueued: boolean;
}

export interface GameProgressResponse {
	playerId: string;
	percentCompleted: number;
	wpm: number;
}

export interface GameStartResponse {
	countdown: number;
}

export interface PlayerFinishedResponse {
	placements: string[];
}

export interface GameFinishedResponse {
	placements: string[];
}

export interface NextGameResponse {
	snippet: SnippetResponse;
	gameProgress: GameProgress[];
	queuedPlayers: GameProgress[];
	placements: string[];
}

export interface LeaveGameResponse {
	playerId: string;
	placements: string[];
}

