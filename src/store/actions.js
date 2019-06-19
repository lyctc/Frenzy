export default {
  updateTitle({ commit }, Title) {
    commit('updateTitle', Title);
  },
  updatePage({ commit }, Page) {
    commit('updatePage', Page);
  },
  updateUID({ commit }, UID) {
    commit('updateUID', UID);
  },
  updatePID({ commit }, PID) {
    commit('updatePID', PID);
  },
  updateTokenString({ commit }, TokenString) {
    commit('updateTokenString', TokenString);
  },
  updatePosA({ commit }, posA) {
    commit('updatePosA', posA);
  },
  updateMode({ commit }, {mode, modesub}) {
    commit('updateMode', {mode, modesub});
  },
  updatePlanA({ commit }, PlanA) {
    commit('updatePlanA', PlanA);
  },
  updateItemA({ commit }, itemA) {
    commit('updateItemA', itemA);
  },
}
