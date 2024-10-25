// gameSet.ts

import type {User} from "@/models/user";

export class GameSet {
    id: string;
    ownerId: string;
    status: GameSetStatus;
    players: User[] = [];

    constructor(id: string, ownerId: string, status: GameSetStatus) {
        this.id = id;
        this.ownerId = ownerId;
        this.status = status;
    }

    setPlayers(players: User[]) {
        this.players = players;
    }
}

export enum GameSetStatus {
    GAME_SET_STATUS_WAITED_FOR_PLAYERS = "waiting_for_players",
    GAME_SET_STATUS_READY_TO_START = "ready_to_start",
    GAME_SET_STATUS_PLAYING = "playing",
}