import React from 'react';
import { useHistory } from 'react-router';
import Typist from 'react-typist';

import { Box, Button, Container, Grid, makeStyles, Theme, Typography } from '@material-ui/core';
import PersonIcon from '@material-ui/icons/Person';
import KeyboardIcon from '@material-ui/icons/Keyboard';
import { blurb } from './constants';
import { Navbar } from '../navbar/Navbar';

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

	return (
		<Container maxWidth='sm'>
			<Grid container>
				<Navbar isHome={true} />
				<Grid item xs={12}>
					<Typist>
						<Box mt='5vh' display='flex' width='100%' justifyContent='center' alignItems='center'>
							<KeyboardIcon className={classes.titleIcon} color='primary' />
							<Typist.Delay ms={1000} />
							<Typography component='span' variant='h3' color='primary'> 
								CodeRacers
							</Typography>
						</Box>
					</Typist>
				</Grid>
				<Grid container item xs={12}>
					<Typography variant='body1' color='primary' align='center'>
							{blurb}
					</Typography>
				</Grid>
				<Grid item xs={12}>
					<Box mt='5vh' display='flex' width='100%' justifyContent='center' alignItems='center'>
						<Typography component='span' variant='h4' color='primary'> 
							Modes
						</Typography>
					</Box>
				</Grid>
				<Grid item xs={12} spacing={3}>
					<Box display='flex' width='100%' justifyContent='center' mt="3vh">
						<Button variant='contained' size='large' endIcon={<PersonIcon />} onClick={onPracticeClick}>Practice</Button>
					</Box>
				</Grid>
			</Grid>
			<Box display='flex' width='100%' mt="5vh" justifyContent='center'>
				<Typography color='primary' align='center'>Made with &#10084;&#65039; by Joseph and David</Typography>
			</Box>
		</Container>
	);
}
