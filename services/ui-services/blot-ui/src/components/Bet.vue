<template>
  <!-- Show existing bet -->
  <div class="trump">
    <div v-if="existingBet">
      <p><strong>Team ID:</strong> {{ existingBet.teamId }}</p>
      <p><strong>Amount:</strong> {{ existingBet.amount }}</p>
      <Icon :type="existingBet.trump"/>
      <div>{{existingBet.teamName}}</div>
    </div>
  </div>

  <!-- Button to open dialog for new bet -->
  <div style="margin-top: 20px;">
    <v-btn @click="dialog = true" color="primary" v-if="!betIsSet">Open Bet Form</v-btn>
  </div>

  <!-- Dialog for setting new bet -->
  <v-dialog v-model="dialog" max-width="500px">
    <v-card>
      <v-card-title class="headline">Set Bet</v-card-title>
      <v-card-text>
        <v-row class="trump-options">
          <v-col cols="3" class="trump-option" @click="selectTrump('spades')" :class="{ selected: trump === 'spades' }">
            <Icon type="spades"/>
          </v-col>
          <v-col cols="3" class="trump-option" @click="selectTrump('hearts')" :class="{ selected: trump === 'hearts' }">
            <Icon type="hearts"/>
          </v-col>
          <v-col cols="3" class="trump-option" @click="selectTrump('diamonds')" :class="{ selected: trump === 'diamonds' }">
            <Icon type="diamonds"/>
          </v-col>
          <v-col cols="3" class="trump-option" @click="selectTrump('clubs')" :class="{ selected: trump === 'clubs' }">
            <Icon type="clubs"/>
          </v-col>
        </v-row>
        <v-text-field
            label="Amount *"
            v-model="amount"
            type="number"
            outlined
            dense
            :min="8"
            :max="50"
            @input="validateAmount"
            style="margin-bottom: 10px"
            :error-messages="amountError"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn text @click="dialog = false">Cancel</v-btn>
        <v-btn
            color="primary"
            @click="handleSetBet"
            :disabled="!isBetFormValid"
        >
          Set Bet
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {computed, defineProps, ref} from 'vue';
import Icon from "@/components/Icon.vue";

interface Bet {
  teamId: string;
  amount: number;
  trump: string;
  teamName: string;
}

// Props
interface Props {
  existingBet?: Bet;
  setBet: (trump: string, amount: number) => void;
}

const props = defineProps<Props>();

// Dialog visibility
const dialog = ref(false);

// Reactive variables for setting a new bet
const trump = ref<string>('');
const amount = ref<number | null>(8);
const amountError = ref<string | null>(null);

// Computed property to check if a bet is already set
const betIsSet = computed(() => {
  console.log('Existing Bet:', props.existingBet);
  return props.existingBet
});

const validateAmount = () => {
  if (!isValidAmount(amount.value)) {
    amountError.value = 'Amount must be between 8 and 50';
  } else {
    amountError.value = null;
  }
};

function isValidAmount(amount: number | null): boolean {
  return amount !== null && amount >= 8 && amount <= 50;
}

function isValidTrump(trump: string): boolean {
  return ['spades', 'hearts', 'diamonds', 'clubs'].includes(trump);
}

// Computed property to determine if the bet is valid
const isBetFormValid = computed(() => {
  return isValidAmount(amount.value) && isValidTrump(trump.value);
});

// Function to select a trump
const selectTrump = (selectedTrump: string) => {
  trump.value = selectedTrump;
};

// Function to handle setting a new bet
const handleSetBet = () => {
  if (!betIsSet.value) {
    if (!isBetFormValid.value) {
      console.warn('Not all fields are filled correctly.');
      return;
    }
    console.log('New Bet:', {
      trump: trump.value,
      amount: amount.value,
    });
    console.log('Setting bet...');
    props.setBet(trump.value, amount.value);
    dialog.value = false; // Close the dialog after setting the bet
  } else {
    console.warn('Bet is already set. Cannot set a new bet.');
  }
};
</script>

<style scoped>
.trump-options{
  margin-top: 10px;
  margin-bottom: 10px;
}

.trump-option {
  max-width: 100%;
  cursor: pointer;
  padding: 5px;
}

.trump-option:hover {
  border: 1px solid;
  border-color: var(--card-border, #333333);
  border-radius: 8px;
  box-shadow: var(--card-shadow, 0 10px 10px rgba(255, 255, 255, 0.1));
  background-color: var(--card-background, white);
  opacity: 0.7;
}

.trump-option.selected {
  border: 1px solid;
  border-color: var(--card-border, #333333);
  border-radius: 8px;
  box-shadow: var(--card-shadow, 0 10px 10px rgba(255, 255, 255, 0.1));
  background-color: var(--card-background, white);
}
</style>
