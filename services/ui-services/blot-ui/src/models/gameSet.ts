// gameSet.ts

import type {User} from "@/models/user";

export class GameSet {
    id: string;
    firstPlayerId: string;
    status: GameSetStatus;
    players: User[] = [];

    constructor(id: string, firstPlayerId: string, status: GameSetStatus) {
        this.id = id;
        this.firstPlayerId = firstPlayerId;
        this.status = status;
    }

    setPlayers(players: User[]) {
        this.players = players;
    }
}

export enum GameSetStatus {
    GAME_SET_STATUS_WAITED_FOR_PLAYERS = "waiting_for_players",
}