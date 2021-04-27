import React from 'react';
import { useHistory } from 'react-router';
import Typist from 'react-typist';

import { Box, Button, makeStyles, Typography } from '@material-ui/core';
import PersonIcon from '@material-ui/icons/Person';
import KeyboardIcon from '@material-ui/icons/Keyboard';

const useStyles = makeStyles({
	titleIcon: {
		fontSize: "4em",
		marginRight: 10
	},
	test: {
		color: 'white'
	}
});

export const MainMenu = (): JSX.Element => {
	const history = useHistory();
	const classes = useStyles();

	const onPracticeClick = () => {
		history.push('/practice');
	}

	return (
		<Box display='flex' flexDirection='column' justifyContent='center' alignItems='center'>
			<Typist>
				<Box mt={10} display='flex' justifyContent='center' alignContent='center'>
					<KeyboardIcon className={classes.titleIcon} color='primary' />
					<Typist.Delay ms={1000} />
					<Typography component='span' variant='h3' color='primary'> 
						CodeRacers
					</Typography>
				</Box>
			</Typist>
			<Box width='60%' mt={3}>
				<Typography color='primary'>
					CodeRacers is an online typing game inspired by TypeRacer. Type the most popular snippets from supported languages and increase your coding speed!
				</Typography>
			</Box>
			<Box mt={5}>
				<Button variant='contained' endIcon={<PersonIcon />} onClick={onPracticeClick}>
					Practice
				</Button>
			</Box>
		</Box>
	);
}
