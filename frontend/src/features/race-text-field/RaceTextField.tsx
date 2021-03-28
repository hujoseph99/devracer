import { Paper, Typography } from '@material-ui/core';
import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux';
import { fetchNewPracticeRace, selectSnippet } from './raceTextFieldSlice';

export const RaceTextField = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	return (
		<Paper variant='outlined'>
			<Typography color='secondary'>{snippet.raceContent}</Typography>
		</Paper>
	);
};
