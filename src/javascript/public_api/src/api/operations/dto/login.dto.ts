import { IsNotEmpty, IsString } from 'class-validator';

export class OperationsQueryDto {
  @IsString()
  limit: string;

  @IsString()
  page: string;

  @IsString()
  orderBy: string;

  @IsString()
  sortBy: string;
}

export class DeleteOperationDto {
  @IsString()
  @IsNotEmpty()
  id: string;
}
