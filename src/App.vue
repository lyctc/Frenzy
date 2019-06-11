<template>
<div id="app">
  <Dash />
</div>
</template>

<script>
import Dash from './components/Dash.vue'
import { mapActions } from 'vuex';

function normalShortcuts(keyCode, order, tid, dispatch) {
  let pos = order.indexOf(tid);
  if (keyCode === 38) {
    if (pos > 0) {
      dispatch('normalArrow', order[pos - 1]);
    }
  } else if (keyCode === 40) {
    if (pos < order.length - 1) {
      dispatch('normalArrow', order[pos + 1]);
    }
  } else if (keyCode === 13) {
    dispatch('modeSet', {mode: 'selected', modesubHighlight: 'move'});
  }
}

function selectedShortcuts(keyCode, order, tid, modesubHighlight, dispatch) {
  let pos = order.indexOf(tid);
  if (keyCode === 38) {
    dispatch('modeSet', {mode: 'normal', modesubHighlight: ''});
    if (pos > 0) {
      dispatch('normalArrow', order[pos - 1]);
    }
  } else if (keyCode === 40) {
    dispatch('modeSet', {mode: 'normal', modesubHighlight: ''});
    if (pos < order.length - 1) {
      dispatch('normalArrow', order[pos + 1]);
    }
  } else if (keyCode === 37) {
    if (modesubHighlight === 'move') {
      // pass
    } else if (modesubHighlight === 'edit') {
      dispatch('modeSet', {mode: 'selected', modesubHighlight: 'move'});
    } else if (modesubHighlight === 'delete') {
      dispatch('modeSet', {mode: 'selected', modesubHighlight: 'edit'});
    }
  } else if (keyCode === 39) {
    if (modesubHighlight === 'move') {
      dispatch('modeSet', {mode: 'selected', modesubHighlight: 'edit'});
    } else if (modesubHighlight === 'edit') {
      dispatch('modeSet', {mode: 'selected', modesubHighlight: 'delete'});
    } else if (modesubHighlight === 'delete') {
      // pass
    }
  } else if (keyCode === 13) {
    if (modesubHighlight === 'move') {
      dispatch('modeSet', {mode: 'move', modesubHighlight: ''});
    } else if (modesubHighlight === 'edit') {
      dispatch('modeSet', {mode: 'edit', modesubHighlight: ''});
    } else if (modesubHighlight === 'delete') {
      dispatch('deleteTodo', tid);
    }
  }
}

function moveShortcuts(keyCode, order, tid, dispatch) {
  let pos = order.indexOf(tid);
  if (keyCode === 38) {
    if (pos > 0) {
      dispatch('moveOrder', {pos0: pos, pos1: pos - 1, tid});
    }
  } else if (keyCode === 40) {
    if (pos < order.length - 1) {
      dispatch('moveOrder', {pos0: pos, pos1: pos + 1, tid});
    }
  } else if (keyCode === 13) {
    dispatch('modeReset');
  }
}

export default {
  name: 'app',
  components: {
    Dash,
  },
  methods: {
    ...mapActions([
      'modeReset',
      'normalArrow',
    ]),
  },
  computed: {
  },
  mounted() {
    window.addEventListener("keyup", e => {
      if (e.keyCode === 27) {
        this.$store.dispatch('modeReset', 1);
      }
      if (this.$store.state.mode === 'normal') {
        normalShortcuts(
          e.keyCode,
          this.$store.state.order,
          this.$store.state.tid,
          this.$store.dispatch
        )
      } else if (this.$store.state.mode === 'selected') {
        selectedShortcuts(
          e.keyCode,
          this.$store.state.order,
          this.$store.state.tid,
          this.$store.state.modesubHighlight,
          this.$store.dispatch
        )
      } else if (this.$store.state.mode === 'move') {
        moveShortcuts(
          e.keyCode,
          this.$store.state.order,
          this.$store.state.tid,
          this.$store.dispatch
        );
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
