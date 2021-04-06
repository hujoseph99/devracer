import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux';

import { BackgroundEditor } from './BackgroundEditor';
import { fetchNewPracticeRace } from './raceFieldSlice';

export const RaceField = (): JSX.Element => {
	const dispatch = useDispatch();

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	return (
		<BackgroundEditor />
	)
};
