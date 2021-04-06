import React from 'react';
import { useSelector } from 'react-redux';

import { AceEditor } from './AceEditor';
import { selectSnippet } from './raceFieldSlice';

export const BackgroundEditor = (): JSX.Element => {
	const snippet = useSelector(selectSnippet);

	return (
		<AceEditor 
			mode={snippet.language}
			value={snippet.raceContent}
		/>
	);
}
