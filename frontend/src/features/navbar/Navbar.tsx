import React from 'react';

import { Button, Grid, makeStyles } from '@material-ui/core';
import { useHistory } from 'react-router';

const useStyles = makeStyles({
	navbar: {
		marginTop: '1vh',
	}
});

interface NavbarProps {
	isHome?: boolean
}

export const Navbar = ({ 
	isHome = false 
}: NavbarProps): JSX.Element => {
	const classes = useStyles();
	const history = useHistory();

	const onLoginClick = () => {
		history.push('/login');
	}

	const onHomeClick = () => {
		history.push('/');
	}

	return isHome ? (
		<Grid container justify='flex-end' className={classes.navbar}>
			<Grid item>
				<Button variant='outlined' size='small' onClick={onLoginClick}>Login</Button>
			</Grid>
		</Grid>
	) : (
		<Grid container justify='space-between' className={classes.navbar}>
			<Grid item>
				<Button variant='outlined' size='small' onClick={onHomeClick}>Home</Button>
			</Grid>
			<Grid item>
				<Button variant='outlined' size='small' onClick={onLoginClick}>Login</Button>
			</Grid>
		</Grid>
	);
}
