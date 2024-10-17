<template>
  <div class="card-closed" :style="cardStyle">
    <div class="template" :style="templateStyle"></div>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent} from 'vue';

export default defineComponent({
  props: {
    type: {
      type: String,
      required: false
    }
  },
  setup(props) {
    const spriteUrl = '/cards/sprite.png';
    const spriteFrameCounts = {width: 4, height: 5};
    const cardSize = {width: 896, height: 1302};

    const positions: Record<string, { x: number, y: number }> = {
      red: {x: 0, y: 2},
      black: {x: 3, y: 1},
    };
    let type = '';
    switch (props.type) {
      case "red":
        type = "red";
        break;
      case "black":
        type = "black";
        break;
      default:
        type = "black";
    }
    const backgroundPositionTemplateX = positions[type].x*100/(spriteFrameCounts.width-1)
    const backgroundPositionTemplateY = positions[type].y*100/(spriteFrameCounts.height-1)
    const cardStyle = computed(() => ({
      aspectRatio: `${cardSize.width}/${cardSize.height}`,
    }));
    const templateStyle = computed(() => ({
      backgroundImage: `url(${spriteUrl})`,
      backgroundSize: ` ${spriteFrameCounts.width * 100}% auto`,
      backgroundPosition: `${backgroundPositionTemplateX}% ${backgroundPositionTemplateY}%`,
    }));

    return {
      cardStyle,
      templateStyle,
    };
  }
});
</script>

<style scoped>

.card-closed {
  width: 100%;
  border: 1px solid;
  border-color: var(--card-border, #333333);
  border-radius: 8px;
  box-shadow: var(--card-shadow, 0 10px 10px rgba(255, 255, 255, 0.1));
  background-color: var(--card-background, white);
  position: relative;
}

.card-closed .template {
  background-repeat: no-repeat;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

</style>
