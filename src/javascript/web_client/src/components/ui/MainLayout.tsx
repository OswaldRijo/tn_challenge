/**
 * v0 by Vercel.
 * @see https://v0.dev/t/IL82z8XC22T
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import React, { useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useLocation, useNavigate } from 'react-router-dom';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import CalculateIcon from '@mui/icons-material/Calculate';
import ListIcon from '@mui/icons-material/List';
import { BottomNavigation, BottomNavigationAction, Box, Paper } from '@mui/material';

import { Routes } from '@/config/routes';
import { logout } from '@/services/auth';

interface MainLayoutProps {
  children: React.ReactNode;
}

function getInitValue(path: string): number {
  if (path === Routes.operations) {
    return 1;
  }

  return 0;
}

function getRouteFromValue(value: number): string {
  if (value === 1) {
    return Routes.operations;
  }

  return Routes.home;
}

const MainLayout = ({ children }: MainLayoutProps) => {
  const { t } = useTranslation(['client']);
  const location = useLocation();
  const navigate = useNavigate();
  const initialValue = getInitValue(location.pathname);
  const [value, setValue] = React.useState<number | null>(initialValue);

  useEffect(() => {
    navigate(getRouteFromValue(value));
  }, [value]);

  return (
    <Box>
      <Box sx={{ width: '100%', height: '100%', pt: 5 }} display="flex" justifyContent="center" alignItems="center">
        {children}
      </Box>

      <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
        <BottomNavigation
          showLabels
          value={value}
          onChange={(event, newValue) => {
            if (newValue < 2) {
              setValue(newValue);
            } else {
              logout().then(() => navigate(Routes.login));
            }
          }}
        >
          <BottomNavigationAction label={t('home.label')} icon={<CalculateIcon />} />
          <BottomNavigationAction label={t('operations.label')} icon={<ListIcon />} />
          <BottomNavigationAction label={t('logout.label')} icon={<AccountCircleIcon />} />
        </BottomNavigation>
      </Paper>
    </Box>
  );
};

export default MainLayout;
