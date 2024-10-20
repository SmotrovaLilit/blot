<template>
  <v-container v-if="user">
    <v-card v-if="gameSet">
      <v-card-title>Game Details</v-card-title>
      <v-card-text>
        <v-list>
          <v-list-item>
            <v-list-item-title>Game ID: {{ gameSet.id }}</v-list-item-title>
          </v-list-item>
          <!--          <v-list-item>-->
          <!--            <v-list-item-title>First Player: {{-->
          <!--                gameSet.firstPlayerId === user.id ? 'You' : gameSet.firstPlayerId-->
          <!--              }}-->
          <!--            </v-list-item-title>-->
          <!--          </v-list-item>-->
          <v-list-item>
            <v-list-item-title v-if="gameSet.players">
              Players: {{
                gameSet.players.map((p) => {
                  return p.id === user.id ? 'You' : p.name
                }).join(', ')
              }}
            </v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>Status: {{ gameSet.status }}</v-list-item-title>
          </v-list-item>
          <v-list-item>
            <!--            <v-list-item-title>Created At: {{ gameSet.createdAt }}</v-list-item-title>-->
          </v-list-item>
        </v-list>
        <v-btn color="primary" @click="goBack">Back to Games List</v-btn>
      </v-card-text>
      <v-card-actions>
        <v-btn v-if="canJoinGame" color="primary" @click="joinGameSet">Join Game</v-btn>
        <v-alert v-if="joinGameError.length!=0" type="error">{{ joinGameError }}</v-alert>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
import {computed, onMounted, ref} from 'vue';
import {useRoute, useRouter} from 'vue-router';
import {useUserStore} from "@/stores/userStore";
import {User} from "@/models/user";
import gameSetRemoteRepository from "@/repo/repositores";
import {GameSet} from "@/models/gameSet";

const route = useRoute();
const router = useRouter();
const userStore = useUserStore();

const gameSetId = String(route.params.gameSetId)
const gameSet = ref<GameSet>();
const user = computed(() => userStore.user as User);

onMounted(async () => {
  const playerId: string = userStore.userId;
  if (!playerId) {
    // TODO push to login
    return;
  }
  gameSet.value = await gameSetRemoteRepository.get(gameSetId, playerId);
});

const joinGameError = ref<string[]>([]);

interface CanUserJoinGameResult {
  canJoin: boolean;
  error: string | null;
}

const canUserJoinGame = (): CanUserJoinGameResult => {
  const playerId = userStore.userId;
  if (!playerId) return { canJoin: false, error: 'User should be logged in to join a game' };
  if (!gameSet.value) return { canJoin: false, error: 'Game not found' };
  if (gameSet.value.players.length >= 4) return { canJoin: false, error: 'Game is full' };
  if (gameSet.value.players.find(p => p.id === playerId)) return { canJoin: false, error: 'You are already in the game' };

  return { canJoin: true, error: null };
};

const canJoinGame = computed(() => {
  const { canJoin } = canUserJoinGame();
  return canJoin;
});

const joinGameSet = async () => {
  joinGameError.value = [];

  const validation = canUserJoinGame();
  if (!validation.canJoin) {
    joinGameError.value = [validation.error!];
    throw new Error(validation.error!);
  }

  if (!userStore.user) {
    // TODO push to login
    return;
  }
  await gameSetRemoteRepository.join(gameSetId, userStore.user);
  gameSet.value = await gameSetRemoteRepository.get(gameSetId, userStore.userId);
};


const goBack = () => {
  router.push({name: 'gameSets'});
};


</script>

<style scoped>
</style>
