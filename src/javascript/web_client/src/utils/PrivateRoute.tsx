import React from 'react';
import { useNavigate } from 'react-router-dom';

import AccessDeniedPage from '@/components/ui/AccessDenied';
import CircularLoader from '@/components/ui/CircularLoader';
import { logout, whoami } from '@/services/auth';
import { dispatch } from '@/store';
import { setCurrentUser } from '@/store/slices/user';

const PrivateRoute: React.FC<{
  children: any;
  withRedirect?: boolean;
  redirectURL: string;
}> = ({ children, withRedirect = true, redirectURL }) => {
  const navigate = useNavigate();
  const [hasPermission, setHasPermission] = React.useState<boolean | null>(null);
  const [isError, setIsError] = React.useState<boolean | null>(null);
  const [performLogout, setPerformLogout] = React.useState<boolean | null>(null);

  function loadWhoAmI() {
    whoami()
      .then((user) => {
        setHasPermission(true);
        setIsError(false);
        dispatch(setCurrentUser(user));
      })
      .catch((e) => {
        if (e.response.status === 401 || e.response.status === 403) {
          setHasPermission(false);
          setPerformLogout(true);
        } else {
          setIsError(true);
        }
      });
  }

  function logUserOut() {
    logout()
      .then(() => {
        setIsError(null);
      })
      .catch((e) => {
        setIsError(true);
      })
      .finally(() => {
        if (withRedirect) {
          navigate(redirectURL);
        }
        setHasPermission(null);
        setPerformLogout(null);
      });
  }

  React.useEffect(() => {
    if (performLogout) {
      logUserOut();
    }
  }, [performLogout]);
  React.useEffect(() => {
    loadWhoAmI();
  }, []);
  if (isError) {
    return <>ERROR</>;
  }
  if (hasPermission !== null && hasPermission) {
    return <>{children}</>;
  }
  if (hasPermission !== null && !hasPermission) {
    return <AccessDeniedPage />;
  }
  // Session is being fetched, or no user.
  // If no user, useEffect() will redirect.
  return <CircularLoader />;
};

export default PrivateRoute;
