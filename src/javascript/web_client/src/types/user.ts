export enum Status {
  ACTIVE,
  INACTIVE
}

export interface User {
  id: number;
  username: string;
  createdAt: Date;
  updatedAt: Date;
  status: Status;
}

export interface UserStateProps {
  current: Partial<User> | null;
  error: object | string | null;
  isLoading: boolean;
  isFetching: boolean;
  isUpdating: boolean;
  isCreating: boolean;
}
