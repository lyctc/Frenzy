<template>
<div class="titleBar" @mouseover="buttonF=true">
  <div style="width: 30%; text-align: left;">
    <div v-if="UID !== -1">
      <Planlist style="display:inline-block;" />
      <div v-if="buttonF" style="display:inline-block;">
        <button @click="handleDelete">Delete Plan</button>
        <button @click="handleEditTitle">Edit Title</button>
      </div>
    </div>
  </div>
  <div style="width: 40%;">
    <div v-if="!editF && this.$store.state.Title !== ''">
      <h1 @click="handleEditTitle">{{this.$store.state.Title}}</h1>
    </div>
    <div v-if="editF || this.$store.state.Title === ''" style="padding-top: 10px">
      <input v-select type="text" v-model="titleNew" v-on:keyup.enter="handleEditTitleSubmit" placeholder="Title" />
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
</template>

<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueCookies from 'vue-cookies'
import { mapActions } from 'vuex';
Vue.use(VueCookies)
Vue.use(VueAxios, axios)

import Planlist from './Planlist.vue'

export default {
  name: 'TitleView',
  data: function () {
    return {
      buttonF: false,
      editF: false,
      titleNew: this.$store.state.Title,
    }
  },
  components: {
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
      'updateTitle',
    ]),
    handleEditTitle: function () {
      this.editF = true;
    },
    handleEditTitleSubmit: function () {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      params.append('Title', this.titleNew);
      params.append('PID', this.$store.state.PID);
      axios.post('http://localhost:3030/title', params)
        .then(response => {
          this.$store.dispatch('updateTitle', this.titleNew);
          this.showF = true;
        })
      this.editF = false;
    },
    handleDelete: function () {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      params.append('PID', this.$store.state.PID);
      axios.post('http://localhost:3030/delete', params)
        .then(response => {
          location.reload();
        })
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
  directives: {
    select: {
      inserted: function (el) {
        el.select()
      }
    }
  },
}
</script>

<style>
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
