// stores/gameStore.ts
import {GrpcWebFetchTransport} from '@protobuf-ts/grpcweb-transport'
import {defineStore} from 'pinia';
import {
    Bet as BetInResponse,
    Card as CardInResponse,
    Game as GameInResponse,
    GameStatus as GameStatusInResponse,
    GetGameForPlayerRequest, type GetGameForPlayerResponse,
    Player as PlayerInResponse,
    PlayerCard as PlayerCardInResponse,
    Rank as RankInResponse,
    Round as RoundInResponse,
    RoundStatus as RoundStatusInResponse,
    Suit as SuitInResponse,
    Team as TeamInResponse,
} from '../generated/blotservice/v1beta1/blotservice'
import {BlotServiceClient} from '@/generated/blotservice/v1beta1/blotservice.client'
import {useUserStore} from './userStore';
// Assuming the response structure matches your ProtoBuf definitions
import {Bet, Card, Game, Player, PlayerCard, Round, Team} from '@/models/game';

// @ts-ignore
export const PROXY_URL: string = import.meta.env.VITE_PROXY_URL || 'http://127.0.0.1:8080' // grpc-web-proxy
const TIMEOUT_MILLISECS: number = 5 * 1000


// Function to convert the gRPC response to your model
function convertToGameModel(response: GetGameForPlayerResponse): Game {
    const responseGame = response.game;
    const convertCard = (card: CardInResponse): Card => ({
        rank: convertRank(card.rank),
        suit: convertSuit(card.suit),
    });

    const convertPlayerCard = (playerCard: PlayerCardInResponse): PlayerCard => ({
        player_id: playerCard.player_id,
        card: convertCard(playerCard.card),
    });

    const convertRound = (round: RoundInResponse): Round => ({
        number: round.number,
        table_cards: round.table_cards.map(convertPlayerCard),
        status: convertRoundStatus(round.status),
        current_player_id: round.current_player_id,
    });

    const convertTeam = (team: TeamInResponse): Team => ({
        id: team.id,
        name: team.name,
    });

    const convertBet = (bet: BetInResponse): Bet => ({
        trump: convertSuit(bet.trump),
        team_id: bet.team_id,
        amount: bet.amount,
    });

    const convertPlayer = (player: PlayerInResponse): Player => ({
        id: player.id,
        name: player.name,
        hand_cards: player.hand_cards.map(convertCard),
        discard_stack: player.discard_stack.map(convertCard),
        team_id: player.team_id,
    });

    const convertRank = (rank: RankInResponse): Rank => {
        switch (rank) {
            case RankInResponse.UNSPECIFIED:
                return 'unspecified';
            case RankInResponse.ACE:
                return 'a';
            case RankInResponse.SEVEN:
                return '7';
            case RankInResponse.EIGHT:
                return '8';
            case RankInResponse.NINE:
                return '9';
            case RankInResponse.TEN:
                return '10';
            case RankInResponse.JACK:
                return 'j';
            case RankInResponse.QUEEN:
                return 'q';
            case RankInResponse.KING:
                return 'k';
            default:
                return 'unspecified';
        }
    };

    const convertSuit = (suit: SuitInResponse): Suit => {
        switch (suit) {
            case SuitInResponse.UNSPECIFIED:
                return 'unspecified';
            case SuitInResponse.HEARTS:
                return 'hearts';
            case SuitInResponse.DIAMONDS:
                return 'diamonds';
            case SuitInResponse.CLUBS:
                return 'clubs';
            case SuitInResponse.SPADES:
                return 'spades';
            default:
                return 'unspecified';
        }
    };

    const convertRoundStatus = (status: RoundStatusInResponse): RoundStatus => {
        switch (status) {
            case RoundStatusInResponse.UNSPECIFIED:
                return 'unspecified';
            case RoundStatusInResponse.STARTED:
                return 'started';
            case RoundStatusInResponse.FINISHED:
                return 'finished';
            default:
                return 'unspecified';
        }
    };

    const convertGameStatus = (status: GameStatusInResponse): GameStatus => {
        switch (status) {
            case GameStatusInResponse.UNSPECIFIED:
                return 'unspecified';
            case GameStatusInResponse.CREATED:
                return 'created';
            case GameStatusInResponse.STARTED:
                return 'started';
            case GameStatusInResponse.TALKING:
                return 'talking';
            case GameStatusInResponse.BET_PLACED:
                return 'bet_placed';
            case GameStatusInResponse.FINISHED:
                return 'finished';
            default:
                return 'unspecified';
        }
    };

    // Convert the response Game object to your Game model
    return {
        id: responseGame.id,
        status: convertGameStatus(responseGame.status),
        leftPlayer: convertPlayer(response.left_player),
        rightPlayer: convertPlayer(response.right_player),
        allyPlayer: convertPlayer(response.ally_player),
        currentPlayer: convertPlayer(response.current_player),
        round: convertRound(responseGame.round),
        bet: convertBet(responseGame.bet),
        teams: responseGame.teams.map(convertTeam),
    };
}

export const useGameStore = defineStore('game', {
    state: () => ({
        game: null as Game | null,
        loading: false,
        error: null as string | null,
        proxyURL: PROXY_URL,
        client: new BlotServiceClient(new GrpcWebFetchTransport({baseUrl: PROXY_URL})),
    }),

    getters: {
        tableCards: (state) => {
            return state.game?.round.table_cards.map((playerCard) => playerCard.card);
        },
        betTeam: (state) => {
            return state.game?.teams.find((team) => team.id === state.game?.bet?.team_id);
        }
    },

    actions: {
        async fetchGameForPlayer(gameId: string, playerId: string) {
            this.loading = true;
            this.error = null;

            const request: GetLogsRequest = GetGameForPlayerRequest.create();
            request.game_id = gameId;
            request.player_id = playerId;

            try {
                const {response} = await this.client.getGameForPlayer(request, {
                    meta: {},
                    timeout: TIMEOUT_MILLISECS,
                });

                this.game = convertToGameModel(response);
                console.log('gRPC Response fetching game:', this.game);
            } catch (err) {
                console.error('Error fetching game:', err);
                this.error = 'Error fetching game data';
            } finally {
                this.loading = false;
            }
        },

        async fetchGame(gameId: string) {
            const userStore = useUserStore();
            await this.fetchGameForPlayer(gameId, userStore.playerId);
        },
    },
});

