<template>
  <div class="player">
    {{player}}:
  </div>
  <v-row no-gutters>
    <v-col class="rank">
      {{rank}}
    </v-col>
    <v-col>
      <div class="icon" :style="iconStyle">
        <div class="template" :style="templateStyle"></div>
      </div>
    </v-col>
  </v-row>
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
    },
    player: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const spriteUrl = '/icons/sprite.png';
    const spriteFrameCounts = {width: 5, height: 3};
    const size = {width: 311, height: 285};

    const positions: Record<string, { x: number, y: number }> = {
      spades: {x: 0, y: 0},
      clubs: {x: 1, y: 0},
      hearts: {x: 2, y: 0},
      diamonds: {x: 3, y: 0}
    };
    const type = props.suit;
    if (!type) {
      throw new Error('Suit is required');
    }
    const backgroundPositionTemplateX = positions[type].x * 100 / (spriteFrameCounts.width - 1)
    const backgroundPositionTemplateY = positions[type].y * 100 / (spriteFrameCounts.height - 1)
    const iconStyle = computed(() => ({
      aspectRatio: `${size.width}/${size.height}`,
    }));
    const templateStyle = computed(() => ({
      backgroundImage: `url(${spriteUrl})`,
      backgroundSize: ` ${spriteFrameCounts.width * 100}% auto`,
      backgroundPosition: `${backgroundPositionTemplateX}% ${backgroundPositionTemplateY}%`,
    }));

    return {
      iconStyle,
      templateStyle,
    };
  }
});
</script>

<style scoped>

.icon {
  width: 100%;
  position: relative;
}

.icon .template {
  background-repeat: no-repeat;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.player {
  font-size: 0.5em;
}
.rank {
  font-size: 0.8em;
}

</style>
