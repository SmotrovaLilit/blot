// @generated by protobuf-ts 2.9.4 with parameter use_proto_field_name
// @generated from protobuf file "blotservice/v1beta1/blotservice.proto" (package "blotservice.v1beta1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message blotservice.v1beta1.GetGameForPlayerRequest
 */
export interface GetGameForPlayerRequest {
    /**
     * @generated from protobuf field: string game_id = 1;
     */
    game_id: string;
    /**
     * @generated from protobuf field: string player_id = 2;
     */
    player_id: string;
}
/**
 * @generated from protobuf message blotservice.v1beta1.GetGameForPlayerResponse
 */
export interface GetGameForPlayerResponse {
    /**
     * @generated from protobuf field: blotservice.v1beta1.Game game = 1;
     */
    game?: Game;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Player current_player = 2;
     */
    current_player?: Player;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Player ally_player = 3;
     */
    ally_player?: Player;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Player left_player = 4;
     */
    left_player?: Player;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Player right_player = 5;
     */
    right_player?: Player;
}
/**
 * @generated from protobuf message blotservice.v1beta1.Game
 */
export interface Game {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: blotservice.v1beta1.GameStatus status = 2;
     */
    status: GameStatus;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Round round = 3;
     */
    round?: Round;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Bet bet = 4;
     */
    bet?: Bet;
    /**
     * @generated from protobuf field: repeated blotservice.v1beta1.Team teams = 5;
     */
    teams: Team[];
}
/**
 * @generated from protobuf message blotservice.v1beta1.Bet
 */
export interface Bet {
    /**
     * @generated from protobuf field: blotservice.v1beta1.Suit trump = 1;
     */
    trump: Suit;
    /**
     * @generated from protobuf field: string team_id = 2;
     */
    team_id: string;
    /**
     * @generated from protobuf field: int32 amount = 3;
     */
    amount: number;
}
/**
 * @generated from protobuf message blotservice.v1beta1.Team
 */
export interface Team {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: string name = 2;
     */
    name: string;
}
/**
 * @generated from protobuf message blotservice.v1beta1.Player
 */
export interface Player {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * @generated from protobuf field: repeated blotservice.v1beta1.Card hand_cards = 4;
     */
    hand_cards: Card[];
    /**
     * @generated from protobuf field: repeated blotservice.v1beta1.Card discard_stack = 5;
     */
    discard_stack: Card[];
    /**
     * @generated from protobuf field: string team_id = 6;
     */
    team_id: string;
}
/**
 * @generated from protobuf message blotservice.v1beta1.Round
 */
export interface Round {
    /**
     * @generated from protobuf field: int32 number = 1;
     */
    number: number;
    /**
     * @generated from protobuf field: repeated blotservice.v1beta1.PlayerCard table_cards = 2;
     */
    table_cards: PlayerCard[];
    /**
     * @generated from protobuf field: blotservice.v1beta1.RoundStatus status = 3;
     */
    status: RoundStatus;
    /**
     * @generated from protobuf field: string current_player_id = 4;
     */
    current_player_id: string;
}
/**
 * @generated from protobuf message blotservice.v1beta1.PlayerCard
 */
export interface PlayerCard {
    /**
     * @generated from protobuf field: string player_id = 1;
     */
    player_id: string;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Card card = 2;
     */
    card?: Card;
}
/**
 * @generated from protobuf message blotservice.v1beta1.Card
 */
export interface Card {
    /**
     * @generated from protobuf field: blotservice.v1beta1.Rank rank = 1;
     */
    rank: Rank;
    /**
     * @generated from protobuf field: blotservice.v1beta1.Suit suit = 2;
     */
    suit: Suit;
}
/**
 * @generated from protobuf enum blotservice.v1beta1.RoundStatus
 */
export enum RoundStatus {
    /**
     * @generated from protobuf enum value: ROUND_STATUS_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: ROUND_STATUS_STARTED = 1;
     */
    STARTED = 1,
    /**
     * @generated from protobuf enum value: ROUND_STATUS_FINISHED = 2;
     */
    FINISHED = 2
}
/**
 * @generated from protobuf enum blotservice.v1beta1.GameStatus
 */
export enum GameStatus {
    /**
     * @generated from protobuf enum value: GAME_STATUS_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: GAME_STATUS_CREATED = 1;
     */
    CREATED = 1,
    /**
     * @generated from protobuf enum value: GAME_STATUS_STARTED = 2;
     */
    STARTED = 2,
    /**
     * @generated from protobuf enum value: GAME_STATUS_TALKING = 3;
     */
    TALKING = 3,
    /**
     * @generated from protobuf enum value: GAME_STATUS_BET_PLACED = 4;
     */
    BET_PLACED = 4,
    /**
     * @generated from protobuf enum value: GAME_STATUS_FINISHED = 5;
     */
    FINISHED = 5
}
/**
 * @generated from protobuf enum blotservice.v1beta1.Rank
 */
