import React, { createRef, useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { AceEditor } from './AceEditor';
import { selectSnippet } from './raceFieldSlice';

export const BackgroundEditor = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);

	return (
		<AceEditor 
			mode='plain_text'
			value={snippet.raceContent}
			readOnly={true}
			highlightActiveLine={false}
			ref={aceEditor}
		/>
	);
}
