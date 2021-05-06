import React, { KeyboardEvent, SyntheticEvent, useEffect, useState } from 'react'
import { useDispatch } from 'react-redux';

import { Box } from '@material-ui/core';

import { BackgroundEditor } from './BackgroundEditor';
import { ForegroundEditor } from './ForegroundEditor';
import { fetchNewPracticeRace } from './raceFieldSlice';


const navkeys = ["ArrowDown",
	"ArrowLeft",
	"ArrowRight",
	"ArrowUp",
	"End",
	"Home",
	"PageDown",
	"PageUp"];


function filterMouseEvents(e: SyntheticEvent) {
	e.stopPropagation()
	e.preventDefault()
	return true
}

function filterKeyboardEvents(e: KeyboardEvent) {
	console.log(e.key)

	if (navkeys.includes(e.key)){
		e.stopPropagation()
		e.preventDefault()
		return true
	}
}

export const RaceField = (): JSX.Element => {
	const dispatch = useDispatch();
	const [focus, setFocus] = useState(false);

	useEffect(() => {
		dispatch(fetchNewPracticeRace());
	}, [dispatch])


	const onfocus = (e: SyntheticEvent) => {
		setFocus(true);
	}

	const onblur = () => {
		setFocus(false);
	}

	return (
		<Box style={{ height: "1000px", width: "1000px" }}
			onKeyDownCapture={filterKeyboardEvents}
			onKeyPressCapture={filterKeyboardEvents}
			onKeyUpCapture={filterKeyboardEvents}
			onClickCapture={onfocus}
			onMouseDownCapture={filterMouseEvents}
			onMouseMoveCapture={filterMouseEvents}
			onFocusCapture={filterMouseEvents}
			onChangeCapture={filterMouseEvents}
			onBlurCapture={filterMouseEvents}
			onMouseUpCapture={filterMouseEvents}>
			<BackgroundEditor />
			{/* elements that appear later are on top */}
			<ForegroundEditor focus={focus} onblur={onblur} />
		</Box>
	)
};
