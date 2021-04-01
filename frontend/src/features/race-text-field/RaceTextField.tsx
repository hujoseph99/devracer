import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux';
import { MonacoEditor } from './MonacoEditor';
import { fetchNewPracticeRace, selectSnippet } from './raceTextFieldSlice';

export const RaceTextField = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	// TODO: Figure out how to stop text highlighting on load ... it's because
	// 	the text initially starts out as "" because have no gotten it back from backend
	// Potential fix: Only mount the component when we have a value to put into it
	return (
		<MonacoEditor value={snippet.raceContent} />
	)
};
