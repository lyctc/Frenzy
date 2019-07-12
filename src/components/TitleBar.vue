<template>
<div class="titleBar" @mouseover="buttonF=true">
  <div style="width: 30%; text-align: left;">
    <div v-if="UID !== -1">
      <Planlist style="display:inline-block;" />
      <div v-if="buttonF" style="display:inline-block;">
        <button @click="handleDelete" class="buttonPlanAction">Delete Plan</button>
        <button @click="handleEditTitle" class="buttonPlanAction">Edit Title</button>
      </div>
    </div>
  </div>
  <div style="width: 40%;">
    <div v-if="!editF && this.$store.state.Title !== ''">
      <h1 @click="handleEditTitle">{{this.$store.state.Title}}</h1>
    </div>
    <div v-if="editF || this.$store.state.Title === ''">
      <input
        v-select
        v-model="titleNew"
        v-on:keyup.enter="handleEditTitleSubmit"
        type="text"
        placeholder="Title"
      />
    </div>
  </div>
  <div style="width: 30%; height: 100%; text-align: right;">
    <div v-if="UID === -1">
      <button @click="handleLogin" class="loginButton">Log In</button>
      <button @click="handleSignup" class="signupButton">Sign Up</button>
    </div>
    <div v-if="UID > -1">
      <button @click="handleLogout" class="logoutButton">Log Out</button>
    </div>
  </div>
</div>
</template>

<script>
import Vue from 'vue';
import axios from 'axios';
import VueAxios from 'vue-axios';
import VueCookies from 'vue-cookies';
import { mapActions } from 'vuex';
import { defaultItemA, rebalanceItemA } from '../lib';

import Planlist from './Planlist.vue';

Vue.use(VueCookies);
Vue.use(VueAxios, axios);

export default {
  name: 'TitleView',
  data() {
    return {
      buttonF: false,
      editF: false,
      titleNew: this.$store.state.Title,
    };
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
    handleEditTitle() {
      this.editF = true;
    },
    handleEditTitleSubmit() {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      params.append('Title', this.titleNew);
      params.append('PID', this.$store.state.PID);
      axios.post(`${process.env.VUE_APP_AXIOS}title`, params)
        .then(() => {
          this.$store.dispatch('updateTitle', this.titleNew);
          document.title = `${this.titleNew} | Frenzy`;
          this.showF = true;
        });
      this.editF = false;
    },
    handleDelete() {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      params.append('PID', this.$store.state.PID);
      axios.post(`${process.env.VUE_APP_AXIOS}delete`, params)
        .then(() => {
          window.location.reload();
        });
    },
    handleLogin() {
      this.$store.dispatch('updatePage', 'login');
    },
    handleSignup() {
      this.$store.dispatch('updatePage', 'signup');
    },
    handleLogout() {
      window.$cookies.set('TokenString', '');
      this.$store.dispatch('updateTokenString', '');
      this.$store.dispatch('updatePage', 'login');
      this.$store.dispatch('updateUID', -1);
      this.$store.dispatch('updatePID', -1);
      this.$store.dispatch('updateTitle', 'Plan Your Next Frenzy');
      document.title = 'Plan Your Next Frenzy | Frenzy';
      const r = rebalanceItemA([], [], defaultItemA(), []);
      this.$store.dispatch('updateItemA', { itemA: r.itemA, dispA: r.dispA });
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
      inserted(el) {
        el.select();
      },
    },
  },
};
</script>

<style scoped>
.titleBar {
  width: 100%;
  text-align: center;
  border-bottom: 2px solid var(--baseTableBorder);
  margin-bottom: 10px;
  height: 42px;
  line-height: 42px;
  vertical-align: middle;
  display: flex;
}

.titleBar h1 {
  font-size: 1.4em;
  padding: 0;
  margin: 0;
}

.titleBar button {
  padding: 4px 10px;
  margin-left: 6px;
  font-weight: 600;
}

.buttonPlanAction {
  background-color: var(--buttonPlanActionBG);
  color: var(--buttonPlanActionFG);
  border: 2px solid var(--buttonPlanActionBO);
  border-radius: 5px;
}

.loginButton {
  background-color: var(--buttonLoginBG);
  color: var(--buttonLoginFG);
  border: 2px solid var(--buttonLoginBO);
  border-radius: 5px;
}

.signupButton {
  color: var(--buttonSignupFG);
  background-color: var(--buttonSignupBG);
  border: 2px solid var(--buttonSignupBO);
  border-radius: 5px;
}

.logoutButton {
  background-color: var(--buttonLogoutBG);
  color: var(--buttonLogoutFG);
  border: 2px solid var(--buttonLogoutBO);
  border-radius: 5px;
}
</style>
