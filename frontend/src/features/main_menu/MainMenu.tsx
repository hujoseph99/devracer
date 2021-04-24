import React from 'react';
import { useHistory } from 'react-router';

import { Box, Button, Typography } from '@material-ui/core';
import PersonIcon from '@material-ui/icons/Person';

export const MainMenu = (): JSX.Element => {
	const history = useHistory();

	const onPracticeClick = () => {
		history.push('/practice');
	}

	return (
		<Box justifyContent='center'>
			<Typography variant='h3' color='primary'>Code Racers</Typography>
			<Button variant='contained' endIcon={<PersonIcon />} onClick={onPracticeClick}>
				Practice
			</Button>
		</Box>
	);
}
