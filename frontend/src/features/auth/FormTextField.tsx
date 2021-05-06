import React from 'react';

import { Box, TextField, TextFieldProps } from '@material-ui/core';

export const FormTextField = ({
	...props
}: TextFieldProps): JSX.Element => (
	<Box width='100%'>
		<TextField 
			variant='outlined' 
			margin='normal'
			fullWidth
			{...props}
		/>
	</Box>
)
