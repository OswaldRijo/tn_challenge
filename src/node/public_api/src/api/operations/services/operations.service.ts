import { OperationsServiceClient } from '@/pb';
import { getOperationsServiceTransport } from '@/rpc_clients/grpc-transport';
import { Injectable } from '@nestjs/common';

@Injectable()
export class OperationsService extends OperationsServiceClient {
  constructor() {
    const transport = getOperationsServiceTransport();
    super(transport);
  }
}
