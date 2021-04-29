import React from 'react';

import { Box, Button, Grid, makeStyles } from '@material-ui/core';

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

	return isHome ? (
		<Grid container item xs={12} justify='flex-end' className={classes.navbar}>
			<Grid item>
				<Button variant='outlined' size='small' color='primary'>Login</Button>
			</Grid>
		</Grid>
	) : (
		<Grid container item xs={12} justify='space-between' className={classes.navbar}>
			<Grid item>
				<Button variant='outlined' size='small' color='primary'>Home</Button>
			</Grid>
			<Grid item>
				<Button variant='outlined' size='small' color='primary'>Login</Button>
			</Grid>
		</Grid>
	);
}
