import { ApplyOperationRequest, ApplyOperationResponse } from '@/pb';
import { extractRpcErrorMessage } from '@/utils/rpc-errors';
import {
  BadRequestException,
  Body,
  Controller,
  Post,
  Req,
  UseGuards,
} from '@nestjs/common';
import { OperationsService } from '@/api/operations/services/operations.service';
import { AuthGuard } from '@/api/auth/guards/jwt-auth.guard';

@Controller('api/operations')
export class OperationsController {
  constructor(private readonly operationsService: OperationsService) {}

  @Post('apply')
  @UseGuards(AuthGuard)
  async create(
    @Body() applyOpBody: ApplyOperationRequest,
    @Req() req,
  ): Promise<ApplyOperationResponse> {
    try {
      const { response } = await this.operationsService.applyOperation({
        ...applyOpBody,
        userId: req.user.userId,
      });
      return response;
    } catch (e) {
      throw new BadRequestException({
        code: 'BAD_REQUEST',
        message: extractRpcErrorMessage(e),
      });
    }
  }
}
