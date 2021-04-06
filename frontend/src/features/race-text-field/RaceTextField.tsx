import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux';

import { BackgroundEditor } from './BackgroundEditor';
import { fetchNewPracticeRace } from './raceTextFieldSlice';

export const RaceTextField = (): JSX.Element => {
	const dispatch = useDispatch();

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	return (
		<BackgroundEditor />
	)
};
