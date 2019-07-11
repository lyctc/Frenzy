<template>
<div id="app">
  <Login v-if="Page == 'login'" />
  <Signup v-if="Page == 'signup'" />

  <div v-if="Page == ''" style="margin: 0px 14px;z-index: 0;">
    <TitleBar />
    <div id="layers" style="display: flex; overflow: auto">
      <Layer
        v-for="layer in layerA"
        v-bind:key="layer.key"
        v-bind:step="layer.step"
        v-bind:parentPathItemA="layer.parentPathItemA"
        style="width: 300px; flex: none;"
      />
    </div>
  </div>
</div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueCookies from 'vue-cookies'
Vue.use(VueCookies)
Vue.use(VueAxios, axios)

import Layer from './components/Layer.vue'
import Signup from './components/Signup.vue'
import Login from './components/Login.vue'
import TitleBar from './components/TitleBar.vue'
import { mapActions } from 'vuex';
import { rebalanceItemA, viewLabelHelper, labelItemHelper, moveItemHelper, deleteItemHelper, defaultItemA } from './lib.js'

function normalShortcuts(e, posA, itemA, dispA, viewLabelA, dispatch) {
  let i;
  let r;
  let posA0 = Array.from(posA);
  let itemA0 = Array.from(itemA);
  let dispAN = Array.from(dispA); // last layer
  for (i = 0; i < posA.length - 1; i += 1) {
    dispAN = (dispAN !== []) ? dispAN[posA[i]].childA : [];
  }

  if (e.keyCode === 38) { // up
    if (posA0[posA0.length - 1] > 0) {
      posA0[posA0.length - 1] -= 1;
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 40) { // down
    if (posA0[posA0.length - 1] < dispAN.length) {
      posA0[posA0.length - 1] += 1;
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 37) { // left
    if (posA0.length > 1) {
      posA0.pop();
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 39) { // right
    if (dispAN.length > 0 && posA0[posA0.length - 1] < dispAN.length) {
      posA0.push(0);
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 13) {
    if (posA0[posA0.length - 1] < dispAN.length) {
      dispatch('updateMode', {mode: 'selected', modesub: 'move'});
    } else if (posA0[posA0.length - 1] === dispAN.length) {
      dispatch('updateMode', {mode: 'add', modesub: ''});
    }
  } else if (e.keyCode === 49 && e.shiftKey) { // Shift-1
    itemA0 = labelItemHelper(1, itemA0, posA0, 0);
    r = rebalanceItemA([], [], itemA0, viewLabelA);
    dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
  } else if (e.keyCode === 50 && e.shiftKey) { // Shift-2
    itemA0 = labelItemHelper(2, itemA0, posA0, 0);
    r = rebalanceItemA([], [], itemA0, viewLabelA);
    dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
  } else if (e.keyCode === 51 && e.shiftKey) { // Shift-3
    itemA0 = labelItemHelper(3, itemA0, posA0, 0);
    r = rebalanceItemA([], [], itemA0, viewLabelA);
    dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
  } else if (e.keyCode === 49 && !e.shiftKey) { // 1
    viewLabelA = viewLabelHelper(viewLabelA, 1);
    dispatch('updateViewLabelA', viewLabelA);
    r = rebalanceItemA([], [], itemA0, viewLabelA);
    dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA});
  } else if (e.keyCode === 50 && !e.shiftKey) { // 2
    viewLabelA = viewLabelHelper(viewLabelA, 2)
    dispatch('updateViewLabelA', viewLabelA);
    r = rebalanceItemA([], [], itemA0, viewLabelA);
    dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA});
  } else if (e.keyCode === 51 && !e.shiftKey) { // 3
    viewLabelA = viewLabelHelper(viewLabelA, 3)
    dispatch('updateViewLabelA', viewLabelA);
    r = rebalanceItemA([], [], itemA0, viewLabelA);
    dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA});
  }
}

