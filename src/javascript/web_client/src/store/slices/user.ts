import { createSlice } from '@reduxjs/toolkit';

import { API_ENDPOINTS } from '@/client/api-endpoints';
import { dispatch } from '@/store';
import { DefaultRootStateProps } from '@/types';
import axios, { ServerError } from '@/utils/axios';

// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['user'] = {
  error: null,
  current: null,
  isLoading: null,
  isFetching: true,
  isUpdating: false,
  isCreating: false
};

const slice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    hasError(state, action) {
      state.error = action.payload;
    },
    setLoading(state, action) {
      state.isLoading = action.payload.isLoading;
    },
    setIsCreating(state, action) {
      state.isCreating = action.payload.isCreating;
    },
    setIsFetching(state, action) {
      state.isFetching = action.payload.isFetching;
    },
    setIsUpdating(state, action) {
      state.isUpdating = action.payload.isUpdating;
    },
    setCurrentUserSuccess(state, action) {
      state.current = action.payload.user;
    },
    setCurrentUser(state, action) {
      state.current = {
        id: action.payload.id,
        username: action.payload.username,
        status: action.payload.status
      };
    },
    resetState(state) {
      state = initialState;
    }
  }
});

// Reducer
export default slice.reducer;

// ----------------------------------------------------------------------

export function createCustomer(user: { username: string; password?: string }) {
  dispatch(slice.actions.setLoading({ isLoading: true }));
  dispatch(slice.actions.setIsFetching({ isFetching: true }));
  return async () => {
    try {
      const response = await axios.post(API_ENDPOINTS.USERS_SIGN_UP, user);
      return response.data;
    } catch (e) {
      const error = e as ServerError;
      dispatch(slice.actions.hasError(error.response.data.message));
      throw error.response.data;
    } finally {
      dispatch(slice.actions.setLoading({ isLoading: false }));
      dispatch(slice.actions.setIsFetching({ isFetching: false }));
    }
  };
}

export function login(body: { username: string; password: string }) {
  dispatch(slice.actions.setLoading({ isLoading: true }));
  dispatch(slice.actions.setIsFetching({ isFetching: true }));
  return async () => {
    try {
      await axios.post(API_ENDPOINTS.AUTH_LOGIN, body);
    } catch (e) {
      const error = e as ServerError;
      dispatch(slice.actions.hasError(error.response.data.message));
      throw error.response.data;
    } finally {
      dispatch(slice.actions.setLoading({ isLoading: false }));
      dispatch(slice.actions.setIsFetching({ isFetching: false }));
    }
  };
}

export const { hasError, setCurrentUser, resetState } = slice.actions;
