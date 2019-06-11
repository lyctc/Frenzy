import Vue from 'vue';
import Vuex from 'vuex';
import { mutations } from './mutations';
import actions from './actions';

Vue.use(Vuex);

export default new Vuex.Store( {
  state: {
    tid: -1,
    mode: 'normal', // 'normal', 'selected', 'move', 'edit'
    modesubHighlight: '',
    modesub: '',
    order: ['0', '1', '2', '3'],
    todos: {
      0: {
        title: 'Feed the dog',
        labels: [],
      },
      1: {
        title: 'Get the milk',
        labels: [],
      },
      2: {
        title: 'Do the laundry',
        labels: [],
      },
      3: {
        title: 'Meet with John',
        labels: [],
      }
    }
  },
  actions, 
  mutations,
});
