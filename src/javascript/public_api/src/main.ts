import { NestFactory } from '@nestjs/core';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';

import { AppModule } from './app.module';
import * as cookieParser from 'cookie-parser';
import { AppLogger } from '@/logger';
import { Logger } from '@nestjs/common';

async function bootstrap() {

  const app = await NestFactory.create(AppModule, {
    logger: new AppLogger(),
  });
  app.use(cookieParser());
  const logger = new Logger();

  const mid = (options) => (req, res, next) => {
    const methodName = `${req.method}`;
    if (methodName !== 'OPTIONS') {
      const url = `${req.url}`;
      const startTime = new Date();
      logger.log({
        httpVerb: methodName,
        method: url,
        startTime,
        params: req.params,
        query: req.query,
        body: req.body,
      });
    }
    return next();
  };
  app.use(mid({}));

  app.enableCors({
    credentials: true,
    origin: 'http://localhost:3000',
  });
  logger.log(`LISTENING ON PORT ${process.env.PORT || 8080}`);

  const config = new DocumentBuilder()
    .setTitle('True north public API')
    .setDescription('The public API')
    .setVersion('0.0.1')
    .addTag('publicapi')
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('swagger', app, document);

  await app.listen(process.env.PORT || 8080);
}

bootstrap();
