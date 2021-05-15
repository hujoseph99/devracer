import React from 'react'
import { Ace } from 'ace-builds';
import { AceEditor } from './AceEditor';

interface BackgroundEditorProps {
	text?: string;
}

export const BackgroundEditor = ({
	text = "",
}: BackgroundEditorProps): JSX.Element => {

	const handleLoad = (editor: Ace.Editor) => {
		// settings to make the text editor look like it's disabled, some of the things are
		// not yet supported in typescript so have to ignore them
		// editor.getSession().selection.on('changeSelection', function ()
		// {
		// 	editor.getSession().selection.clearSelection();
		// });
		// @ts-ignore
		editor.commands.commmandKeyBinding={}
		// @ts-ignore
		editor.textInput.getElement().disabled=true
		// @ts-ignore
		editor.renderer.$cursorLayer.element.style.display = "none";
	}

	return (
		<AceEditor 
			mode='plain_text'
			className='backgroundEditor'
			value={text}
			readOnly={true}
			onLoad={handleLoad}
		/>
	);
}
