import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

let transport: any | null = null;
export const getTransport = () => {
  if (transport === null) {
    transport = new GrpcTransport({
      host: process.env.USERS_SERVICE_PATH,
      channelCredentials: ChannelCredentials.createInsecure(),
    });
  }
  return transport;
};

let operationsTransport: any | null = null;
export const getOperationsTransport = () => {
  if (operationsTransport === null) {
    operationsTransport = new GrpcTransport({
      host: process.env.OPERATIONS_SERVICE_PATH,
      channelCredentials: ChannelCredentials.createInsecure(),
    });
  }
  return operationsTransport;
};
