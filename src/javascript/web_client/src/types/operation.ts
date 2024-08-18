export enum OperationType {
  ADDITION = 0,
  SUBTRACTION = 1,
  MULTIPLICATION = 2,
  DIVISION = 3,
  SQUARE_ROOT = 4,
  RANDOM_STRING = 5,
  SWITCH,
  EQ,
  DOT,
  DEL
}

export const OperatorTypeToSignMap = {
  [OperationType.ADDITION]: '+',
  [OperationType.SUBTRACTION]: '-',
  [OperationType.MULTIPLICATION]: 'x',
  [OperationType.DIVISION]: '/',
  [OperationType.SQUARE_ROOT]: '√',
  [OperationType.RANDOM_STRING]: 'Rand str'
};

export interface ApplyOperation {
  operationType: OperationType;
  args: number[];
}

export interface Balance {
  id: number;
  userId: number;
  currentBalance: number;
  createdAt: Date;
  updatedAt: Date;
}

export interface Operation {
  id: number;
  userId: number;
  operationType: OperationType;
  cost: number;
  args: string[];
  createdAt: Date;
  updatedAt: Date;
}

export interface Record {
  id: number;
  operationId: number;
  userBalance: number;
  deleted: boolean;
  operationResponse: string;
  createdAt: Date;
  updatedAt: Date;
  operation: Operation;
}

export interface OperationStateProps {
  currentUserBalance: Balance | null;
  records: Record[] | null;
  totalRecords: number | null;
  error: object | string | null;
  isLoading: boolean;
  isFetching: boolean;
  isUpdating: boolean;
  isCreating: boolean;
}
