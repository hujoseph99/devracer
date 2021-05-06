import React, { useEffect, useState } from 'react';
import {
	Route,
	BrowserRouter as Router,
	Switch,
  Redirect
} from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';

import { CssBaseline, MuiThemeProvider } from '@material-ui/core';

import { theme } from './theme';
import { MainMenu } from './features/main-menu/MainMenu';
import { RaceField } from './features/race-text-field/RaceField';
import { LoginPage } from './features/auth/LoginPage';
import { RegisterPage } from './features/auth/RegisterPage';
import { refresh, selectRefreshToken } from './features/auth/authSlice';

const App = (): JSX.Element => {
  const dispatch = useDispatch();
  const refreshToken = useSelector(selectRefreshToken);
  const [firstTime, setFirstTime] = useState(true);

  // check if logged in on first load of web app
  useEffect(() => {
    if (firstTime && refreshToken) {
      dispatch(refresh({ refreshToken }));
      setFirstTime(false);
    }
  }, [firstTime, setFirstTime, refreshToken, dispatch]);

  return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <Switch>
          <Route path='/practice'>
            <RaceField />
          </Route>
          <Route path='/login'>
            <LoginPage />
          </Route>
          <Route path='/register'>
            <RegisterPage />
          </Route>
          <Route path='/'>
            <MainMenu />
          </Route>
          <Redirect to='/' />
        </Switch>
      </Router>
    </MuiThemeProvider>
  )
};

export default App;
