<template>
<div class="item">
  <div v-if="mode !== 'edit' || tid !== tid0">
    {{todos[tid].title}}
  </div>
  <div v-if="mode === 'selected' && tid === tid0">
    <button :class="{ 'modesubHighlight' : modesubHighlight === 'move' }">Move</button>
    <button :class="{ 'modesubHighlight' : modesubHighlight === 'edit' }">Edit</button>
    <button :class="{ 'modesubHighlight' : modesubHighlight === 'delete' }">Delete</button>
  </div>
  <div v-if="mode === 'edit' && tid === tid0">
    <input v-select type="text" :value="todos[tid].title" @keyup.enter="addTodo" />
  </div>

</div>
</template>

<script>
export default {
  name: 'item',
  props: ['tid'],
  directives: {
    select: {
      inserted: function (el) {
        el.select()
      }
    }
  },
  methods: {
    addTodo (e) {
      this.$store.dispatch('addTodo', {tid: this.$props.tid, title: e.target.value})
    },
  },
  computed: {
    todos() {
      return this.$store.state.todos;
    },
    tid0() {
      return this.$store.state.tid;
    },
    mode() {
      return this.$store.state.mode;
    },
    modesubHighlight() {
      return this.$store.state.modesubHighlight;
    },
  },

}
</script>

<style scoped>
.item {
  margin: 5px 0px;
  width: 200px;
  padding: 5px;
}

.modesubHighlight {
  background-color: #654321;
}
</style>
