import { SyntheticEvent } from 'react';
import CloseIcon from '@mui/icons-material/Close';
import { Alert, Button, Fade, Grow, IconButton, Slide, SlideProps } from '@mui/material';
import MuiSnackbar from '@mui/material/Snackbar';

import { dispatch, useSelector } from '@/store';
import { closeSnackbar } from '@/store/slices/snackbar';
import { KeyObject } from '@/types/common';

function TransitionSlideLeft(props: SlideProps) {
  return <Slide {...props} direction="left" />;
}

function TransitionSlideUp(props: SlideProps) {
  return <Slide {...props} direction="up" />;
}

function TransitionSlideRight(props: SlideProps) {
  return <Slide {...props} direction="right" />;
}

function TransitionSlideDown(props: SlideProps) {
  return <Slide {...props} direction="down" />;
}

function GrowTransition(props: SlideProps) {
  return <Grow {...props} />;
}

// animation options
const animation: KeyObject = {
  SlideLeft: TransitionSlideLeft,
  SlideUp: TransitionSlideUp,
  SlideRight: TransitionSlideRight,
  SlideDown: TransitionSlideDown,
  Grow: GrowTransition,
  Fade
};

const Snackbar = () => {
  const snackbar = useSelector((state) => state.snackbar);
  const { actionButton, anchorOrigin, alert, close, message, open, transition, variant, closeColor } = snackbar;

  const handleClose = (event: SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    dispatch(closeSnackbar());
  };

  return (
    <>
      {variant === 'default' && (
        <MuiSnackbar
          anchorOrigin={anchorOrigin}
          open={open}
          autoHideDuration={6000}
          onClose={handleClose}
          message={message}
          TransitionComponent={animation[transition]}
          action={
            <>
              <Button color="secondary" size="small" onClick={handleClose}>
                UNDO
              </Button>
              <IconButton size="small" aria-label="close" color="inherit" onClick={handleClose} sx={{ mt: 0.25 }}>
                <CloseIcon color="inherit" fontSize="small" />
              </IconButton>
            </>
          }
        />
      )}

      {variant === 'alert' && (
        <MuiSnackbar
          TransitionComponent={animation[transition]}
          anchorOrigin={anchorOrigin}
          open={open}
          autoHideDuration={6000}
          onClose={handleClose}
        >
          <Alert
            variant={alert.variant}
            color={alert.color}
            action={
              <>
                {actionButton !== false && (
                  <Button size="small" onClick={handleClose} sx={{ color: 'background.paper' }}>
                    UNDO
                  </Button>
                )}
                {close !== false && (
                  <IconButton sx={{ color: closeColor }} size="small" aria-label="close" onClick={handleClose}>
                    <CloseIcon fontSize="small" />
                  </IconButton>
                )}
              </>
            }
            sx={{
              ...(alert.variant === 'outlined' && {
                bgcolor: 'background.paper'
              })
            }}
          >
            {message}
          </Alert>
        </MuiSnackbar>
      )}
    </>
  );
};

export default Snackbar;
