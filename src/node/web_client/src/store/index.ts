import { TypedUseSelectorHook, useSelector as useAppSelector } from 'react-redux';
import { configureStore } from '@reduxjs/toolkit';

// project imports
import rootReducer from './reducer';

// ==============================|| REDUX - MAIN STORE ||============================== //

const store: any = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware({ serializableCheck: false, immutableCheck: false })
});

export type RootState = ReturnType<typeof rootReducer>;

const useSelector: TypedUseSelectorHook<RootState> = useAppSelector;

const { dispatch } = store;

export { dispatch, store, useSelector };
