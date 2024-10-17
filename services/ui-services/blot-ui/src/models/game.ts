// game.ts

export type RoundStatus = 'unspecified' | 'started' | 'finished';
export type GameStatus = 'unspecified' | 'created' | 'started' | 'talking' | 'bet_placed' | 'finished';
export type Rank = 'unspecified' | 'a' | '7' | '8' | '9' | '10' | 'j' | 'q' | 'k';
export type Suit = 'unspecified' | 'hearts' | 'diamonds' | 'clubs' | 'spades';

export class Card {
    constructor(public rank: Rank, public suit: Suit) {}
}

export class PlayerCard {
    constructor(public player_id: string, public card: Card) {}
}

export class Round {
    public table_cards: PlayerCard[] = [];

    constructor(public number: number, public status: RoundStatus, public current_player_id: string) {}
}

export class Team {
    constructor(public id: string, public name: string) {}
}

export class Bet {
    constructor(public trump: Suit, public team_id: string, public amount: number) {}
}

export class Player {
    public hand_cards: Card[] = [];
    public discard_stack: Card[] = [];

    constructor(public id: string, public name: string, public team_id: string) {}
}

export class Game {
    public round: Round;
    public bet?: Bet;
    public currentPlayer: Player;
    public leftPlayer: Player;
    public rightPlayer: Player;
    public allyPlayer: Player;

    constructor(public id: string, public status: GameStatus, public teams: Team[]) {
    }
}
