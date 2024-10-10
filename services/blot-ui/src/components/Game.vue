<template>
  <div>
    <h1 v-if="game">Game = {{ game.id }}</h1>
    <p v-if="loading">Loading...</p>
    <p v-if="error">{{ error }}</p>
    <div v-if="game">
      <h2>Players:</h2>
      <ul>
        <li v-for="player in game.players" :key="player.id">{{ player.name }}</li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, onMounted, ref} from 'vue';
import {useGameStore} from '@/stores/gameStore.ts';
import {toRefs} from 'vue';


export default defineComponent({
  setup() {

    const store = useGameStore();
    const {game, loading, error} = toRefs(store); // Use toRefs for reactivity
    onMounted(() => {
      const gameId = '1';
      const playerId = '1';
      store.fetchGameForPlayer(gameId, playerId);

    });

    return {
      game: game,
      loading: loading,
      error: error,
    };
  },
});
</script>
