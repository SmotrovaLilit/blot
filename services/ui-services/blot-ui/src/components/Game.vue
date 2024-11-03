<template>
  <div>
    <h1 v-if="user">Current player = {{ user.name }}</h1>
    <!--    <h1 v-if="game">Team = {{ game.currentPlayer.team_id }}</h1>-->
  </div>
  <div v-if="game">
    <div class="game-stat">
      <Bet :existingBet="game.bet" :set-bet="setBet"/>
      <GameStatistics :rounds="game.rounds"/>
      <div class="round">
        <!--        R {{ game.round.number }}-->
      </div>
      <!--      <div v-if="betTeam" class="bet">-->
      <!--        {{ betTeam.name }} ({{ game.bet.amount }})-->
      <!--      </div>-->
    </div>
    <div v-if="!game.isFinished" class="player-container top-player-container"
         :class="{ selected: game.allyPlayer.isCurrentTurn }">
      <div class="player-name">{{ game.allyPlayer.name }}</div>
      <div class="players-cards">
        <ClosedDeck :cards-count="game.allyPlayer.handCards.length"/>
      </div>
    </div>
    <div v-if="!game.isFinished"  class="player-container left-player-container"
         :class="{ selected: game.leftPlayer.isCurrentTurn }">
      <div class="player-name">{{ game.leftPlayer.name }}</div>
      <div class="players-cards">
        <ClosedDeck :cards-count="game.leftPlayer.handCards.length"/>
      </div>
    </div>
    <div v-if="!game.isFinished"  class="player-container right-player-container"
         :class="{ selected: game.rightPlayer.isCurrentTurn }">
      <div class="player-name">{{ game.rightPlayer.name }}</div>
      <div class="players-cards">
        <ClosedDeck :cards-count="game.rightPlayer.handCards.length"/>
      </div>
    </div>
    <div class="game-container">
      <div class="game-top">

      </div>
      <div class="game-middle">
        <div class="left-bar">
        </div>
        <div class="middle">
          <div v-if="!game.isFinished"   class="game-table" >
            <div class="table-cards">
              <TableDeck :cards="game.tableCards"/>
            </div>
          </div>
        </div>
        <div class="right-bar">
        </div>
      </div>
      <div class="game-bottom">
        <div v-if="!game.isFinished" class="your-cards" :class="{ selected: game.currentPlayer.isCurrentTurn }">
          <HandDeck :cards="game.currentPlayer.handCards"
                    :is-your-turn="game.currentPlayer.isCurrentTurn"
                    :play-card="playCard"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import HandDeck from "@/components/HandDeck.vue";
import ClosedDeck from "@/components/ClosedDeck.vue";
import {useRoute, useRouter} from "vue-router";
import {useUserStore} from "@/stores/userStore";
import {computed, onMounted, ref} from "vue";
import {
  Card,
  type GameSet,
  GameSetStatus,
  GameStatus,
  PlayerState
} from "@/models/gameSet";
import type {User} from "@/models/user";
import gameSetRemoteRepository from "@/repo/repositores";
import Bet from "@/components/Bet.vue";
import TableDeck from "@/components/TableDeck.vue";
import GameStatistics from "@/components/GameStatistics.vue";

const route = useRoute();
const router = useRouter();
const userStore = useUserStore();

const gameSetId = String(route.params.gameSetId)
const gameSet = ref<GameSet>();
const user = computed(() => userStore.user as User);
const game = computed(() => {
  if (!gameSet.value) return null;
  const currentPlayerPosition = gameSet.value.game?.playerStates.findIndex(p => p.playerId === userStore.userId);
  const leftPlayerPosition = (currentPlayerPosition + 1) % 4;
  const rightPlayerPosition = (currentPlayerPosition + 3) % 4;
  const allyPlayerPosition = (currentPlayerPosition + 2) % 4;
  const g = {
    currentPlayer: enrichPlayer(gameSet.value.game?.playerStates[currentPlayerPosition]),
    leftPlayer: enrichPlayer(gameSet.value.game?.playerStates[leftPlayerPosition]),
    rightPlayer: enrichPlayer(gameSet.value.game?.playerStates[rightPlayerPosition]),
    allyPlayer: enrichPlayer(gameSet.value.game?.playerStates[allyPlayerPosition]),
    bet: enrichBet(gameSet.value),
    currentTurn: gameSet.value.game?.currentTurnPlayerId,
    tableCards: enrichTableCards(gameSet.value),
    rounds: enrichRounds(gameSet.value),
    isFinished: gameSet.value.game?.status === GameStatus.GAME_STATUS_FINISHED,
  };
  console.log("game", g);
  return g;
});

