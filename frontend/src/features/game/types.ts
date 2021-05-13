export type language = 'plain_text' | 'c_cpp' | 'golang' | 'javascript' | 'python';

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
	snippet: SnippetState;
	isQueued: boolean;
	gameProgress: GameProgress[];
	queuedPlayers: GameProgress[];
	placements: string[];
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
