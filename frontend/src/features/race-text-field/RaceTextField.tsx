import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux';
// import { MonacoEditor } from './MonacoEditor';
import { fetchNewPracticeRace, selectSnippet } from './raceTextFieldSlice';

export const RaceTextField = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	return (
		// <MonacoEditor value={snippet.raceContent} />
		<div></div>
	)
};
