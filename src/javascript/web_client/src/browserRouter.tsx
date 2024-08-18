import * as React from 'react';
import { createBrowserRouter } from 'react-router-dom';

import { Routes } from '@/config/routes';
import LoginPage from '@/pages/auth/login';
import SignupPage from '@/pages/auth/signup';
import OperationsPage from '@/pages/operations';
import PrivateRoute from '@/utils/PrivateRoute';

import MainLayout from './components/ui/MainLayout';
import SuccessfulSignupPage from './pages/auth/successful_signup';
import HomePage from './pages/home';

const router = createBrowserRouter([
  {
    path: Routes.home,
    element: (
      <PrivateRoute redirectURL={Routes.login}>
        <MainLayout>
          <HomePage />
        </MainLayout>
      </PrivateRoute>
    )
  },
  {
    path: Routes.operations,
    element: (
      <PrivateRoute redirectURL={Routes.login}>
        <MainLayout>
          <OperationsPage />
        </MainLayout>
      </PrivateRoute>
    )
  },
  {
    path: Routes.signUpSuccessful,
    Component: SuccessfulSignupPage
  },
  {
    path: Routes.signUp,
    Component: SignupPage
  },
  {
    path: Routes.login,
    Component: LoginPage
  },
  {
    path: '*',
    element: <>NOT FOUND </>
  }
]);

export default router;
