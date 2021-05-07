import React, { useEffect } from 'react';

import { Button, Grid, makeStyles } from '@material-ui/core';
import { useHistory } from 'react-router';
import { logout, resetStatus, selectIsLoggedIn, selectRefreshToken } from '../auth/authSlice';
import { useDispatch, useSelector } from 'react-redux';

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
	const dispatch = useDispatch();
	const isLoggedIn = useSelector(selectIsLoggedIn);
	const refreshToken = useSelector(selectRefreshToken);

	const onLoginClick = () => {
		history.push('/login');
	}
	
	const onLogoutClick = () => {
		dispatch(logout({ refreshToken }));
		dispatch(resetStatus);
		history.push('/');
	};

	const onHomeClick = () => {
		history.push('/');
	}

	const loginLogoutButton = isLoggedIn ? (
		<Button variant='outlined' size='small' onClick={onLogoutClick}>Logout</Button>
	) : (
		<Button variant='outlined' size='small' onClick={onLoginClick}>Login</Button>
	)

	return isHome ? (
		<Grid container justify='flex-end' className={classes.navbar}>
			<Grid item>
				{loginLogoutButton}
			</Grid>
		</Grid>
	) : (
		<Grid container justify='space-between' className={classes.navbar}>
			<Grid item>
				<Button variant='outlined' size='small' onClick={onHomeClick}>Home</Button>
			</Grid>
			<Grid item>
				{loginLogoutButton}
			</Grid>
		</Grid>
	);
}
