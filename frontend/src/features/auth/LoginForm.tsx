import { Avatar, Box, Button, Checkbox, Container, Divider, FormControl, FormControlLabel, Grid, Link, makeStyles, Paper, TextField, TextFieldProps, Theme, Typography } from '@material-ui/core';
import { cyan } from '@material-ui/core/colors';
import { LockOutlined } from '@material-ui/icons';
import React from 'react';
import { theme } from '../../theme';

const FormTextField = ({
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

const useStyles = makeStyles<Theme>(theme => ({
	avatar: {
		margin: theme.spacing(1),
		backgroundColor: theme.palette.secondary.main
	},
	submit: {
		margin: theme.spacing(3, 0, 2)
	}
}));

export const LoginForm = (): JSX.Element => {
	const classes = useStyles();

	return (
		<Grid container justify='center'>
			<Grid item xs={10} sm={8}>
				<Box display='flex' flexDirection='column' alignItems='center' width='100%'>
					<Avatar className={classes.avatar}>
						<LockOutlined />
					</Avatar>
					<Typography variant="h4" align='center'>Sign in</Typography>
					<Box width='100%'>
						<FormTextField autoFocus label='Username' />
						<FormTextField label='Password' />
						<FormControlLabel
							control={<Checkbox value="remember" color="primary" />}
							label="Remember me"
						/>
						<Button
							type='submit'
							fullWidth
							variant='contained'
							className={classes.submit}
						>
							Sign In
						</Button>
						<Box display='flex' flexDirection='row-reverse'>
							<Link href="#" variant="body2">
								Don't have an account? Sign Up
							</Link>
						</Box>
					</Box>
				</Box>
			</Grid>
		</Grid>
		// <Box mt={5}>
		// 	<Grid container justify='center'>
		// 		<Grid item xs={12} sm={10}>
		// 			<Paper >
		// 				<Grid container justify='center' spacing={4}>
		// 					<Grid item xs={8}>
		// 						<Box mt={4}>
		// 							<Typography variant='h4' align='center' gutterBottom>Login</Typography>
		// 						</Box>
		// 					</Grid>
		// 					<Grid item xs={10} sm={8}>
		// 						<FormTextField label='Email' />
		// 						<FormTextField label='Username' />
		// 						<FormTextField label='Password' />
		// 						<FormTextField label='Confirm Password' />
		// 						<FormTextField label='Nickname' />
		// 					</Grid>
		// 				</Grid>
		// 			</Paper>
		// 		</Grid>
		// 	</Grid>
		// </Box>
	)
}
