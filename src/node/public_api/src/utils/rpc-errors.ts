import { RpcError } from '@protobuf-ts/runtime-rpc';

export function extractRpcErrorMessage(e): string {
  if (e instanceof RpcError && e.message.includes('desc')) {
    return e.message.slice(
      e.message.indexOf('desc =', 0) + 7,
      e.message.length,
    );
  }

  return e.message;
}
