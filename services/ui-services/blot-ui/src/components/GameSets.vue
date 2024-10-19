<template>
  <v-container v-if="userName">
    <v-card>
      <v-card-title>My Games</v-card-title>
      <v-card-text>
        <CreateGameSet/>
        <v-list>
            <v-list-item
                v-for="gameSet in typedGameSets"
                :key="gameSet.id"
                @click="goToDetails(gameSet.id)"
            >
              <v-list-item-title>{{ gameSet.id }}</v-list-item-title>

              <v-list-item-action>
                <v-btn icon @click.stop="deleteGameSet(gameSet.id)">
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
              </v-list-item-action>
            </v-list-item>
        </v-list>
        <!--        <v-list>-->
        <!--          <v-list-item v-for="gameSet in typedGameSets" :key="gameSet.id">-->
        <!--            <router-link :to="{ name: 'gameSet', params: { gameSetId: gameSet.id } }">-->
        <!--              <v-card class="game-card" outlined>-->
        <!--                <v-card-title>{{ gameSet.firstPlayer }}</v-card-title>-->
        <!--                <v-card-subtitle>ID: {{ gameSet.id }}</v-card-subtitle>-->
        <!--                <v-card-actions>-->
        <!--                  <v-btn color="primary">View Game</v-btn>-->
        <!--                </v-card-actions>-->
        <!--              </v-card>-->
        <!--            </router-link>-->
        <!--          </v-list-item>-->
        <!--        </v-list>-->
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
import {computed, onMounted} from 'vue';
import {useGameSetsStore} from '@/stores/gameSetsStore';
import {storeToRefs} from 'pinia';
import type {GameSet} from "@/models/gameSet";
import CreateGameSet from "@/components/CreateGameSet.vue";
import {useRouter} from 'vue-router';
import {useUserStore} from "@/stores/userStore";


const gameSetsStore = useGameSetsStore();
const userStore = useUserStore();
const {gameSets} = storeToRefs(gameSetsStore);

const typedGameSets = computed(() => gameSets.value as GameSet[]);
const router = useRouter();

const userName = computed(() => userStore.userName);

onMounted(() => {
  gameSetsStore.loadGameSets();
});

const goToDetails = (id: string) => {
  router.push({name: 'gameSet', params: {gameSetId: id}});
};

const deleteGameSet = (id: string) => {
  gameSetsStore.deleteGameSet(id);
};
</script>

<style scoped>
.game-card {
  margin: 10px 0;
}
</style>
