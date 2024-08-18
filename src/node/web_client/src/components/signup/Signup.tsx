import React from 'react';
import { useTranslation } from 'react-i18next';
import { ArrowBack, PersonAddRounded } from '@mui/icons-material';
import { Box, Button, Card, Grid, Link, Stack, Typography, useMediaQuery } from '@mui/material';
import { useTheme } from '@mui/material/styles';

import { Routes } from '@/config/routes';

import SignupForm from './SignupForm';

const Signup = () => {
  const theme = useTheme();
  const { t } = useTranslation(['client']);
  const matchDownSM = useMediaQuery(theme.breakpoints.down('md'));

  return (
    <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
      <Card variant="outlined" sx={{ minHeight: '50vh', maxWidth: matchDownSM ? '80vw' : '50vw', borderRadius: 5 }}>
        <Grid container direction="column" justifyContent="flex-end">
          <Grid item xs={12}>
            <Grid container justifyContent="center" alignItems="center" sx={{ minHeight: 'calc(100vh - 68px)' }}>
              <Grid item sx={{ m: { xs: 1, sm: 3 }, mb: 0 }}>
                <Grid container spacing={2} alignItems="center" justifyContent="center">
                  <Grid item xs={1} sx={{ mb: 4 }}>
                    <Button component={Link} href={Routes.login} aria-label="theme-logo">
                      <ArrowBack color="secondary" />
                    </Button>
                  </Grid>
                  <Grid item xs={11} sx={{ mb: 3 }} container spacing={2} alignItems="center" justifyContent="center">
                    <Link href={Routes.home} aria-label="theme-logo">
                      <PersonAddRounded sx={{ ml: -5, width: 100, height: 100 }} />
                    </Link>
                  </Grid>
                  <Grid item xs={12}>
                    <Grid container direction={matchDownSM ? 'column-reverse' : 'row'} alignItems="center" justifyContent="center">
                      <Grid item>
                        <Stack alignItems="center" justifyContent="center" spacing={1}>
                          <Typography color={theme.palette.secondary.main} gutterBottom variant={matchDownSM ? 'h3' : 'h2'}>
                            {t('signup.new_account')}
                          </Typography>
                        </Stack>
                      </Grid>
                    </Grid>
                  </Grid>
                  <Grid item xs={12}>
                    <SignupForm />
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

export default Signup;
