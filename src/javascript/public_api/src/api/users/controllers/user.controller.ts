import { CreateUserRequest } from '@/pb';
import { throwErrorBasedOnType } from '@/utils/rpc-errors';
import {
  Body,
  Controller,
  HttpCode,
  HttpStatus,
  Post,
  Req,
  UnauthorizedException,
  UseGuards,
} from '@nestjs/common';
import { AuthGuard } from '@/api/auth/guards/jwt-auth.guard';

import { UsersService } from '@/api/users/services/users.service';

@Controller('api/users')
export class UserController {
  constructor(private readonly usersService: UsersService) {}

  @Post('whoami')
  @HttpCode(HttpStatus.OK)
  @UseGuards(AuthGuard)
  async whoami(@Req() req) {
    if (!req.user) {
      throw new UnauthorizedException();
    }
    return req.user;
  }

  @Post('sign-up')
  @HttpCode(HttpStatus.OK)
  async createUser(@Body() body: CreateUserRequest): Promise<object> {
    try {
      const { response } = await this.usersService.createUser(
        CreateUserRequest.create(body),
      );
      return response;
    } catch (e) {
      throwErrorBasedOnType(e);
    }
  }
}
