import Vue from 'vue'

export const mutations = {
  normalArrow(state, tid) {
    state.tid = tid
  },
  modeReset(state) {
    state.mode = 'normal';
    state.modesubHighlight = '';
  },
  modeSet(state, {mode, modesubHighlight}) {
    state.mode = mode;
    state.modesubHighlight = modesubHighlight;
  },
  moveOrder(state, {pos0, pos1, tid}) { // tid starts at pos0
    state.order.splice(pos0, 1, state.order[pos1]);
    state.order.splice(pos1, 1, tid);
  },
  addTodo(state, {tid, title}) {
    Vue.set(state.todos, tid, {title: title});
    state.mode = 'normal';
  },
  deleteTodo(state, tid) {
    state.order = state.order.filter(function (i) {
      return i !== tid
    })
    delete state.todos[tid];
    state.mode = 'normal';
  },
};
