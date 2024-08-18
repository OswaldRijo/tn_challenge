import React from 'react';
import { Paper } from '@mui/material';

import OperationsTable from '@/components/operations/OperationsTable';

const OperationsPage = () => {
  return (
    <Paper sx={{ p: 3, borderRadius: 2, mb: 15 }} elevation={3}>
      <OperationsTable />
    </Paper>
  );
};

export default OperationsPage;
