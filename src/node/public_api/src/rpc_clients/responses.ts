import { RpcError } from '@protobuf-ts/runtime-rpc';
import { StatusCodes } from '@/types/grpc';

export const isErrorNotFound = (error: any) => {
  if (error instanceof RpcError && error.code === StatusCodes.NOT_FOUND)
    return true;
  return false;
};

export const isInvalid = (error: any) => {
  if (error instanceof RpcError && error.code === StatusCodes.INVALID_ARGUMENT)
    return true;
  return false;
};
