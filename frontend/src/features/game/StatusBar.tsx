import React from 'react';
import { useSelector } from 'react-redux';

import { Box, CircularProgress, Grid, IconButton, Typography } from '@material-ui/core';
import LinkIcon from '@material-ui/icons/Link';

import { selectCountdown, selectIsHost, selectState } from './gameSlice';

const COUNTDOWN_START = 5;

interface StatusBarProps {
	handleOpenLinkDialog: () => void;
}

export const StatusBar = ({ handleOpenLinkDialog }: StatusBarProps): JSX.Element => {
	const state = useSelector(selectState);
	const isHost = useSelector(selectIsHost);
	const countdown = useSelector(selectCountdown);

	const progressBarValue = countdown === 0 ? 100 : Math.round(COUNTDOWN_START - countdown / COUNTDOWN_START * 100);

	let status = '';
	if (state === 'waiting') {
		if (isHost) {
			status = 'The lobby is ready. Feel free to start the game whenever you want.';
		} else {
			status = 'The lobby is ready. Please wait for the host to start the game.';
		}
	} else if (state === 'countdown') {
		status = 'Get ready! The game is starting.'
	} else if (state === 'inProgress') {
		status = 'The game is in progress!';
	} else if (state === 'finished') {
		if (isHost) {
			status = 'The race is finished! Please fetch the next game.';
		} else {
			status = 'The race is finished! Please wait for the host to move onto the next game.';
		}
	}

	return (
		<Grid container justify='space-between' alignItems='center' spacing={1}>
			<Grid item>
				<IconButton onClick={handleOpenLinkDialog}>
					<LinkIcon />
				</IconButton>
			</Grid>
			<Grid item>
				<Box display='flex' alignItems='center'>
					<Box mr={1}>
						<Typography component='span' variant='caption'>{status}</Typography>
					</Box>
					{state === 'countdown' ? (
						<Box position="relative" display="inline-flex">
							<CircularProgress variant="determinate" value={progressBarValue} size={30} />
							<Box
								top={0}
								left={0}
								bottom={0}
								right={0}
								position="absolute"
								display="flex"
								alignItems="center"
								justifyContent="center"
							>
								<Typography variant="caption" component="div" color="textSecondary">{countdown}</Typography>
							</Box>
						</Box>
					) : null}
				</Box>
			</Grid>
		</Grid>
	)
}
