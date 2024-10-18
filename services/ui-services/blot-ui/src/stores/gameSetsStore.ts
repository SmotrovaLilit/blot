import { defineStore } from 'pinia';
import apiClient from '@/api/grpcClient';
import { GameSet } from '@/models/gameSet';
import { useUserStore } from './userStore';

export const useGameSetsStore = defineStore('gameSets', {
    state: () => ({
        gameSets: new Map<string, GameSet>(),
    }),
    actions: {
        async loadGameSets() {
            const savedGameSets = localStorage.getItem('gameSets');
            if (savedGameSets) {
                const parsedGameSets: { [key: string]: GameSet } = JSON.parse(savedGameSets);
                this.gameSets = new Map(Object.entries(parsedGameSets));
            }
        },
        async fetchGameSet(id: string) {
            const userStore = useUserStore();
            const playerName = userStore.playerName;

            try {
                const gameSet = await apiClient.getGameSetForPlayer(id, playerName);
                this.addOrUpdateGameSet(gameSet);
            } catch (error) {
                console.error('Error fetching game set:', error);
            }
        },
        async createGameSet(id: string) {
            const userStore = useUserStore();
            const playerName = userStore.playerName;

            try {
                await apiClient.createGameSet(id, playerName);
                const newGameSet: GameSet = { id, firstPlayer: playerName};
                this.addGameSet(newGameSet);
            } catch (error) {
                console.error('Error creating game set:', error);
            }
        },
        addOrUpdateGameSet(gameSet: GameSet) {
            this.gameSets.set(gameSet.id, gameSet);
            this.saveGameSetsToLocalStorage();
        },
        addGameSet(gameSet: GameSet) {
            this.gameSets.set(gameSet.id, gameSet);
            this.saveGameSetsToLocalStorage();
        },
        saveGameSetsToLocalStorage() {
            const gameSetsArray = Array.from(this.gameSets.entries()).reduce((acc, [key, value]) => {
                acc[key] = value;
                return acc;
            }, {} as { [key: string]: GameSet });

            localStorage.setItem('gameSets', JSON.stringify(gameSetsArray));
        },
        clearGameSets() {
            this.gameSets.clear();
            localStorage.removeItem('gameSets');
        },
    },
});
