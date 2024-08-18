import React from 'react';
import { Box } from '@mui/material';

import AuthLogin from '@/components/login/AuthLogin';

const LoginPage = () => {
  return (
    <Box title="Login">
      <AuthLogin />
    </Box>
  );
};

export default LoginPage;
