// grpcClient.ts
import {BlotServiceClient} from '@/generated/blotservice/v1beta1/blotservice.client';
import {
    CreateGameSetRequest,
    GameSet as GameSetResp,
    GameSetStatus as GameSetStatusResp,
    GetGameSetForPlayerRequest, JoinGameSetRequest, LeaveGameSetRequest,
    Player as PlayerResp
} from '@/generated/blotservice/v1beta1/blotservice';
import {GameSet, GameSetStatus} from '@/models/gameSet';
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import {User} from "@/models/user";
// @ts-ignore
const PROXY_URL: string = import.meta.env.VITE_PROXY_URL || 'http://127.0.0.1:8080' // grpc-web-proxy
const TIMEOUT_MILLISECS: number = 5 * 1000

export interface GameSetRepository {
    get(id: string, playerId: string): Promise<GameSet>;
    getPlayerGameSets(playerId: string): Promise<GameSet[]>

    create(id: string, player: User): Promise<void>;
    join(gameSetId: string, player: User): Promise<void>;
    leave(id: string, playerId: string): Promise<void>;
}

export class GrpcGameSetRepository implements GameSetRepository {
    private client: BlotServiceClient;

    constructor() {
        this.client = new BlotServiceClient(new GrpcWebFetchTransport({baseUrl: PROXY_URL}));
    }

    public async create(id: string, player: User) {
        const request = CreateGameSetRequest.create();
        request.id = id;
        request.player_id = player.id;
        request.player_name = player.name;


        console.log('createGameSet started', request);

        await this.client.createGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });

        console.log('createGameSet ended');

        return;
    }

    public async getPlayerGameSets(playerId: string): Promise<GameSet[]> {
        const request = GetGameSetForPlayerRequest.create();
        request.player_id = playerId;
        console.log('getGameSetsForPlayer started', request);
        const {response} = await this.client.getGameSetsForPlayer(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        if (!response || !response.game_sets) {
            throw new Error('Empty response');
        }
        console.log('getGameSetsForPlayer ended', response.game_sets);
        return response.game_sets.map(convertToGameSet);
    }

    public async get(id: string, playerId: string): Promise<GameSet> {
        const request = GetGameSetForPlayerRequest.create();
        request.id = id;
        request.player_id = playerId;
        console.log('getGameSetForPlayer started', request);
        const {response} = await this.client.getGameSetForPlayer(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        if (!response || !response.game_set) {
            throw new Error('Empty response');
        }
        console.log('getGameSetForPlayer ended', response.game_set);
        return convertToGameSet(response.game_set!);
    }

    public async leave(id: string, playerId: string): Promise<void> {
        const request = LeaveGameSetRequest.create();
        request.id = id;
        request.player_id = playerId;
        console.log('leaveGameSet started', request);
        await this.client.leaveGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        console.log('leaveGameSet ended');
    }
    public async join(gameSetId: string, player: User): Promise<void> {
        const request = JoinGameSetRequest.create();
        request.id = gameSetId;
        request.player_id = player.id;
        request.player_name = player.name;
        console.log('joinGameSet started', request);
        await this.client.joinGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        // TODO: handle errors
        console.log('joinGameSet ended');
    }
}

function convertToGameSet(resp: GameSetResp): GameSet {
    const g = new GameSet(resp.id, resp.owner_id, convertToGameSetStatus(resp.status));

    g.setPlayers(convertToUsers(resp.players));
    console.log('convertToGameSet', g);
    return g;
}

function convertToUsers(players: PlayerResp[]): User[] {
    if (players) {
        return players.map(p => {
            return new User(p.id, p.name);
        });
    }
    return [];
}

function convertToGameSetStatus(status: GameSetStatusResp): GameSetStatus {
    switch (status) {
        case GameSetStatusResp.WAITED_FOR_PLAYERS:
            return GameSetStatus.GAME_SET_STATUS_WAITED_FOR_PLAYERS;
        case GameSetStatusResp.READY_TO_START:
            return GameSetStatus.GAME_SET_STATUS_READY_TO_START;
        default:
            throw new Error('Unknown game set status');
    }
}
