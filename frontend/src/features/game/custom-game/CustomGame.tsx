import React, { useEffect, useRef, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RouteComponentProps } from 'react-router';

import { Box, Button, Container, Grid, TextField } from '@material-ui/core';

import * as CONSTANTS from './constants'
import { CreateGameResponse, ErrorResponse, GameProgressResponse, JoinGameResponse, NewPlayerResponse } from '../types';
import { createGameAction, gameProgressAction, joinGameAction, newPlayerAction, selectLangauge, selectRaceContent } from '../gameSlice';
import { Footer } from '../../footer/Footer';
import { Navbar } from '../../navbar/Navbar';
import { RaceField } from '../../race-text-field/RaceField';
import { selectDisplayName } from '../../user/userSlice';

import "../../race-text-field/editor.css"
import { UserProgress } from '../UserProgress';

interface MatchParams {
	lobby?: string;
}

export const CustomGame = (props : RouteComponentProps<MatchParams>): JSX.Element => {
	const ws = useRef<WebSocket | undefined>(undefined);
	const dispatch = useDispatch();

	const displayName = useSelector(selectDisplayName);
	const raceContent = useSelector(selectRaceContent);
	const language = useSelector(selectLangauge);
	
	console.log(raceContent);

	const lobbyId = props.match.params.lobby ?? "";

	// connect to websocket
	useEffect(() => {
		ws.current = new WebSocket(`ws://localhost:8080/custom?name=${displayName}`);
		ws.current?.addEventListener('open', handleConnectedToWebsocket);
		ws.current.addEventListener('message', event => handleNewMessage(event))
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
		}
	}

	const handleErrorResponse = (payload: ErrorResponse) => {
		console.log(payload);
	}

	const handleCreateGameResponse = (payload: CreateGameResponse) => {
		dispatch(createGameAction(payload));
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

	return (
		<Container maxWidth='sm'>
			<Box minHeight='100vh' display='flex' flexDirection='column' py={5}>
				<Navbar />
				<Grid container justify='center'>
					<Grid item xs={12}>
						<UserProgress />
					</Grid>
					<Grid item className="aceEditorContainer">
						<RaceField snippet={raceContent} language={language} />
					</Grid>
				</Grid>
				<Footer />
			</Box>
		</Container>
	)
}

