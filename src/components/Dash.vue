<template>
<div>
  <Item
    v-for="tid in order"
    v-bind:tid="tid"
    v-bind:key="tid"
    :class="{ 'selectedTodo' : tid === tid0 }"
  />
</div>
</template>

<script>
import Item from './Item.vue'

export default {
  name: 'app',
  components: {
    Item,
  },
  computed: {
    order() {
      return this.$store.state.order;
    },
    tid0() {
      return this.$store.state.tid;
    }
  },
};
</script>

<style scoped>
.selectedTodo {
  background-color: #CCFFFF;
}
</style>
