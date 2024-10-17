<template>
  <div>
    <h1 v-if="game">Current player = {{ game.currentPlayer.name }}</h1>
    <h1 v-if="game">Team = {{ game.currentPlayer.team_id }}</h1>
    <p v-if="loading">Loading...</p>
    <p v-if="error">{{ error }}</p>
  </div>
  <div v-if="game">
    <div class="game-stat">
      <div class="trump">
        <Icon :type="game.bet.trump"/>
      </div>
      <div class="round">
        R {{ game.round.number }}
      </div>
      <div v-if="betTeam" class="bet">
        {{ betTeam.name }} ({{ game.bet.amount }})
      </div>
    </div>
    <div class="player-container top-player-container">
      <div class="player-name">{{ game.allyPlayer.name }}</div>
      <div class="players-cards">
        <ClosedDeck :cards-count="game.allyPlayer.hand_cards.length"/>
      </div>
    </div>
    <div class="player-container left-player-container">
      <div class="player-name">{{ game.leftPlayer.name }}</div>
      <div class="players-cards">
        <ClosedDeck :cards-count="game.leftPlayer.hand_cards.length"/>
      </div>
    </div>
    <div class="player-container right-player-container">
      <div class="player-name">{{ game.rightPlayer.name }}</div>
      <div class="players-cards">
        <ClosedDeck :cards-count="game.rightPlayer.hand_cards.length"/>
      </div>
    </div>
    <div class="game-container">
      <div class="game-top">

      </div>
      <div class="game-middle">
        <div class="left-bar">
        </div>
        <div class="middle">
          <div class="game-table">
            <div class="table-cards">
              <TableDeck :cards="tableCards"/>
            </div>
          </div>
        </div>
        <div class="right-bar">
        </div>
      </div>
      <div class="game-bottom">
        <div class="your-cards">
          <HandDeck :cards="game.currentPlayer.hand_cards"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, onMounted, watch} from 'vue';
import {storeToRefs} from 'pinia'
import HandDeck from "@/components/HandDeck.vue";
import TableDeck from "@/components/TableDeck.vue";
import ClosedCard from "@/components/ClosedCard.vue";
import ClosedDeck from "@/components/ClosedDeck.vue";
import {useGameStore} from '@/stores/gameStore.ts';
import {useRoute} from 'vue-router';
import Icon from "@/components/Icon.vue";

export default defineComponent({
  components: {
    Icon,
    ClosedDeck,
    ClosedCard,
    TableDeck,
    HandDeck,
  },
  setup() {
    const gameStore = useGameStore();
    const {
      game,
      loading,
      error,
      tableCards,
      betTeam,
    } = storeToRefs(gameStore);
    const route = useRoute();
    watch(game, (newGame) => {
      console.log('count', game.value.bet);
    });
    onMounted(() => {
      gameStore.fetchGame(route.params.gameId);

    });
    return {
      loading: loading,
      error: error,
      game: game,
      tableCards: tableCards,
      betTeam: betTeam,
    };
  }
});
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
</style>