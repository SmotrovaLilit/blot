<template>
  <div>
    <v-btn color="primary" @click="openDialog">View Game Statistics</v-btn>

    <v-dialog v-model="dialog" max-width="700px">
      <v-card>
        <v-card-title class="headline">Game Statistics</v-card-title>
        <v-list>
          <v-list-item v-for="(teamScore, index) in teamsScore" :key="index">
            <v-list-item-title>{{ teamScore.team }}</v-list-item-title>
            <v-list-item-subtitle>{{ teamScore.score }}</v-list-item-subtitle>
          </v-list-item>
        </v-list>
        <v-card-text>
          <v-table>
            <thead>
            <tr>
              <th>Round</th>
              <th>Score</th>
              <th>Winner</th>
              <th>Winner Team</th>
              <th>Played cards</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(round, index) in displayedRounds" :key="index">
              <td>{{ round.number }}</td>
              <td>{{ round.score }}</td>
              <td>{{ round.winner }}</td>
              <td>{{ round.winnerTeam }}</td>
              <td>
                <v-row>
                  <v-col cols="3" v-for="(pCard, index) in round.table" :key="index">
                    <CardIcon :rank="pCard.card.rank" :suit="pCard.card.suit"
                              :player="pCard.player.name"/>
                  </v-col>
                </v-row>
              </td>
            </tr>
            </tbody>
          </v-table>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn @click="closeDialog">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts" setup>
import {computed, defineProps, ref} from 'vue';
import {
  VBtn,
  VCard,
  VCardActions,
  VCardText,
  VCardTitle,
  VDialog,
  VSpacer,
  VTable
} from 'vuetify/components';
import CardIcon from "@/components/CardIcon.vue";

interface PlayerCard {
  player: {
    id: string;
    name: string;
  }
  card: { rank: string; suit: string };
}

interface RoundStats {
  number: number;
  score: number;
  winner: string;
  winnerTeam: string;
  table: PlayerCard[];
}

interface Props {
  rounds: RoundStats[];
}

const props = defineProps<Props>();

// Transform the rounds data to be displayed in the table
const displayedRounds = computed(() => props.rounds.map((round) => {
  return {
    number: round.number,
    score: round.score,
    winner: round.winner,
    winnerTeam: round.winnerTeam,
    table: round.table,
  };
}));

interface TeamScore {
  team: string;
  score: number;
}

const teamsScore = computed(() => {
  const teamScores: TeamScore[] = [];
  props.rounds.forEach((round) => {
    const team = round.winnerTeam;
    const score = round.score;
    const teamScore = teamScores.find((ts) => ts.team === team);
    if (teamScore) {
      teamScore.score += score;
    } else {
      teamScores.push({team, score});
    }
  });
  return teamScores;
});

const dialog = ref(false);

const openDialog = () => {
  dialog.value = true;
};

const closeDialog = () => {
  dialog.value = false;
};
</script>

<style scoped>
/* Add custom styling if needed */
</style>
