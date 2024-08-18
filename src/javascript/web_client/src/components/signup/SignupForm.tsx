import React from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import { Visibility, VisibilityOff } from '@mui/icons-material';
import { LoadingButton } from '@mui/lab';
import { Box, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput, TextField } from '@mui/material';

import { Routes } from '@/config/routes';
import { dispatch } from '@/store';
import { openSnackbar } from '@/store/slices/snackbar';
import { createCustomer } from '@/store/slices/user';

const SignupForm = () => {
  const [showPassword, setShowPassword] = React.useState(false);
  const [isLoading, setIsLoading] = React.useState(false);
  const [username, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');

  const navigate = useNavigate();
  const { t } = useTranslation(['client']);

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleMouseDownPassword = (event: React.MouseEvent) => {
    event.preventDefault();
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    dispatch(
      createCustomer({
        username,
        password
      })
    )
      .then(async () => {
        navigate(Routes.signUpSuccessful);
      })
      .catch((e) => {
        setIsLoading(false);

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

  return (
    <form onSubmit={onSubmit}>
      <Box sx={{ display: 'flex', flexWrap: 'wrap' }}>
        <TextField
          sx={{ m: 1 }}
          label={t('signup.input.label.username')}
          type="text"
          value={username}
          onChange={(e) => setEmail(e.currentTarget.value)}
          fullWidth
        />
        <FormControl fullWidth sx={{ m: 1 }} variant="outlined">
          <InputLabel htmlFor="password">{t('signup.input.label.password')}</InputLabel>
          <OutlinedInput
            label={t('signup.input.label.password')}
            id="password"
            type={showPassword ? 'text' : 'password'}
            onChange={(e) => setPassword(e.currentTarget.value)}
            value={password}
            endAdornment={
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle password visibility"
                  onClick={handleClickShowPassword}
                  onMouseDown={handleMouseDownPassword}
                  edge="end"
                  size="large"
                >
                  {showPassword ? <Visibility /> : <VisibilityOff />}
                </IconButton>
              </InputAdornment>
            }
          />
        </FormControl>
      </Box>
      <Box sx={{ mt: 2 }}>
        <LoadingButton color="primary" disabled={isLoading} fullWidth size="large" type="submit" variant="contained">
          {t('signup.button.label')}
        </LoadingButton>
      </Box>
    </form>
  );
};

export default SignupForm;
