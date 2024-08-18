import { StatusCodes } from '@/types/grpc';

export const isErrorNotFound = (error: any) => {
  return error.code === StatusCodes.NOT_FOUND;
};

export const isInvalid = (error: any) => {
  return error.code === StatusCodes.INVALID_ARGUMENT;
};
