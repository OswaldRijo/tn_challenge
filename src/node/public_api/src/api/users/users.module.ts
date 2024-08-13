import { Module } from '@nestjs/common';
import { UsersService } from '@/api/users/services/users.service';
import { UserController } from '@/api/users/controllers/user.controller';

@Module({
  providers: [UsersService],
  controllers: [UserController],
  exports: [UsersService],
})
export class UsersModule {}
