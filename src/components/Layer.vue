<template>
<div>
  <div class="layerLeftBorder">
    <Item
      v-for="disp in dispAN"
      v-bind:key="disp.pathDispA.toString()"
      v-bind:disp="disp"
      style="margin: 0px 0px 10px 0px; position: relative;"
    />

    <div class="disp dispSelectedButton" v-if="mode === 'normal' && posA[step] === dispAN.length">
      ADD ITEM
    </div>

    <div class="disp" v-if="mode === 'add' && posA[step] === dispAN.length">
      <input v-select type="text" @keyup.enter="addItem" />
    </div>

  </div>
</div>
</template>

<script>
import { mapActions } from 'vuex';
import { rebalanceItemA, addItemHelper } from '../lib.js'
import Item from './Item.vue'

export default {
  name: 'Layer',
  props: ['step', 'parentPathItemA'],
  components: {
    Item,
  },
  directives: {
    select: {
      inserted: function (el) {
        el.select()
      }
    }
  },
  methods: {
    ...mapActions([
      'updateItemA',
    ]),
    addItem (e) {
      let r;
      let itemA0 = addItemHelper(e.target.value, this.$store.state.itemA, this.$store.state.viewLabelA, this.$props.parentPathItemA, 0);
      r = rebalanceItemA([], [], itemA0, this.$store.state.viewLabelA);
      this.$store.dispatch('updateItemA', {itemA: r.itemA, dispA: r.dispA})
      // mode is automatically updated to 'normal' in the window listener
    },
  },
  computed: {
    mode() {
      return this.$store.state.mode;
    },
    modesub() {
      return this.$store.state.modesub;
    },
    viewLabelA() {
      return this.$store.state.viewLabelA;
    },
    dispAN() {
      let i;
      let dispAN = Array.from(this.$store.state.dispA);
      for (i = 0; i < this.$props.step; i += 1) {
        if (dispAN.length > 0 && this.$store.state.posA[i] < dispAN.length) {
          dispAN = dispAN[this.$store.state.posA[i]].childA;
        } else { // hit add new disp
          dispAN = [];
        }
      }
      return dispAN;
    },
    posA() {
      return this.$store.state.posA;
    }
  },
};
</script>

<style scoped>
.layerLeftBorder {
  border-left: 2px solid var(--layerLeftBorder);
  height: 100%;
}

input {
  margin: 1px 5px;
  width: 95%;
}

.disp {
  height: 36px;
  line-height: 36px;
  vertical-align: middle;
  width: 95%;
  border: 1px solid var(--baseTableBorder);
  border-left: none;
}

.dispSelectedButton {
  text-align: center;
  font-weight: 600;
  font-size: 0.8em;
  color: var(--dispSelectedButtonFG);
  background-color: var(--dispSelectedButtonBG);
}

</style>
