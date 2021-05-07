import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { Avatar, Box, Button, Grid, IconButton, InputAdornment, Link, makeStyles, Paper, Theme, Typography } from '@material-ui/core';
import { LockOutlined, Visibility, VisibilityOff } from '@material-ui/icons';
import { red } from '@material-ui/core/colors';

import { FormTextField } from './FormTextField';
import { register, resetStatus, selectStatus } from './authSlice';
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
	email: string;
	nickname: string;
}

export const RegisterForm = (): JSX.Element => {
	const history = useHistory();
	const dispatch = useDispatch();
	const classes = useStyles();
	const [formState, setFormState] = useState<FormState>({
		username: '',
		password: '',
		email: '',
		nickname: '',
	});
	const [showPassword, setShowPassword] = useState(false);
	const status = useSelector(selectStatus);

	useEffect(() => {
		dispatch(resetStatus());
		return () => {
			dispatch(resetStatus());
		}
	}, [dispatch]);

	useEffect(() => {
		if (status === 'succeeded')  {
			dispatch(resetStatus());
			history.push('/');
		} 
	}, [dispatch, history, status]);

	const handleClickShowPassword = () => {
		setShowPassword(prev => !prev);
	}

	const handleChange = (key: keyof FormState) => (event: React.ChangeEvent<HTMLInputElement>) => {
		setFormState(prev => ({ ...prev, [key]: event.target.value }));
	}

	const handleSubmit = () => {
		dispatch(register(formState));
	}

	const handleLogin = () => {
		dispatch(resetStatus());
		history.push('/login');
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
									<Typography variant="h4" align='center' gutterBottom>Register</Typography>
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
										<FormTextField 
											label='Email' 
											value={formState.email} 
											error={status === 'failed'}
											onChange={handleChange('email')}
										/>
										<FormTextField 
											label='Nickname' 
											value={formState.nickname} 
											error={status === 'failed'}
											onChange={handleChange('nickname')}
											helperText={status === 'failed' ? 'The username or email has already been taken.' : ''}
										/>
										<Button
											fullWidth
											variant='contained'
											className={classes.submit}
											onClick={handleSubmit}
										>
											Register
										</Button>
										<Box display='flex' flexDirection='row-reverse'>
											<Link variant="body2" onClick={handleLogin}>
												Already have an account? Log In
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
