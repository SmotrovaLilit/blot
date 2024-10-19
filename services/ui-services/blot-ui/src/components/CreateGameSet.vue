<template>
  <v-form @submit.prevent="submitForm">
    <v-btn type="submit" color="primary">Create Game</v-btn>
  </v-form>
</template>

<script setup lang="ts">
import {useRouter} from 'vue-router';
import {useMyGameSetsStore} from '@/stores/myGameSetsStore';
import {v4 as uuidv4} from 'uuid';

const router = useRouter();
const gameSetsStore = useMyGameSetsStore();

const submitForm = async () => {
  const Id = uuidv4()
  await gameSetsStore.createGameSet(Id);
  await router.push({name: 'gameSet', params: {gameSetId: Id}});
};

</script>
<style scoped>

</style>
