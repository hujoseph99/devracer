import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";
import { CreateGameResponse, GameFinishedResponse, GameProgress, GameProgressResponse, GameStartResponse, GameState, JoinGameResponse, NewPlayerResponse, PlayerFinishedResponse } from "./types";
import { transformSnippetResponse } from "./utils";

// redux prefix for this slice
const GAME_SLICE_NAME = 'game';

const initialState: GameState = {
	playerId: '',
	lobbyId: '',
	isHost: false,
	state: 'waiting',
	snippet: {
		id: '',
		snippet: '',
		tokenCount: 0,
		language: 'plain_text',
		dateCreated: new Date(),
	},
	isQueued: false,
	gameProgress: [],
	queuedPlayers: [],
	placements: [],
	countdown: 0,
}

const gameSlice = createSlice({
	name: GAME_SLICE_NAME,
	initialState,
	reducers: {
		createGameAction: (state, action: PayloadAction<CreateGameResponse>) => {
			const payload = action.payload;
			state.state = 'waiting';
			state.isHost = true;
			state.playerId = payload.playerId;
			state.lobbyId = payload.lobbyId;
			state.snippet = transformSnippetResponse(payload.snippet);
		},
		joinGameAction: (state, action: PayloadAction<JoinGameResponse>) => {
			const payload = action.payload;
			state.state = 'waiting';
			state.playerId = payload.playerId;
			state.snippet = transformSnippetResponse(payload.snippet);
			state.gameProgress = payload.gameProgress;
			state.queuedPlayers = payload.queuedPlayers;
			state.placements = payload.placements;
			state.isQueued = payload.wasQueued;
		},
		newPlayerAction: (state, action: PayloadAction<NewPlayerResponse>) => {
			const payload = action.payload;
			const gameProgress: GameProgress = {
				playerId: payload.playerId,
				displayName: payload.displayName,
				percentCompleted: payload.percentCompleted,
				wpm: 0,
			}
			if (payload.wasQueued) {
				state.queuedPlayers.push(gameProgress);
			} else {
				state.gameProgress.push(gameProgress)
			}
		},
		gameProgressAction: (state, action: PayloadAction<GameProgressResponse>) => {
			const payload = action.payload;
			// should only update the game progress if they are actually particpating in the game
			for (let i = 0; i < state.gameProgress.length; i++) {
				if (state.gameProgress[i].playerId === payload.playerId) {
					state.gameProgress[i].percentCompleted = payload.percentCompleted;
					state.gameProgress[i].wpm = payload.wpm;
				}
			}
		},
		gameStartAction: (state, action: PayloadAction<GameStartResponse>) => {
			const payload = action.payload;
			if (payload.countdown === 0) {
				state.state = 'inProgress'
			} else {
				state.state = 'countdown'
			}
			state.countdown = payload.countdown;
		},
		playerFinishedAction: (state, action: PayloadAction<PlayerFinishedResponse>) => {
			const payload = action.payload;
			state.placements = payload.placements;
		},
		gameFinishedAction: (state, action: PayloadAction<GameFinishedResponse>) => {
			const payload = action.payload;
			state.state = 'finished';
			state.placements = payload.placements;
		}
	},
})

export default gameSlice.reducer;

export const { 
	createGameAction, 
	gameFinishedAction,
	gameProgressAction, 
	gameStartAction, 
	joinGameAction, 
	playerFinishedAction, 
	newPlayerAction 
} = gameSlice.actions;

export const selectState = (state: RootState) => state.game.state;
export const selectIsHost = (state: RootState) => state.game.isHost;
export const selectLobbyId = (state: RootState) => state.game.lobbyId;
export const selectPlayerId = (state: RootState) => state.game.playerId;

export const selectRaceContent = (state: RootState) => state.game.snippet.snippet;
export const selectLangauge = (state: RootState) => state.game.snippet.language;
export const selectGameProgress = (state: RootState) => state.game.gameProgress;
export const selectCountdown = (state: RootState) => state.game.countdown;
export const selectPlacements = (state: RootState) => state.game.placements;
