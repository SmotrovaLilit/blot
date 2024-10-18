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
            <v-list-item-title>First Player: {{
                gameSet.firstPlayer
              }}
            </v-list-item-title>
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

<script setup lang="ts">
import {onMounted, ref} from 'vue';
import {useRoute, useRouter} from 'vue-router';
import {useGameSetsStore} from '@/stores/gameSetsStore';
import type {GameSet} from "@/models/gameSet";


const route = useRoute();
const router = useRouter();
const gameSetsStore = useGameSetsStore();
const gameSet = ref<GameSet | null>(null);


onMounted(async () => {
  const gameSetId = String(route.params.gameSetId)
  await gameSetsStore.loadGameSets();
  gameSet.value = await gameSetsStore.fetchGameSet(gameSetId);
  console.log(gameSet.value); // должен вывести объект GameSet
});

const goBack = () => {
  router.push({name: 'gameSets'});
};


</script>

<style scoped>
</style>
