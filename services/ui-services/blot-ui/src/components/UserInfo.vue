<template>
  <v-card v-if="userName">
    <v-card-title>
      <h1>Welcome, {{ userName }}</h1>
    </v-card-title>
  </v-card>
  <v-card v-else>
    <CreateUser/>
  </v-card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useUserStore } from '@/stores/userStore';
import CreateUser from "@/components/CreateUser.vue";

const userStore = useUserStore();

const userName = computed(() => userStore.userName);

</script>

<style scoped>
</style>