function selectedShortcuts(e, posA, itemA, modesub, viewLabelA, dispatch) {
  let i;
  let posA0 = Array.from(posA);
  let itemA0 = Array.from(itemA);
  let itemAN = Array.from(itemA); // last layer
  for (i = 0; i < posA.length - 1; i += 1) {
    itemAN = (itemAN !== []) ? itemAN[posA[i]].childA : [];
  }
  if (e.keyCode === 38) { // up
    dispatch('updateMode', {mode: 'normal', modesub: ''});
    if (posA[posA.length - 1] > 0) {
      posA0[posA0.length - 1] -= 1;
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 40) { // down
    dispatch('updateMode', {mode: 'normal', modesub: ''});
    if (posA[posA.length - 1] < itemAN.length) {
      posA0[posA0.length - 1] += 1;
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 37) {
    if (modesub === 'move') {
      // pass
    } else if (modesub === 'edit') {
      dispatch('updateMode', {mode: 'selected', modesub: 'move'});
    } else if (modesub === 'delete') {
      dispatch('updateMode', {mode: 'selected', modesub: 'edit'});
    }
  } else if (e.keyCode === 39) {
    if (modesub === 'move') {
      dispatch('updateMode', {mode: 'selected', modesub: 'edit'});
    } else if (modesub === 'edit') {
      dispatch('updateMode', {mode: 'selected', modesub: 'delete'});
    } else if (modesub === 'delete') {
      // pass
    }
  } else if (e.keyCode === 13) {
    if (modesub === 'move') {
      viewLabelA = []
      dispatch('updateViewLabelA', viewLabelA);
      let r = rebalanceItemA([], [], itemA0, viewLabelA);
      dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA});
      dispatch('updateMode', {mode: 'move', modesub: ''});
    } else if (modesub === 'edit') {
      dispatch('updateMode', {mode: 'edit', modesub: ''});
    } else if (modesub === 'delete') {
      itemA0 = deleteItemHelper(itemA0, posA, 0);
      let r = rebalanceItemA([], [], itemA0, viewLabelA);
      dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
    }
  }
}

