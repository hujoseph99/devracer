import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';
import raceFieldReducer from '../features/race-text-field/raceFieldSlice'

export const store = configureStore({
  reducer: {
    raceField: raceFieldReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
