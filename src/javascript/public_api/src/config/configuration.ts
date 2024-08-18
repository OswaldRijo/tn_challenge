import { registerAs } from '@nestjs/config';

export default registerAs('config', () => ({
  port: parseInt(process.env.PORT, 10) || 3000,
  jwt: {
    secret: process.env.JWT_SECRET,
  },
  aws: {
    region: process.env.AWS_REGION || 'us-east-2',
  },
  client: {
    host: process.env.CLIENT_HOST,
  },
  usersService: {
    path: process.env.USERS_SERVICE_PATH,
  },
  operationsService: {
    path: process.env.OPERATIONS_SERVICE_PATH,
  },
}));
