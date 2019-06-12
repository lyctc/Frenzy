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
        title: 'Navigate with arrow keys',
        labelA: [],
        childA: [
          {
            pathA: [0, 0],
            title: 'Up/down to move within layers',
            labelA: [],
            childA: [],
          },
          {
            pathA: [0, 1],
            title: 'Left/right to move between layers',
            labelA: [],
            childA: [],
          },
          {
            pathA: [0, 2],
            title: 'Press enter, then left/right to edit',
            labelA: [],
            childA: [],
          },
          {
            pathA: [0, 3],
            title: 'Press escape to cancel',
            labelA: [],
            childA: [],
          },
        ],
      },
      {
        pathA: [1],
        title: 'Filter by custom tags',
        labelA: [],
        childA: [
          {
            pathA: [1, 0],
            title: 'Press 1 to highlight the new features',
            labelA: [],
            childA: [],
          },
          {
            pathA: [1, 1],
            title: 'Press 1-1 to only show new features',
            labelA: [],
            childA: [],
          },
          {
            pathA: [1, 2],
            title: 'Press 1-1-1 to cancel',
            labelA: [],
            childA: [],
          },
        ],
      },
      {
        pathA: [2],
        title: 'Share with access controls',
        labelA: [],
        childA: [],
      },
      {
        pathA: [3],
        title: 'Plan multiple frenzies',
        labelA: [],
        childA: [],
      },
      {
        pathA: [4],
        title: 'Get started for free',
        labelA: [],
        childA: [],
      },
    ],
  },
  actions, 
  mutations,
});
