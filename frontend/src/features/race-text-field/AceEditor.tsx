import React from 'react';

import AceEditorComponent, { IEditorProps } from "react-ace";

// all currently supported languages
import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/mode-plain_text";

// just a simple implementation for now
interface AceEditorProps extends IEditorProps {
	value?: string
}

export const AceEditor = ({ value = "", ...props }: AceEditorProps): JSX.Element => {
	return (
		<AceEditorComponent 
			value={value}
			{...props}
		/>
	);
}
