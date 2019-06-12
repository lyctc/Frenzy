import Vue from 'vue'

export const mutations = {
  updatePosA(state, posA) {
    Vue.set(state, 'posA', posA);
  },
  updateMode(state, {mode, modesub}) {
    state.mode = mode;
    state.modesub = modesub;
  },
  updateItemA(state, itemA) {
    Vue.set(state, 'itemA', itemA);
  },
};
