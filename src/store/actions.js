export default {
  updatePosA({ commit }, posA) {
    commit('updatePosA', posA);
  },
  updateMode({ commit }, {mode, modesub}) {
    commit('updateMode', {mode, modesub});
  },
  updateItemA({ commit }, itemA) {
    commit('updateItemA', itemA);
  },
}
