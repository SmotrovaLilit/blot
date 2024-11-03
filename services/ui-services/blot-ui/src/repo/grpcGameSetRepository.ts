// grpcClient.ts
import {BlotServiceClient} from '@/generated/blotservice/v1beta1/blotservice.client';
import {
    Card as CardResp,
    CreateGameSetRequest,
    Game as GameResp,
    GameSet as GameSetResp,
    GameSetStatus as GameSetStatusResp,
    GameStatus as GameStatusResp,
    GetGameSetForPlayerRequest,
    JoinGameSetRequest,
    LeaveGameSetRequest,
    Player as PlayerResp,
    PlayerStateInGame as PlayerStateResp,
    StartGameRequest,
    Team as TeamResp,
    Suit as SuitResp,
    Rank as RankResp,
    Bet as BetResp,
    Round as RoundResp,
    PlayedCard as PlayedCardResp, SetBetRequest,
} from '@/generated/blotservice/v1beta1/blotservice';
import {
    Bet,
    Card,
    Game,
    GameSet,
    GameSetStatus,
    GameStatus, PlayerCard,
    PlayerState, Round,
    Team
} from '@/models/gameSet';
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import {User} from "@/models/user";
// @ts-ignore
const PROXY_URL: string = import.meta.env.VITE_PROXY_URL || 'http://127.0.0.1:8080' // grpc-web-proxy
const TIMEOUT_MILLISECS: number = 5 * 1000

export interface GameSetRepository {
    get(id: string, playerId: string): Promise<GameSet>;

    getPlayerGameSets(playerId: string): Promise<GameSet[]>

    create(id: string, player: User): Promise<void>;

    startGame(id: string, playerId: string, gameId: string): Promise<void>;

    join(id: string, player: User): Promise<void>;

    setBet(id: string, playerId: string, amount: number, trump: string): Promise<void>;

    leave(id: string, playerId: string): Promise<void>;
}

export class GrpcGameSetRepository implements GameSetRepository {
    private client: BlotServiceClient;

    constructor() {
        this.client = new BlotServiceClient(new GrpcWebFetchTransport({baseUrl: PROXY_URL}));
    }

