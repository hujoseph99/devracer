import React, { useCallback } from 'react';
import { useSelector } from 'react-redux'

import { Box, LinearProgress, Typography } from '@material-ui/core'

import { selectGameProgress } from './gameSlice';
import { GameProgress } from './types';

export const UserProgress = (): JSX.Element => {
	const progress = useSelector(selectGameProgress);

	const progresses = useCallback(() => {
		let res: JSX.Element[] = [];
		for (let i = 0; i < progress.length; i++) {
			res.push(<Progress progress={progress[i]} />)
		}
		return res;
	}, [progress]);

	return (
		<Box>
			{progresses()}
		</Box>
	);
}

interface ProgressProps {
	progress: GameProgress;
}

const Progress = ({ progress }: ProgressProps): JSX.Element => {
	return (
		<Box mb={1}>
			<Typography>{progress.displayName} | {progress.wpm}</Typography>
			<LinearProgress variant='determinate' value={progress.percentCompleted * 100} />
		</Box>
	)
}
