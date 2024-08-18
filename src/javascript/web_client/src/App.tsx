import { Provider } from 'react-redux';
import { RouterProvider } from 'react-router-dom';
import router from 'browserRouter';

import Snackbar from '@/components/ui/Snackbar';
import { store } from '@/store';

import './i18n/config';

function MyApp() {
  return (
    <Provider store={store}>
      <>
        <RouterProvider router={router} />
        <Snackbar />
      </>
    </Provider>
  );
}

export default MyApp;
