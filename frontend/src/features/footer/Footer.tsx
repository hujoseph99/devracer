import { Box, Typography } from '@material-ui/core';
import React from 'react';

export const Footer = (): JSX.Element => {
	return (
		<Box display='flex' width='100%' mt={4} justifyContent='center'>
			<Typography align='center'>Made with &#10084;&#65039; by Joseph and David</Typography>
		</Box>
	);
}
