import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux';

import { Box } from '@material-ui/core';

import { BackgroundEditor } from './BackgroundEditor';
import { ForegroundEditor } from './ForegroundEditor';
import { fetchNewPracticeRace } from './raceFieldSlice';

export const RaceField = (): JSX.Element => {
	const dispatch = useDispatch();

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	return (
		<Box>
			<BackgroundEditor />
			<ForegroundEditor />
		</Box>
	)
};