    public async create(id: string, player: User) {
        const request = CreateGameSetRequest.create();
        request.id = id;
        request.player_id = player.id;
        request.player_name = player.name;


        console.log('createGameSet started', request);

        await this.client.createGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });

        console.log('createGameSet ended');

        return;
    }

    public async getPlayerGameSets(playerId: string): Promise<GameSet[]> {
        const request = GetGameSetForPlayerRequest.create();
        request.player_id = playerId;
        console.log('getGameSetsForPlayer started', request);
        const {response} = await this.client.getGameSetsForPlayer(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        if (!response || !response.game_sets) {
            throw new Error('Empty response');
        }
        console.log('getGameSetsForPlayer ended', response.game_sets);
        return response.game_sets.map(convertToGameSet);
    }

    public async get(id: string, playerId: string): Promise<GameSet> {
        const request = GetGameSetForPlayerRequest.create();
        request.id = id;
        request.player_id = playerId;
        console.log('getGameSetForPlayer started', request);
        const {response} = await this.client.getGameSetForPlayer(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        if (!response || !response.game_set) {
            throw new Error('Empty response');
        }
        console.log('getGameSetForPlayer ended', response.game_set);
        return convertToGameSet(response.game_set!);
    }

    public async leave(id: string, playerId: string): Promise<void> {
        const request = LeaveGameSetRequest.create();
        request.id = id;
        request.player_id = playerId;
        console.log('leaveGameSet started', request);
        await this.client.leaveGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        console.log('leaveGameSet ended');
    }

    public async join(id: string, player: User): Promise<void> {
        const request = JoinGameSetRequest.create();
        request.id = id;
        request.player_id = player.id;
        request.player_name = player.name;
        console.log('joinGameSet started', request);
        await this.client.joinGameSet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        // TODO: handle errors
        console.log('joinGameSet ended');
    }

    public async setBet(id: string, playerId: string, amount: number, trump: string): Promise<void> {
        const request = SetBetRequest.create();
        request.game_set_id = id;
        request.player_id = playerId;
        request.amount = amount;
        request.trump = convertToSuitResp(trump);
        console.log('setBet started', request);
        await this.client.setBet(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        console.log('setBet ended');
    }

    public async startGame(id: string, playerId: string, gameId: string): Promise<void> {
        const request = StartGameRequest.create();
        request.game_set_id = id;
        request.game_id = gameId;
        request.player_id = playerId;
        console.log('startGame started', request);
        await this.client.startGame(request, {
            meta: {},
            timeout: TIMEOUT_MILLISECS,
        });
        console.log('startGame ended');
    }
}

function convertToGameSet(resp: GameSetResp): GameSet {
    const g = new GameSet(resp.id, resp.owner_id, convertToGameSetStatus(resp.status));

    g.setPlayers(convertToUsers(resp.players));

    if (resp.game) {
        g.setGame(convertToGame(resp.game));
    }
    console.log('convertToGameSet', g);
    return g;
}

function convertToUsers(players: PlayerResp[]): User[] {
    if (players) {
        return players.map(p => {
            return new User(p.id, p.name);
        });
    }
    return [];
}

function convertToGameSetStatus(status: GameSetStatusResp): GameSetStatus {
    switch (status) {
        case GameSetStatusResp.WAITED_FOR_PLAYERS:
            return GameSetStatus.GAME_SET_STATUS_WAITED_FOR_PLAYERS;
        case GameSetStatusResp.READY_TO_START:
            return GameSetStatus.GAME_SET_STATUS_READY_TO_START;
        case GameSetStatusResp.PLAYING:
            return GameSetStatus.GAME_SET_STATUS_PLAYING;
        default:
            throw new Error('Unknown game set status: ' + status);
    }
}

function convertToGameStatus(status: GameStatusResp): GameStatus {
    switch (status) {
        case GameStatusResp.BETTING:
            return GameStatus.GAME_STATUS_BETTING;
        case GameStatusResp.PLAYING:
            return GameStatus.GAME_STATUS_PLAYING;
        default:
            throw new Error('Unknown game status: ' + status);
    }
}

function convertToTeam(team?: TeamResp): Team {
    if (!team) {
        throw Error('Team is empty');
    }
    return new Team(team.id, team.player1, team.player2);
}

function convertToSuit(suit: SuitResp): string {
    switch (suit) {
        case SuitResp.CLUBS:
            return "clubs";
        case SuitResp.DIAMONDS:
            return "diamonds";
        case SuitResp.HEARTS:
            return "hearts";
        case SuitResp.SPADES:
            return "spades";
    }
    throw new Error('Unknown suit: ' + suit);
}

function convertToRank(rank: RankResp): string {
    switch (rank) {
        case RankResp.ACE:
            return "a";
        case RankResp.SEVEN:
            return "7";
        case RankResp.EIGHT:
            return "8";
        case RankResp.NINE:
            return "9";
        case RankResp.TEN:
            return "10";
        case RankResp.JACK:
            return "j";
        case RankResp.QUEEN:
            return "q";
        case RankResp.KING:
            return "k";
    }
    throw new Error('Unknown rank: ' + rank);
}

function convertToCard(c: CardResp): Card {
    return new Card(convertToRank(c.rank), convertToSuit(c.suit));
}

function convertToPlayerState(state: PlayerStateResp): PlayerState {
    return new PlayerState(state.id, state.hand_cards.map(c => convertToCard(c)));
}

function convertToGame(game: GameResp): Game {
    const g = new Game(game.id, convertToGameStatus(game.status), convertToTeam(game.team1), convertToTeam(game.team2));
    g.playerStates = game.player_states.map(convertToPlayerState);
    g.bet = convertToBet(game.bet);
    g.rounds = game.rounds.map(convertToRound);
    g.currentTurnPlayerId = game.current_turn_player_id;
    return g;
}

function convertToBet(bet?: BetResp): Bet | undefined {
    if (!bet) {
        return undefined;
    }
    return new Bet(bet.team_id, bet.amount, convertToSuit(bet.trump));
}

function convertToRound(round: RoundResp): Round {
   const r = new Round(round.number);
   r.score = round.score;
   r.winnerPlayerId = round.winner_id;
   r.table = round.table_cards.map(convertToPlayerCard);
   return r;
}

function convertToPlayerCard(pCard: PlayedCardResp): PlayerCard {
    if (pCard.card == null) {
        throw new Error('Card is empty');
    }
    return new PlayerCard(pCard.player_id, convertToCard(pCard.card));
}

function convertToSuitResp(suit: string): SuitResp {
    switch (suit) {
        case "clubs":
            return SuitResp.CLUBS;
        case "diamonds":
            return SuitResp.DIAMONDS;
        case "hearts":
            return SuitResp.HEARTS;
        case "spades":
            return SuitResp.SPADES;
    }
    throw new Error('Unknown suit: ' + suit);
}