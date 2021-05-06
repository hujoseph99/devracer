import { Box, Container } from '@material-ui/core';
import React from 'react';
import { Footer } from '../footer/Footer';
import { Navbar } from '../navbar/Navbar';
import { RegisterForm } from './RegisterForm';

export const RegisterPage = (): JSX.Element => {
	return (
		<Container maxWidth='sm'>
			<Box minHeight='100vh' display='flex' flexDirection='column' justifyContent='space-between' py={5}>
				<Navbar />
				<RegisterForm />
				<Footer />
			</Box>
		</Container>
	)
}
