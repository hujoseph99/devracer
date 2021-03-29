import { Paper, Typography } from '@material-ui/core';
import React, { useEffect } from 'react'
import MonacoEditor from 'react-monaco-editor';
import { useDispatch, useSelector } from 'react-redux';
import { fetchNewPracticeRace, selectSnippet } from './raceTextFieldSlice';



export const RaceTextField = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	// return (
	// 	<Paper variant='outlined'>
	// 		<Typography color='secondary'>{snippet.raceContent}</Typography>
	// 	</Paper>
	// );
	return (
		<MonacoEditor 
			width="800"
			height="600"
			language="javascript"
			value={snippet.raceContent}
			theme="vs-dark"
			options={{
				minimap: {
					enabled: false
				},
				scrollbar: {
					vertical: 'hidden',
					verticalHasArrows: false
				},
				overviewRulerLanes: 0,
				hideCursorInOverviewRuler: true,
				overviewRulerBorder: false
			}}
		/>
	)
};
