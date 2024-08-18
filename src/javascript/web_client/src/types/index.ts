import { OperationStateProps } from '@/types/operation';
import { UserStateProps } from '@/types/user';

export interface DefaultRootStateProps {
  user: UserStateProps;
  operation: OperationStateProps;
}
