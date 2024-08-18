import { OperationsServiceClient } from '@/pb';
import { ChannelCredentials } from '@grpc/grpc-js';
import { Injectable } from '@nestjs/common';

@Injectable()
export class OperationsService extends OperationsServiceClient {
  constructor() {
    super(
      process.env.OPERATIONS_SERVICE_PATH,
      ChannelCredentials.createInsecure(),
    );
  }
}
