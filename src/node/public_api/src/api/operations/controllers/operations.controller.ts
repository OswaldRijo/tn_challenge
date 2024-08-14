import {ApplyOperationRequest, ApplyOperationResponse} from "@/pb";
import {
  BadRequestException,
  Body,
  Controller,
  Post,
  UseGuards,
} from '@nestjs/common';
import { OperationsService } from '@/api/operations/services/operations.service';
import { AuthGuard } from '@/api/auth/guards/jwt-auth.guard';

@Controller('api/operations')
export class OperationsController {
  constructor(private readonly operationsService: OperationsService) {}

  @Post()
  @UseGuards(AuthGuard)
  async create(@Body() applyOpBody: ApplyOperationRequest): Promise<ApplyOperationResponse> {
    try {
      const { response } = await this.operationsService.applyOperation(applyOpBody);
      return response;
    } catch (e) {
      throw new BadRequestException({
        code: 'BAD_REQUEST',
        message: e.message,
      });
    }
  }
}
