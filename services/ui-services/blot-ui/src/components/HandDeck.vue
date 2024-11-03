<template>
  <draggable
      :list="cardsComputed"
      :disabled="false"
      item-key="id"
      class="hand-deck"
      ghost-class="ghost"
      :move="checkMove"
      @end="onDragEnd"
      @start="onDragStart"
  >
    <template #item="{ element }">
      <div class="card-container" :class="{ 'not-draggable': false }"
           :onclick="playCard(element.rank, element.suit)"
           :style="element.style">
        <CardView :rank="element.rank" :suit="element.suit"/>
      </div>
    </template>
  </draggable>
</template>

<script setup lang="ts">
import {computed} from 'vue';
import CardView from './Card.vue';
import draggable from "vuedraggable";
import {Card} from "@/models/gameSet";

const props = defineProps<{
  cards: Array<{ rank: string; suit: string }>;
  isYourTurn: boolean;
  playCard: (card: Card) => void;
}>();

function calculateAngle(cardIndex: number, totalCards: number) {
  const angle = 10;
  const half = Math.floor(totalCards / 2);
  const distance = cardIndex - half;
  return distance * angle;
}

function calculateMarginTop(angle: number) {
  return Math.abs(Math.tan(angle * Math.PI / 180) * 100);
}

function calculateNewPositions(
    length: number,
    currentIndex: number,
    futureIndex: number
): number[] {
  if (currentIndex < 0 || futureIndex < 0 || currentIndex >= length || futureIndex >= length) {
    return [];
  }
  let res: number[] = new Array(length).fill(0);
  for (let i = 0; i < length; i++) {
    if (i === currentIndex) {
      res[currentIndex] = futureIndex;
    } else if (currentIndex < futureIndex) { // moving right
      if (i > currentIndex && i <= futureIndex) { // changed range
        res[i] = i - 1;
      } else { // no changed range
        res[i] = i;
      }
    } else { // moving left
      if (i >= futureIndex && i < currentIndex) { // changed range
        res[i] = i + 1;
      } else {
        res[i] = i  // no changed range
      }
    }
  }
  return res;

}

// const cardsComputed = ref<Array<{
//   id: string;
//   rank: string;
//   suit: string;
//   style: {}
// }>>([]);
const cardsComputed = computed(() => {
  return props.cards.map((card) => {
    const id = `${card.rank}-${card.suit}`;
    return {
      id,
      rank: card.rank,
      suit: card.suit,
      style: {
        transform: ``,
        marginTop: ``,
        zIndex: 0,
        top: ``,
      },
    };
  });
});
const rotateCards = () => {
  cardsComputed.value.forEach((card, index) => {
    const angle = calculateAngle(index, props.cards.length);
    const marginTop = calculateMarginTop(angle);
    card.style.transform = `rotate(${angle}deg)`;
    // card.style.transform = `rotate(${angle}deg) translateY(${marginTop}px)`;
    // card.style.marginTop = `${marginTop}px`;
    // card.style.paddingTop = `${marginTop}px`;
    card.style.top = `${marginTop}px`;
    card.style.zIndex = index;
  });
};
// cardsComputed.value = props.cards.map((card) => {
//   const id = `${card.rank}-${card.suit}`;
//   return {
//     id,
//     rank: card.rank,
//     suit: card.suit,
//     style: {
//       transform: ``,
//       marginTop: ``,
//       zIndex: 0,
//     },
//   };
// });
rotateCards();
const onDragEnd = (e) => {
  console.log('hand card onDragEnd', e);
  rotateCards();
};

const checkMove = (e) => {
  console.log('hand card checkMove moving', e);
  if (e.draggedContext.futureIndex === undefined) {
    return;
  }
  // console.log('Future index: ' + e.draggedContext.futureIndex);
  // console.log('index: ' + e.draggedContext.index);
  //
  // console.log('Old array: ' + cardsComputed.value[0].rank + ' ' + cardsComputed.value[0].style.zIndex)
  // console.log('Old array: ' + cardsComputed.value[1].rank + ' ' + cardsComputed.value[1].style.zIndex)
  // console.log('old array: ' + cardsComputed.value[2].rank + ' ' + cardsComputed.value[2].style.zIndex)

  const newPositions = calculateNewPositions(cardsComputed.value.length, e.draggedContext.index, e.draggedContext.futureIndex);
  for (let i = 0; i < newPositions.length; i++) {
    const angle = calculateAngle(newPositions[i], cardsComputed.value.length);
    // const marginTop = e.draggedContext.index == i ?  calculateMarginTop(angle): calculateMarginTop(angle);
    const marginTop = e.draggedContext.index == i ? -50 + calculateMarginTop(angle) : calculateMarginTop(angle);
    cardsComputed.value[i].style.zIndex = newPositions[i];
    cardsComputed.value[i].style.transform = `rotate(${angle}deg)`;
    // cards.value[i].style.transform = `rotate(${angle}deg) translateY(${marginTop}px)`;
    cardsComputed.value[i].style.top = `${marginTop}px`;
    // cards.value[i].style.marginTop = `${marginTop}px`;
    // cards.value[i].style.paddingTop = `${marginTop}px`;
  }
  // console.log(cardsComputed.value[0].rank + ' ' + cardsComputed.value[0].style.zIndex)
  // console.log(cardsComputed.value[1].rank + ' ' + cardsComputed.value[1].style.zIndex)
  // console.log(cardsComputed.value[2].rank + ' ' + cardsComputed.value[2].style.zIndex)
};
const onDragStart = (e) => {
  console.log('hand card onDragStart!', e);
  const draggedElement = e.item;
  draggedElement.style.zIndex = cardsComputed.value[e.oldIndex].style.zIndex;
  draggedElement.style.transform = cardsComputed.value[e.oldIndex].style.transform;
  const angle = calculateAngle(e.oldIndex, cardsComputed.value.length);
  const marginTop = -100 + calculateMarginTop(angle);
  cardsComputed.value[e.oldIndex].style.top = `${marginTop}px`;
};

const playCard = (rank: string, suit: string) => {
  return function () {
    if (!props.isYourTurn) {
      console.warn('Not your turn!');
      return;
    }
    console.log('Playing card:', rank, suit);
    props.playCard(new Card(rank, suit));
  };
};

</script>

<style scoped>
.hand-deck .ghost {
  opacity: 0.5;
}

.hand-deck {
  display: flex;
  width: 100%;
  align-items: center;
  transform: translateX(-7%);
}

.hand-deck .card-container {
  width: 100%;
  cursor: move;
  transition: transform 0.2s ease;
  position: relative;
  margin-right: -10%;
}

.hand-deck .not-draggable {
  cursor: no-drop;
}
</style>
