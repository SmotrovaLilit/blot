<template>
  <v-container>
    <v-card>
      <v-card-title>My Games</v-card-title>
      <v-card-text>
        <v-list>
          <v-list-item v-for="gameSet in typedGameSets" :key="gameSet.id">
            <router-link :to="{ name: 'gameSet', params: { gameSetId: gameSet.id } }">
              <v-card class="game-card" outlined>
                <v-card-title>{{ gameSet.firstPlayer }}</v-card-title>
                <v-card-subtitle>ID: {{ gameSet.id }}</v-card-subtitle>
                <v-card-actions>
                  <v-btn color="primary">View Game</v-btn>
                </v-card-actions>
              </v-card>
            </router-link>
          </v-list-item>
        </v-list>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script setup lang="ts" >
import { onMounted, computed } from 'vue';
import { useGameSetsStore } from '@/stores/gameSetsStore';
import { storeToRefs } from 'pinia';
import type {GameSet} from "@/models/gameSet";

const gameSetsStore = useGameSetsStore();
const { gameSets } = storeToRefs(gameSetsStore);

const typedGameSets = computed(() => gameSets.value as GameSet[]);

onMounted(() => {
  gameSetsStore.loadGameSets();
});

</script>

<style scoped>
.game-card {
  margin: 10px 0;
}
</style>
