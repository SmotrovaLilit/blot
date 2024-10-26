// gameSet.ts

import type {User} from "@/models/user";

export class GameSet {
    id: string;
    ownerId: string;
    status: GameSetStatus;
    players: User[] = [];
    game?: Game

    constructor(id: string, ownerId: string, status: GameSetStatus) {
        this.id = id;
        this.ownerId = ownerId;
        this.status = status;
    }

    setGame(game: Game) {
        this.game = game;
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

export class Game {
    id: string;
    status: GameStatus;
    team1: Team
    team2: Team
    playerStates: PlayerState[] = []

    constructor(id: string, status: GameStatus, team1: Team, team2: Team) {
        this.id = id;
        this.status = status;
        this.team1 = team1;
        this.team2 = team2;
    }
}

export enum GameStatus {
    GAME_STATUS_BETTING = "betting",
}

export class Team {
    player1: string
    player2: string

    constructor(player1: string, player2: string) {
        this.player1 = player1;
        this.player2 = player2;
    }
}

export class PlayerState {
    playerId: string
    handCards: Card[] = []

    constructor(playerId: string, handCards: Card[]) {
        this.playerId = playerId;
        this.handCards = handCards;
    }
}

export class Card {
    rank: string;
    suit: string;

    constructor(rank: string, suit: string) {
        this.rank = rank;
        this.suit = suit;
    }
}