import React, { useEffect, useRef, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RouteComponentProps } from 'react-router';

import { Box, Button, Container, Grid, TextField } from '@material-ui/core';

import * as CONSTANTS from './constants'
import { 
	CreateGameResponse, 
	ErrorResponse, 
	GameFinishedResponse, 
	GameProgressResponse, 
	GameStartResponse, 
	JoinGameResponse, 
	NewPlayerResponse, 
	PlayerFinishedResponse 
} from '../types';
import { 
	createGameAction, 
	gameFinishedAction,
	gameProgressAction, 
	gameStartAction, 
	joinGameAction, 
	newPlayerAction, 
	playerFinishedAction, 
	selectIsHost, 
	selectLangauge, 
	selectPlacements, 
	selectPlayerId, 
	selectRaceContent, 
	selectState 
} from '../gameSlice';
import { Footer } from '../../footer/Footer';
import { Navbar } from '../../navbar/Navbar';
import { RaceField } from '../../race-text-field/RaceField';
import { selectDisplayName } from '../../user/userSlice';

import "../../race-text-field/editor.css"
import { UserProgress } from '../UserProgress';
import { selectStatus } from '../../auth/authSlice';
import { StatusBar } from '../StatusBar';
import { LinkDialog } from './LinkDialog';
import { checkPlayerFinished } from '../utils';

interface MatchParams {
	lobby?: string;
}

export const CustomGame = (props : RouteComponentProps<MatchParams>): JSX.Element => {
	const ws = useRef<WebSocket | undefined>(undefined);
	const dispatch = useDispatch();

	const [showLink, setShowLink] = useState(false);

	const displayName = useSelector(selectDisplayName);
	const raceContent = useSelector(selectRaceContent);
	const language = useSelector(selectLangauge);
	const isHost = useSelector(selectIsHost);
	const state = useSelector(selectState);
	const playerId = useSelector(selectPlayerId);
	const placements = useSelector(selectPlacements);
	
	const lobbyId = props.match.params.lobby ?? "";

	// connect to websocket
	useEffect(() => {
		ws.current = new WebSocket(`ws://localhost:8080/custom?name=${displayName}`);
		ws.current?.addEventListener('open', handleConnectedToWebsocket);
		ws.current?.addEventListener('message', event => handleNewMessage(event))
		return () => {
			ws.current?.send(JSON.stringify({
				action: CONSTANTS.LEAVE_GAME_ACTION,
			}));
		}
	}, [])

	const handleConnectedToWebsocket = () => {
		if (lobbyId) {
			ws.current?.send(JSON.stringify({
				action: CONSTANTS.JOIN_GAME_ACTION,
				lobbyId: lobbyId,
			}));
		} else {
			ws.current?.send(JSON.stringify({
				action: CONSTANTS.CREATE_GAME_ACTION,
			}));
		}
	}

	const handleNewMessage = (event: MessageEvent) => {
		const message: { action: string, payload: any } = JSON.parse(event.data)
		switch (message.action) {
			case CONSTANTS.ERROR_RESPONSE:
				handleErrorResponse(message.payload as ErrorResponse);
				break;
			case CONSTANTS.CREATE_GAME_RESPONSE:
				handleCreateGameResponse(message.payload as CreateGameResponse);
				break
			case CONSTANTS.JOIN_GAME_RESPONSE:
				handleJoinGameResponse(message.payload as JoinGameResponse);
				break;
			case CONSTANTS.NEW_PLAYER_RESPONSE:
				handleNewPlayerResponse(message.payload as NewPlayerResponse);
				break;
			case CONSTANTS.GAME_PROGRESS_RESPONSE:
				handleGameProgressResponse(message.payload as GameProgressResponse);
				break;
			case CONSTANTS.GAME_START_RESPONSE:
				handleGameStartResponse(message.payload as GameStartResponse);
				break;
			case CONSTANTS.PLAYER_FINISHED_RESPONSE:
				handlePlayerFinishedResponse(message.payload as PlayerFinishedResponse);
				break
			case CONSTANTS.GAME_FINISHED_RESPONSE:
				handleGameFinishedResponse(message.payload as GameFinishedResponse);
				break

		}
	}

	const handleErrorResponse = (payload: ErrorResponse) => {
		console.log(payload);
	}

	const handleCreateGameResponse = (payload: CreateGameResponse) => {
		dispatch(createGameAction(payload));
		setShowLink(true);
	}

	const handleJoinGameResponse = (payload: JoinGameResponse) => {
		dispatch(joinGameAction(payload));
	}
	
	const handleNewPlayerResponse = (payload: NewPlayerResponse) => {
		dispatch(newPlayerAction(payload));
	}

	const handleGameProgressResponse = (payload: GameProgressResponse) => {
		dispatch(gameProgressAction(payload));
	}

	const handleGameStartResponse = (payload: GameStartResponse) => {
		dispatch(gameStartAction(payload));
	}

	const handlePlayerFinishedResponse = (payload: PlayerFinishedResponse) => {
		dispatch(playerFinishedAction(payload));
	}

	const handleGameFinishedResponse = (payload: GameFinishedResponse) => {
		dispatch(gameFinishedAction(payload));
	}

	const handleStartGameClick = () => {
		ws.current?.send(JSON.stringify({
			action: CONSTANTS.START_GAME_ACTION,
		}));
	}

	const handleNextGameClick = () => {
		ws.current?.send(JSON.stringify({
			action: CONSTANTS.NEXT_GAME_ACTION
		}));
	}
	
	const handleRaceFieldChange = (text: string) => {
		ws.current?.send(JSON.stringify({
			action: CONSTANTS.GAME_PROGRESS_ACTION,
			payload: text,
		}));
	}

	const handleLinkOpen = () => {
		setShowLink(true);
	}

	const handleLinkClose = () => {
		setShowLink(false);
	}

	return (
		<Container maxWidth='sm'>
			<Box minHeight='100vh' display='flex' flexDirection='column' py={5}>
				<Navbar />
				<Grid container justify='center'>
					<Grid item xs={12}>
						<Box mt={2}>
							<StatusBar handleOpenLinkDialog={handleLinkOpen} />
						</Box>
					</Grid>
					<Grid item xs={12}>
						<UserProgress />
					</Grid>
					<Grid item className="aceEditorContainer">
						<RaceField 
							snippet={raceContent} 
							language={language} 
							disabled={state !== 'inProgress' || checkPlayerFinished(placements, playerId)} 
							onChange={handleRaceFieldChange}
						/>
					</Grid>
					{ isHost ? (
						<Grid container item xs={12} justify='flex-end'>
							<Grid item>
								<Box mt={2}>
									{state === 'finished' ? (
										<Button  variant='contained' onClick={handleNextGameClick}>Next Game</Button>
									) : (
										<Button 
											variant='contained'
											onClick={handleStartGameClick}
											disabled={state !== 'waiting'}
										>
											Start Game
										</Button>
									)}
								</Box>
							</Grid>
						</Grid>
					) : null
					}
				</Grid>
				<Footer />
			</Box>
			<LinkDialog open={showLink} handleClose={handleLinkClose} />
		</Container>
	)
}

