import React from 'react';
import { CssBaseline, MuiThemeProvider, Typography } from '@material-ui/core';
import { theme } from './theme';

const App = (): JSX.Element => {
  return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <Typography variant='h1'>Hi</Typography>
    </MuiThemeProvider>
  )
};

export default App;
