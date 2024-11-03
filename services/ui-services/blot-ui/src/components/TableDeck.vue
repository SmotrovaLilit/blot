<template>
  <div id="round-deck-table" class="round-deck-table">
    <div class="round-deck">
      <draggable-resizable-vue
          v-for="(card) in displayedCards"
          :key="card.id"
          class="card-container"
          v-model:x="card.x"
          v-model:y="card.y"
          v-model:w="card.width"
          :resizable="false"
          parent="#round-deck-table"
          :z="card.zIndex"
      >
        <div :style="card.style">
          <div class="player-name">{{card.player.name}}</div>
          <CardView :rank="card.rank" :suit="card.suit"/>
        </div>
      </draggable-resizable-vue>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineProps, onBeforeUnmount, onMounted, ref, watch} from 'vue';
import DraggableResizableVue from 'draggable-resizable-vue3';
import CardView from './Card.vue';

interface PlayerCard {
  player: { id: string; name: string };
  card: { rank: string; suit: string };
}


interface Props {
  cards: { type: PlayerCard[], required: false, default: [] };
}

const props = defineProps<Props>()

function generateRandomNumberBetween(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function randomAngle() {
  return Math.floor(Math.random() * 20) - 10;
}

const displayedCards = ref<Array<{
  id: string;
  player: { id: string; name: string };
  rank: string;
  suit: string;
  style: {};
  x: number;
  y: number;
  zIndex: number;
  width: number;
}>>([]);

interface ParentContainer {
  width: number;
  height: number;
}

const parentContainer = ref<ParentContainer>({
  width: 0,
  height: 0,
});

const calculateCards = () => {
  console.log('props.cards', props.cards);
  if (!parentContainer.value.width || !parentContainer.value.height) {
    return [];
  }
  const parentWidth = parentContainer.value.width;
  const parentHeight = parentContainer.value.height;
  const cardWidth = 0.25 * parentWidth; // 20% of the parent container width
  const cardX = (parentWidth - cardWidth) / 2; // Center horizontally
  const cardY = (parentHeight - cardWidth) / 2; // Center vertically
  const tolerance = 0.10 * parentWidth; // 5% of the parent container width
  const previousCards = displayedCards.value;
  displayedCards.value = props.cards.map((pCard, index) => {
    const card = pCard.card;
    const id = `${card.rank}-${card.suit}`;
    const angle = randomAngle();
    if (previousCards != undefined) {
      const prevValues = previousCards.find((c) => c.id === id && c.rank === card.rank && c.suit === card.suit);
      if (prevValues) {
        return prevValues;
      }
    }
    return {
      id,
      player: {
        id: pCard.player.id,
        name: pCard.player.name,
      },
      rank: card.rank,
      suit: card.suit,
      style: {
        transform: `rotate(${angle}deg)`,
      },
      zIndex: index,
      x: card.x ? card.x : generateRandomNumberBetween(cardX - tolerance, cardX + tolerance), // Centered X position;
      y: card.y ? card.y : generateRandomNumberBetween(cardY - tolerance, cardY + tolerance / 2), // Centered Y position
      width: cardWidth,
    };
  });
  console.log('displayedCards', displayedCards.value);
};

watch(() => props.cards, () => {
  calculateCards();
});

const calculateParentContainer = () => {
  const container = document.getElementById('round-deck-table');
  if (container) {
    parentContainer.value.width = container.clientWidth;
    parentContainer.value.height = container.clientHeight;
  }
}

onMounted(() => {
  calculateParentContainer();
  calculateCards();
  // window.addEventListener('resize', calculateCardSizes);
});

onBeforeUnmount(() => {
  //window.removeEventListener('resize', calculateCardSizes);
});

</script>

<style scoped>
.round-deck-table {
  width: 100%;
}

.round-deck {
  display: flex;
  position: relative;
  width: 100%;
  height: 100vh;
}

.round-deck .card-container {
  width: 60%;
  position: absolute;
  top: 0;
  left: 0;
  height: auto;
  cursor: move;
  //transition: transform 1s ease;
  border: none;
}

</style>
