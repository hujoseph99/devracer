import React from 'react';
import { useHistory } from 'react-router';

import { Box, Button, Container, Grid, makeStyles, Theme, Typography } from '@material-ui/core';

import PersonIcon from '@material-ui/icons/Person';
import KeyboardIcon from '@material-ui/icons/Keyboard';
import { blurb } from './constants';
import { Navbar } from '../navbar/Navbar';
import { Footer } from '../footer/Footer';

const useStyles = makeStyles<Theme>(theme => ({
	titleIcon: {
		fontSize: theme.typography.h2.fontSize,
		marginRight: 10
	}
}));

export const MainMenu = (): JSX.Element => {
	const history = useHistory();
	const classes = useStyles();

	const onPracticeClick = () => {
		history.push('/practice');
	}

	const onCustomClick = () => {
		history.push('/custom');
	}

	return (
		<Container maxWidth='sm'>
			<Box minHeight='100vh' display='flex' flexDirection='column' justifyContent='space-between' py={5}>
				<Navbar isHome />
				<Grid container>
					<Grid item xs={12}>
						<Box display='flex' width='100%' justifyContent='center' alignItems='center' mb={2}>
							<KeyboardIcon className={classes.titleIcon} />
							<Typography component='span' variant='h3'> 
								DevRacer
							</Typography>
						</Box>
					</Grid>
					<Grid container item xs={12}>
						<Typography variant='body1' align='center'>
							{blurb}
						</Typography>
					</Grid>
					<Grid item xs={12} spacing={3}>
						<Box display='flex' width='100%' justifyContent='center' mt={8}>
							<Button variant='contained' size='large' endIcon={<PersonIcon />} onClick={onCustomClick}>Start a Lobby</Button>
						</Box>
					</Grid>
				</Grid>
				<Footer />
			</Box>
		</Container>
	);
}
