import { Button } from '@mui/material';

interface DigitButtonProps {
  text: string | number;
  handleClick: () => void;
  enabled?: boolean;
}

export const DigitButton = ({ handleClick, text, enabled = true }: DigitButtonProps) => {
  return (
    <Button disabled={!enabled} sx={{ mr: 1 }} variant="outlined" size="large" onClick={handleClick}>
      {text}
    </Button>
  );
};

export default DigitButton;
