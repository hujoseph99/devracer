import React from 'react';

import { 
	Button, 
	Dialog, 
	DialogActions, 
	DialogContent, 
	DialogContentText, 
	DialogProps, 
	DialogTitle, 
} from '@material-ui/core';

interface LinkDialogProps extends DialogProps {
	handleClose?: () => void;
}

export const LobbyClosedDialog = ({
	handleClose = () => {},
	...props
}: LinkDialogProps) => {
	return (
		<Dialog
			open={props.open}
			onClose={handleClose}
			aria-labelledby="responsive-dialog-title"
		>
			<DialogTitle id="responsive-dialog-title">Lobby closed</DialogTitle>
			<DialogContent>
				<DialogContentText gutterBottom>
					The lobby has closed because the host left the game. 
					Feel free to join another lobby or practice on your own!
				</DialogContentText>
			</DialogContent>
			<DialogActions>
				<Button autoFocus onClick={handleClose} color="primary">
					Close
				</Button>
			</DialogActions>
		</Dialog>
	)
}
