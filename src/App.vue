<template>
<div id="app">
  <Dash v-bind:store="store" />
</div>
</template>

<script>
import Vue from 'vue'
import Dash from './components/Dash.vue'

const store = {
  state: {
    selected: -1,
    mode: 'move', // 'move', 'selected', 'edit'
    todos: [
      {
        tid: 0,
        title: 'Feed the dog',
        labels: [],
      }, {
        tid: 1,
        title: 'Get the milk',
        labels: [],
      }, {
        tid: 2,
        title: 'Do the laundry',
        labels: [],
      }, {
        tid: 3,
        title: 'Meet with John',
        labels: [],
      }
    ],
  },
  setState(key, value) {
    Vue.set(this.state, key, value);
  },
}

export default {
  name: 'app',
  components: {
    Dash,
  },
  methods: {
  },
  data: function() {
    return {
      store: store,
    };
  },
  mounted() {
    window.addEventListener("keyup", e => {
      if (e.keyCode === 38) { // up arrow
        if (store.state.selected > 0) {
          store.setState('selected', store.state.selected - 1);
        }
      }
      if (e.keyCode === 40) { // down arrow
        if (store.state.selected < store.state.todos.length - 1) {
          store.setState('selected', store.state.selected + 1);
        }
      }
      if (e.keyCode === 13) { // enter
      }
    });
  },
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 40px;
}
</style>
