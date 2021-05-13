import { Box, Button, Container, Grid, TextField } from '@material-ui/core';
import React, { useEffect, useRef, useState } from 'react';
import { useSelector } from 'react-redux';
import { RouteComponentProps } from 'react-router';
import { Footer } from '../footer/Footer';
import { Navbar } from '../navbar/Navbar';
import { RaceField } from '../race-text-field/RaceField';
import { selectDisplayName } from '../user/userSlice';
import * as CONSTANTS from './constants'

import "../race-text-field/editor.css"
import { CreateGameResponse, ErrorResponse, JoinGameResponse } from './types';
import { isConstructorDeclaration } from 'typescript';


interface MatchParams {
	lobby?: string;
}

export const CustomGame = (props : RouteComponentProps<MatchParams>): JSX.Element => {
	const ws = useRef<WebSocket | undefined>(undefined);
	const displayName = useSelector(selectDisplayName);
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
		}
	}

	const handleErrorResponse = (payload: ErrorResponse) => {
		console.log(payload);
	}

	const handleCreateGameResponse = (payload: CreateGameResponse) => {
		console.log(payload);
	}

	const handleJoinGameResponse = (payload: JoinGameResponse) => {
		console.log(payload);
	}


	return (
		<Container maxWidth='sm'>
			<Box minHeight='100vh' display='flex' flexDirection='column' py={5}>
				<Navbar />
				<Grid container justify='center'>
					<Grid item className="aceEditorContainer">
						<RaceField />
					</Grid>
				</Grid>
				<Footer />
			</Box>
		</Container>
	)
}

