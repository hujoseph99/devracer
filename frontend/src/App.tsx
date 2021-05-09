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
import { refresh, selectIsLoggedIn, selectRefreshToken, selectUserID } from './features/auth/authSlice';
import { fetchUserData } from './features/user/userSlice';
import { CustomGame } from './features/custom-game/CustomGame';

const App = (): JSX.Element => {
  const dispatch = useDispatch();
  const refreshToken = useSelector(selectRefreshToken);
  const [firstTime, setFirstTime] = useState(true);

  const loggedIn = useSelector(selectIsLoggedIn);
  const userid = useSelector(selectUserID);

  // check if logged in on first load of web app
  useEffect(() => {
    if (firstTime && refreshToken) {
      dispatch(refresh({ refreshToken }));
      setFirstTime(false);
    }
  }, [firstTime, setFirstTime, refreshToken, dispatch]);

  useEffect(() => {
    if (loggedIn) {
      dispatch(fetchUserData({ userid }))
    }
  }, [dispatch, loggedIn, userid])

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
          <Route path='/custom'>
            <CustomGame />
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
