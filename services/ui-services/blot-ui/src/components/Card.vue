<template>
  <div class="card" :style="cardStyle">
    <div class="template" :style="templateStyle"></div>
    <div class="rank" :style="rankStyle"></div>
    <div class="rank-text-top">{{textInTheTop}}</div>
    <div class="rank-text-middle">{{textInTheMiddle}}</div>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent} from 'vue';

export default defineComponent({
  props: {
    rank: {
      type: String,
      required: true
    },
    suit: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const spriteUrl = '/cards/sprite.png';
    const spriteFrameCounts = {width: 4, height: 5};
    const cardSize = {width: 896, height: 1302};

    const suitTemplatePositions: Record<string, { x: number, y: number }> = {
      spades: {x: 0, y: 0},
      clubs: {x: 1, y: 0},
      hearts: {x: 2, y: 0},
      diamonds: {x: 3, y: 0}
    };

    const suitPositions: Record<string, { x: number, y: number }> = {
      spades: {x: 0, y: 3},
      clubs: {x: 1, y: 3},
      hearts: {x: 2, y: 3},
      diamonds: {x: 3, y: 3}
    };
    const rankPositions: Record<string, { x: number, y: number }> = {
      k: {x: 0, y: 1},
      j: {x: 1, y: 1},
      q: {x: 2, y: 1}
    };
    const backgroundPositionSuitTemplateX = suitTemplatePositions[props.suit].x*100/(spriteFrameCounts.width-1)
    const backgroundPositionSuitTemplateY = suitTemplatePositions[props.suit].y*100/(spriteFrameCounts.height-1)
    const cardStyle = computed(() => ({
      aspectRatio: `${cardSize.width}/${cardSize.height}`,
    }));
    const templateStyle = computed(() => ({
      backgroundImage: `url(${spriteUrl})`,
      backgroundSize: ` ${spriteFrameCounts.width * 100}% auto`,
      backgroundPosition: `${backgroundPositionSuitTemplateX}% ${backgroundPositionSuitTemplateY}%`,
    }));
    let rankStyle = {};
    const textInTheTop = props.rank.toUpperCase();
    let textInTheMiddle = '';
    switch (props.rank) {
      case 'k':
      case 'q':
      case 'j':
        const backgroundPositionRankX = rankPositions[props.rank].x*100/(spriteFrameCounts.width-1)
        const backgroundPositionRankY = rankPositions[props.rank].y*100/(spriteFrameCounts.height-1)
        rankStyle = computed(() => ({
          backgroundImage: `url(${spriteUrl})`,
          backgroundSize: ` ${spriteFrameCounts.width * 100}% auto`,
          backgroundPosition: `${backgroundPositionRankX}% ${backgroundPositionRankY}%`,
        }));

        break;
      case 'a':
      case '6': // TODO нужно ли 6,
      case '7':
      case '8':
      case '9':
      case '10':
        textInTheMiddle = props.rank.toUpperCase();
        break;
      default:
        break;
    }



    return {
      cardStyle,
      templateStyle,
      rankStyle,
      textInTheTop,
      textInTheMiddle
    };
  }
});
</script>

<style scoped>

.card {
  width: 100%;
  border: 1px solid;
  border-color: var(--card-border, #333333);
  border-radius: 8px;
  box-shadow: var(--card-shadow, 0 10px 10px rgba(255, 255, 255, 0.1));
  background-color: var(--card-background, white);
  position: relative;
}

.template, .rank {
  background-repeat: no-repeat;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.rank-text-top{
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  position: absolute;
  top: 13%;
  left: 18%;
  transform: translate(-50%, -50%);
  font-size: 1.5rem;
  color: #9f865b;
  font-family: "URW Chancery L", "Comic Sans MS", cursive;
  font-weight: bold;
  text-shadow: 2px 2px 2px black;
}

.rank-text-middle {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  position: absolute;
  bottom: 50%;
  right: 50%;
  transform: translate(50%, 50%);
  font-size: 4rem;
  color: #9f865b;
  font-family: "URW Chancery L", "Comic Sans MS", cursive;
  font-weight: bold;
  text-shadow: 2px 2px 2px black;
}

</style>
