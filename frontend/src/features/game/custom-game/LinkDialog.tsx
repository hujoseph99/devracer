import React from 'react';
import { useSelector } from 'react-redux';

import { 
	Button, 
	Dialog, 
	DialogActions, 
	DialogContent, 
	DialogContentText, 
	DialogProps, 
	DialogTitle, 
	IconButton, 
	InputAdornment, 
	TextField 
} from '@material-ui/core';
import FileCopyIcon from '@material-ui/icons/FileCopy';

import { selectLobbyId } from '../gameSlice';

interface LinkDialogProps extends DialogProps {
	handleClose?: () => void;
}

export const LinkDialog = ({
	handleClose = () => {},
	...props
}: LinkDialogProps) => {
	const lobbyId = useSelector(selectLobbyId);
	const url = 'http://localhost:3000/custom/' + lobbyId;

	const handleCopyClick = () => {
		navigator.clipboard.writeText(url);
	}

	return (
		<Dialog
			open={props.open}
			onClose={handleClose}
			aria-labelledby="responsive-dialog-title"
		>
		<DialogTitle id="responsive-dialog-title">{"Invite your friends"}</DialogTitle>
		<DialogContent>
			<DialogContentText gutterBottom>
				Copy the link below and send it to your friends to have them join you in this race!
			</DialogContentText>
			<TextField 
				disabled 
				value={url} 
				variant='outlined' 
				fullWidth 
				InputProps={{
					endAdornment: (
						<InputAdornment position='end'>
							<IconButton onClick={handleCopyClick}>
								<FileCopyIcon />
							</IconButton>
						</InputAdornment>
					)
				}}
			/>
		</DialogContent>
		<DialogActions>
			<Button autoFocus onClick={handleClose} color="primary">
				Close
			</Button>
		</DialogActions>
	 </Dialog>
	)
}
