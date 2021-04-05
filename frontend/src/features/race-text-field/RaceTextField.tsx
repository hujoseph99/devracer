import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux';

import AceEditor from "react-ace";
import "ace-builds/src-noconflict/mode-java";
import "ace-builds/src-noconflict/theme-monokai";

import { fetchNewPracticeRace, selectSnippet } from './raceTextFieldSlice';

export const RaceTextField = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	return (
		// <MonacoEditor value={snippet.raceContent} />
		<AceEditor
			mode="javascript"
			theme="monokai"
			value={snippet.raceContent}
		/>
	)
};
