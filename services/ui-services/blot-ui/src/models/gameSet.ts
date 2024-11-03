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
    bet?: Bet
    rounds: Round[] = []
    currentTurnPlayerId?: string

    constructor(id: string, status: GameStatus, team1: Team, team2: Team) {
        this.id = id;
        this.status = status;
        this.team1 = team1;
        this.team2 = team2;
    }
}

export enum GameStatus {
    GAME_STATUS_BETTING = "betting",
    GAME_STATUS_PLAYING = "playing",
    GAME_STATUS_FINISHED = "finished",
}

export class Team {
    id: string
    player1: string
    player2: string

    constructor(id: string, player1: string, player2: string) {
        this.id = id;
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

export class Bet {
    teamId: string
    amount: number
    trump: string

    constructor(teamId: string, amount: number, trumpSuit: string) {
        this.teamId = teamId;
        this.amount = amount;
        this.trump = trumpSuit;
    }
}

export class Round {
    number: number
    table: PlayerCard[] = []
    winnerPlayerId?: string
    score?: number

    constructor(number: number) {
        this.number = number;
    }
}

export class PlayerCard {
    playerId: string
    card: Card

    constructor(playerId: string, card: Card) {
        this.playerId = playerId;
        this.card = card;
    }
}
