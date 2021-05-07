import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { Avatar, Box, Button, Checkbox, FormControlLabel, Grid, IconButton, InputAdornment, Link, makeStyles, Paper, Theme, Typography } from '@material-ui/core';
import { LockOutlined, Visibility, VisibilityOff } from '@material-ui/icons';
import { red } from '@material-ui/core/colors';

import { FormTextField } from './FormTextField';
import { login, resetStatus, selectStatus } from './authSlice';
import { useHistory } from 'react-router';

const useStyles = makeStyles<Theme>(theme => ({
	avatar: {
		margin: theme.spacing(1),
		backgroundColor: red[400]
	},
	submit: {
		margin: theme.spacing(3, 0, 2)
	}
}));

interface FormState {
	username: string;
	password: string;
	rememberMe: boolean;
}

export const LoginForm = (): JSX.Element => {
	const history = useHistory();
	const dispatch = useDispatch();
	const classes = useStyles();
	const [formState, setFormState] = useState<FormState>({
		username: '',
		password: '',
		rememberMe: false,
	});
	const [showPassword, setShowPassword] = useState(false);
	const status = useSelector(selectStatus);

	// reset it on load too
	useEffect(() => {
		dispatch(resetStatus());
		return () => {
			dispatch(resetStatus());
		}
	}, [dispatch]);

	// TODO: BUG WITH THIS - SOMETIMES IT WILL COME IN AS SUCCEEDED (NAMELY AFTER LOGGING OUT)
	useEffect(() => {
		if (status === 'succeeded') {
			history.push('/');
			dispatch(resetStatus());
		} 
	}, [dispatch, history, status])

	const handleClickShowPassword = () => {
		setShowPassword(prev => !prev);
	}

	const handleClickRememberMe = () => {
		setFormState(prev => ({ ...prev, rememberMe: !prev.rememberMe}));
	}

	const handleChange = (key: keyof FormState) => (event: React.ChangeEvent<HTMLInputElement>) => {
		setFormState(prev => ({ ...prev, [key]: event.target.value }));
	}

	const handleSubmit = () => {
		dispatch(login(formState));
	}

	const handleRegister = () => {
		history.push('/register');
	}
	
	return (
		<Grid container justify='center'>
			<Grid item xs={12} sm={10}>
				<Paper>
					<Box py={5} mt={4}>  { /* because footer also has mt={4} */} 
						<Grid container justify='center'>
							<Grid item xs={10} sm={8}>
								<Box display='flex' flexDirection='column' alignItems='center' width='100%'>
									<Avatar className={classes.avatar}>
										<LockOutlined />
									</Avatar>
									<Typography variant="h4" align='center' gutterBottom>Sign in</Typography>
									<Box width='100%' mt={2}>
										<FormTextField 
											autoFocus 
											label='Username' 
											value={formState.username} 
											error={status === 'failed'}
											onChange={handleChange('username')}
										/>
										<FormTextField 
											label='Password' 
											value={formState.password}
											type={showPassword ? 'text' : 'password'}
											error={status === 'failed'}
											onChange={handleChange('password')}
											helperText={status === 'failed' ? 'The username and password were incorrect' : ''}
											InputProps={{
												endAdornment: (
													<InputAdornment position='end'>
														<IconButton
															onClick={handleClickShowPassword}
															disableFocusRipple
															disableRipple
															disableTouchRipple
														>
															{showPassword ? <VisibilityOff /> : <Visibility />}
														</IconButton>
													</InputAdornment>
												)
											}
										}/>
										<FormControlLabel
											control={(
												<Checkbox 
													checked={formState.rememberMe} 
													onClick={handleClickRememberMe} 
													color="primary" 
												/>
											)}
											label="Remember me"
										/>
										<Button
											fullWidth
											variant='contained'
											className={classes.submit}
											onClick={handleSubmit}
										>
											Sign In
										</Button>
										<Box display='flex' flexDirection='row-reverse'>
											<Link variant="body2" onClick={handleRegister}>
												Don't have an account? Sign Up
											</Link>
										</Box>
									</Box>
								</Box>
							</Grid>
						</Grid>
					</Box>
				</Paper>
			</Grid>
		</Grid>
	)
}
