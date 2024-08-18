import { IsNotEmpty, IsString } from 'class-validator';

export class OperationsQueryDto {
  @IsString()
  limit: string;

  @IsString()
  page: string;
}

export class DeleteOperationDto {
  @IsString()
  @IsNotEmpty()
  id: string;
}
