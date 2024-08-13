// @generated by protobuf-ts 2.9.4 with parameter generate_dependencies,long_type_string,optimize_code_size
// @generated from protobuf file "operations/operations_service.proto" (package "pb.operations", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { Balance } from "./structures";
import { Operation } from "./structures";
import { Record } from "./structures";
import { OperationType } from "./structures";
/**
 * @generated from protobuf message pb.operations.ApplyOperationRequest
 */
export interface ApplyOperationRequest {
    /**
     * @generated from protobuf field: pb.operations.OperationType operationType = 1;
     */
    operationType: OperationType;
    /**
     * @generated from protobuf field: repeated double args = 2;
     */
    args: number[];
    /**
     * @generated from protobuf field: int64 user_id = 3;
     */
    userId: string;
}
/**
 * @generated from protobuf message pb.operations.ApplyOperationResponse
 */
export interface ApplyOperationResponse {
    /**
     * @generated from protobuf field: pb.operations.Record record = 1;
     */
    record?: Record;
    /**
     * @generated from protobuf field: pb.operations.Operation operation = 2;
     */
    operation?: Operation;
    /**
     * @generated from protobuf field: pb.operations.Balance current_user_balance = 3;
     */
    currentUserBalance?: Balance;
}
/**
 * @generated from protobuf message pb.operations.GetUserBalanceRequest
 */
export interface GetUserBalanceRequest {
    /**
     * @generated from protobuf field: int64 user_id = 1;
     */
    userId: string;
}
/**
 * @generated from protobuf message pb.operations.GetUserBalanceResponse
 */
export interface GetUserBalanceResponse {
    /**
     * @generated from protobuf field: pb.operations.Balance balance = 1;
     */
    balance?: Balance;
}
/**
 * @generated from protobuf message pb.operations.FilterRecordsRequest
 */
export interface FilterRecordsRequest {
    /**
     * @generated from protobuf field: optional double balance = 1;
     */
    balance?: number;
    /**
     * @generated from protobuf field: optional string limit = 2;
     */
    limit?: string;
    /**
     * @generated from protobuf field: optional bool page = 3;
     */
    page?: boolean;
}
/**
 * @generated from protobuf message pb.operations.FilterRecordsResponse
 */
export interface FilterRecordsResponse {
    /**
     * @generated from protobuf field: repeated pb.operations.Record records = 1;
     */
    records: Record[];
}
/**
 * @generated from protobuf message pb.operations.DeleteRecordsRequest
 */
export interface DeleteRecordsRequest {
    /**
     * @generated from protobuf field: repeated int64 record_ids = 1;
     */
    recordIds: string[];
    /**
     * @generated from protobuf field: int64 user_id = 2;
     */
    userId: string;
}
/**
 * @generated from protobuf message pb.operations.DeleteRecordsResponse
 */
export interface DeleteRecordsResponse {
    /**
     * @generated from protobuf field: repeated pb.operations.Record records = 1;
     */
    records: Record[];
}
// @generated message type with reflection information, may provide speed optimized methods
class ApplyOperationRequest$Type extends MessageType<ApplyOperationRequest> {
    constructor() {
        super("pb.operations.ApplyOperationRequest", [
            { no: 1, name: "operationType", kind: "enum", T: () => ["pb.operations.OperationType", OperationType] },
            { no: 2, name: "args", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 3, name: "user_id", kind: "scalar", T: 3 /*ScalarType.INT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.ApplyOperationRequest
 */
export const ApplyOperationRequest = new ApplyOperationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ApplyOperationResponse$Type extends MessageType<ApplyOperationResponse> {
    constructor() {
        super("pb.operations.ApplyOperationResponse", [
            { no: 1, name: "record", kind: "message", T: () => Record },
            { no: 2, name: "operation", kind: "message", T: () => Operation },
            { no: 3, name: "current_user_balance", kind: "message", T: () => Balance }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.ApplyOperationResponse
 */
export const ApplyOperationResponse = new ApplyOperationResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetUserBalanceRequest$Type extends MessageType<GetUserBalanceRequest> {
    constructor() {
        super("pb.operations.GetUserBalanceRequest", [
            { no: 1, name: "user_id", kind: "scalar", T: 3 /*ScalarType.INT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.GetUserBalanceRequest
 */
export const GetUserBalanceRequest = new GetUserBalanceRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetUserBalanceResponse$Type extends MessageType<GetUserBalanceResponse> {
    constructor() {
        super("pb.operations.GetUserBalanceResponse", [
            { no: 1, name: "balance", kind: "message", T: () => Balance }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.GetUserBalanceResponse
 */
export const GetUserBalanceResponse = new GetUserBalanceResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FilterRecordsRequest$Type extends MessageType<FilterRecordsRequest> {
    constructor() {
        super("pb.operations.FilterRecordsRequest", [
            { no: 1, name: "balance", kind: "scalar", opt: true, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "limit", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "page", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.FilterRecordsRequest
 */
export const FilterRecordsRequest = new FilterRecordsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FilterRecordsResponse$Type extends MessageType<FilterRecordsResponse> {
    constructor() {
        super("pb.operations.FilterRecordsResponse", [
            { no: 1, name: "records", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Record }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.FilterRecordsResponse
 */
export const FilterRecordsResponse = new FilterRecordsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteRecordsRequest$Type extends MessageType<DeleteRecordsRequest> {
    constructor() {
        super("pb.operations.DeleteRecordsRequest", [
            { no: 1, name: "record_ids", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 3 /*ScalarType.INT64*/ },
            { no: 2, name: "user_id", kind: "scalar", T: 3 /*ScalarType.INT64*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.DeleteRecordsRequest
 */
export const DeleteRecordsRequest = new DeleteRecordsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteRecordsResponse$Type extends MessageType<DeleteRecordsResponse> {
    constructor() {
        super("pb.operations.DeleteRecordsResponse", [
            { no: 1, name: "records", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Record }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message pb.operations.DeleteRecordsResponse
 */
export const DeleteRecordsResponse = new DeleteRecordsResponse$Type();
/**
 * @generated ServiceType for protobuf service pb.operations.OperationsService
 */
export const OperationsService = new ServiceType("pb.operations.OperationsService", [
    { name: "ApplyOperation", options: {}, I: ApplyOperationRequest, O: ApplyOperationResponse },
    { name: "GetUserBalance", options: {}, I: GetUserBalanceRequest, O: GetUserBalanceResponse },
    { name: "FilterRecords", options: {}, I: FilterRecordsRequest, O: FilterRecordsResponse },
    { name: "DeleteRecords", options: {}, I: DeleteRecordsRequest, O: DeleteRecordsResponse }
]);
