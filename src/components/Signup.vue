<template>
<div>
  <h4>Sign Up</h4>
  <form @submit="handleSubmit" method="post">
    <b>Email</b><br />
    <input type="text" v-model="email" placeholder="Email"><br />
    <b>Password</b><br />
    <input type="password" v-model="password" placeholder="Password"><br />
    <input type="submit" value="Submit" />
  </form>
</div>
</template>

<script>
import Vue from 'vue'
import { mapActions } from 'vuex';
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)

export default {
  name: 'signup',
  data: function () {
    return {
      email: '',
      password: '',
    }
  },
  props: [],
  components: {},
  directives: {
    select: {
      inserted: function (el) {
        el.select()
      }
    }
  },
  methods: {
    handleSubmit: function (event) {
      event.preventDefault();
      const params = new URLSearchParams();
      params.append("email", this.email);
      params.append("password", this.password);
      axios.post("http://localhost:3030/signup", params)
        .then(response => {
          window.$cookies.set("tokenString", response['data']['tokenString'])
          location.reload();
        })
      .catch(err => console.log(err));
    }
  },
  computed: {
  },
};
</script>

<style scoped>
</style>
