import {
  Body,
  Controller,
  HttpCode,
  HttpStatus,
  Inject,
  Post,
  Req,
  Res,
  UnauthorizedException,
} from '@nestjs/common';
import { AuthService } from '@/api/auth/services/auth.service';
import { LoginDto } from '@/api/auth/dto/login.dto';
import { Response } from 'express';

import config from '@/config/configuration';
import { ConfigType } from '@nestjs/config';

@Controller('api/auth')
export class AuthController {
  constructor(
    private authService: AuthService,
    @Inject(config.KEY) private configService: ConfigType<typeof config>,
  ) {}

  @HttpCode(HttpStatus.OK)
  @Post('login')
  async signIn(@Body() signInDto: LoginDto, @Res() res: Response) {
    try {
      const { accessToken, refreshToken } = await this.authService.signIn(
        signInDto.username,
        signInDto.password,
      );

      res.cookie('accessToken', accessToken, {
        maxAge: 2592000000,
        sameSite: 'none',
        secure: true,
      });
      res.cookie('lastLogin', new Date().getTime(), {
        maxAge: 2592000000,
        sameSite: 'none',
        secure: true,
      });

      res.cookie('refreshToken', refreshToken, {
        maxAge: 2592000000,
        sameSite: 'none',
        secure: true,
      });
      res.send();
    } catch (e) {
      throw new UnauthorizedException({
        code: 'UNAUTHORIZED',
        message: e.message,
      });
    }
  }

  @HttpCode(HttpStatus.OK)
  @Post('logout')
  logout(@Req() req, @Res() res: Response) {
    res.clearCookie('accessToken');
    res.clearCookie('lastLogin');
    res.clearCookie('refreshToken');
    return res.end();
  }
}
