import { Module } from '@nestjs/common';
import { AuthService } from './services/auth.service';
import { AuthController } from '@/api/auth/controllers/auth.controller';
import { UsersModule } from '@/api/users/users.module';
import { JwtModule } from '@nestjs/jwt';
import { jwtConstants } from '@/api/auth/constants';
import { ConfigModule } from '@nestjs/config';
import configuration from '@/config/configuration';
@Module({
  imports: [
    UsersModule,
    JwtModule.register({
      global: true,
      secret: jwtConstants.secret,
      signOptions: { expiresIn: '86400s' },
    }),
    ConfigModule.forRoot({
      load: [configuration],
    }),
  ],
  providers: [AuthService],
  controllers: [AuthController],
  exports: [],
})
export class AuthModule {}
