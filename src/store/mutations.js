import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)

function updateItemADB(TokenString, uid, pid, itemA) {
  if (uid !== -1) {
    const params = new URLSearchParams();
    params.append('TokenString', TokenString);
    params.append('UID', uid);
    params.append('PID', pid);
    params.append('ItemA', itemA);
    axios.post(process.env.VUE_APP_AXIOS + 'save', params)
  }
}

export const mutations = {
  updateTitle(state, Title) {
    Vue.set(state, 'Title', Title);
  },
  updatePage(state, Page) {
    Vue.set(state, 'Page', Page);
  },
  updateUID(state, UID) {
    Vue.set(state, 'UID', UID);
  },
  updatePID(state, PID) {
    Vue.set(state, 'PID', PID);
  },
  updateTokenString(state, TokenString) {
    Vue.set(state, 'TokenString', TokenString);
  },
  updatePosA(state, posA) {
    Vue.set(state, 'posA', posA);
  },
  updateMode(state, {mode, modesub}) {
    state.mode = mode;
    state.modesub = modesub;
  },
  updatePlanA(state, PlanA) {
    Vue.set(state, 'PlanA', PlanA);
  },
  updateItemA(state, itemA) {
    updateItemADB(state.TokenString, state.UID, state.PID, JSON.stringify(itemA))
    Vue.set(state, 'itemA', itemA);
  },
};
