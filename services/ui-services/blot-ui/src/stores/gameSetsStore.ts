import {defineStore} from 'pinia';
import repository from '@/api/grpcRepository';
import {GameSet} from '@/models/gameSet';
import {useUserStore} from './userStore';

export const useGameSetsStore = defineStore('gameSets', {
    state: () => ({
        gameSets: [] as Array<GameSet>
    }),
    actions: {
        async loadGameSets() {
            const savedGameSets = localStorage.getItem('gameSets');
            if (savedGameSets) {
                this.gameSets = JSON.parse(savedGameSets);
            }
        },
        async fetchGameSet(id: string): Promise<GameSet> {
            const userStore = useUserStore();
            const playerName = userStore.playerName;

            const gameSet = await repository.getGameSetForPlayer(id, playerName);
            this.addOrUpdateGameSet(gameSet);
            return gameSet;
        },
        async createGameSet(id: string) {
            const userStore = useUserStore();
            const playerName = userStore.playerName;
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
        clearGameSets() {
            this.gameSets = [];
            localStorage.removeItem('gameSets');
        },
    },
});
