import React from 'react';
import ReactMonacoEditor, { 
	EditorConstructionOptions, 
	MonacoEditorProps 
} from 'react-monaco-editor';

interface MonacoEditorTypes extends MonacoEditorProps {
	options?: EditorConstructionOptions;
	value?: string;
}

export const MonacoEditor = ({ 
	options =  {},
	value = "",
	...rest 
}: MonacoEditorTypes): JSX.Element => {
	return (
		<ReactMonacoEditor
			width="800"
			height="600"
			theme="vs-dark"
			language="javascript"
			value={value}
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
				overviewRulerBorder: false,
				...options
			}}
			{ ...rest }
		/>
	)
};
