import React, { useEffect, useState } from 'react';
import {
  Route,
  Switch,
  Redirect,
  HashRouter
} from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';

import { CssBaseline, MuiThemeProvider } from '@material-ui/core';

import { theme } from './theme';
import { MainMenu } from './features/main-menu/MainMenu';
import { LoginPage } from './features/auth/LoginPage';
import { RegisterPage } from './features/auth/RegisterPage';
import { refresh, selectIsLoggedIn, selectRefreshToken, selectUserID } from './features/auth/authSlice';
import { fetchUserData } from './features/user/userSlice';
import { CustomGame } from './features/game/custom-game/CustomGame';
import { GitHubCallback } from './features/auth/GitHubCallback';

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
      <HashRouter>
        <Switch>
          <Route path='/auth/githubCallback'>
            <Title title="Logging in..." />
            <GitHubCallback />
          </Route>
          <Route path='/login'>
            <Title title="Login - DevRacer" />
            <LoginPage />
          </Route>
          <Route path='/register'>
            <Title title="Register - DevRacer" />
            <RegisterPage />
          </Route>
          <Route path='/custom/:lobby?'>
            <Title title="Custom Game - DevRacer" />
            <CustomGame />
          </Route>
          <Route path='/'>
            <Title title="DevRacer" />
            <MainMenu />
          </Route>
          <Redirect to='/' />
        </Switch>
      </HashRouter>
    </MuiThemeProvider>
  )
};

const Title = (props: any): JSX.Element => {
  useEffect(() => {
    document.title = props.title;
  }, [props.title]);
  return <></>;
};

export default App;
