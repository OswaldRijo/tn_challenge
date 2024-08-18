import { FunctionComponent, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { Avatar, Box, Typography } from '@mui/material';
import { useTheme } from '@mui/material/styles';

import UserBalanceSkeleton from '@/components/balance/UserBalanceSkeleton';
import { dispatch, useSelector } from '@/store';
import { getUserBalance } from '@/store/slices/operation';
import { openSnackbar } from '@/store/slices/snackbar';

interface UserBalanceProps {}

const UserBalance: FunctionComponent<UserBalanceProps> = () => {
  const theme = useTheme();
  const { t } = useTranslation(['client', 'server']);
  const userSelector = useSelector((s) => s.user);
  const operationSelector = useSelector((s) => s.operations);
  const getCurrentBalance = () => {
    dispatch(getUserBalance()).catch((e) => {
      dispatch(
        openSnackbar({
          open: true,
          message: t(`server:${e.message}`),
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
    });
  };
  useEffect(() => {
    getCurrentBalance();
  }, []);

  if (operationSelector.isFetching) return <UserBalanceSkeleton />;

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
            <Avatar>{userSelector.current.username[0]}</Avatar>
          </Box>
          <Box>
            <Typography component={'p'} fontSize={12} color={theme.palette.grey['400']} sx={{ fontWeight: '900' }}>
              {userSelector.current.username}
            </Typography>
          </Box>
        </Box>
        <Box
          sx={{
            display: 'block',
            textAlign: 'center'
          }}
        >
          <Typography component={'p'} fontSize={12} color={theme.palette.grey['400']} sx={{ fontWeight: '500' }}>
            {t('client:common.balance.label')}
          </Typography>
          <Typography component={'p'} fontSize={16} color={theme.palette.grey['400']} sx={{ fontWeight: '900' }}>
            ${operationSelector.currentUserBalance?.currentBalance.toLocaleString()}
          </Typography>
        </Box>
      </Box>
    </div>
  );
};

export default UserBalance;
