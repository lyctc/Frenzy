<template>
<div id="app">
  <Login v-if="Page == 'login'" />
  <Signup v-if="Page == 'signup'" />
  <Planlist v-if="Page == ''"/>

  <div style="margin: 0px 14px;">
    <div class="titleBar">
      <div style="width: 30%; text-align: left;">
        <button @click="handleList" style="margin-top: 6px; margin-left: 0px; border: 2px solid #000; background-color: #EEE; color: #000; font-weight: 600; border-radius: 3px; padding: 4px 10px;">Plans &#9662;</button>
      </div>
      <div style="width: 40%;">
        <div v-if="mode !== 'title' && Title !== ''">
          <h1 @click="handleTitle">{{Title}}</h1>
        </div>
        <div v-if="mode === 'title' || Title === ''">
          <input type="text" v-model="title" placeholder="Title">
        </div>
      </div>
      <div style="width: 30%; height: 100%; text-align: right;">
        <div v-if="UID === -1">
          <button @click="handleLogin" class="loginButton">Log In</button>
          <button @click="handleSignup" class="signupButton">Sign Up</button>
        </div>
        <div v-if="UID !== -1">
          <button @click="handleLogout" class="loginButton">Log Out</button>
        </div>
      </div>
    </div>
   
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
import Planlist from './components/Planlist.vue'
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
  name: 'app',
  components: {
    Layer,
    Signup,
    Login,
    Planlist,
  },
  methods: {
    ...mapActions([
      'updatePage',
      'updateUID',
      'updatePID',
      'updateTokenString',
      'updateMode',
      'updatePosA',
      'updatePlanA',
    ]),
    handleList: function () {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      axios.post('http://localhost:3030/list', params)
        .then(response => {
          this.$store.dispatch('updatePlanA', response.data.PlanA)
        })
    },
    handleTitle: function () {
      alert('hello world');
    },
    handleLogin: function () {
      this.$store.dispatch('updatePage', 'login');
    },
    handleSignup: function () {
      this.$store.dispatch('updatePage', 'signup');
    },
    handleLogout: function () {
      window.$cookies.set('TokenString', '');
      this.$store.dispatch('updateTokenString', '');
      this.$store.dispatch('updatePage', 'login');
      this.$store.dispatch('updateUID', -1);
      this.$store.dispatch('updatePID', -1);
      this.$store.dispatch('updateTitle', 'Plan Your Next Frenzy');
      this.$store.dispatch('updateItemA', []);
    },
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
    Title() {
      return this.$store.state.Title;
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
      axios.post('http://localhost:3030/refresh', params)
        .then(response => {
          this.$store.dispatch('updateUID', response.data.UID)
          this.$store.dispatch('updatePID', response.data.PID)
          this.$store.dispatch('updateTitle', response.data.Title)
          this.$store.dispatch('updateItemA', JSON.parse(response.data.ItemA))
          this.$store.dispatch('updatePosA', [0])
          this.pid = response.data.PID;
        })
        .catch(err => {
          window.$cookies.set('TokenString', '');
          this.$store.dispatch('updateUID', -1)
        });
    } else if (this.$store.state.UID == -1) {
      this.$store.dispatch('updateItemA', defaultItemA());
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

<style>
:root {
  --baseBG: #FFFFFF;
  --baseFG: #000000;
  --baseTableBorder: #CCCCCC;
  --layerLeftBorder: #000000;
  --itemAncestorBG: #EFEFEF;
  --itemAncestorFG: #000000;
  --itemSelectedBG: #B7FFFF;
  --itemSelectedFG: #000000;
  --itemSelectedButtonBG: #B7FFFF;
  --itemSelectedButtonFG: #888888;
  --modeSelectedBG: #FFFFB7;
  --modeSelectedFG: #000000;
  --modeUnselectedBG: #EEEEEE;
  --modeUnselectedFG: #000000;
  --buttonLoginFG: #EEEEEE;
  --buttonLoginBG: #555555;
  --buttonSignupFG: #EEEEEE;
  --buttonSignupBG: #555555;
}

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

.titleBar {
  width: 100%;
  text-align: center;
  border-bottom: 2px solid var(--baseTableBorder);
  margin-bottom: 10px;
  height: 40px;
  display: flex;
}

.titleBar h1 {
  font-size: 1.4em;
  padding: 0;
  margin-top: 12px;
}

.titleBar button {
  padding: 4px 12px;
  border: none;
  border-radius: 5px;
  margin-top: 7px;
  margin-left: 6px;
  font-weight: 600;
  font-size: 0.7em;
}

.titleBar .loginButton {
  color: var(--buttonLoginFG);
  background-color: var(--buttonLoginBG);
}

.titleBar .signupButton {
  color: var(--buttonSignupFG);
  background-color: var(--buttonSignupBG);
}

</style>
