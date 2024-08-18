import { Button } from '@mui/material';

interface ActionButtonProps {
  text: string;
  handleClick: () => void;
  enabled?: boolean;
}

const ActionButton = ({ handleClick, text, enabled = true }: ActionButtonProps) => {
  return (
    <Button disabled={!enabled} sx={{ mr: 1 }} variant="contained" size="large" onClick={handleClick}>
      {text}
    </Button>
  );
};

export default ActionButton;
