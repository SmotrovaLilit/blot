<template>
  <v-container>
    <v-card>
      <v-card-title>Game Details</v-card-title>
      <v-card-text v-if="gameSet">
        <v-list>
          <v-list-item>
            <v-list-item-title>Game ID: {{ gameSet.id }}</v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>First Player: {{ gameSet.firstPlayerName }}</v-list-item-title>
          </v-list-item>
<!--          <v-list-item>-->
<!--            <v-list-item-title>Second Player: {{ game.secondPlayer.name }}</v-list-item-title>-->
<!--          </v-list-item>-->
          <v-list-item>
            <v-list-item-title>Status: {{ gameSet.status }}</v-list-item-title>
          </v-list-item>
          <v-list-item>
<!--            <v-list-item-title>Created At: {{ gameSet.createdAt }}</v-list-item-title>-->
          </v-list-item>
        </v-list>
        <v-btn color="primary" @click="goBack">Back to Games List</v-btn>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useGameSetsStore } from '@/stores/gameSetsStore';
import { storeToRefs } from 'pinia';


export default defineComponent({
  setup() {
    const route = useRoute();
    const router = useRouter();
    const gameSetsStore = useGameSetsStore();

    const { gameSet } = storeToRefs(gameSetsStore);

    onMounted(() => {
      gameSetsStore.loadGameSets();
      gameSetsStore.fetchGameSet(route.params.gameSetId);
      console.log(gameSet);
    });

    const goBack = () => {
      router.push({ name: 'gameSets' });
    };

    return {
      gameSet,
      goBack,
    };
  },
});
</script>

<style scoped>
</style>
