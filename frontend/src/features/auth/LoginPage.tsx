import { Box, Container } from '@material-ui/core';
import React from 'react';
import { Footer } from '../footer/Footer';
import { Navbar } from '../navbar/Navbar';
import { LoginForm } from './LoginForm';

export const LoginPage = (): JSX.Element => {
	return (
		<Container maxWidth='sm'>
			<Box minHeight='100vh' display='flex' flexDirection='column' justifyContent='space-between' py={5}>
				<Navbar />
				<LoginForm />
				<Footer />
			</Box>
		</Container>
	)
}
