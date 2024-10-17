// @generated by protobuf-ts 2.9.4 with parameter use_proto_field_name
// @generated from protobuf file "blotservice/v1beta1/blotservice.proto" (package "blotservice.v1beta1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { BlotService } from "./blotservice";
import type { GetGameSetForPlayerResponse } from "./blotservice";
import type { GetGameSetForPlayerRequest } from "./blotservice";
import type { CreateGameSetResponse } from "./blotservice";
import type { CreateGameSetRequest } from "./blotservice";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetGameForPlayerResponse } from "./blotservice";
import type { GetGameForPlayerRequest } from "./blotservice";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service blotservice.v1beta1.BlotService
 */
export interface IBlotServiceClient {
    /**
     * @generated from protobuf rpc: GetGameForPlayer(blotservice.v1beta1.GetGameForPlayerRequest) returns (blotservice.v1beta1.GetGameForPlayerResponse);
     */
    getGameForPlayer(input: GetGameForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameForPlayerRequest, GetGameForPlayerResponse>;
    /**
     * @generated from protobuf rpc: CreateGameSet(blotservice.v1beta1.CreateGameSetRequest) returns (blotservice.v1beta1.CreateGameSetResponse);
     */
    createGameSet(input: CreateGameSetRequest, options?: RpcOptions): UnaryCall<CreateGameSetRequest, CreateGameSetResponse>;
    /**
     * @generated from protobuf rpc: GetGameSetForPlayer(blotservice.v1beta1.GetGameSetForPlayerRequest) returns (blotservice.v1beta1.GetGameSetForPlayerResponse);
     */
    getGameSetForPlayer(input: GetGameSetForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameSetForPlayerRequest, GetGameSetForPlayerResponse>;
}
/**
 * @generated from protobuf service blotservice.v1beta1.BlotService
 */
export class BlotServiceClient implements IBlotServiceClient, ServiceInfo {
    typeName = BlotService.typeName;
    methods = BlotService.methods;
    options = BlotService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: GetGameForPlayer(blotservice.v1beta1.GetGameForPlayerRequest) returns (blotservice.v1beta1.GetGameForPlayerResponse);
     */
    getGameForPlayer(input: GetGameForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameForPlayerRequest, GetGameForPlayerResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetGameForPlayerRequest, GetGameForPlayerResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CreateGameSet(blotservice.v1beta1.CreateGameSetRequest) returns (blotservice.v1beta1.CreateGameSetResponse);
     */
    createGameSet(input: CreateGameSetRequest, options?: RpcOptions): UnaryCall<CreateGameSetRequest, CreateGameSetResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateGameSetRequest, CreateGameSetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetGameSetForPlayer(blotservice.v1beta1.GetGameSetForPlayerRequest) returns (blotservice.v1beta1.GetGameSetForPlayerResponse);
     */
    getGameSetForPlayer(input: GetGameSetForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameSetForPlayerRequest, GetGameSetForPlayerResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetGameSetForPlayerRequest, GetGameSetForPlayerResponse>("unary", this._transport, method, opt, input);
    }
}
