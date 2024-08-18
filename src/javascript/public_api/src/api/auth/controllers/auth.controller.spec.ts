import { JwtService } from '@nestjs/jwt';
import { AuthController } from './auth.controller';
import { UsersService } from '@/api/users/services/users.service';
import { AuthService } from '@/api/auth/services/auth.service';
import { Test } from '@nestjs/testing';
import { Response } from 'express';

const verifyCustomerEmailMock = jest.fn();
const authServiceMock = jest.fn(() => ({}) as AuthService);
const usersServiceMock = jest.fn(() => ({
  verifyCustomerEmail: verifyCustomerEmailMock,
}));
const jwtServiceMock = jest.fn(() => ({}));
const customersServiceMock = jest.fn(() => ({}));

describe('AuthController', () => {
  let controller: AuthController;

  beforeEach(async () => {
    const moduleRef = await Test.createTestingModule({
      controllers: [AuthController],
      providers: [
        {
          provide: AuthService,
          useValue: authServiceMock(),
        },
        {
          provide: UsersService,
          useValue: usersServiceMock(),
        },
        {
          provide: JwtService,
          useValue: jwtServiceMock(),
        },
        {
          provide: 'CONFIGURATION(config)',
          useValue: {},
        },
      ],
    }).compile();

    controller = moduleRef.get<AuthController>(AuthController);
  });

  it('should login', async () => {
    const dto = {
      username: 'cacaroto',
      password: '123456',
    };

    const expectedRes = { message: 'Verified' };

    usersServiceMock().verifyCustomerEmail.mockResolvedValue({
      response: {},
    } as any);
    const response = {} as  Response;

    // act
    const result = await controller.signIn(dto, response);

    // assert
    expect(result).toEqual(expectedRes);
  });
});
