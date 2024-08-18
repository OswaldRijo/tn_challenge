import {
  DeleteOperationDto,
  OperationsQueryDto,
} from '@/api/operations/dto/login.dto';
import {
  ApplyOperationRequest,
  DeleteRecordsRequest,
  FilterRecordsRequest,
  GetUserBalanceRequest,
} from '@/pb';
import { throwErrorBasedOnType } from '@/utils/rpc-errors';
import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Query,
  Req,
  UseGuards,
} from '@nestjs/common';
import { OperationsService } from '@/api/operations/services/operations.service';
import { AuthGuard } from '@/api/auth/guards/jwt-auth.guard';

@Controller('api/operations')
export class OperationsController {
  constructor(private readonly operationsService: OperationsService) {}

  @Post()
  @UseGuards(AuthGuard)
  async create(
    @Body() applyOpBody: ApplyOperationRequest,
    @Req() req,
  ): Promise<object> {
    try {
      const response = await this.operationsService.ApplyOperation(
        ApplyOperationRequest.fromObject({
          ...applyOpBody,
          userId: req.user.userId,
        }),
      );
      return response.toObject();
    } catch (e) {
      throwErrorBasedOnType(e);
    }
  }

  @Get()
  @UseGuards(AuthGuard)
  async filter(
    @Query() params: OperationsQueryDto,
    @Req() req,
  ): Promise<object> {
    try {
      const response = await this.operationsService.FilterRecords(
        FilterRecordsRequest.fromObject({
          limit: parseInt(params.limit || '10'),
          page: parseInt(params.page || '0'),
          userId: req.user.userId,
        }),
      );
      return response.toObject();
    } catch (e) {
      throwErrorBasedOnType(e);
    }
  }

  @Get('balance')
  @UseGuards(AuthGuard)
  async userBalance(@Req() req): Promise<object> {
    try {
      const response = await this.operationsService.GetUserBalance(
        GetUserBalanceRequest.fromObject({
          userId: req.user.userId,
        }),
      );
      return response.toObject();
    } catch (e) {
      throwErrorBasedOnType(e);
    }
  }

  @Delete(':id')
  @UseGuards(AuthGuard)
  async delete(
    @Param() params: DeleteOperationDto,
    @Req() req,
  ): Promise<object> {
    try {
      const response = await this.operationsService.DeleteRecords(
        DeleteRecordsRequest.fromObject({
          userId: req.user.userId,
          recordIds: [req.params.id],
        }),
      );
      return response.toObject();
    } catch (e) {
      throwErrorBasedOnType(e);
    }
  }
}
