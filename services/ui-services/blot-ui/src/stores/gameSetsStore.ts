// src/stores/gamesStore.ts
import { defineStore } from 'pinia';
import {CreateGameSetRequest} from "@/generated/blotservice/v1beta1/blotservice";
import {BlotServiceClient} from "@/generated/blotservice/v1beta1/blotservice.client";
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import {PROXY_URL, TIMEOUT_MILLISECS} from "@/stores/gameStore";

export const useGameSetsStore = defineStore('gameSetsStore', {
    state: () => ({
        gameSets: [] as Array<any>,
        errorMessage: '' as string | null,
        client: new BlotServiceClient(new GrpcWebFetchTransport({baseUrl: PROXY_URL})),
        gameSet: {},
    }),
    actions: {
        initializeGameSets() {
            const savedGameSets = localStorage.getItem('gameSets');
            if (savedGameSets) {
                this.gameSets = JSON.parse(savedGameSets);
            }
        },
        saveGameSets() {
            localStorage.setItem('gameSets', JSON.stringify(this.gameSets));
        },
        async createGameSetAPI(Id: string, playerName: string) {
            this.loading = true;
            this.error = null;

            const request: CreateGameSetRequest = CreateGameSetRequest.create();
            request.id = Id;
            request.first_player = playerName;

            try {
                await this.client.createGameSet(request, {
                    meta: {},
                    timeout: TIMEOUT_MILLISECS,
                });

                console.log('createGameSet finished');
            } catch (err) {
                console.error('Error creating game set:', err);
                this.error = 'Error creating game set';
            } finally {
                this.loading = false;
            }
        },

        async createGameSet(Id, playerName :string) {
            await this.createGameSetAPI(Id, playerName);
            this.gameSets.push({
                id: Id,
                firstPlayerName: playerName,
                status: 'waiting_for_players',
            })
            this.saveGameSets();
        },

        async fetchGameSet(Id  :string) {
            this.gameSet = this.gameSets.find((gameSet) => gameSet.id === Id);
        },
    },
});
