import { Box, Button, TextField } from '@material-ui/core';
import React, { useEffect, useRef, useState } from 'react';

export const CustomGame = (): JSX.Element => {
	const ws = useRef<WebSocket | undefined>(undefined);
	const [messages, setMessages] = useState<string[]>([]);
	const [currMessage, setCurrMessage] = useState("");

	// connect to websocket
	useEffect(() => {
		ws.current = new WebSocket("ws://localhost:8080/custom");
		ws.current.addEventListener('open', () => { onWebsocketOpen() })
		ws.current.addEventListener('message', event => handleNewMessage(event))
	}, [])

	const handleNewMessage = (event: MessageEvent) => {
		let data = event.data;
		console.log(data);
		let dataArr = data.split(/\r?\n/);
		for (let i = 0; i < dataArr.length; i++) {
			let msg = JSON.parse(dataArr[i]);
			setMessages(messages => [...messages, msg.message]);
		}
	}

	const sendMessage = () => {
		if (currMessage !== "") {
			ws.current?.send(JSON.stringify({ message: currMessage }));
			setCurrMessage("");
		}
	}

	const handleChangeMessage = (e: React.ChangeEvent<HTMLInputElement>) => {
		setCurrMessage(e.target.value);
	}

	const onWebsocketOpen = () => {
		console.log("connected to WS!");
	}

	const messagesDisplay = () => {
		let messagesConverted = [];
		for (let i = 0; i < messages.length; i++) {
			messagesConverted.push(<p>{messages[i]}</p>)
		}
		return messagesConverted;
	}

	return (
		<Box>
			{messagesDisplay()}
			<TextField onChange={handleChangeMessage} value={currMessage} />
			<Button onClick={sendMessage} variant='contained'>SEND</Button>
		</Box>
	)
}

