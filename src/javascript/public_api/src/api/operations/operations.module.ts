import { Module } from '@nestjs/common';
import { OperationsService } from '@/api/operations/services/operations.service';
import { OperationsController } from '@/api/operations/controllers/operations.controller';

@Module({
  providers: [OperationsService],
  controllers: [OperationsController],
  exports: [OperationsService],
})
export class OperationsModule {}
