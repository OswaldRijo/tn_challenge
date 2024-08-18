import React from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import { LoadingButton } from '@mui/lab';
import { Grid, Typography, useMediaQuery, useTheme } from '@mui/material';

import { Routes } from '@/config/routes';

const Success = () => {
  const { t } = useTranslation(['client']);
  const navigate = useNavigate();
  const theme = useTheme();
  const matchDownSM = useMediaQuery(theme.breakpoints.down('md'));
  const handleGoToHomepage = () => {
    navigate(Routes.home);
  };

  return (
    <Grid container direction="column" justifyContent="flex-end">
      <Grid item xs={12}>
        <Grid container justifyContent="center" alignItems="center" spacing={2}>
          <Grid item xs={12}>
            <Typography color={theme.palette.secondary.main} gutterBottom variant={matchDownSM ? 'h3' : 'h2'} textAlign="center">
              {t('signup.successful.label')}
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <Typography variant={matchDownSM ? 'h4' : 'h3'} textAlign="center">
              {t('signup.welcome.label')}
            </Typography>
          </Grid>
          <Grid container item xs={12} justifyContent="center" alignItems="center">
            <LoadingButton onClick={() => handleGoToHomepage()} size="large" variant="outlined">
              {t('signup.button.label.go_to_login')}
            </LoadingButton>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default Success;
