import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { LoadingButton } from '@mui/lab';
import { Box, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput, TextField } from '@mui/material'; // material-ui

import { Routes } from '@/config/routes';
import { dispatch } from '@/store';
import { openSnackbar } from '@/store/slices/snackbar';
import { login } from '@/store/slices/user';

const AuthLoginForm = ({ loginProp, ...others }: { loginProp?: number }) => {
  const [username, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');

  const [showPassword, setShowPassword] = React.useState(false);
  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleMouseDownPassword = (event: React.MouseEvent) => {
    event.preventDefault()!;
  };

  const { t } = useTranslation(['client']);
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);

  function onSubmit(e) {
    e.preventDefault();
    setIsLoading(true);
    dispatch(
      login({
        username,
        password
      })
    )
      .then((r) => {
        navigate(Routes.home);
      })
      .catch((e) => {
        dispatch(
          openSnackbar({
            open: true,
            message: e.message,
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
      })
      .finally(() => setIsLoading(false));
  }

  return (
    <form onSubmit={onSubmit}>
      <Box sx={{ display: 'flex', flexWrap: 'wrap' }}>
        <TextField
          sx={{ m: 1 }}
          label={t('login.input.label.username')}
          type="text"
          value={username}
          onChange={(e) => setEmail(e.currentTarget.value)}
          fullWidth
        />
        <FormControl fullWidth sx={{ m: 1 }} variant="outlined">
          <InputLabel htmlFor="password">{t('login.input.label.password')}</InputLabel>
          <OutlinedInput
            label={t('login.input.label.password')}
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
          {t('login.button.label')}
        </LoadingButton>
      </Box>
    </form>
  );
};

export default AuthLoginForm;
