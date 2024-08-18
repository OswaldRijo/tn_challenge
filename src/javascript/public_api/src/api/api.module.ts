import { Module } from '@nestjs/common';

import { OperationsService } from '@/api/operations/services/operations.service';
import { UsersService } from '@/api/users/services/users.service';

import { AuthModule } from '@/api/auth/auth.module';
import { OperationsModule } from '@/api/operations/operations.module';
import { UsersModule } from '@/api/users/users.module';

@Module({
  imports: [
    AuthModule,
    OperationsModule,
    UsersModule,
  ],
  providers: [
    OperationsService,
    UsersService,
  ],
})
export class ApiModule {}
