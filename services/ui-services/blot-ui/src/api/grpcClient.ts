// grpcClient.ts
import { BlotServiceClient } from '@/generated/blotservice/v1beta1/blotservice.client.ts';
import {
    CreateGameSetRequest,
    GetGameSetForPlayerRequest, type GetGameSetForPlayerResponse
} from '@/generated/blotservice/v1beta1/blotservice.ts';
import { GameSet } from '@/models/gameSet';
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
// @ts-ignore
const PROXY_URL: string = import.meta.env.VITE_PROXY_URL || 'http://127.0.0.1:8080' // grpc-web-proxy
const TIMEOUT_MILLISECS: number = 5 * 1000

export interface ApiClient {
    createGameSet(id: string, currentUserName: string): Promise<void>;
    getGameSetForPlayer(id: string, currentUserName: string): Promise<GameSet>;
}

class GrpcClient implements ApiClient {
    private client: BlotServiceClient;

    constructor(baseUrl: string) {
        this.client = new BlotServiceClient(new GrpcWebFetchTransport({baseUrl: baseUrl}));
    }

    public async createGameSet(id: string, currentUserName: string): Promise<void> {
        const request = CreateGameSetRequest.create();
        request.id = id;
        request.first_player = currentUserName;
        return new Promise((resolve, reject) => {
            console.log('createGameSet started', request);
            this.client.createGameSet(request, {
                meta: {},
                timeout: TIMEOUT_MILLISECS,
            }, (err, response) => {
                console.log('createGameSet ended', err, response);
                if (err) {
                    reject(err);
                } else {
                    resolve();
                }
            });
        });
    }

    public async getGameSetForPlayer(id: string, currentUserName: string): Promise<GameSet> {
        const request = GetGameSetForPlayerRequest.create();
        request.id = id;
        request.player_name = currentUserName;

        return new Promise((resolve, reject) => {
            this.client.getGameSetForPlayer(request, {
                meta: {},
                timeout: TIMEOUT_MILLISECS,
            }, (err, response) => {
                if (err) {
                    reject(err);
                } else {
                    const data = response.toObject();
                    resolve(data.gameSet as GameSet);
                }
            });
        });
    }
}

const apiClient = new GrpcClient(PROXY_URL);
export default apiClient;