interface Player {
  id: string;
  name: string;
  handCards: Card[];
  isCurrentTurn: boolean;
}

const enrichRounds = (gameSet: GameSet) => {
  return gameSet.game?.rounds?.map(r => {
    let winnerTeam = '';
    if (r.winnerPlayerId) {
      if (r.winnerPlayerId === '00000000-0000-0000-0000-000000000000') { // TODO fix this weird id
        winnerTeam = 'None';
      } else {
        const id = findTeamId(gameSet, r.winnerPlayerId)
        winnerTeam = calculateTeamName(gameSet, id)
      }
    }
    return {
      number: r.number,
      winner: gameSet.players?.find(p => p.id === r.winnerPlayerId)?.name || '',
      winnerTeam: winnerTeam,
      score: r.score,
      table: r.table.map(c => {
        return {
          card: c.card,
          player: {
            id: c.playerId,
            name: gameSet.players?.find(p => p.id === c.playerId)?.name || '',
          }
        };
      })
    };
  });
};

const findTeamId = (gameSet: GameSet, playerId: string) => {
  if (gameSet.game?.team1.player1 === playerId || gameSet.game?.team1.player2 === playerId) {
    return gameSet.game?.team1.id;
  }
  if (gameSet.game?.team2.player1 === playerId || gameSet.game?.team2.player2 === playerId) {
    return gameSet.game?.team2.id;
  }
  throw new Error("Team not found for player " + playerId);
};

const enrichBet = (gameSet: GameSet) => {
  return {
    ...gameSet.game?.bet,
    teamName: calculateTeamName(gameSet, gameSet.game?.bet?.teamId),
  }
};

const calculateTeamName = (gameSet, teamId) => {
  if (gameSet.game?.team1.id === teamId) {
    const p1 = gameSet.game?.team1.player1
    const p2 = gameSet.game?.team1.player2
    return gameSet.players?.find(p => p.id === p1)?.name + ' & ' + gameSet.players?.find(p => p.id === p2)?.name;
  }
  if (gameSet.game?.team2.id === teamId) {
    const p1 = gameSet.game?.team2.player1
    const p2 = gameSet.game?.team2.player2
    return gameSet.players?.find(p => p.id === p1)?.name + ' & ' + gameSet.players?.find(p => p.id === p2)?.name;
  }
  throw new Error("Team not found: " + teamId);
};

const enrichTableCards = (gameSet: GameSet) => {
  if (!gameSet.game) return [];
  if (!gameSet.game.rounds) return [];
  const currentRound = gameSet.game.rounds[gameSet.game.rounds.length - 1];
  if (currentRound.table.length == 0) {
    if (gameSet.game.rounds.length == 1) {
      return [];
    }
    return gameSet.game.rounds[gameSet.game.rounds.length - 2].table.map(c => {
      return {
        card: c.card,
        player: {
          id: c.playerId,
          name: gameSet.players?.find(p => p.id === c.playerId)?.name || '',
        }
      };
    });
  }
  return currentRound.table.map(c => {
    return {
      card: c.card,
      player: {
        id: c.playerId,
        name: gameSet.players?.find(p => p.id === c.playerId)?.name || '',
      }
    };
  });
};
const enrichPlayer = (playerState: PlayerState | undefined): Player => {
  if (playerState === undefined) {
    throw new Error("Player is undefined");
  }
  return {
    id: playerState.playerId,
    name: gameSet.value?.players?.find(p => p.id === playerState.playerId)?.name || '',
    handCards: playerState.handCards,
    isCurrentTurn: playerState.playerId === gameSet.value?.game?.currentTurnPlayerId,
  };
};

