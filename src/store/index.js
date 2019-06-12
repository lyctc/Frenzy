import Vue from 'vue';
import Vuex from 'vuex';
import { mutations } from './mutations';
import actions from './actions';

Vue.use(Vuex);

export default new Vuex.Store( {
  state: {
    posA: [0, 0], // path of current position
    pathA: [], // path of selected task
    mode: 'normal', // 'normal', 'selected', 'move', 'edit', 'add'
    modesub: '',
    addButtonVisible: false,
    addInputVisible: false,
    itemA: [
      {
        pathA: [0],
        title: 'Navigate with Arrow Keys',
        labelA: [],
        childA: [
          {
            pathA: [0, 0],
            title: 'Up/Down to Move Within Layers',
            labelA: [],
            childA: [],
          },
          {
            pathA: [0, 1],
            title: 'Left/Right to Move Between Layers',
            labelA: [],
            childA: [],
          }
        ],
      },
      {
        pathA: [1],
        title: 'Filter by Custom Tags',
        labelA: [],
        childA: [
          {
            pathA: [1, 0],
            title: 'Press 1 to Highlight the New Features',
            labelA: [],
            childA: [],
          },
          {
            pathA: [1, 1],
            title: 'Press 1 Again to Only Show New Features',
            labelA: [],
            childA: [],
          },
          {
            pathA: [1, 2],
            title: 'Press 1 Again to Return to Overall Plan',
            labelA: [],
            childA: [],
          },
        ],
      },
      {
        pathA: [2],
        title: 'Share with Access Controls',
        labelA: [],
        childA: [],
      },
      {
        pathA: [3],
        title: 'Plan Multiple Frenzies',
        labelA: [],
        childA: [],
      },
      {
        pathA: [4],
        title: 'Get Started for Free',
        labelA: [],
        childA: [],
      },
    ],
  },
  actions, 
  mutations,
});
