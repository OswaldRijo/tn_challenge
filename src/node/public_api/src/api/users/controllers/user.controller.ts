import {CreateUserResponse, GetUserRequest, GetUserResponse, CreateUserRequest } from "@/pb";
import {
  BadRequestException, Body,
  Controller,
  Get,
  HttpCode,
  HttpStatus,
  Param,
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

  @Get(':id')
  @HttpCode(HttpStatus.OK)
  @UseGuards(AuthGuard)
  async getUser(
    @Param() params: GetUserRequest,
    @Req() req,
  ): Promise<GetUserResponse> {
    if (!req.usern || req.user.id !== params.id ) {
      throw new UnauthorizedException({
        code: 'UNAUTHORIZED',
        message: 'api.unauthorized',
      });
    }

    try {
      const { response } = await this.usersService.getUser({ id: params.id });
      return response;
    } catch (e) {
      throw new BadRequestException({
        code: 'BAD_REQUEST',
        message: e.message,
      });
    }
  }

  @Get('sign-up')
  @HttpCode(HttpStatus.OK)
  @UseGuards(AuthGuard)
  async createUser(
    @Body() body: CreateUserRequest,
  ): Promise<CreateUserResponse> {
    try {
      const { response } = await this.usersService.createUser(body);
      return response;
    } catch (e) {
      throw new BadRequestException({
        code: 'BAD_REQUEST',
        message: e.message,
      });
    }
  }
}
