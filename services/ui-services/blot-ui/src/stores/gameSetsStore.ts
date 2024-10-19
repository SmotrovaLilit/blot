import {defineStore} from 'pinia';
import repository from '@/api/grpcRepository';
import {GameSet} from '@/models/gameSet';
import {useUserStore} from './userStore';

export const useGameSetsStore = defineStore('gameSets', {
    state: () => ({
        gameSets: [] as Array<GameSet>
    }),
    getters: {
        findGameSet: (state) => {
            return (id: string) => state.gameSets.find(gameSet => gameSet.id === id);
        },
    },
    actions: {
        async loadGameSets() {
            const savedGameSets = localStorage.getItem('gameSets');
            if (savedGameSets) {
                this.gameSets = JSON.parse(savedGameSets);
            }
        },
        async fetchGameSet(id: string): Promise<GameSet> {
            const userStore = useUserStore();
            const playerName: string = userStore.userName;

            const gameSet = await repository.getGameSetForPlayer(id, playerName);
            this.addOrUpdateGameSet(gameSet);
            return gameSet;
        },
        async createGameSet(id: string) {
            const userStore = useUserStore();
            const playerName = userStore.userName;
            await repository.createGameSet(id, playerName);
            await this.fetchGameSet(id);
        },
        addOrUpdateGameSet(gameSet: GameSet) {
            let find = false;
            this.gameSets.forEach((existingGameSet, index) => {
                if (existingGameSet.id === gameSet.id) {
                    this.gameSets[index] = gameSet;
                    find = true;
                }
            });
            if (!find) {
                this.gameSets.push(gameSet);
            }
            this.saveGameSetsToLocalStorage();
        },
        saveGameSetsToLocalStorage() {
            localStorage.setItem('gameSets', JSON.stringify(this.gameSets));
        },
        deleteGameSet(gameSetId: string) {
            this.gameSets = this.gameSets.filter(gameSet => gameSet.id !== gameSetId);
            this.saveGameSetsToLocalStorage();
        },
        clearGameSets() {
            this.gameSets = [];
            localStorage.removeItem('gameSets');
        },
    },
});
