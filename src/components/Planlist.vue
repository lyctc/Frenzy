<template>
<div class="planADiv">
  <a href="javascript:void(0)" @click="handleAddPlan">Add New Plan</a>
  <div v-for="plan in PlanA" v-bind:key="plan.pid">
    <a href="javascript:void(0)" @click="handleSelectPlan">{{plan.title}}</a>
  </div>
</div>
</template>

<script>
import Vue from 'vue'
import { mapActions } from 'vuex';

export default {
  name: 'planlist',
  props: [],
  components: {},
  directives: {},
  methods: {
    handleAddPlan: function (event) {
      console.log('-- handleAddPlan');
    },
    handleSelectPlan: function (event) {
      console.log('-- handleSelectPlan');
    },
  },
  computed: {
    PlanA() {
      return this.$store.state.PlanA;
    },
  },
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
}
</style>
