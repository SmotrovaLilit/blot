// @generated by protobuf-ts 2.9.4 with parameter use_proto_field_name
// @generated from protobuf file "blotservice/v1beta1/blotservice.proto" (package "blotservice.v1beta1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { BlotService } from "./blotservice";
import type { GetGameSetsForPlayerResponse } from "./blotservice";
import type { GetGameSetsForPlayerRequest } from "./blotservice";
import type { GetGameSetForPlayerResponse } from "./blotservice";
import type { GetGameSetForPlayerRequest } from "./blotservice";
import type { PlayCardResponse } from "./blotservice";
import type { PlayCardRequest } from "./blotservice";
import type { StartGameResponse } from "./blotservice";
import type { StartGameRequest } from "./blotservice";
import type { LeaveGameSetResponse } from "./blotservice";
import type { LeaveGameSetRequest } from "./blotservice";
import type { JoinGameSetResponse } from "./blotservice";
import type { JoinGameSetRequest } from "./blotservice";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { CreateGameSetResponse } from "./blotservice";
import type { CreateGameSetRequest } from "./blotservice";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service blotservice.v1beta1.BlotService
 */
export interface IBlotServiceClient {
    /**
     * @generated from protobuf rpc: CreateGameSet(blotservice.v1beta1.CreateGameSetRequest) returns (blotservice.v1beta1.CreateGameSetResponse);
     */
    createGameSet(input: CreateGameSetRequest, options?: RpcOptions): UnaryCall<CreateGameSetRequest, CreateGameSetResponse>;
    /**
     * @generated from protobuf rpc: JoinGameSet(blotservice.v1beta1.JoinGameSetRequest) returns (blotservice.v1beta1.JoinGameSetResponse);
     */
    joinGameSet(input: JoinGameSetRequest, options?: RpcOptions): UnaryCall<JoinGameSetRequest, JoinGameSetResponse>;
    /**
     * @generated from protobuf rpc: LeaveGameSet(blotservice.v1beta1.LeaveGameSetRequest) returns (blotservice.v1beta1.LeaveGameSetResponse);
     */
    leaveGameSet(input: LeaveGameSetRequest, options?: RpcOptions): UnaryCall<LeaveGameSetRequest, LeaveGameSetResponse>;
    /**
     * @generated from protobuf rpc: StartGame(blotservice.v1beta1.StartGameRequest) returns (blotservice.v1beta1.StartGameResponse);
     */
    startGame(input: StartGameRequest, options?: RpcOptions): UnaryCall<StartGameRequest, StartGameResponse>;
    /**
     * @generated from protobuf rpc: PlayCard(blotservice.v1beta1.PlayCardRequest) returns (blotservice.v1beta1.PlayCardResponse);
     */
    playCard(input: PlayCardRequest, options?: RpcOptions): UnaryCall<PlayCardRequest, PlayCardResponse>;
    /**
     * @generated from protobuf rpc: GetGameSetForPlayer(blotservice.v1beta1.GetGameSetForPlayerRequest) returns (blotservice.v1beta1.GetGameSetForPlayerResponse);
     */
    getGameSetForPlayer(input: GetGameSetForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameSetForPlayerRequest, GetGameSetForPlayerResponse>;
    /**
     * @generated from protobuf rpc: GetGameSetsForPlayer(blotservice.v1beta1.GetGameSetsForPlayerRequest) returns (blotservice.v1beta1.GetGameSetsForPlayerResponse);
     */
    getGameSetsForPlayer(input: GetGameSetsForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameSetsForPlayerRequest, GetGameSetsForPlayerResponse>;
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
     * @generated from protobuf rpc: CreateGameSet(blotservice.v1beta1.CreateGameSetRequest) returns (blotservice.v1beta1.CreateGameSetResponse);
     */
    createGameSet(input: CreateGameSetRequest, options?: RpcOptions): UnaryCall<CreateGameSetRequest, CreateGameSetResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateGameSetRequest, CreateGameSetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: JoinGameSet(blotservice.v1beta1.JoinGameSetRequest) returns (blotservice.v1beta1.JoinGameSetResponse);
     */
    joinGameSet(input: JoinGameSetRequest, options?: RpcOptions): UnaryCall<JoinGameSetRequest, JoinGameSetResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<JoinGameSetRequest, JoinGameSetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: LeaveGameSet(blotservice.v1beta1.LeaveGameSetRequest) returns (blotservice.v1beta1.LeaveGameSetResponse);
     */
    leaveGameSet(input: LeaveGameSetRequest, options?: RpcOptions): UnaryCall<LeaveGameSetRequest, LeaveGameSetResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<LeaveGameSetRequest, LeaveGameSetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: StartGame(blotservice.v1beta1.StartGameRequest) returns (blotservice.v1beta1.StartGameResponse);
     */
    startGame(input: StartGameRequest, options?: RpcOptions): UnaryCall<StartGameRequest, StartGameResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<StartGameRequest, StartGameResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: PlayCard(blotservice.v1beta1.PlayCardRequest) returns (blotservice.v1beta1.PlayCardResponse);
     */
    playCard(input: PlayCardRequest, options?: RpcOptions): UnaryCall<PlayCardRequest, PlayCardResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<PlayCardRequest, PlayCardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetGameSetForPlayer(blotservice.v1beta1.GetGameSetForPlayerRequest) returns (blotservice.v1beta1.GetGameSetForPlayerResponse);
     */
    getGameSetForPlayer(input: GetGameSetForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameSetForPlayerRequest, GetGameSetForPlayerResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetGameSetForPlayerRequest, GetGameSetForPlayerResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetGameSetsForPlayer(blotservice.v1beta1.GetGameSetsForPlayerRequest) returns (blotservice.v1beta1.GetGameSetsForPlayerResponse);
     */
    getGameSetsForPlayer(input: GetGameSetsForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameSetsForPlayerRequest, GetGameSetsForPlayerResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetGameSetsForPlayerRequest, GetGameSetsForPlayerResponse>("unary", this._transport, method, opt, input);
    }
}
