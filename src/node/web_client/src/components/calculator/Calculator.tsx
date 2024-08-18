/**
 * v0 by Vercel.
 * @see https://v0.dev/t/8NCRQ4iLktk
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */

import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Box, TextField } from '@mui/material';

import Userbalance from '@/components/balance/Userbalance';
import ActionButton from '@/components/calculator/ActionButton';
import DigitButton from '@/components/calculator/DigitButton';
import { dispatch } from '@/store';
import { applyOperation } from '@/store/slices/operation';
import { openSnackbar } from '@/store/slices/snackbar';
import { OperationType } from '@/types/operation';

const MULTI_ARGS_OPS = [
  OperationType.ADDITION,
  OperationType.MULTIPLICATION,
  OperationType.SUBTRACTION,
  OperationType.DIVISION,
  OperationType.EQ
];
const SINGLE_ARGS_OPS = [
  OperationType.SQUARE_ROOT,
  OperationType.SWITCH,
  OperationType.DOT,
  OperationType.RANDOM_STRING,
  OperationType.DEL
];

export default function Calculator() {
  const { t } = useTranslation(['server']);
  const [result, setResult] = useState<string>('0');
  const [memory, setMemory] = useState(0);
  const [operation, setOperation] = useState<OperationType | null>(null);
  const [operationApplied, setOperationApplied] = useState<boolean>(false);
  const [enableNumericOps, setEnableNumericOps] = useState<boolean>(true);
  const handleNumber = (num) => {
    if (!operationApplied) {
      if (result === '0') {
        setResult(num.toString());
      } else {
        setResult(result + num.toString());
      }
    } else {
      setResult(num.toString());
      setOperationApplied(false);
    }
  };
  const handleOperation = (op: OperationType) => {
    if (operation !== null && MULTI_ARGS_OPS.includes(op)) {
      handleMultiArgsOps();
      setOperation(op);
    } else if (SINGLE_ARGS_OPS.includes(op)) {
      handleSingleArgsOps(op);
    } else if (op !== OperationType.EQ) {
      setMemory(parseFloat(result));
      setResult('0');
      setOperation(op);
    }
  };

  function sendOperation(op: OperationType, args: number[]) {
    dispatch(
      applyOperation({
        operationType: op,
        args
      })
    )
      .then((r) => {
        setResult(r.record.operationResponse);
        if (op !== OperationType.RANDOM_STRING) {
          setMemory(parseFloat(r.record.operationResponse));
          setOperationApplied(true);
        } else {
          setEnableNumericOps(false);
        }
      })
      .catch((e) => {
        dispatch(
          openSnackbar({
            open: true,
            message: t(e.message),
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
  }

  const handleMultiArgsOps = () => {
    switch (operation) {
      case OperationType.ADDITION:
      case OperationType.SUBTRACTION:
      case OperationType.MULTIPLICATION:
      case OperationType.DIVISION:
        sendOperation(operation, [memory, parseFloat(result)]);
        break;
      default:
        setMemory(parseFloat(result));
    }
  };

  const handleSingleArgsOps = (op: OperationType) => {
    let newResult;
    switch (op) {
      case OperationType.SQUARE_ROOT:
        sendOperation(op, [parseFloat(result)]);
        setOperation(op);
        break;
      case OperationType.RANDOM_STRING:
        sendOperation(op, []);
        break;
      case OperationType.SWITCH:
        newResult = parseFloat(result) * -1;
        setResult(newResult.toString());
        break;
      case OperationType.DOT:
        if (!result.includes('.')) {
          setResult(result + '.');
        }
        break;
      case OperationType.DEL:
        if (result.length > 1) {
          setResult(result.slice(0, result.length - 1));
        } else {
          setResult('0');
        }

        break;
    }
  };

  const handleClear = () => {
    setResult('0');
    setMemory(0);
    setOperation(null);
    setEnableNumericOps(true);
  };
  return (
    <div>
      <Box sx={{ display: 'flex', flexWrap: 'wrap' }}>
        <Userbalance />
      </Box>

      <Box sx={{ mb: 1 }}>
        <TextField inputProps={{ style: { fontSize: 13 } }} fullWidth id="RESULT" value={result} variant="outlined" />
      </Box>
      <Box sx={{ m: 1 }} display="flex" justifyContent="center">
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.DEL)} text={'<-'} />
        <ActionButton handleClick={() => handleOperation(OperationType.RANDOM_STRING)} text={'R'} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.SWITCH)} text={'-/+'} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.SQUARE_ROOT)} text={'√'} />
      </Box>
      <Box sx={{ m: 1 }} display="flex" justifyContent="center">
        <ActionButton handleClick={() => handleClear()} text={'C'} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.DIVISION)} text={'/'} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.MULTIPLICATION)} text={'*'} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.SUBTRACTION)} text={'-'} />
      </Box>

      <Box sx={{ m: 1 }} display="flex" justifyContent="center">
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(7)} text={7} />
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(8)} text={8} />
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(9)} text={9} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.ADDITION)} text={'+'} />
      </Box>
      <Box sx={{ m: 1 }} display="flex" justifyContent="center">
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(4)} text={4} />
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(5)} text={5} />
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(6)} text={6} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.DOT)} text={'.'} />
      </Box>
      <Box sx={{ m: 1 }} display="flex" justifyContent="center">
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(1)} text={1} />
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(2)} text={2} />
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(3)} text={3} />
        <ActionButton enabled={enableNumericOps} handleClick={() => handleOperation(OperationType.EQ)} text={'='} />
      </Box>
      <Box sx={{ m: 1 }} display="flex" justifyContent="flex-start">
        <DigitButton enabled={enableNumericOps} handleClick={() => handleNumber(0)} text={0} />
      </Box>
    </div>
  );
}
