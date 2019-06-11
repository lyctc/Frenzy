export default {
  normalArrow({ commit }, tid) {
    commit('normalArrow', tid);
  },
  modeReset({ commit }) {
    commit('modeReset');
  },
  modeSet({ commit }, {mode, modesubHighlight}) {
    commit('modeSet', {mode, modesubHighlight});
  },
  moveOrder({ commit }, {pos0, pos1, tid}) {
    commit('moveOrder', {pos0, pos1, tid});
  },
  addTodo ({ commit }, {tid, title}) {
    commit('addTodo', {tid, title})
  },
  deleteTodo ({ commit }, tid) {
    commit('deleteTodo', tid)
  },
}
