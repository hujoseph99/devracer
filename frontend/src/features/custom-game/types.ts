import { RaceSnippet } from '../race-text-field/types'

export interface GameProgress {
	playerId: string;
	displayName: string;
	percentCompleted: number;
	wpm: number;
}

export interface ErrorResponse {
	message: string;
}

export interface CreateGameResponse {
	playerId: string;
	lobbyId: string;
	snippet: RaceSnippet;
}

export interface JoinGameResponse {
	playerId: string;
	snippet: RaceSnippet;
	gameProgress: GameProgress[];
	queuedPlayers: GameProgress[];
	placements: string[];
	wasQueued: boolean;
}
