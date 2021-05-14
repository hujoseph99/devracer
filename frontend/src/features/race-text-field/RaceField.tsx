import React, { SyntheticEvent, useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux';

import { Box } from '@material-ui/core';

import { Ace, Range } from 'ace-builds';

import { BackgroundEditor } from './BackgroundEditor';
import { ForegroundEditor } from './ForegroundEditor';
import { fetchNewPracticeRace } from './raceFieldSlice';
import { selectSnippet } from './raceFieldSlice';
import { RaceSnippet } from './types';
import { language } from '../game/types';


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

interface RaceFieldProps {
	snippet?: string;
	language?: language;
	disabled?: boolean
}

export const RaceField = ({ 
	snippet = '', 
	language = 'plain_text' ,
	disabled = false,
}: RaceFieldProps): JSX.Element => {
	const [focus, setFocus] = useState(false);
	const [foregroundText, setForegroundText] = useState("");
	const [backgroundText, setBackgroundText] = useState("");
	const [markers, setMarkers] = useState<Ace.Range[]>([]);
	const [snippetArray, setSnippetArray] = useState<string[]>([]);

	useEffect(() => {
		setSnippetArray(snippet.replace(/\t/g, ' '.repeat(4)).split('\n'));
		setBackgroundText(snippet);
	}, [snippet])

	const onFocus = (e: SyntheticEvent) => {
		setFocus(true);
	}

	const onBlur = () => {
		setFocus(false);
	}

	const onChange = (playerText: string) => {
		if (playerText === backgroundText){
			// win
		}
		const playerTextArray = playerText.split("\n");

		const backgroundArray: string[] = [];
		const newMarkers: Ace.Range[] = [];

		let i = 0;
		// compare text
		for (; i < playerTextArray.length; i++) {
			let playerLine = playerTextArray[i];

			// if no line in snippet to compare with, the rest of text is wrong. mark wrong for each line
			if (i > snippetArray.length - 1) {
				for (let j = i; j < playerTextArray.length; j++) {
					newMarkers.push(new Range(j, 0, j, playerLine.length));
				}
				break
			}

			// compare with line in snippet
			let snippetLine = snippetArray[i];
			let differenceIndex = 0;
			while (playerLine[differenceIndex] === snippetLine[differenceIndex] && differenceIndex < playerLine.length) differenceIndex++;

			if (differenceIndex < playerLine.length) {
				newMarkers.push(new Range(i, differenceIndex, i, playerLine.length));
			}
			if (differenceIndex <= snippetLine.length) {
				snippetLine = snippetLine.slice(0, differenceIndex) + ' '.repeat(playerLine.length - differenceIndex) + snippetLine.slice(differenceIndex)
			}
			backgroundArray.push(snippetLine)
		}

		// if player_text was shorter than snippet, push rest of snippet into the background.
		for (; i < snippetArray.length; i++) {
			backgroundArray.push(snippetArray[i]);
		}

		setBackgroundText(backgroundArray.join("\n"));
		setMarkers(newMarkers);
		if (foregroundText !== playerText) setForegroundText(playerText);
	}

	return (
		<Box
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
			<BackgroundEditor text={backgroundText} />
			{/* elements that appear later are on top */}
			<ForegroundEditor 
				language={language} 
				text={foregroundText} 
				ranges={markers} 
				focus={focus} 
				onChange={onChange} 
				onBlur={onBlur} 
				disabled={disabled}
			/>
		</Box>
	)
};
