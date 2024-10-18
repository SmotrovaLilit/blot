// grpcClient.ts
import {BlotServiceClient} from '@/generated/blotservice/v1beta1/blotservice.client';
import {
    CreateGameSetRequest,
    GetGameSetForPlayerRequest,
    GameSet as GameSetResp,
} from '@/generated/blotservice/v1beta1/blotservice';
import {GameSet, GameSetStatus} from '@/models/gameSet';
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
// @ts-ignore
const PROXY_URL: string = import.meta.env.VITE_PROXY_URL || 'http://127.0.0.1:8080' // grpc-web-proxy
const TIMEOUT_MILLISECS: number = 5 * 1000

export interface Repository {
    createGameSet(id: string, currentUserName: string): Promise<void>;
    getGameSetForPlayer(id: string, currentUserName: string): Promise<GameSet>;
}

class GrpcRepository implements Repository {
    private client: BlotServiceClient;

    constructor(baseUrl: string) {
        this.client = new BlotServiceClient(new GrpcWebFetchTransport({baseUrl: baseUrl}));
    }

    public async createGameSet(id: string, player: string) {
        const request = CreateGameSetRequest.create();
        request.id = id;
        request.first_player = player;

        console.log('createGameSet started', request);

        await this.client.createGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });

        console.log('createGameSet ended');

        return;
    }

    public async getGameSetForPlayer(id: string, player: string): Promise<GameSet> {
        const request = GetGameSetForPlayerRequest.create();
        request.id = id;
        request.player_name = player;

        const {response} = await this.client.getGameSetForPlayer(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        if (!response || !response.game_set) {
            throw new Error('Empty response');
        }
        return convertToGameSet(response.game_set!);
    }
}

function convertToGameSet(resp: GameSetResp): GameSet {
    return new GameSet(resp.id, resp.first_player, GameSetStatus.GAME_SET_STATUS_WAITED_FOR_PLAYERS);
}

const repository = new GrpcRepository(PROXY_URL);
export default repository;
