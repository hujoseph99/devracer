import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux';

import { AceEditor } from './AceEditor';
import { fetchNewPracticeRace, selectSnippet } from './raceTextFieldSlice';

export const RaceTextField = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	return (
		<AceEditor
			mode={snippet.language}
			value={snippet.raceContent}
		/>
	)
};
