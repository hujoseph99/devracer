import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import { Ace } from 'ace-builds';

import { AceEditor } from './AceEditor';
import { selectSnippet } from './raceFieldSlice';

import "./editor.css"

interface ForegroundEditorProps {
	focus?: boolean;
	onblur: () => void;
}

export const ForegroundEditor = ({
	focus = false,
	onblur = () => {}
}: ForegroundEditorProps): JSX.Element => {
	const snippet = useSelector(selectSnippet);
	const [editor, setEditor] = useState<Ace.Editor | undefined>(undefined);
	
	useEffect(() => {
		if (editor && focus) {
			editor.focus();
		}
	}, [editor, focus])

	const handleLoad = (editor_ref: Ace.Editor) => {
		setEditor(editor_ref)
	}

	return (
		<AceEditor  
			className="foregroundEditor"
			mode={snippet.language}
			value=""
			onLoad={handleLoad}
			onBlur={onblur}
		/>
	);
}
