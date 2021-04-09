import React from 'react';

import AceEditorComponent, { IAceEditorProps } from "react-ace";

import "./editor.css"

// all currently supported languages
import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/mode-plain_text";

// currently supported themes
import "ace-builds/src-noconflict/theme-dracula";

// just a simple implementation for now
interface AceEditorProps extends IAceEditorProps {
	value?: string
}

export const AceEditor = ({ value = "", ...props }: AceEditorProps): JSX.Element => {
	return (
		<AceEditorComponent 
			className="aceEditor"
			value={value}
			theme="dracula"
			showGutter={false}
			highlightActiveLine={false}
			{...props}
		/>
	);
}
