import {defineStore} from 'pinia';
import gameSetRemoteRepository from '@/repo/repositores';
import {GameSet} from '@/models/gameSet';
import {useUserStore} from './userStore';

class GameSetLocalRepository {
    private key: string = 'myGameSets';

    constructor() {
    }

    public save(gameSetsIds: Array<string>) {
        localStorage.setItem(this.key, JSON.stringify(gameSetsIds));
    }

    public load(): Array<string> {
        const savedGameSets = localStorage.getItem(this.key);
        if (savedGameSets) {
            return JSON.parse(savedGameSets);
        }
        return [];
    }

    public add(id: string): void {
        const gameSetsIds = this.load();
        gameSetsIds.push(id);
        this.save(gameSetsIds);
    }

    public delete(id: string): void {
        const gameSetsIds = this.load();
        const index = gameSetsIds.indexOf(id);
        if (index > -1) {
            gameSetsIds.splice(index, 1);
        }
        this.save(gameSetsIds);
    }
}

const gameSetLocalRepository = new GameSetLocalRepository();

// Store saves ids in local storage and fetches game sets from remote API
export const useMyGameSetsStore = defineStore('myGameSets', {
    state: () => ({
        // Games where the player is a participant
        myGameSets: [] as Array<GameSet>
    }),
    getters: {
        findGameSet: (state) => {
            return (id: string) => state.myGameSets.find(gameSet => gameSet.id === id);
        },
    },
    actions: {
        async loadGameSets() {
            const gameSetsIds = gameSetLocalRepository.load();
            const userStore = useUserStore();
            const playerId = userStore.userId;
            const rawGameSets = await Promise.all(gameSetsIds.map(async (id) => {
                try {
                    return await gameSetRemoteRepository.get(id, playerId);
                } catch (e) {
                    console.error('Failed to load game set', e);
                    return null;
                }
            }));
            // TODO remove from local storage game sets that do not exist

            this.myGameSets = rawGameSets.filter((gameSet) => gameSet !== null) as GameSet[];
        },
        addGameSet(gameSet: GameSet) {
            // Check if user is participant in game set
            if (this.findGameSet(gameSet.id)) {
                throw new Error('Game set already exists');
            }
            this.myGameSets.push(gameSet);
            gameSetLocalRepository.add(gameSet.id);
        },
        async createGameSet(id: string) {
            if (this.findGameSet(id)) {
                throw new Error('Game set already exists');
            }
            const userStore = useUserStore();
            if (!userStore.user) {
                throw new Error('User not found');
            }
            await gameSetRemoteRepository.create(id, userStore.user);
            gameSetLocalRepository.add(id);

            const gameSet = await gameSetRemoteRepository.get(id, userStore.userId);
            this.myGameSets.push(gameSet);
        },
        deleteGameSet(gameSetId: string) {
            this.myGameSets = this.myGameSets.filter(gameSet => gameSet.id !== gameSetId);
            // TODO if is an owner, delete game set from remote
            gameSetLocalRepository.delete(gameSetId);
        },
        clearGameSets() {
            this.myGameSets = [];
            gameSetLocalRepository.save([]);
        },
    },
});
