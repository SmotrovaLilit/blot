<template>
  <v-form @submit.prevent="submitForm">
    <v-btn type="submit" color="primary">Create Game</v-btn>
  </v-form>
</template>

<script setup lang="ts">
import {useRouter} from 'vue-router';
import {v4 as uuidv4} from 'uuid';
import {useUserStore} from "@/stores/userStore";
import gameSetRemoteRepository from "@/repo/repositores";

const router = useRouter();
const userStore = useUserStore();

const submitForm = async () => {
  const Id = uuidv4()
  await gameSetRemoteRepository.create(Id, userStore.user);
  await router.push({name: 'gameSet', params: {gameSetId: Id}});
};

</script>
<style scoped>

</style>
