<template>
  <div>
    <button @click="handleList" style="margin-top: 6px; margin-left: 0px; border: 2px solid #000; background-color: #EEE; color: #000; font-weight: 600; border-radius: 3px; padding: 4px 10px;">Plans &#9662;</button>
    <div v-if="this.showF" class="planADiv">
      <div @click="handleAddPlan" class="planOption" style="border-bottom: 2px solid #CCC">
        <p>Add New Plan</p>
      </div>
      <div v-for="plan in PlanA" v-bind:key="plan.PID" @click="handleSelectPlan(plan.PID)" class="planOption">
        <p>{{plan.Title}}</p>
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


export default {
  name: 'Planlist',
  data: function () {
    return {
      showF: false,
      PlanA: [],
    }
  },
  props: [],
  components: {},
  directives: {},
  methods: {
    ...mapActions([
      'updateTitle',
      'updatePID',
      'updateItemA',
    ]),
    handleList: function () {
      if (!this.showF) {
        const params = new URLSearchParams();
        params.append('TokenString', window.$cookies.get('TokenString'));
        axios.post(process.env.VUE_APP_AXIOS + 'list', params)
          .then(response => {
            this.PlanA = response.data.PlanA;
            this.showF = true;
          })
      } else {
          this.showF = false;
      }
    },
    handleAddPlan: function () {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      axios.post(process.env.VUE_APP_AXIOS + 'add', params)
        .then(response => {
          this.$store.dispatch('updatePID', response.data.PID);
          this.$store.dispatch('updateTitle', response.data.Title);
          this.$store.dispatch('updateItemA', JSON.parse(response.data.ItemA));
          this.$store.dispatch('updatePosA', [0]);
          this.showF = false;
        })
    },
    handleSelectPlan: function (PID) {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      params.append('PID', PID);
      axios.post(process.env.VUE_APP_AXIOS + 'load', params)
        .then(response => {
          this.$store.dispatch('updatePID', response.data.PID);
          this.$store.dispatch('updateTitle', response.data.Title);
          this.$store.dispatch('updateItemA', JSON.parse(response.data.ItemA));
          this.$store.dispatch('updatePosA', [0]);
          this.showF = false;
        })
    },
  },
  computed: {},
};
</script>

<style scoped>
.planADiv {
  position: absolute;
  height: 300px;
  width: 240px;
  left: 14px;
  top: 32px;
  background-color: #FFF;
  border: 2px solid black;
  z-index: 500;
}

.planOption {
  width: 100%;
  height: 30px;
  line-height: 30px;
  vertical-align: middle;
  background-color: #FFF;
  cursor: pointer;
}

.planOption p {
  margin: 0;
  padding: 0;
  padding-left: 10px;
}

.planOption:hover {
  background: green;
}
</style>
