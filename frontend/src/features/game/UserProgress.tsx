import React, { useCallback } from 'react';
import { useSelector } from 'react-redux'

import { Box } from '@material-ui/core'

import { selectGameProgress, selectQueuedPlayers } from './gameSlice';
import { Progress } from './Progress';

export const UserProgress = (): JSX.Element => {
	const progress = useSelector(selectGameProgress);
	const queuedPlayers = useSelector(selectQueuedPlayers);

	const progresses = useCallback(() => {
		let res: JSX.Element[] = [];
		for (let i = 0; i < progress.length; i++) {
			res.push(<Progress progress={progress[i]} />)
		}
		for (let i = 0; i < queuedPlayers.length; i++) {
			res.push(<Progress progress={queuedPlayers[i]} queued />)
		}
		return res;
	}, [progress, queuedPlayers]);

	return (
		<Box width='100%'>
			{progresses()}
		</Box>
	);
}
