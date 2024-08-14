import { StatusCodes } from '@/types/grpc';
import {
  BadRequestException,
  InternalServerErrorException,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common';
import { RpcError } from '@protobuf-ts/runtime-rpc';

export function extractRpcErrorMessage(e): string {
  if (e instanceof RpcError && e.message.includes('desc')) {
    return e.message.slice(
      e.message.indexOf('desc =', 0) + 7,
      e.message.length,
    );
  }

  return e.message;
}

export function throwErrorBasedOnType(e): string {
  if (e instanceof RpcError) {
    switch (e.code) {
      case StatusCodes.INVALID_ARGUMENT:
        throw new BadRequestException({
          code: 'BAD_REQUEST',
          message: extractRpcErrorMessage(e),
        });
      case StatusCodes.NOT_FOUND:
        throw new NotFoundException({
          code: StatusCodes.NOT_FOUND,
          message: extractRpcErrorMessage(e),
        });
      case StatusCodes.INTERNAL:
        throw new InternalServerErrorException({
          code: StatusCodes.INTERNAL,
          message: extractRpcErrorMessage(e),
        });
      case StatusCodes.PERMISSION_DENIED:
        throw new UnauthorizedException({
          code: StatusCodes.PERMISSION_DENIED,
          message: extractRpcErrorMessage(e),
        });
    }
  }

  throw new InternalServerErrorException({
    code: StatusCodes.INTERNAL,
    message: e.message,
  });
}
