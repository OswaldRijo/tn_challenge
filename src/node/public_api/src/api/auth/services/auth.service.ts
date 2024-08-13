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
      const { response } = await this.usersService.checkUserCredentials({
        username,
        password,
      });

      const payload = {
        userId: response.user.id,
        username: response.user.username,
        status: response.user.status,
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
