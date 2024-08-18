import { useTranslation } from 'react-i18next';
import HomeTwoToneIcon from '@mui/icons-material/HomeTwoTone';
import { Box, Button, Card, CardContent, CardMedia, Grid, Link, Typography } from '@mui/material';
import { styled } from '@mui/material/styles';

import { Routes } from '@/config/routes';

// styles
const CardMediaWrapper = styled('div')({
  width: '85vh',
  margin: '0 auto',
  position: 'relative'
});

const ErrorWrapper = styled('div')({
  maxWidth: 350,
  margin: '0 auto',
  textAlign: 'center'
});

const ErrorCard = styled(Card)({
  height: 'auto',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center'
});

const AccessDeniedPage = () => {
  const { t } = useTranslation('common');

  return (
    <Box title="Error 500">
      <ErrorCard>
        <CardContent>
          <Grid container justifyContent="center">
            <Grid item xs={12}>
              <CardMediaWrapper>
                <CardMedia component="img" title="Slider5 image" />
              </CardMediaWrapper>
            </Grid>
            <Grid item xs={12} justifyItems={'center'}>
              <ErrorWrapper>
                <Grid container spacing={3}>
                  <Grid item xs={12} justifyItems={'center'}>
                    <Typography variant="h1" component="div">
                      {t('text-access-denied')}
                    </Typography>
                  </Grid>
                  <Grid item xs={12} justifyItems={'center'}>
                    <Typography variant="body2">{t('text-access-denied-message')}</Typography>
                  </Grid>
                  <Grid item xs={12} justifyItems={'center'}>
                    <Button variant="contained" size="large" component={Link} href={Routes.home}>
                      <HomeTwoToneIcon sx={{ fontSize: '1.3rem', mr: 0.75 }} /> {t('text-return-home')}
                    </Button>
                  </Grid>
                </Grid>
              </ErrorWrapper>
            </Grid>
          </Grid>
        </CardContent>
      </ErrorCard>
    </Box>
  );
};

export default AccessDeniedPage;
