import Vue from 'vue';
import Vuex from 'vuex';
import { mutations } from './mutations';
import actions from './actions';

Vue.use(Vuex);

export default new Vuex.Store( {
  state: {
    UID: -1,
    PID: -1,
    PlanA: [],
    Title: '',
    Page: '',
    posA: [0], // path of current position
    pathA: [], // path of selected task
    mode: 'normal', // 'normal', 'selected', 'move', 'edit', 'add'
    modesub: '',
    addButtonVisible: false,
    addInputVisible: false,
    itemA: [],
  },
  actions, 
  mutations,
});
