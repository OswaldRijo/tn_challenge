import { CreateUserRequest, User, UserServiceClient } from '@/pb';
import { getTransport } from '@/rpc_clients/grpc-transport';
import { Injectable } from '@nestjs/common';
import { validateNonNull } from '@/utils/validate-non-null';
import { isErrorNotFound } from '@/rpc_clients/responses';
import { v4 as uuidv4 } from 'uuid';

@Injectable()
export class UsersService extends UserServiceClient {
  constructor() {
    super(getTransport());
  }

  async create(user): Promise<User> {
    try {
      const createUserRequest = CreateUserRequest.create({
        ...user,
        password: uuidv4().toString(),
      });
      const { response } = await this.createUser(createUserRequest);

      validateNonNull(response.user, 'user');
      validateNonNull(response.user.id, 'user.id');
      validateNonNull(response.user.username, 'user.username');

      return response.user;
    } catch (error) {
      throw error;
    }
  }

  async getByUsername(username: string): Promise<User> {
    try {
      const { response } = await this.getUserByUsername({ username });

      validateNonNull(response.user, 'user');
      validateNonNull(response.user.id, 'user.id');
      validateNonNull(response.user.username, 'user.email');

      return response.user;
    } catch (error) {
      if (isErrorNotFound(error)) return null;
      throw error;
    }
  }
}
