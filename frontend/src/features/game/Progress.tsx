import { Box, createStyles, LinearProgress, makeStyles, Theme, Typography } from '@material-ui/core';
import { green, lightGreen } from '@material-ui/core/colors';
import React from 'react';
import { useSelector } from 'react-redux';
import { selectPlacements } from './gameSlice';
import { GameProgress } from './types';
import { checkPlayerFinished } from './utils';


interface StyleProps {
	finished: boolean;
}

const useStyles = makeStyles<Theme, StyleProps>((theme: Theme) => createStyles({
	progress: ({ finished }) => {
		if (finished) {
			return {
				backgroundColor: green[500],
				width: '100%'
			}
		} 
		return {
			width: '100%',
		}
	}
}));

interface ProgressProps {
	progress: GameProgress;
}

export const Progress = ({ progress }: ProgressProps): JSX.Element => {
	const placements = useSelector(selectPlacements);

	const finished = checkPlayerFinished(placements, progress.playerId)
	const styles = useStyles({ finished });

	const value = finished ? 100 : progress.percentCompleted * 100;

	return (
		<Box display='flex'alignItems='center' mb={1} width='100%'>
			<Typography variant='caption'>{progress.displayName}</Typography>
			<Box display='flex' flexGrow={1} mx={1}>
				<LinearProgress 
					variant='determinate' 
					value={value} 
					classes={{ colorPrimary: styles.progress, barColorPrimary: styles.progress }}
				/>
			</Box>
			<Typography variant='caption' align='right'>{progress.wpm} wpm</Typography>
		</Box>
	)
}
