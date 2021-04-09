import { Ace } from 'ace-builds';
import React, { createRef, useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { AceEditor } from './AceEditor';
import { selectSnippet } from './raceFieldSlice';

export const BackgroundEditor = (): JSX.Element => {
	const dispatch = useDispatch();
	const snippet = useSelector(selectSnippet);
	let editor: Ace.Editor | null = null

	const handleLoad = (aceEditor: Ace.Editor) => {
		editor = aceEditor;

		// settings to make the text editor look like it's disabled, some of the things are
		// not yet supported in typescript so have to ignore them
		aceEditor.getSession().selection.on('changeSelection', function ()
		{
			aceEditor.getSession().selection.clearSelection();
		});
		// @ts-ignore
		editor.commands.commmandKeyBinding={}
		// @ts-ignore
		editor.textInput.getElement().disabled=true
		// @ts-ignore
		aceEditor.renderer.$cursorLayer.element.style.display = "none";
	}

	return (
		<AceEditor 
			mode='plain_text'
			value={snippet.raceContent}
			highlightActiveLine={false}
			readOnly={true}
			onLoad={handleLoad}
		/>
	);
}
