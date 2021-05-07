import React, { useEffect, useRef } from 'react';
import { useSelector } from 'react-redux';

import { Ace } from 'ace-builds';

import { AceEditor } from './AceEditor';
import { selectSnippet } from './raceFieldSlice';

import "./editor.css"

interface ForegroundEditorProps {
	focus?: boolean;
	onblur?: () => void;
	ranges: Ace.Range[];
	text: string;
	onchange: (s: string) => void;
}

export const ForegroundEditor = ({
	focus = false,
	onblur = () => {},
	ranges,
	onchange,
	text
}: ForegroundEditorProps): JSX.Element => {
	const snippet = useSelector(selectSnippet);
	const editor = useRef<Ace.Editor>();
	const markers = useRef<number[]>([]);
	
	useEffect(() => {
		if (editor.current !== undefined){
			for (let marker of markers.current){
				editor.current.session.removeMarker(marker);
			}

			let new_markers: number[] = [];
			for (let marker of ranges){
				new_markers.push(editor.current.session.addMarker(marker, "error-marker", "text"));
			}
			markers.current = new_markers;
		}
	}, [ranges])

	useEffect(() => {
		if (editor && focus) {
			editor.current?.focus();
		}
	}, [editor, focus])

	const handleLoad = (editor_ref: Ace.Editor) => {
		editor.current = editor_ref;
	}

	return (
		<AceEditor  
			className="foregroundEditor"
			mode={snippet.language}
			onLoad={handleLoad}
			onBlur={onblur}
			onChange={onchange}
			value={text}
		/>
	);
}
