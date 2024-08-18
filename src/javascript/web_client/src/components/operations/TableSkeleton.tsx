import { FunctionComponent } from 'react';
import { Box, Skeleton } from '@mui/material';

const TableSkeleton: FunctionComponent = () => {
  return (
    <div style={{ width: '100%' }}>
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          p: 1,
          m: 1,
          bgcolor: 'background.paper',
          borderRadius: 1
        }}
      >
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            alignContent: 'center',
            textAlign: 'center'
          }}
        >
          <Box>
            <Skeleton variant="circular" width={40} height={40} />
          </Box>
          <Box
            sx={{
              width: 80
            }}
          >
            <Skeleton variant="text" sx={{ fontSize: '1rem', width: '100%' }} />
          </Box>
        </Box>
        <Box
          sx={{
            width: 80
          }}
        >
          <Skeleton variant="text" sx={{ fontSize: '1rem', width: '100%' }} />
          <Skeleton variant="text" sx={{ fontSize: '1rem', width: '100%' }} />
        </Box>
      </Box>
    </div>
  );
};

export default TableSkeleton;
