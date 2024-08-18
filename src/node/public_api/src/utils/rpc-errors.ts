import { StatusCodes, StatusCodesStr } from '@/types/grpc';
import {
  BadRequestException,
  InternalServerErrorException,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common';

export function throwErrorBasedOnType(e): string {
  switch (e.code) {
    case StatusCodes.INVALID_ARGUMENT:
      throw new BadRequestException({
        code: 'BAD_REQUEST',
        message: e.details,
      });
    case StatusCodes.NOT_FOUND:
      throw new NotFoundException({
        code: StatusCodesStr.NOT_FOUND,
        message: e.details,
      });
    case StatusCodes.INTERNAL:
      throw new InternalServerErrorException({
        code: StatusCodesStr.INTERNAL,
        message: e.details,
      });
    case StatusCodes.PERMISSION_DENIED:
      throw new UnauthorizedException({
        code: StatusCodesStr.PERMISSION_DENIED,
        message: e.details,
      });
  }

  throw new InternalServerErrorException({
    code: StatusCodesStr.INTERNAL,
    message: e.message,
  });
}
