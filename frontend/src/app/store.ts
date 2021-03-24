import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';
import RaceTextFieldReducer from '../features/race-text-field/raceTextFieldSlice'

export const store = configureStore({
  reducer: {
    RaceTextField: RaceTextFieldReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
