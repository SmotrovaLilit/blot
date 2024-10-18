// gameSet.ts

export class GameSet {
    id: string;
    firstPlayer: string;
    status: GameSetStatus;
}

export enum GameSetStatus {
    GAME_SET_STATUS_WAITED_FOR_PLAYERS = "waiting_for_players",
}