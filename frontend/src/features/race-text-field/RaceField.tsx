import React, { SyntheticEvent, useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux';

import { Box } from '@material-ui/core';

import { Ace, Range } from 'ace-builds';

import { BackgroundEditor } from './BackgroundEditor';
import { ForegroundEditor } from './ForegroundEditor';
import { fetchNewPracticeRace } from './raceFieldSlice';
import { selectSnippet } from './raceFieldSlice';


// const navkeys = ["ArrowDown",
// 	"ArrowLeft",
// 	"ArrowRight",
// 	"ArrowUp",
// 	"End",
// 	"Home",
// 	"PageDown",
// 	"PageUp"];


// function filterMouseEvents(e: SyntheticEvent) {
// 	e.stopPropagation()
// 	e.preventDefault()
// 	return true
// }

export const RaceField = (): JSX.Element => {
	const snippet = useSelector(selectSnippet);
	const dispatch = useDispatch();
	const [focus, setFocus] = useState(false);
	const [foregroundText, setForegroundText] = useState("");
	const [backgroundText, setBackgroundText] = useState("");
	const [markers, setMarkers] = useState<Ace.Range[]>([]);
	const [snippetArray, setSnippetArray] = useState<string[]>([]);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])

	useEffect(() => {
		setSnippetArray(snippet.raceContent.replace(/\t/g, ' '.repeat(4)).split('\n'));
		setBackgroundText(snippet.raceContent);
	}, [snippet])


	const onFocus = (e: SyntheticEvent) => {
		setFocus(true);
	}

	const onBlur = () => {
		setFocus(false);
	}

	const onChange = (text: string) => {
		if (text === backgroundText){
			// win
		}
		const textArray = text.split("\n");

		const backgroundArray: string[] = [];
		const markerArray: Ace.Range[] = [];

		let i = 0;
		for (; i < textArray.length; i++) {
			if (i > snippetArray.length - 1) {
				break
			}
			let userLine = textArray[i];
			let snippet_line = snippetArray[i];

			var difference_index = 0;
			while (userLine[difference_index] === snippet_line[difference_index] && difference_index < userLine.length) difference_index++;

			if (difference_index < userLine.length) {
				markerArray.push(new Range(i, difference_index, i, userLine.length));
			}
			if (difference_index <= snippet_line.length) {
				snippet_line = snippet_line.slice(0, difference_index) + ' '.repeat(userLine.length - difference_index) + snippet_line.slice(difference_index)
			}
			backgroundArray.push(snippet_line)
		}

		for (; i < snippetArray.length; i++) {
			let snippetLine = snippetArray[i];
			backgroundArray.push(snippetLine);
		}
		setBackgroundText(backgroundArray.join("\n"));
		setMarkers(markerArray);
		if (foregroundText !== text) setForegroundText(text);
	}

	return (
		<Box style={{ height: "1000px", width: "1000px" }}
			// onKeyDownCapture={filterKeyboardEvents}
			// onKeyPressCapture={filterKeyboardEvents}
			// onKeyUpCapture={filterKeyboardEvents}
			onClickCapture={onFocus}>
			{/* onMouseDownCapture={filterMouseEvents}
			onMouseMoveCapture={filterMouseEvents}
			onFocusCapture={filterMouseEvents}
			onChangeCapture={filterMouseEvents}
			onBlurCapture={filterMouseEvents}
			onMouseUpCapture={filterMouseEvents}> */}
			<BackgroundEditor text={backgroundText}/>
			{/* elements that appear later are on top */}
			<ForegroundEditor text={foregroundText} ranges={markers} focus={focus} onChange={onChange} onBlur={onBlur} />
		</Box>
	)
};
