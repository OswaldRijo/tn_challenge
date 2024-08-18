import React from 'react';
import { PersonAdd } from '@mui/icons-material';
import { Grid, Link } from '@mui/material';

import Success from '@/components/signup/Success';
import { Routes } from '@/config/routes';

const SuccessfulSignupPage = () => {
  return (
    <Grid container direction="column" justifyContent="flex-end" sx={{ minHeight: '100vh' }}>
      <Grid item xs={12}>
        <Grid container justifyContent="center" alignItems="center" sx={{ minHeight: 'calc(100vh - 68px)' }}>
          <Grid item sx={{ m: { xs: 1, sm: 3 }, mb: 0 }}>
            <Grid container spacing={2} alignItems="center" justifyContent="center">
              <Grid item sx={{ mb: 3 }}>
                <Link href={Routes.login} aria-label="theme-logo">
                  <PersonAdd sx={{ width: 100, height: 100 }} />
                </Link>
              </Grid>
              <Grid item xs={12}>
                <Success />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default SuccessfulSignupPage;
