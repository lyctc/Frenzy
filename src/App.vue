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
        style="width: 240px; flex: none;"
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
import { rebalancePathA, moveItemHelper, deleteItemHelper, defaultItemA } from './lib.js'

function normalShortcuts(keyCode, posA, itemA, dispatch) {
  let i;
  let posA0 = Array.from(posA);
  let itemAN = Array.from(itemA); // last layer
  for (i = 0; i < posA.length - 1; i += 1) {
    itemAN = (itemAN !== []) ? itemAN[posA[i]].childA : [];
  }

  if (keyCode === 38) { // up
    if (posA0[posA0.length - 1] > 0) {
      posA0[posA0.length - 1] -= 1;
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 40) { // down
    if (posA0[posA0.length - 1] < itemAN.length) {
      posA0[posA0.length - 1] += 1;
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 37) { // left
    if (posA0.length > 1) {
      posA0.pop();
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 39) { // right
    if (itemAN.length > 0 && posA0[posA0.length - 1] < itemAN.length) {
      posA0.push(0);
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 13) {
    if (posA0[posA0.length - 1] < itemAN.length) {
      dispatch('updateMode', {mode: 'selected', modesub: 'move'});
    } else if (posA0[posA0.length - 1] === itemAN.length) {
      dispatch('updateMode', {mode: 'add', modesub: ''});
    }
  }
}

function selectedShortcuts(keyCode, posA, itemA, modesub, dispatch) {
  let i;
  let _;
  let posA0 = Array.from(posA);
  let itemA0 = Array.from(itemA);
  let itemAN = Array.from(itemA); // last layer
  for (i = 0; i < posA.length - 1; i += 1) {
    itemAN = (itemAN !== []) ? itemAN[posA[i]].childA : [];
  }
  if (keyCode === 38) { // up
    dispatch('updateMode', {mode: 'normal', modesub: ''});
    if (posA[posA.length - 1] > 0) {
      posA0[posA0.length - 1] -= 1;
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 40) { // down
    dispatch('updateMode', {mode: 'normal', modesub: ''});
    if (posA[posA.length - 1] < itemAN.length) {
      posA0[posA0.length - 1] += 1;
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 37) {
    if (modesub === 'move') {
      // pass
    } else if (modesub === 'edit') {
      dispatch('updateMode', {mode: 'selected', modesub: 'move'});
    } else if (modesub === 'delete') {
      dispatch('updateMode', {mode: 'selected', modesub: 'edit'});
    }
  } else if (keyCode === 39) {
    if (modesub === 'move') {
      dispatch('updateMode', {mode: 'selected', modesub: 'edit'});
    } else if (modesub === 'edit') {
      dispatch('updateMode', {mode: 'selected', modesub: 'delete'});
    } else if (modesub === 'delete') {
      // pass
    }
  } else if (keyCode === 13) {
    if (modesub === 'move') {
      dispatch('updateMode', {mode: 'move', modesub: ''});
    } else if (modesub === 'edit') {
      dispatch('updateMode', {mode: 'edit', modesub: ''});
    } else if (modesub === 'delete') {
      itemA0 = deleteItemHelper(itemA0, posA, 0);
      _, itemA0 = rebalancePathA([], itemA0);
      dispatch('updateItemA', itemA0)
    }
  }
}

function moveShortcuts(keyCode, posA, itemA, dispatch) {
  let i;
  let _;
  let posA0 = Array.from(posA);
  let itemA0 = Array.from(itemA);
  let itemAN = Array.from(itemA); // last layer
  for (i = 0; i < posA.length - 1; i += 1) {
    itemAN = (itemAN !== []) ? itemAN[posA[i]].childA : [];
  }
  if (keyCode === 38) {
    if (posA0[posA0.length - 1] > 0) {
      itemA0 = moveItemHelper('up', itemA0, posA0, 0);
      _, itemA0 = rebalancePathA([], itemA0);
      dispatch('updateItemA', itemA0)
      posA0[posA0.length - 1] -= 1;
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 40) {
    if (posA0[posA0.length - 1] < itemAN.length - 1) {
      itemA0 = moveItemHelper('down', itemA0, posA0, 0);
      _, itemA0 = rebalancePathA([], itemA0);
      dispatch('updateItemA', itemA0)
      posA0[posA0.length - 1] += 1;
      dispatch('updatePosA', posA0);
    }
  } else if (keyCode === 13) {
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
      'updateTitle',
    ]),
  },
  computed: {
    layerA() {
      let i;
      let layerA = [];
      let itemAN = Array.from(this.$store.state.itemA);
      for (i = 0; i < this.$store.state.posA.length; i += 1) {
        layerA.push({key: `layer${i}`, step: i});
        if (itemAN.length > 0 && this.$store.state.posA[i] < itemAN.length) {
          itemAN = itemAN[this.$store.state.posA[i]].childA;
        } else { // hit add new item
          itemAN = [];
        }
      }
      if (itemAN.length > 0) { // preview
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
    if (this.$store.state.UID == -1 && window.$cookies.get('TokenString')) {
      // automatically log back in if cookie is set
      this.$store.dispatch('updateTokenString', window.$cookies.get('TokenString'))
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      axios.post(process.env.VUE_APP_AXIOS + 'refresh', params)
        .then(response => {
          this.$store.dispatch('updateUID', response.data.UID)
          this.$store.dispatch('updatePID', response.data.PID)
          this.$store.dispatch('updateTitle', response.data.Title)
          this.$store.dispatch('updateItemA', JSON.parse(response.data.ItemA))
          this.$store.dispatch('updatePosA', [0])
          this.pid = response.data.PID;
        })
        .catch(() => {
          window.$cookies.set('TokenString', '');
          this.$store.dispatch('updateUID', -1)
        });
    } else if (this.$store.state.UID == -1) {
      this.$store.dispatch('updateItemA', defaultItemA());
      this.$store.dispatch('updateTitle', "Plan Your Next Frenzy");
      this.$store.dispatch('updatePosA', [0, 0]);
    }
    window.addEventListener('keyup', e => {
      if (this.$store.state.Page === '') {
        if (e.keyCode === 27) {
          this.$store.dispatch('updateMode', {mode: 'normal', modesub: ''});
        }
        if (this.$store.state.mode === 'normal') {
          normalShortcuts(
            e.keyCode,
            this.$store.state.posA,
            this.$store.state.itemA,
            this.$store.dispatch,
          )
        } else if (this.$store.state.mode === 'selected') {
          selectedShortcuts(
            e.keyCode,
            this.$store.state.posA,
            this.$store.state.itemA,
            this.$store.state.modesub,
            this.$store.dispatch,
          )
        } else if (this.$store.state.mode === 'move') {
          moveShortcuts(
            e.keyCode,
            this.$store.state.posA,
            this.$store.state.itemA,
            this.$store.dispatch
          );
        }
      }
    });
  },
};
</script>

<style scoped>
body {
  background-color: var(--baseBG);
  color: var(--baseFG);
  margin: 0px;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 12px;
  width: 100%;
}

</style>
