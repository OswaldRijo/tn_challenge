import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import { SupervisedUserCircleRounded } from '@mui/icons-material';
import { Box, Card, Grid, Link, Stack, Typography, useMediaQuery } from '@mui/material';
import { useTheme } from '@mui/material/styles';

import { Routes } from '@/config/routes';
import { dispatch } from '@/store';
import { openSnackbar } from '@/store/slices/snackbar';

import AuthLoginForm from './AuthLoginForm';

const AuthLogin = () => {
  const theme = useTheme();
  const navigate = useNavigate();
  const { t } = useTranslation(['client']);

  const matchDownSM = useMediaQuery(theme.breakpoints.down('md'));
  const [errorOnLogin] = useState(false);

  useEffect(() => {
    if (errorOnLogin) {
      dispatch(
        openSnackbar({
          open: true,
          message: t('common:error-login'),
          anchorOrigin: { vertical: 'top', horizontal: 'right' },
          variant: 'alert',
          closeColor: 'white',
          alert: {
            color: 'error'
          },
          close: true,
          transition: 'SlideLeft'
        })
      );
    }
  }, [errorOnLogin]);

  return (
    <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
      <Card variant="outlined" sx={{ minHeight: '50vh', maxWidth: matchDownSM ? '80vw' : '50vw', borderRadius: 5 }}>
        <Grid container direction="column" justifyContent="flex-end">
          <Grid item xs={12}>
            <Grid container justifyContent="center" alignItems="center" sx={{ minHeight: 'calc(100vh - 68px)' }}>
              <Grid item sx={{ m: { xs: 1, sm: 3 }, mb: 0 }}>
                <Grid container spacing={2} alignItems="center" justifyContent="center">
                  <Grid item sx={{ mb: 3 }}>
                    <Link href={Routes.home} aria-label="theme-logo">
                      <SupervisedUserCircleRounded sx={{ width: 100, height: 100 }} />
                    </Link>
                  </Grid>
                  <Grid item xs={12}>
                    <Grid container alignItems="center" justifyContent="center">
                      <Grid item>
                        <Stack alignItems="center" justifyContent="center" spacing={1}>
                          <Typography color={theme.palette.secondary.main} gutterBottom variant={matchDownSM ? 'h3' : 'h2'}>
                            {t('login.greeting')}
                          </Typography>
                          <Typography variant="caption" fontSize="16px" textAlign={matchDownSM ? 'center' : 'inherit'}>
                            {t('login.greetings_2')}
                          </Typography>
                        </Stack>
                      </Grid>
                    </Grid>
                  </Grid>
                  <Grid item xs={12}>
                    <AuthLoginForm />
                  </Grid>
                  <Grid item xs={12}>
                    <Typography align="center">
                      <Link
                        component="button"
                        onClick={() => {
                          navigate(Routes.signUp);
                        }}
                      >
                        {t('login.no_account')}
                      </Link>
                    </Typography>
                  </Grid>
                </Grid>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Card>
    </Box>
  );
};

export default AuthLogin;
