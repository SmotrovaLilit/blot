<template>
  <v-container>
    <v-card class="pa-5" max-width="500">
      <h2 class="text-h5">Create a New Game</h2>
      <v-card-text>
        <v-form @submit.prevent="submitForm">
          <v-text-field
              v-model="playerName"
              label="Enter your name"
              required
              :error-messages="nameError"
          />
          <v-btn type="submit" color="primary">Create Game</v-btn>
        </v-form>
        <v-alert v-if="errorMessage" type="error" dismissible>{{ errorMessage }}</v-alert>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { defineComponent } from 'vue';
import { useGameSetsStore } from '@/stores/gameSetsStore';
import { v4 as uuidv4 } from 'uuid';

export default defineComponent({
  setup() {
    const playerName = ref<string>('');
    const nameError = ref<string | null>(null);
    const errorMessage = ref<string | null>(null);
    const router = useRouter();
    const gameSetsStore = useGameSetsStore();

    const submitForm = async () => {
      if (!playerName.value) {
        nameError.value = 'Name is required';
        return;
      }
      nameError.value = null;

      try {
        const Id = uuidv4()
        await gameSetsStore.createGameSet(Id, playerName.value);
        await router.push({ name: 'gameSet', params: { gameSetId: Id } });
      } catch (error) {
        console.error(error);
      }
    };

    return {
      playerName,
      nameError,
      submitForm,
      errorMessage
    };
  },
});
</script>

<style scoped>
.v-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.v-btn {
  margin-top: 20px;
}
</style>
