import React from 'react';
import { CssBaseline, MuiThemeProvider } from '@material-ui/core';
import { theme } from './theme';
import { RaceField } from './features/race-text-field/RaceField';

const App = (): JSX.Element => {
  return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <RaceField />
    </MuiThemeProvider>
  )
};

export default App;
