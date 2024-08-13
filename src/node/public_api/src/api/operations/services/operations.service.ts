import { OperationsServiceClient } from "@/pb";
import { getUsersServiceTransport } from '@/rpc_clients/grpc-transport';
import { Injectable } from '@nestjs/common';

@Injectable()
export class OperationsService extends OperationsServiceClient {
  constructor() {
    const transport = getUsersServiceTransport();
    super(transport);
  }
}
