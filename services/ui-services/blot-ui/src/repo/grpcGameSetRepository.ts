// grpcClient.ts
import {BlotServiceClient} from '@/generated/blotservice/v1beta1/blotservice.client';
import {
    CreateGameSetRequest,
    GameSet as GameSetResp,
    GameSetStatus as GameSetStatusResp,
    GetGameSetForPlayerRequest,
    Player as PlayerResp
} from '@/generated/blotservice/v1beta1/blotservice';
import {GameSet, GameSetStatus} from '@/models/gameSet';
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import {User} from "@/models/user";
// @ts-ignore
const PROXY_URL: string = import.meta.env.VITE_PROXY_URL || 'http://127.0.0.1:8080' // grpc-web-proxy
const TIMEOUT_MILLISECS: number = 5 * 1000

export interface GameSetRepository {
    create(id: string, player: User): Promise<void>;
    get(id: string, playerId: string): Promise<GameSet>;
}

export class GrpcGameSetRepository implements GameSetRepository {
    private client: BlotServiceClient;

    constructor() {
        this.client = new BlotServiceClient(new GrpcWebFetchTransport({baseUrl: PROXY_URL}));
    }

    public async create(id: string, player: User) {
        const request = CreateGameSetRequest.create();
        request.id = id;
        request.first_player_id = player.id;
        request.first_player_name = player.name;


        console.log('createGameSet started', request);

        await this.client.createGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });

        console.log('createGameSet ended');

        return;
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
}

function convertToGameSet(resp: GameSetResp): GameSet {
    const g = new GameSet(resp.id, resp.first_player_id, convertToGameSetStatus(resp.status));

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
        default:
            throw new Error('Unknown game set status');
    }
}
