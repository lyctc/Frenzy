<template>
<div>
  <div
    class="disp"
    :class="{
      'dispAncestor': posA.join(',').startsWith(disp.pathDispA.join(',')) && posA.join(',') !== disp.pathDispA.join(','),
      'dispSelected': posA.join(',') === disp.pathDispA.join(','),
      'dispMove': posA.join(',') === disp.pathDispA.join(',') && mode === 'move',
      'dispExpand': posA.join(',').startsWith(disp.pathDispA.join(',')) && !(posA.join(',') === disp.pathDispA.join(',') && disp.childA.length === 0)
    }"
  >
    <div v-if="!(mode === 'edit' && posA.join(',').startsWith(disp.pathDispA.join(',')) && disp.pathDispA.toString() === posA.toString())" class="dispTitle">
      {{disp.title}}
      <div v-if="disp.childA.length !== 0" class="dispChildren">
        {{disp.childA.length}}
      </div>
    </div>
    <div v-if="mode === 'edit' && posA.join(',').startsWith(disp.pathDispA.join(',')) && disp.pathDispA.toString() === posA.toString()">
      <input v-select type="text" :value="disp.title" @keyup.enter="updateItem" />
    </div>
  </div>
  <div v-if="mode === 'selected' && posA.join(',').startsWith(disp.pathDispA.join(',')) && disp.pathDispA.toString() === posA.toString()" class="modesubBar">
    <div class="modeUnselected" :class="{ 'modeSelected' : modesub === 'move' }">Move</div>
    <div class="modeUnselected" :class="{ 'modeSelected' : modesub === 'edit' }">Edit</div>
    <div class="modeUnselected" :class="{ 'modeSelected' : modesub === 'delete' }">Delete</div>
  </div>
  <div style="display: flex;width: 95%;">
    <div class="dispLabel" :class="{'dispLabel1': disp.labelA.indexOf(1) !== -1 }"></div>
    <div class="dispLabel" :class="{'dispLabel2': disp.labelA.indexOf(2) !== -1 }"></div>
    <div class="dispLabel" :class="{'dispLabel3': disp.labelA.indexOf(3) !== -1 }"></div>
  </div>
</div>
</template>

<script>
import { mapActions } from 'vuex';
import { rebalanceItemA, updateItemHelper } from '../lib.js'

export default {
  name: 'Item',
  props: ['disp'],
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
      'updatePosA',
    ]),
    updateItem (e) {
      let itemA0 = Array.from(this.$store.state.itemA);
      itemA0 = updateItemHelper(e.target.value, itemA0, this.$props.disp.pathItemA, 0);
      let r = rebalanceItemA([], [], itemA0, this.$store.state.viewLabelA);
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
    posA() {
      return this.$store.state.posA;
    }
  },
};
</script>

<style scoped>
.disp {
  height: 36px;
  line-height: 36px;
  vertical-align: middle;
  width: 95%;
  border: 1px solid var(--baseTableBorder);
  border-left: none;
}

.dispTitle {
  padding-left: 7px;
}

.dispChildren {
  position: absolute;
  display: inline-block;
  left: 260px;
  color: var(--dispChildren);
  font-weight: 600;
}

.dispSelected {
  color: var(--dispSelectedFG);
  background-color: var(--dispSelectedBG);
}

.dispMove {
  color: var(--dispMoveFG);
  background-color: var(--dispMoveBG);
}

.dispAncestor {
  color: var(--dispAncestorFG);
  background-color: var(--dispAncestorBG);
}

.dispLabel {
  width: 33.3333%;
  height: 3px;
}

.dispLabel1 {
  background-color: red;
}

.dispLabel2 {
  background-color: green;
}

.dispLabel3 {
  background-color: blue;
}

.dispExpand {
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
  font-size: 1.0em;
}

.modeSelected {
  color: var(--modeSelectedFG);
  background-color: var(--modeSelectedBG);
}
</style>
