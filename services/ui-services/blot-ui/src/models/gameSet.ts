// gameSet.ts

export class GameSet {
    id: string;
    firstPlayer: string;
    status: GameSetStatus;

    constructor(id: string, firstPlayer: string, status: GameSetStatus) {
        this.id = id;
        this.firstPlayer = firstPlayer;
        this.status = status;
    }
}

export enum GameSetStatus {
    GAME_SET_STATUS_WAITED_FOR_PLAYERS = "waiting_for_players",
}