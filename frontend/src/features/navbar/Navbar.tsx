import React from 'react';

import { Box, Button, Grid, Typography } from '@material-ui/core';
import { useHistory } from 'react-router';
import { logout, resetStatus, selectIsLoggedIn, selectRefreshToken } from '../auth/authSlice';
import { useDispatch, useSelector } from 'react-redux';
import { selectDisplayName } from '../user/userSlice';

interface NavbarProps {
	isHome?: boolean
}

export const Navbar = ({ 
	isHome = false 
}: NavbarProps): JSX.Element => {
	const history = useHistory();
	const dispatch = useDispatch();

	const isLoggedIn = useSelector(selectIsLoggedIn);
	const refreshToken = useSelector(selectRefreshToken);
	const displayName = useSelector(selectDisplayName)

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

	const heading = isLoggedIn ? "Welcome back, " + displayName : "Playing as Guest";

	const loginLogoutButton = isLoggedIn ? (
		<Button variant='outlined' size='small' onClick={onLogoutClick}>Logout</Button>
	) : (
		<Button variant='outlined' size='small' onClick={onLoginClick}>Login</Button>
	)

	return isHome ? (
		<Grid container>
			<Grid item xs={3}></Grid>
			<Grid item xs={6} justify='center'>
				<Box height='100%' display='flex' justifyContent='center' alignItems='center'>
					<Typography component='span' align="center">{heading}</Typography>
				</Box>
			</Grid>
			<Grid item xs={3} justify='flex-end'>
				<Box display='flex' justifyContent='flex-end' alignContent='center'>
					{loginLogoutButton}
				</Box>
			</Grid>
		</Grid>
	) : (
		<Grid container justify='space-between'>
			<Grid item xs={3}>
				<Box display='flex'>
					<Button variant='outlined' size='small' onClick={onHomeClick}>Home</Button>
				</Box>
			</Grid>
			<Grid item xs={6}>
				<Box height='100%' display='flex' justifyContent='center' alignItems='center'>
					<Typography component='span' align="center">{heading}</Typography>
				</Box>
			</Grid>
			<Grid item xs={3}>
			</Grid>
		</Grid>
	);
}
