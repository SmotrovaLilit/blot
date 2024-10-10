<template>
  <div class="closed-deck">
    <div class="card-container" v-for="card in cards"
         :style="card.style">
      <ClosedCard :type="cardsTemplate"/>
    </div>
  </div>
</template>


<script lang="ts">
import {defineComponent, ref, watch} from 'vue';
import ClosedCard from './ClosedCard.vue';

function calculateAngle(cardIndex: number, totalCards: number) {
  const angle = 2;
  const half = Math.floor(totalCards / 2);
  const distance = cardIndex - half;
  return distance * angle;
}

function calculateMarginTop(angle: number) {
  return Math.abs(Math.tan(angle * Math.PI / 180) * 100);
}

export default defineComponent({
  components: {
    ClosedCard
  },
  props: {
    cardsCount: {
      type: Number,
      required: true
    },
    cardsTemplate: {
      type: String,
      required: false
    }
  },
  setup(props) {
    const cards = ref<Array<{
      style: {};
    }>>([]);
    const rotateCards = () => {
      cards.value.forEach((card, index) => {
        const angle = calculateAngle(index, props.cardsCount);
        const marginTop = calculateMarginTop(angle);
        card.style.transform = `rotate(${angle}deg)`;
        // card.style.transform = `rotate(${angle}deg) translateY(${marginTop}px)`;
        // card.style.marginTop = `${marginTop}px`;
        // card.style.paddingTop = `${marginTop}px`;
        card.style.top = `${marginTop}px`;
        card.style.zIndex = index;
      });
    };
    for (let i = 0; i < props.cardsCount; i++) {
      cards.value.push({style: {}});
    }
    rotateCards();
    return {
      cards,
      cardsTemplate: props.cardsTemplate,
    };
  }
});
</script>

<style scoped>

.closed-deck {
  display: flex;
  width: 100%;
  align-items: center;
  //padding-left: 100px;
  //padding-right: 200px;
}

.closed-deck .card-container {
  width: 100%;
  transition: transform 0.2s ease;
  position: relative;
  margin-right: -10%;
}

</style>
