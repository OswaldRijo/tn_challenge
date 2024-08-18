import { CheckUserCredentialsRequest } from '@/pb';
import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UsersService } from '@/api/users/services/users.service';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
  ) {}

  async signIn(username: string, password: string) {
    try {
      const response = await this.usersService.CheckUserCredentials(
        CheckUserCredentialsRequest.fromObject({
          username,
          password,
        }),
      );
      const res = response.toObject();
      const payload = {
        userId: res.user.id,
        username: res.user.username,
        status: res.user.status,
      };
      const refreshTokenPayload = {};
      const [accessToken, refreshToken] = await Promise.all([
        this.jwtService.signAsync(payload),
        this.jwtService.signAsync(refreshTokenPayload),
      ]);

      return {
        accessToken,
        refreshToken,
      };
    } catch (e) {
      throw new UnauthorizedException();
    }
  }
}
