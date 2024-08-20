import { OperationsServiceClient } from '@/pb';
import { getOperationsTransport } from '@/rpc_clients/grpc-transport';
import { Injectable } from '@nestjs/common';

@Injectable()
export class OperationsService extends OperationsServiceClient {
  constructor() {
    super(getOperationsTransport());
  }
}
