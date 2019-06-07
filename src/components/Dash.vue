<template>
<div>
  <Item
    v-for="todo in store.state.todos"
    v-bind:todo="todo"
    v-bind:key="todo.tid"
    :class="{ 'selectedTodo' : store.state.selected === todo.tid }"
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
  props: [
    'store',
  ],
};
</script>

<style scoped>
.selectedTodo {
  background-color: #CCFFFF;
}
</style>
