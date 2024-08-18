import React from 'react';
import { Paper } from '@mui/material';

import Calculator from '@/components/calculator/Calculator';

const HomePage = () => {
  return (
    <Paper sx={{ p: 3, borderRadius: 2 }} elevation={3}>
      <Calculator />
    </Paper>
  );
};

export default HomePage;
