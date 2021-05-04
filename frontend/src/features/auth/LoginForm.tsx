import { Avatar, Box, Button, Checkbox, Container, Divider, FormControl, FormControlLabel, Grid, IconButton, InputAdornment, Link, makeStyles, Paper, TextField, TextFieldProps, Theme, Typography } from '@material-ui/core';
import { cyan, red } from '@material-ui/core/colors';
import { LockOutlined, PersonalVideo, Visibility, VisibilityOff } from '@material-ui/icons';
import React, { useState } from 'react';
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
		backgroundColor: red[400]
	},
	submit: {
		margin: theme.spacing(3, 0, 2)
	}
}));

interface FormState {
	username: string;
	password: string;
}

export const LoginForm = (): JSX.Element => {
	const classes = useStyles();
	const [formState, setFormState] = useState<FormState>({
		username: '',
		password: ''
	});
	const [showPassword, setShowPassword] = useState(false)

	const handleClickShowPassword = () => {
		setShowPassword(prev => !prev);
	}

	const handleChange = (key: keyof FormState) => (event: React.ChangeEvent<HTMLInputElement>) => {
		setFormState(prev => ({ ...prev, [key]: event.target.value }))
	}

	return (
		<Grid container justify='center'>
			<Grid item xs={10} sm={8}>
				<Box display='flex' flexDirection='column' alignItems='center' width='100%'>
					<Avatar className={classes.avatar}>
						<LockOutlined />
					</Avatar>
					<Typography variant="h4" align='center'>Sign in</Typography>
					<Box width='100%'>
						<FormTextField 
							autoFocus 
							label='Username' 
							value={formState.username} 
							onChange={handleChange('username')}
						/>
						<FormTextField 
							label='Password' 
							type={showPassword ? 'text' : 'password'}
							value={formState.password}
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
	)
}
