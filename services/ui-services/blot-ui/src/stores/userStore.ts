import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        playerId: '1',
    }),
    actions: {
        setPlayerId(id: string) {
            this.playerId = id;
        },
    },
});