<template>
<div class="loginBG">
  <div class="loginFG">
    <div class="form-title">Log In</div>
    <form @submit="handleSubmit" method="post">
      <div v-if="error != ''"><b>{{error}}</b></div>
      <div class="form-label">
        <p>Email</p>
        <input class="form-input" type="text" v-model="email" placeholder="Email" autofocus>
      </div>
      <div class="form-label">
        <p>Password</p>
        <input class="form-input" type="password" v-model="password" placeholder="Password">
      </div>
      <div class="form-label">
        <input class="form-button" type="submit" value="Submit" />
      </div>
    </form>
  </div>
</div>
</template>

<script>
import Vue from 'vue';
import { mapActions } from 'vuex';
import axios from 'axios';
import VueAxios from 'vue-axios';
import { rebalanceItemA } from '../lib';

Vue.use(VueAxios, axios);

export default {
  name: 'Login',
  data() {
    return {
      error: '',
      email: '',
      password: '',
    };
  },
  props: [],
  components: {},
  directives: {},
  methods: {
    ...mapActions([
      'updateUID',
      'updatePID',
      'updateTokenString',
      'updateItemA',
      'updatePage',
    ]),
    handleSubmit(event) {
      event.preventDefault();
      const params = new URLSearchParams();
      params.append('email', this.email);
      params.append('password', this.password);
      axios.post(`${process.env.VUE_APP_AXIOS}login`, params)
        .then((response) => {
          window.$cookies.set('TokenString', response.data.TokenString);
          this.$store.dispatch('updateTokenString', response.data.TokenString);
          this.$store.dispatch('updateUID', response.data.UID);
          this.$store.dispatch('updatePID', response.data.PID);
          this.$store.dispatch('updateTitle', response.data.Title);
          document.title = `${response.data.Title} | Frenzy`;
          const r = rebalanceItemA([], [], JSON.parse(response.data.ItemA), []);
          this.$store.dispatch('updateItemA', { itemA: r.itemA, dispA: r.dispA });
          this.$store.dispatch('updatePosA', [0]);
          this.$store.dispatch('updatePage', '');
        })
        .catch((err) => {
          this.error = err;
        });
    },
  },
  computed: {
  },
};
</script>

<style scoped>
.form-title {
  margin-top: 10px;
  text-align: center;
  font-weight: 600;
  font-size: 1.2em;
}

.form-label {
  margin: 12px 0px 15px 0px;
}

.form-label p {
  font-weight: 600;
  margin: 0px 0px 7px 0px;
  font-size: 1.0em;
}

.loginBG {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #FFF;
}

.loginFG {
  margin: 0 auto;
  margin-top: 30px;
  text-align: left;
  width: 250px;
}
</style>
