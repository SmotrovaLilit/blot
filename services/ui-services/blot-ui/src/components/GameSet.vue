<template>
  <v-container v-if="user">
    <v-card v-if="gameSet">
      <v-card-title>Game Details</v-card-title>
      <v-card-text>
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
import {computed, onMounted} from 'vue';
import {useRoute, useRouter} from 'vue-router';
import {useGameSetsStore} from '@/stores/gameSetsStore';
import {useUserStore} from "@/stores/userStore";

const route = useRoute();
const router = useRouter();
const gameSetsStore = useGameSetsStore();
const userStore = useUserStore();

const gameSetId = String(route.params.gameSetId)
const gameSet = computed(() => gameSetsStore.findGameSet(gameSetId));
const user = computed(() => userStore.userName);

onMounted(async () => {
  await gameSetsStore.loadGameSets();
  await gameSetsStore.fetchGameSet(gameSetId);
});

const goBack = () => {
  router.push({name: 'gameSets'});
};


</script>

<style scoped>
</style>
