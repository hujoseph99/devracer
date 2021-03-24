import React from 'react';
import { CssBaseline, MuiThemeProvider } from '@material-ui/core';
import { theme } from './theme';
import { RaceTextField } from './features/race-text-field/RaceTextField';

const App = (): JSX.Element => {
  return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <RaceTextField />
    </MuiThemeProvider>
  )
};

export default App;
