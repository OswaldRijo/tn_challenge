import { createSlice } from '@reduxjs/toolkit';

import { API_ENDPOINTS } from '@/client/api-endpoints';
import { dispatch } from '@/store';
import { DefaultRootStateProps } from '@/types';
import { ApplyOperation, Operation, Record } from '@/types/operation';
import axios, { ServerError } from '@/utils/axios';

// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['operation'] = {
  error: null,
  currentUserBalance: null,
  records: null,
  totalRecords: null,
  isLoading: null,
  isFetching: true,
  isUpdating: false,
  isCreating: false
};

function processOperation(operation): Operation {
  return { ...operation, args: JSON.parse(operation.args).args };
}

function processRecords(records: Record[]) {
  return records.map((record) => {
    const op = record.operation ? processOperation(record.operation) : null;
    const r = {
      ...record,
      createdAt: record.createdAt ? new Date((record.createdAt as any).seconds * 1000) : null,
      updatedAt: record.updatedAt ? new Date((record.updatedAt as any).seconds * 1000) : null,
      operation: op
    };
    return r;
  });
}

const slice = createSlice({
  name: 'operations',
  initialState,
  reducers: {
    hasError(state, action) {
      state.error = action.payload;
    },
    setLoading(state, action) {
      state.isLoading = action.payload.isLoading;
    },
    setCurrentUserBalanceSuccess(state, action) {
      state.currentUserBalance = action.payload.balance;
      state.currentUserBalance.createdAt = new Date(action.payload.balance.createdAt.seconds * 1000);
      state.currentUserBalance.updatedAt = new Date(action.payload.balance.updatedAt.seconds * 1000);
    },
    setApplyOperationSuccess(state, action) {
      state.currentUserBalance = action.payload.currentUserBalance;
      state.currentUserBalance.createdAt = new Date(action.payload.currentUserBalance.createdAt.seconds * 1000);
      state.currentUserBalance.updatedAt = new Date(action.payload.currentUserBalance.updatedAt.seconds * 1000);
    },
    setFilterOperationsSuccess(state, action) {
      state.records = processRecords(action.payload.records);
      state.totalRecords = action.payload.totalCount;
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
    resetState(state) {
      state = initialState;
    }
  }
});

// Reducer
export default slice.reducer;

export function getUserBalance() {
  dispatch(slice.actions.setLoading({ isLoading: true }));
  dispatch(slice.actions.setIsFetching({ isFetching: true }));
  return async () => {
    try {
      const response = await axios.get(API_ENDPOINTS.USER_BALANCE);
      dispatch(slice.actions.setCurrentUserBalanceSuccess(response.data));
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

export function applyOperation(operation: ApplyOperation) {
  dispatch(slice.actions.setLoading({ isLoading: true }));
  dispatch(slice.actions.setIsFetching({ isFetching: true }));
  return async () => {
    try {
      const response = await axios.post(API_ENDPOINTS.OPERATIONS, operation);
      dispatch(slice.actions.setApplyOperationSuccess(response.data));
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

export function filterOperations({
  limit = 10,
  page = 0,
  sortBy = '',
  orderBy = ''
}: {
  limit?: number;
  page?: number;
  orderBy?: string;
  sortBy?: string;
}) {
  dispatch(slice.actions.setLoading({ isLoading: true }));
  dispatch(slice.actions.setIsFetching({ isFetching: true }));
  return async () => {
    try {
      const response = await axios.get(API_ENDPOINTS.OPERATIONS, { params: { limit, page, orderBy, sortBy } });
      dispatch(slice.actions.setFilterOperationsSuccess(response.data));
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

export function deleteOperation(id: number) {
  dispatch(slice.actions.setLoading({ isLoading: true }));
  dispatch(slice.actions.setIsFetching({ isFetching: true }));
  return async () => {
    try {
      await axios.delete(`${API_ENDPOINTS.OPERATIONS}/${id}`);
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

export const { hasError, resetState } = slice.actions;