function moveShortcuts(e, posA, itemA, viewLabelA, dispatch) {
  let i;
  let posA0 = Array.from(posA);
  let itemA0 = Array.from(itemA);
  let itemAN = Array.from(itemA); // last layer
  for (i = 0; i < posA.length - 1; i += 1) {
    itemAN = (itemAN !== []) ? itemAN[posA[i]].childA : [];
  }
  if (e.keyCode === 38) {
    if (posA0[posA0.length - 1] > 0) {
      itemA0 = moveItemHelper('up', itemA0, posA0, 0);
      let r = rebalanceItemA([], [], itemA0, viewLabelA);
      dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
      posA0[posA0.length - 1] -= 1;
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 40) {
    if (posA0[posA0.length - 1] < itemAN.length - 1) {
      itemA0 = moveItemHelper('down', itemA0, posA0, 0);
      let r = rebalanceItemA([], [], itemA0, viewLabelA);
      dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
      posA0[posA0.length - 1] += 1;
      dispatch('updatePosA', posA0);
    }
  } else if (e.keyCode === 13) {
    dispatch('updateMode', {mode: 'normal', modesub: ''});
  }
}

export default {
  name: 'App',
  components: {
    Layer,
    Signup,
    Login,
    TitleBar,
  },
  methods: {
    ...mapActions([
      'updatePage',
      'updateUID',
      'updatePID',
      'updateTokenString',
      'updateMode',
      'updatePosA',
      'updateViewLabelA',
      'updateTitle',
    ]),
  },
  computed: {
    layerA() {
      let i;
      let layerA = [];
      let dispAN = Array.from(this.$store.state.dispA);
      let parentPathItemA = [];
      for (i = 0; i < this.$store.state.posA.length; i += 1) {
        layerA.push({key: `layer${i}`, step: i, parentPathItemA});
        if (dispAN.length > 0 && this.$store.state.posA[i] < dispAN.length) {
          parentPathItemA = dispAN[this.$store.state.posA[i]].pathItemA;
          dispAN = dispAN[this.$store.state.posA[i]].childA;
        } else { // hit add new item
          dispAN = [];
        }
      }
      if (dispAN.length > 0) { // preview
        layerA.push({
          key: `layer${this.$store.state.posA.length}`,
          step: this.$store.state.posA.length
        });
      }
      return layerA;
    },
    mode() {
      return this.$store.state.mode;
    },
    posA() {
      return this.$store.state.posA;
    },
    viewLabelA() {
      return this.$store.state.viewLabelA;
    },
    Page() {
      return this.$store.state.Page;
    },
    UID() {
      return this.$store.state.UID;
    },
    PID() {
      return this.$store.state.PID;
    },
  },
  mounted() {
    if (window.$cookies.get('TokenString')) {
      // automatically log back in if cookie is set
      this.$store.dispatch('updateTokenString', window.$cookies.get('TokenString'))
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      axios.post(process.env.VUE_APP_AXIOS + 'refresh', params)
        .then(response => {
          this.$store.dispatch('updateUID', response.data.UID)
          this.$store.dispatch('updatePID', response.data.PID)
          this.$store.dispatch('updateTitle', response.data.Title)
          document.title = response.data.Title + ' | Frenzy';
          let r = rebalanceItemA([], [], JSON.parse(response.data.ItemA), this.$store.state.viewLabelA);
          this.$store.dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
          this.$store.dispatch('updatePosA', [0])
          this.pid = response.data.PID;
        })
        .catch(() => {
          window.$cookies.set('TokenString', '');
          this.$store.dispatch('updateUID', -1)
        });
    } else {
      this.$store.dispatch('updateUID', -1)
      let r = rebalanceItemA([], [], defaultItemA(), this.$store.state.viewLabelA);
      this.$store.dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
      this.$store.dispatch('updateTitle', 'Plan Your Next Frenzy');
      document.title = 'Plan Your Next Frenzy | Frenzy';
      this.$store.dispatch('updatePosA', [0, 0]);
    }
    window.addEventListener('keyup', e => {
      if (this.$store.state.Page === '') {
        if (e.keyCode === 27) {
          this.$store.dispatch('updateMode', {mode: 'normal', modesub: ''});
        }
        if (this.$store.state.mode === 'normal') {
          normalShortcuts(
            e,
            this.$store.state.posA,
            this.$store.state.itemA,
            this.$store.state.dispA,
            this.$store.state.viewLabelA,
            this.$store.dispatch,
          )
        } else if (this.$store.state.mode === 'selected') {
          selectedShortcuts(
            e,
            this.$store.state.posA,
            this.$store.state.itemA,
            this.$store.state.modesub,
            this.$store.state.viewLabelA,
            this.$store.dispatch,
          )
        } else if (this.$store.state.mode === 'move') {
          moveShortcuts(
            e,
            this.$store.state.posA,
            this.$store.state.itemA,
            this.$store.state.viewLabelA,
            this.$store.dispatch
          );
        } else if (['add', 'edit'].indexOf(this.$store.state.mode) !== -1 && e.keyCode === 13) {
          this.$store.dispatch('updateMode', {mode: 'normal', modesub: ''});
        }
      }
    });
  },
};
</script>

<style>
:root {
  --baseBG: #FCFCFC;
  --baseFG: #373737;
  --baseTableBorder: #9099A2;
  --layerLeftBorder: #373737;
  --dispAncestorBG: #EFEFEF;
  --dispAncestorFG: #000000;
  --dispSelectedBG: #E1F5FE;
  --dispSelectedFG: #000000;
  --dispMoveBG: #D0E4ED;
  --dispMoveFG: #000000;
  --dispSelectedButtonBG: #E1F5FE;
  --dispSelectedButtonFG: #7B7B7B;
  --dispChildren: #999999;
  --modeSelectedBG: #FFFDE7;
  --modeSelectedFG: #000000;
  --modeUnselectedBG: #EEEEEE;
  --modeUnselectedFG: #000000;
  --buttonLoginBG: #E8F5E9;
  --buttonLoginFG: #373737;
  --buttonLoginBO: #33691E;
  --buttonSignupBG: #0277BD;
  --buttonSignupFG: #FFFFFF;
  --buttonSignupBO: #01579B;
  --buttonLogoutBG: #EFEFEF;
  --buttonLogoutFG: #373737;
  --buttonLogoutBO: #9099A2;
  --buttonPlansBG: #EFEFEF;
  --buttonPlansFG: #373737;
  --buttonPlansBO: #9099A2;
  --buttonFormBG: #0277BD;
  --buttonFormFG: #FFFFFF;
  --buttonFormBO: #01579B;
  --buttonPlanActionBG: #EFEFEF;
  --buttonPlanActionFG: #373737;
  --buttonPlanActionBO: #9099A2;
  --plansSelectedBG: #E1F5FE;
  --plansSelectedFG: #595959;
  --inputFormAccent: #0277BD;
}

body {
  background-color: var(--baseBG);
  color: var(--baseFG);
  margin: 0px;
  padding: 0px;
}

.form-input {
  box-sizing: border-box;
  height: 34px;
  padding-left: 5px;
  border: 2px solid var(--baseTableBorder);
  width: 100%;
}

.form-input:focus {
  border: 2px solid var(--inputFormAccent);
}

.form-button {
  height: 30px;
  font-weight: 600;
  background-color: var(--buttonFormBG);
  color: var(--buttonFormFG);
  border: 2px solid var(--buttonFormBO);
  border-radius: 4px;
  width: 100%;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 1.0em;
  width: 100%;
}
</style>