export enum Rank {
    /**
     * @generated from protobuf enum value: RANK_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: RANK_ACE = 1;
     */
    ACE = 1,
    /**
     * @generated from protobuf enum value: RANK_SEVEN = 2;
     */
    SEVEN = 2,
    /**
     * @generated from protobuf enum value: RANK_EIGHT = 3;
     */
    EIGHT = 3,
    /**
     * @generated from protobuf enum value: RANK_NINE = 4;
     */
    NINE = 4,
    /**
     * @generated from protobuf enum value: RANK_TEN = 5;
     */
    TEN = 5,
    /**
     * @generated from protobuf enum value: RANK_JACK = 6;
     */
    JACK = 6,
    /**
     * @generated from protobuf enum value: RANK_QUEEN = 7;
     */
    QUEEN = 7,
    /**
     * @generated from protobuf enum value: RANK_KING = 8;
     */
    KING = 8
}
/**
 * @generated from protobuf enum blotservice.v1beta1.Suit
 */
export enum Suit {
    /**
     * @generated from protobuf enum value: SUIT_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: SUIT_HEARTS = 1;
     */
    HEARTS = 1,
    /**
     * @generated from protobuf enum value: SUIT_DIAMONDS = 2;
     */
    DIAMONDS = 2,
    /**
     * @generated from protobuf enum value: SUIT_CLUBS = 3;
     */
    CLUBS = 3,
    /**
     * @generated from protobuf enum value: SUIT_SPADES = 4;
     */
    SPADES = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class GetGameForPlayerRequest$Type extends MessageType<GetGameForPlayerRequest> {
    constructor() {
        super("blotservice.v1beta1.GetGameForPlayerRequest", [
            { no: 1, name: "game_id", kind: "scalar", localName: "game_id", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "player_id", kind: "scalar", localName: "player_id", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<GetGameForPlayerRequest>): GetGameForPlayerRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.game_id = "";
        message.player_id = "";
        if (value !== undefined)
            reflectionMergePartial<GetGameForPlayerRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetGameForPlayerRequest): GetGameForPlayerRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string game_id */ 1:
                    message.game_id = reader.string();
                    break;
                case /* string player_id */ 2:
                    message.player_id = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetGameForPlayerRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string game_id = 1; */
        if (message.game_id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.game_id);
        /* string player_id = 2; */
        if (message.player_id !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.player_id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.GetGameForPlayerRequest
 */
export const GetGameForPlayerRequest = new GetGameForPlayerRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetGameForPlayerResponse$Type extends MessageType<GetGameForPlayerResponse> {
    constructor() {
        super("blotservice.v1beta1.GetGameForPlayerResponse", [
            { no: 1, name: "game", kind: "message", T: () => Game },
            { no: 2, name: "current_player", kind: "message", localName: "current_player", T: () => Player },
            { no: 3, name: "ally_player", kind: "message", localName: "ally_player", T: () => Player },
            { no: 4, name: "left_player", kind: "message", localName: "left_player", T: () => Player },
            { no: 5, name: "right_player", kind: "message", localName: "right_player", T: () => Player }
        ]);
    }
    create(value?: PartialMessage<GetGameForPlayerResponse>): GetGameForPlayerResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetGameForPlayerResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetGameForPlayerResponse): GetGameForPlayerResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* blotservice.v1beta1.Game game */ 1:
                    message.game = Game.internalBinaryRead(reader, reader.uint32(), options, message.game);
                    break;
                case /* blotservice.v1beta1.Player current_player */ 2:
                    message.current_player = Player.internalBinaryRead(reader, reader.uint32(), options, message.current_player);
                    break;
                case /* blotservice.v1beta1.Player ally_player */ 3:
                    message.ally_player = Player.internalBinaryRead(reader, reader.uint32(), options, message.ally_player);
                    break;
                case /* blotservice.v1beta1.Player left_player */ 4:
                    message.left_player = Player.internalBinaryRead(reader, reader.uint32(), options, message.left_player);
                    break;
                case /* blotservice.v1beta1.Player right_player */ 5:
                    message.right_player = Player.internalBinaryRead(reader, reader.uint32(), options, message.right_player);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetGameForPlayerResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* blotservice.v1beta1.Game game = 1; */
        if (message.game)
            Game.internalBinaryWrite(message.game, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* blotservice.v1beta1.Player current_player = 2; */
        if (message.current_player)
            Player.internalBinaryWrite(message.current_player, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* blotservice.v1beta1.Player ally_player = 3; */
        if (message.ally_player)
            Player.internalBinaryWrite(message.ally_player, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* blotservice.v1beta1.Player left_player = 4; */
        if (message.left_player)
            Player.internalBinaryWrite(message.left_player, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* blotservice.v1beta1.Player right_player = 5; */
        if (message.right_player)
            Player.internalBinaryWrite(message.right_player, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.GetGameForPlayerResponse
 */
export const GetGameForPlayerResponse = new GetGameForPlayerResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Game$Type extends MessageType<Game> {
    constructor() {
        super("blotservice.v1beta1.Game", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "status", kind: "enum", T: () => ["blotservice.v1beta1.GameStatus", GameStatus, "GAME_STATUS_"] },
            { no: 3, name: "round", kind: "message", T: () => Round },
            { no: 4, name: "bet", kind: "message", T: () => Bet },
            { no: 5, name: "teams", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Team }
        ]);
    }
    create(value?: PartialMessage<Game>): Game {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "";
        message.status = 0;
        message.teams = [];
        if (value !== undefined)
            reflectionMergePartial<Game>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Game): Game {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* blotservice.v1beta1.GameStatus status */ 2:
                    message.status = reader.int32();
                    break;
                case /* blotservice.v1beta1.Round round */ 3:
                    message.round = Round.internalBinaryRead(reader, reader.uint32(), options, message.round);
                    break;
                case /* blotservice.v1beta1.Bet bet */ 4:
                    message.bet = Bet.internalBinaryRead(reader, reader.uint32(), options, message.bet);
                    break;
                case /* repeated blotservice.v1beta1.Team teams */ 5:
                    message.teams.push(Team.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Game, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* blotservice.v1beta1.GameStatus status = 2; */
        if (message.status !== 0)
            writer.tag(2, WireType.Varint).int32(message.status);
        /* blotservice.v1beta1.Round round = 3; */
        if (message.round)
            Round.internalBinaryWrite(message.round, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* blotservice.v1beta1.Bet bet = 4; */
        if (message.bet)
            Bet.internalBinaryWrite(message.bet, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* repeated blotservice.v1beta1.Team teams = 5; */
        for (let i = 0; i < message.teams.length; i++)
            Team.internalBinaryWrite(message.teams[i], writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.Game
 */
export const Game = new Game$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Bet$Type extends MessageType<Bet> {
    constructor() {
        super("blotservice.v1beta1.Bet", [
            { no: 1, name: "trump", kind: "enum", T: () => ["blotservice.v1beta1.Suit", Suit, "SUIT_"] },
            { no: 2, name: "team_id", kind: "scalar", localName: "team_id", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "amount", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<Bet>): Bet {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.trump = 0;
        message.team_id = "";
        message.amount = 0;
        if (value !== undefined)
            reflectionMergePartial<Bet>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Bet): Bet {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* blotservice.v1beta1.Suit trump */ 1:
                    message.trump = reader.int32();
                    break;
                case /* string team_id */ 2:
                    message.team_id = reader.string();
                    break;
                case /* int32 amount */ 3:
                    message.amount = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Bet, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* blotservice.v1beta1.Suit trump = 1; */
        if (message.trump !== 0)
            writer.tag(1, WireType.Varint).int32(message.trump);
        /* string team_id = 2; */
        if (message.team_id !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.team_id);
        /* int32 amount = 3; */
        if (message.amount !== 0)
            writer.tag(3, WireType.Varint).int32(message.amount);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.Bet
 */
export const Bet = new Bet$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Team$Type extends MessageType<Team> {
    constructor() {
        super("blotservice.v1beta1.Team", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Team>): Team {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "";
        message.name = "";
        if (value !== undefined)
            reflectionMergePartial<Team>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Team): Team {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Team, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.Team
 */
export const Team = new Team$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Player$Type extends MessageType<Player> {
    constructor() {
        super("blotservice.v1beta1.Player", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "hand_cards", kind: "message", localName: "hand_cards", repeat: 1 /*RepeatType.PACKED*/, T: () => Card },
            { no: 5, name: "discard_stack", kind: "message", localName: "discard_stack", repeat: 1 /*RepeatType.PACKED*/, T: () => Card },
            { no: 6, name: "team_id", kind: "scalar", localName: "team_id", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Player>): Player {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.id = "";
        message.name = "";
        message.hand_cards = [];
        message.discard_stack = [];
        message.team_id = "";
        if (value !== undefined)
            reflectionMergePartial<Player>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Player): Player {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                case /* repeated blotservice.v1beta1.Card hand_cards */ 4:
                    message.hand_cards.push(Card.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated blotservice.v1beta1.Card discard_stack */ 5:
                    message.discard_stack.push(Card.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* string team_id */ 6:
                    message.team_id = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Player, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* repeated blotservice.v1beta1.Card hand_cards = 4; */
        for (let i = 0; i < message.hand_cards.length; i++)
            Card.internalBinaryWrite(message.hand_cards[i], writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* repeated blotservice.v1beta1.Card discard_stack = 5; */
        for (let i = 0; i < message.discard_stack.length; i++)
            Card.internalBinaryWrite(message.discard_stack[i], writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* string team_id = 6; */
        if (message.team_id !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.team_id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.Player
 */
export const Player = new Player$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Round$Type extends MessageType<Round> {
    constructor() {
        super("blotservice.v1beta1.Round", [
            { no: 1, name: "number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "table_cards", kind: "message", localName: "table_cards", repeat: 1 /*RepeatType.PACKED*/, T: () => PlayerCard },
            { no: 3, name: "status", kind: "enum", T: () => ["blotservice.v1beta1.RoundStatus", RoundStatus, "ROUND_STATUS_"] },
            { no: 4, name: "current_player_id", kind: "scalar", localName: "current_player_id", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Round>): Round {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.number = 0;
        message.table_cards = [];
        message.status = 0;
        message.current_player_id = "";
        if (value !== undefined)
            reflectionMergePartial<Round>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Round): Round {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 number */ 1:
                    message.number = reader.int32();
                    break;
                case /* repeated blotservice.v1beta1.PlayerCard table_cards */ 2:
                    message.table_cards.push(PlayerCard.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* blotservice.v1beta1.RoundStatus status */ 3:
                    message.status = reader.int32();
                    break;
                case /* string current_player_id */ 4:
                    message.current_player_id = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Round, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 number = 1; */
        if (message.number !== 0)
            writer.tag(1, WireType.Varint).int32(message.number);
        /* repeated blotservice.v1beta1.PlayerCard table_cards = 2; */
        for (let i = 0; i < message.table_cards.length; i++)
            PlayerCard.internalBinaryWrite(message.table_cards[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* blotservice.v1beta1.RoundStatus status = 3; */
        if (message.status !== 0)
            writer.tag(3, WireType.Varint).int32(message.status);
        /* string current_player_id = 4; */
        if (message.current_player_id !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.current_player_id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.Round
 */
export const Round = new Round$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PlayerCard$Type extends MessageType<PlayerCard> {
    constructor() {
        super("blotservice.v1beta1.PlayerCard", [
            { no: 1, name: "player_id", kind: "scalar", localName: "player_id", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "card", kind: "message", T: () => Card }
        ]);
    }
    create(value?: PartialMessage<PlayerCard>): PlayerCard {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.player_id = "";
        if (value !== undefined)
            reflectionMergePartial<PlayerCard>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PlayerCard): PlayerCard {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string player_id */ 1:
                    message.player_id = reader.string();
                    break;
                case /* blotservice.v1beta1.Card card */ 2:
                    message.card = Card.internalBinaryRead(reader, reader.uint32(), options, message.card);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: PlayerCard, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string player_id = 1; */
        if (message.player_id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.player_id);
        /* blotservice.v1beta1.Card card = 2; */
        if (message.card)
            Card.internalBinaryWrite(message.card, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.PlayerCard
 */
export const PlayerCard = new PlayerCard$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Card$Type extends MessageType<Card> {
    constructor() {
        super("blotservice.v1beta1.Card", [
            { no: 1, name: "rank", kind: "enum", T: () => ["blotservice.v1beta1.Rank", Rank, "RANK_"] },
            { no: 2, name: "suit", kind: "enum", T: () => ["blotservice.v1beta1.Suit", Suit, "SUIT_"] }
        ]);
    }
    create(value?: PartialMessage<Card>): Card {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.rank = 0;
        message.suit = 0;
        if (value !== undefined)
            reflectionMergePartial<Card>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Card): Card {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* blotservice.v1beta1.Rank rank */ 1:
                    message.rank = reader.int32();
                    break;
                case /* blotservice.v1beta1.Suit suit */ 2:
                    message.suit = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Card, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* blotservice.v1beta1.Rank rank = 1; */
        if (message.rank !== 0)
            writer.tag(1, WireType.Varint).int32(message.rank);
        /* blotservice.v1beta1.Suit suit = 2; */
        if (message.suit !== 0)
            writer.tag(2, WireType.Varint).int32(message.suit);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message blotservice.v1beta1.Card
 */
export const Card = new Card$Type();
/**
 * @generated ServiceType for protobuf service blotservice.v1beta1.BlotService
 */
export const BlotService = new ServiceType("blotservice.v1beta1.BlotService", [
    { name: "GetGameForPlayer", options: {}, I: GetGameForPlayerRequest, O: GetGameForPlayerResponse }
]);
