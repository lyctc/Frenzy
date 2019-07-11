import Vue from 'vue';
import Vuex from 'vuex';
import { mutations } from './mutations';
import actions from './actions';

Vue.use(Vuex);

export default new Vuex.Store( {
  state: {
    UID: null,
    PID: null,
    PlanA: null,
    Title: null,
    Page: '',
    posA: [0], // path of current position
    pathItemA: [], // path of selected task
    mode: 'normal', // 'normal', 'selected', 'move', 'edit', 'add'
    modesub: '',
    addButtonVisible: false,
    addInputVisible: false,
    itemA: [],
    dispA: [],
    viewLabelA: [],
  },
  actions, 
  mutations,
});
