// third-party
import { combineReducers } from 'redux';

import operationReducer from './slices/operation';
import snackbar from './slices/snackbar';
import userReducer from './slices/user';

// ==============================|| COMBINE REDUCER ||============================== //

const reducer = combineReducers({
  user: userReducer,
  operations: operationReducer,
  snackbar: snackbar
});

export default reducer;
