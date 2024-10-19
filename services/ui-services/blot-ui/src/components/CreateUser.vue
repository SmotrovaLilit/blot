<template>
  <v-form @submit.prevent="submitForm">
    <v-text-field
        v-model="userName"
        label="Enter your name"
        required
        :error-messages="nameError"
    />
    <v-btn type="submit" color="primary">Set your name</v-btn>
  </v-form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUserStore } from '@/stores/userStore';

const userStore = useUserStore();

const userName = ref('');
const nameError = ref<string[]>([]);

const submitForm = async () => {
  if (!userName.value) {
    nameError.value = ['Name is required'];
    return;
  }
  try {
    await userStore.createUser(userName.value);
    nameError.value = [];
    window.location.reload();
  } catch (error) {
    nameError.value = [error.message || 'Failed to set name'];
  }
};
</script>

<style scoped>

</style>
