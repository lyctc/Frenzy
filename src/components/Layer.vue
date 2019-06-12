<template>
<div>
 <div class="layerLeftBorder">
    <div v-for="item in itemAN"
      v-bind:key="item.pathA.toString()"
      style="margin: 0px 0px 10px 0px;"
    >
      <div
        class="item"
        :class="{
          'itemAncestor': posA.join(',').startsWith(item.pathA.join(',')) && posA.join(',') !== item.pathA.join(','),
          'itemSelected': posA.join(',') === item.pathA.join(','),
          'itemExpand': posA.join(',').startsWith(item.pathA.join(',')) && !(posA.join(',') === item.pathA.join(',') && item.childA.length === 0)
        }"
      >
        <div v-if="!(mode === 'edit' && posA.join(',').startsWith(item.pathA.join(',')) && item.pathA.toString() === posA.toString())" class="itemTitle">
          {{item.title}}
        </div>
        <div v-if="mode === 'edit' && posA.join(',').startsWith(item.pathA.join(',')) && item.pathA.toString() === posA.toString()">
          <input v-select type="text" :value="item.title" @keyup.enter="updateItem" />
        </div>
      </div>
      <div v-if="mode === 'selected' && posA.join(',').startsWith(item.pathA.join(',')) && item.pathA.toString() === posA.toString()" class="modesubBar">
        <div class="modeUnselected" :class="{ 'modeSelected' : modesub === 'move' }">Move</div>
        <div class="modeUnselected" :class="{ 'modeSelected' : modesub === 'edit' }">Edit</div>
        <div class="modeUnselected" :class="{ 'modeSelected' : modesub === 'delete' }">Delete</div>
      </div>
    </div>

    <div class="item itemSelectedButton" v-if="mode === 'normal' && posA[step] === itemAN.length">
      ADD ITEM
    </div>

    <div class="item" v-if="mode === 'add' && posA[step] === itemAN.length">
      <input v-select type="text" @keyup.enter="addItem" />
    </div>

  </div>
</div>
</template>

<script>
import { mapActions } from 'vuex';
import { rebalancePathA, addItemHelper, updateItemHelper } from '../lib.js'

export default {
  name: 'app',
  props: ['step'],
  components: {},
  directives: {
    select: {
      inserted: function (el) {
        el.select()
      }
    }
  },
  methods: {
    ...mapActions([
      'updateMode',
      'updatePosA',
    ]),
    addItem (e) {
      let _;
      let itemA0 = Array.from(this.$store.state.itemA);
      itemA0 = addItemHelper(e.target.value, itemA0, this.$store.state.posA, 0);
      this.$store.dispatch('updateMode', {mode: 'normal', modesub: ''});
      _, itemA0 = rebalancePathA([], itemA0);
      this.$store.dispatch('updateItemA', itemA0)
    },
    updateItem (e) {
      let _;
      let itemA0 = Array.from(this.$store.state.itemA);
      itemA0 = updateItemHelper(e.target.value, itemA0, this.$store.state.posA, 0);
      this.$store.dispatch('updateMode', {mode: 'normal', modesub: ''});
      _, itemA0 = rebalancePathA([], itemA0);
      this.$store.dispatch('updateItemA', itemA0)
    },
  },
  computed: {
    mode() {
      return this.$store.state.mode;
    },
    modesub() {
      return this.$store.state.modesub;
    },
    itemAN() {
      let i;
      let itemAN = Array.from(this.$store.state.itemA);
      for (i = 0; i < this.$props.step; i += 1) {
        if (itemAN.length > 0 && this.$store.state.posA[i] < itemAN.length) {
          itemAN = itemAN[this.$store.state.posA[i]].childA;
        } else { // hit add new item
          itemAN = [];
        }
      }
      return itemAN;
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

.item {
  height: 36px;
  line-height: 36px;
  vertical-align: middle;
  width: 95%;
  border: 1px solid var(--baseTableBorder);
  border-left: none;
}

.itemTitle {
  padding-left: 7px;
}

.itemSelected {
  color: var(--itemSelectedFG);
  background-color: var(--itemSelectedBG);
}

.itemAncestor {
  color: var(--itemAncestorFG);
  background-color: var(--itemAncestorBG);
}

.itemSelectedButton {
  text-align: center;
  font-weight: 600;
  font-size: 0.8em;
  color: var(--itemSelectedButtonFG);
  background-color: var(--itemSelectedButtonBG);
}

.itemExpand {
  width: 100%;
  border-right: none;
}

input {
  margin: 1px 5px;
  width: 95%;
}

.modesubBar {
  display: flex;
  border-right: 1px solid var(--baseTableBorder);
  border-bottom: 1px solid var(--baseTableBorder);
  z-index: 100;
  width: 95%;
  height: 16px;
}

.modeUnselected {
  text-align: center;
  width: 33.333%;
  z-index: 100;
  color: var(--modeUnselectedFG);
  background-color: var(--modeUnselectedBG);
  line-height: 16px;
  font-size: 12px;
}

.modeSelected {
  color: var(--modeSelectedFG);
  background-color: var(--modeSelectedBG);
}
</style>
