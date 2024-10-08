import { Injectable, NestMiddleware, Logger } from '@nestjs/common';
import { ConsoleLogger } from '@nestjs/common';

import { Request, Response, NextFunction } from 'express';

@Injectable()
export class AppLoggerMiddleware implements NestMiddleware {
  private logger = new Logger('HTTP');

  use(request: Request, response: Response, next: NextFunction): void {
    const { ip, method, baseUrl: url } = request;
    const userAgent = request.get('user-agent') || '';

    response.on('close', () => {
      const { statusCode } = response;
      const contentLength = response.get('content-length');

      this.logger.log(
        `METHOD: ${method} | Path: ${url} | Response status code: ${statusCode} | Response Size: ${contentLength} - ${userAgent}| IP: ${ip}`,
      );
    });

    next();
  }
}

export class AppLogger extends ConsoleLogger {
  log(message: any) {
    console.log(JSON.stringify(message));
  }
}
