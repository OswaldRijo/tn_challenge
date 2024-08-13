import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

let transport: any | null = null;
export const getUsersServiceTransport = () => {
  if (transport === null) {
    transport = new GrpcTransport({
      host: process.env.USERS_SERVICE_PATH,
      channelCredentials: ChannelCredentials.createInsecure(),
      interceptors: [],
    });
  }
  return transport;
};

let notificationsTransport: any | null = null;
export const getOperationsServiceTransport = () => {
  if (notificationsTransport === null) {
    notificationsTransport = new GrpcTransport({
      host: process.env.OPERATIONS_SERVICE_PATH,
      channelCredentials: ChannelCredentials.createInsecure(),
      interceptors: [],
    });
  }
  return notificationsTransport;
};
