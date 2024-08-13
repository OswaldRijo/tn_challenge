import { Module } from '@nestjs/common';
import { OperationsService } from '@/api/operations/services/operations.service';
import { BrandsController } from '@/api/operations/controllers/operations.controller';

@Module({
  providers: [OperationsService],
  controllers: [BrandsController],
  exports: [OperationsService],
})
export class OperationsModule {}
