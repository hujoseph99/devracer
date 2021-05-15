import React from 'react';
import { useSelector } from 'react-redux';

import { Box, createStyles, LinearProgress, makeStyles, Theme, Typography } from '@material-ui/core';
import { green, grey } from '@material-ui/core/colors';

import { selectPlacements } from './gameSlice';
import { GameProgress } from './types';
import { checkPlayerFinished } from './utils';


interface StyleProps {
	queued: boolean;
	finished: boolean;
}

const useStyles = makeStyles<Theme, StyleProps>((theme: Theme) => createStyles({
	progress: ({ finished, queued }) => {
		if (finished) {
			return {
				backgroundColor: green[500],
				width: '100%',
			}
		} 
		if (queued) {
			return {
				backgroundColor: grey[800],
				width: '100%',
			}
		}
		return {
			width: '100%',
		}
	},
	typography: ({ queued }) => {
		if (queued) {
			return {
				color: grey[400],
			};
		}
		return {};
	}
}));

interface ProgressProps {
	progress: GameProgress;
	queued?: boolean;
}

export const Progress = ({ progress, queued = false }: ProgressProps): JSX.Element => {
	const placements = useSelector(selectPlacements);

	const finished = checkPlayerFinished(placements, progress.playerId);
	const classes = useStyles({ finished, queued });

	const value = finished ? 100 : progress.percentCompleted * 100;

	return (
		<Box display='flex'alignItems='center' mb={1} width='100%'>
			<Typography variant='caption' className={classes.typography}>{progress.displayName}</Typography>
			<Box display='flex' flexGrow={1} mx={1}>
				<LinearProgress 
					variant='determinate' 
					value={value} 
					classes={{ colorPrimary: classes.progress, barColorPrimary: classes.progress }}
				/>
			</Box>
			<Typography variant='caption' align='right' className={classes.typography}>{progress.wpm} wpm</Typography>
		</Box>
	)
}
