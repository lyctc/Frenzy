<template>
  <div>
    <button @click="handleList" class="buttonPlans">Plans &#9662;</button>
    <div v-if="this.showF" class="planADiv">
      <div @click="handleAddPlan" class="planOption" style="border-bottom: 2px solid #CCC">
        <p>Add New Plan</p>
      </div>
      <div
        v-for="plan in PlanA"
        v-bind:key="plan.PID"
        @click="handleSelectPlan(plan.PID)"
        class="planOption"
      >
        <p>{{plan.Title}}</p>
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
import { rebalanceItemA } from '../lib';

Vue.use(VueCookies);
Vue.use(VueAxios, axios);

export default {
  name: 'Planlist',
  data() {
    return {
      showF: false,
      PlanA: [],
    };
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
    handleList() {
      if (!this.showF) {
        const params = new URLSearchParams();
        params.append('TokenString', window.$cookies.get('TokenString'));
        axios.post(`${process.env.VUE_APP_AXIOS}list`, params)
          .then((response) => {
            this.PlanA = response.data.PlanA;
            this.showF = true;
          });
      } else {
        this.showF = false;
      }
    },
    handleAddPlan() {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      axios.post(`${process.env.VUE_APP_AXIOS}add`, params)
        .then((response) => {
          this.$store.dispatch('updatePID', response.data.PID);
          this.$store.dispatch('updateTitle', response.data.Title);
          document.title = `${response.data.Title} | Frenzy`;
          const r = rebalanceItemA([], [], JSON.parse(response.data.ItemA), []);
          this.$store.dispatch('updateItemA', { itemA: r.itemA, dispA: r.dispA });
          this.$store.dispatch('updatePosA', [0]);
          this.showF = false;
        });
    },
    handleSelectPlan(PID) {
      const params = new URLSearchParams();
      params.append('TokenString', window.$cookies.get('TokenString'));
      params.append('PID', PID);
      axios.post(`${process.env.VUE_APP_AXIOS}load`, params)
        .then((response) => {
          this.$store.dispatch('updatePID', response.data.PID);
          this.$store.dispatch('updateTitle', response.data.Title);
          document.title = `${response.data.Title} | Frenzy`;
          const r = rebalanceItemA([], [], JSON.parse(response.data.ItemA), []);
          this.$store.dispatch('updateItemA', { itemA: r.itemA, dispA: r.dispA });
          this.$store.dispatch('updatePosA', [0]);
          this.showF = false;
        });
    },
  },
  computed: {},
};
</script>

<style scoped>
.buttonPlans {
  background-color: var(--buttonPlansBG);
  border: 2px solid var(--buttonPlansBO);
  color: var(--buttonPlansFG);
  font-weight: 600;
  border-radius: 5px;
  padding: 4px 10px;
  outline:none;
}

.planADiv {
  position: absolute;
  height: 300px;
  width: 240px;
  left: 14px;
  top: 36px;
  background-color: #FFF;
  border: 2px solid #888;
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
  background-color: var(--plansSelectedBG);
  color: var(--plansSelectedFG);
}
</style>
