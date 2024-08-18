import React from 'react';
import { CircularProgress, Grid } from '@mui/material';

// ==============================|| LOADER ||============================== //

const Loader = () => (
  <Grid container spacing={0} direction="column" alignItems="center" justifyContent="center" sx={{ minHeight: '100vh' }}>
    <Grid item xs={3}>
      <CircularProgress aria-label="progress" />
    </Grid>
  </Grid>
);

export default Loader;
