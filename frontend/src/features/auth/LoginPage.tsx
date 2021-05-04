import { Box, Container, Grid, makeStyles } from '@material-ui/core';
import React from 'react';
import { Footer } from '../footer/Footer';
import { Navbar } from '../navbar/Navbar';
import { LoginForm } from './LoginForm';

const useStyles = makeStyles({
	container: {
		minHeight: '100vh'
	}
})

export const LoginPage = (): JSX.Element => {
	const classes = useStyles();

	return (
		// <Container maxWidth='sm' className={classes.container}>
		<Container maxWidth='sm'>
			<Box minHeight='100vh' display='flex' flexDirection='column' justifyContent='space-between' py={5}>
				<Navbar />
				<LoginForm />
				<Footer />
			</Box>
		</Container>
	)
}
