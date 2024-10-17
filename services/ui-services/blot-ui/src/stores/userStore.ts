import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        playerId: '1',
        playerName: 'Player 1',
    }),
    actions: {
        setPlayerId(id: string) {
            this.playerId = id;
        },
        setPlayerName(name: string) {
            this.playerName = name;
        }
    },
});