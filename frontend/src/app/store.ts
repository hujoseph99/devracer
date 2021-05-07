import { configureStore, ThunkAction, Action, getDefaultMiddleware, combineReducers } from '@reduxjs/toolkit';
import raceFieldReducer from '../features/race-text-field/raceFieldSlice'
import authReducer from '../features/auth/authSlice';
import userReducer from '../features/user/userSlice';

import storage from 'redux-persist/lib/storage';
import {
  persistStore,
  persistReducer,
  FLUSH,
  REHYDRATE,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
} from 'redux-persist'

const persistConfig = {
  key: 'root',
  version: 1,
  storage,
  whitelist: ['auth']
}

const rootReducer = combineReducers({
  raceField: raceFieldReducer,
  user: userReducer,
  auth: authReducer,
});

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: getDefaultMiddleware({
    serializableCheck: {
      ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
    },
  }),
});

export const persister = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
