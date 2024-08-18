import { Test, TestingModule } from '@nestjs/testing';
import { expect } from '@jest/globals';
import { AuthService } from './auth.service';
import { SessionsService } from '@/api/sessions/services/sessions.service';
import { JwtModule } from '@nestjs/jwt';
import { jwtConstants } from '@/api/auth/constants';

const validateCredentialsMock = jest.fn();
const sessionServiceMock = {
  validateCredentials: validateCredentialsMock,
};

describe('AuthService', () => {
  let service: AuthService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        AuthService,
        {
          provide: SessionsService,
          useValue: sessionServiceMock,
        },
      ],
      imports: [
        JwtModule.register({
          global: true,
          secret: jwtConstants.secret,
          signOptions: { expiresIn: '86400s' },
        }),
      ],
    }).compile();

    service = module.get<AuthService>(AuthService);
  });

  it('should signIn user', async () => {
    // arrange
    const expectedUsername = 'username';
    const expectedPassword = 'pass';
    sessionServiceMock.validateCredentials.mockResolvedValue({
      response: {
        user: {
          id: 23,
          username: 'username',
          userUuid: 'userUuid',
          lastName: 'lastName',
          roles: [],
          isActive: 'isActive',
          emailVerified: 'emailVerified',
        },
      },
    });

    const { accessToken, refreshToken } = await service.signIn(
      expectedUsername,
      expectedPassword,
    );

    expect(accessToken).not.toHaveLength(0);
    expect(refreshToken).not.toHaveLength(0);
    expect(sessionServiceMock.validateCredentials).toHaveBeenCalledWith({
      username: expectedUsername,
      password: expectedPassword,
    });
  });
});
