<template>
<div class="signupBG">
  <div class="signupFG">
    <div class="form-title">
      Sign Up
    </div>
    <form @submit="handleSubmit" method="post">
      <div v-if="error != ''">
        <b>{{error}}</b><br />
      </div>
      <div class="form-label">
        <p>Email</p>
        <input type="text" v-model="email" placeholder="Email">
      </div>
      <div class="form-label">
        <p>Password</p>
        <input type="password" v-model="password" placeholder="Password">
      </div>
      <div class="form-label">
        <input type="submit" value="Submit" />
      </div>
    </form>
  </div>
</div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)

export default {
  name: 'Signup',
  data: function () {
    return {
      email: '',
      password: '',
      error: '',
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
      params.append('email', this.email);
      params.append('password', this.password);
      axios.post(process.env.VUE_APP_AXIOS + 'signup', params)
        .then(response => {
          window.$cookies.set('TokenString', response.data.TokenString);
          this.$store.dispatch('updateTokenString', response.data.TokenString)
          this.$store.dispatch('updatePage', '')
          this.$store.dispatch('updateUID', response.data.UID)
          this.$store.dispatch('updatePID', response.data.PID)
          this.$store.dispatch('updateTitle', response.data.Title)
          this.$store.dispatch('updateItemA', JSON.parse(response.data.ItemA))
          this.$store.dispatch('updatePosA', [0])
        })
        .catch(err => {
          this.error = err;
        });
    }
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
  font-size: 16px;
}

.form-label {
  margin: 10px 0px 15px 0px;
}

.form-label p {
  font-weight: 600;
  margin: 0px 0px 5px 0px;
}

.signupBG {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #FFF;
}

.signupFG {
  margin: 0 auto;
  margin-top: 30px;
  text-align: left;
  width: 250px;
}

.signupFG input {
  z-index: 500;
  width: 100%;
}

</style>
