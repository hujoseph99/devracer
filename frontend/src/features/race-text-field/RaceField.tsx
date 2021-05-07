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


	const onfocus = (e: SyntheticEvent) => {
		setFocus(true);
	}

	const onblur = () => {
		setFocus(false);
	}

	const onchange = (text: string) => {
		if (text === backgroundText){
			// win
		}
		const text_array = text.split("\n");

		const background_array: string[] = [];
		const marker_array: Ace.Range[] = [];

		let i = 0;
		for (; i < text_array.length; i++) {
			if (i > snippetArray.length - 1) {
				break
			}
			let user_line = text_array[i];
			let snippet_line = snippetArray[i];

			var difference_index = 0;
			while (user_line[difference_index] === snippet_line[difference_index] && difference_index < user_line.length) difference_index++;

			if (difference_index < user_line.length) {
				marker_array.push(new Range(i, difference_index, i, user_line.length));
			}
			if (difference_index <= snippet_line.length) {
				snippet_line = snippet_line.slice(0, difference_index) + ' '.repeat(user_line.length - difference_index) + snippet_line.slice(difference_index)
			}
			background_array.push(snippet_line)
		}

		for (; i < snippetArray.length; i++) {
			let snippet_line = snippetArray[i];
			background_array.push(snippet_line);
		}
		setBackgroundText(background_array.join("\n"));
		setMarkers(marker_array);
		if (foregroundText !== text) setForegroundText(text);
	}

	return (
		<Box style={{ height: "1000px", width: "1000px" }}
			// onKeyDownCapture={filterKeyboardEvents}
			// onKeyPressCapture={filterKeyboardEvents}
			// onKeyUpCapture={filterKeyboardEvents}
			onClickCapture={onfocus}>
			{/* onMouseDownCapture={filterMouseEvents}
			onMouseMoveCapture={filterMouseEvents}
			onFocusCapture={filterMouseEvents}
			onChangeCapture={filterMouseEvents}
			onBlurCapture={filterMouseEvents}
			onMouseUpCapture={filterMouseEvents}> */}
			<BackgroundEditor text={backgroundText}/>
			{/* elements that appear later are on top */}
			<ForegroundEditor text={foregroundText} ranges={markers} focus={focus} onchange={onchange} onblur={onblur} />
		</Box>
	)
};
