import React from 'react';
import { useSelector } from 'react-redux';

import { Ace } from 'ace-builds';

import { AceEditor } from './AceEditor';
import { selectSnippet } from './raceFieldSlice';

import "./editor.css"

export const ForegroundEditor = (): JSX.Element => {
	const snippet = useSelector(selectSnippet);

	return (
		<AceEditor  
			className="foregroundEditor"
			mode={snippet.language}
			value={snippet.raceContent}
		/>
	);
}
