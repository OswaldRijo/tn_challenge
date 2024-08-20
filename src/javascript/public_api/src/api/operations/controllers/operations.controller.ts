import { AuthGuard } from '@/api/auth/guards/jwt-auth.guard';
import {
  DeleteOperationDto,
  OperationsQueryDto,
} from '@/api/operations/dto/login.dto';
import { OperationsService } from '@/api/operations/services/operations.service';
import {
  ApplyOperationRequest,
  DeleteRecordsRequest,
  FilterRecordsRequest,
  GetUserBalanceRequest,
  OrderBy,
  OrderFieldsEnum,
  OrderTypeEnum,
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
      const { response } = await this.operationsService.applyOperation(
        ApplyOperationRequest.create({
          ...applyOpBody,
          userId: req.user.userId,
        }),
      );
      return response;
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
      const orderByFields = this.setUpOrderByFields(
        params.orderBy,
        params.sortBy,
      );
      const { response } = await this.operationsService.filterRecords(
        FilterRecordsRequest.create({
          limit: parseInt(params.limit || '10'),
          page: parseInt(params.page || '0'),
          userId: req.user.userId,
          orderByFields: orderByFields,
        }),
      );
      return response;
    } catch (e) {
      throwErrorBasedOnType(e);
    }
  }

  @Get('balance')
  @UseGuards(AuthGuard)
  async userBalance(@Req() req): Promise<object> {
    try {
      const { response } = await this.operationsService.getUserBalance(
        GetUserBalanceRequest.create({
          userId: req.user.userId,
        }),
      );
      return response;
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
      const { response } = await this.operationsService.deleteRecords(
        DeleteRecordsRequest.create({
          userId: req.user.userId,
          recordIds: [req.params.id],
        }),
      );
      return response;
    } catch (e) {
      throwErrorBasedOnType(e);
    }
  }
  private setUpOrderByFields(orderBy: string, sortBy: string): OrderBy[] {
    if (orderBy === 'id') {
      if (sortBy === 'asc') {
        return [
          OrderBy.create({
            orderType: OrderTypeEnum.ASC,
            orderField: OrderFieldsEnum.ID,
          }),
        ];
      }
      return [
        OrderBy.create({
          orderType: OrderTypeEnum.DESC,
          orderField: OrderFieldsEnum.ID,
        }),
      ];
    }
    if (orderBy === 'operation') {
      if (sortBy === 'asc') {
        return [
          OrderBy.create({
            orderType: OrderTypeEnum.ASC,
            orderField: OrderFieldsEnum.OPERATION,
          }),
        ];
      }
      return [
        OrderBy.create({
          orderType: OrderTypeEnum.DESC,
          orderField: OrderFieldsEnum.OPERATION,
        }),
      ];
    }
    if (orderBy === 'createdAt') {
      if (sortBy === 'asc') {
        return [
          OrderBy.create({
            orderType: OrderTypeEnum.ASC,
            orderField: OrderFieldsEnum.CREATED_AT,
          }),
        ];
      }
      return [
        OrderBy.create({
          orderType: OrderTypeEnum.DESC,
          orderField: OrderFieldsEnum.CREATED_AT,
        }),
      ];
    }
    return [];
  }
}
