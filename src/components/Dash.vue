<template>
<div>
  <p>{{store.state.selected}}</p>
  <p>Completed Tasks: {{store.state.todos.filter(todo => {return todo.done === true}).length}}</p>
  <p>Pending Tasks: {{store.state.todos.filter(todo => {return todo.done === false}).length}}</p>
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
    font-weight: 600;
  }
</style>