onMounted(async () => {
  const playerId: string = userStore.userId;
  if (!playerId) {
    // TODO push to login
    return;
  }
  gameSet.value = await gameSetRemoteRepository.get(gameSetId, playerId);
  if (gameSet.value.status != GameSetStatus.GAME_SET_STATUS_PLAYING) {
    await router.push({name: 'GameSet', params: {gameSetId: gameSetId}});
  }
  console.log("gameSet-game-onMounted", gameSet.value);
});

const setBet = async (trump: string, amount: number) => {
  if (!gameSet.value) return;
  await gameSetRemoteRepository.setBet(gameSetId, userStore.userId, amount, trump);
  gameSet.value = await gameSetRemoteRepository.get(gameSetId, userStore.userId);
};

const playCard = async (card: Card) => {
  if (!gameSet.value) return;
  await gameSetRemoteRepository.playCard(gameSetId, userStore.userId, card);
  gameSet.value = await gameSetRemoteRepository.get(gameSetId, userStore.userId);
};
</script>

<style scoped>
.game-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.game-top {
  display: flex;
  justify-content: center;
  height: 20%;
  //background-color: blue;
}

.game-middle {
  display: flex;
  flex-direction: row;
  height: 50%;
  //background-color: red;
}

.left-bar {
  width: 20%;
  //display: flex;
  //justify-content: center;
  //background-color: white;
}

.middle {
  width: 60%;
  display: flex;
  justify-content: center;
  padding: 20px;
  //background-color: yellow;
}

.right-bar {
  width: 20%;
  display: flex;
  justify-content: center;
  //background-color: green;
}

.game-bottom {
  display: flex;
  //position: fixed;
  //bottom: 0;
  //left: 0;
  //width: 100%;
  //height: 30%;
  padding-bottom: 30px;
  box-sizing: border-box;
  //background-color: gray;
}

.game-table {
  display: flex;
  justify-content: center;
  width: 35%;
  aspect-ratio: 1/1;
  border-radius: 50%; /* Make the div circular */
  background: radial-gradient(circle, #822624 0%, #4f0c09 100%);
  box-shadow: 0 20px 30px rgba(25, 24, 24, 0.5),
  inset 0 -10px 15px rgba(49, 48, 48, 0.4),
  inset 0 10px 15px rgba(255, 255, 255, 0.1);
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translateY(-50%) translateX(-50%);
  z-index: 1;
  //margin: 50px auto;
}

.table-cards {
  width: 100%;
  display: flex;
  justify-content: center;
}

.your-cards {
  width: 40%;
  display: flex;
  position: fixed;
  bottom: 0;
  padding-bottom: 60px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 2;
}

.player-container {
  width: 20%;
  position: fixed;
  z-index: 2;
}

.players-cards {
  width: 100%;
}

.left-player-container {
  left: 0;
  top: 50%;
  transform: translateY(-50%);
}

.left-player-container .players-cards {
  transform: rotate(90deg);
}

.top-player-container {
  top: 0;
  left: 50%;
  transform: translateX(-50%);
}

.top-player-container .players-cards {
  transform: rotate(-180deg);
}


.right-player-container {
  right: 0;
  top: 50%;
  transform: translateY(-50%);
}

.right-player-container .players-cards {
  transform: rotate(-90deg);
}

.player-name {
  position: absolute;
  top: 50%;
  z-index: 2;
  color: white;
  font-size: 20px;
  font-weight: bold;
  text-align: center;
  align-items: center;
}

.game-stat {
  position: fixed;
  top: 0;
  right: 50px;
  z-index: 2;
  text-align: center;
  align-items: center;
  padding: 20px;
  width: 10%;
  max-width: 100px;
  display: flex;
  flex-direction: column;
}

.game-stat .trump {
  width: 100%;
}

.game-stat .round {
  width: 100%;
  color: white;
  font-size: 14px;
  font-weight: bold;
}

.game-stat .bet {
  width: 100%;
  color: white;
  font-size: 14px;
  font-weight: bold;
}

.player-container.selected {
  border: 2px solid white;
}

.your-cards.selected {
  border: 2px solid white;
}
</style>