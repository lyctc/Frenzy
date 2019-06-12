<template>
<div>
  <div style="border-left: 2px solid #000;">
    <div v-for="item in itemA"
      v-bind:key="item.pathA.toString()"
      style="margin: 0px 0px 10px 0px;"
    >
      <div
        class="item"
        :class="{ 'itemBlue' : posA.join(',').startsWith(item.pathA.join(',')), 'itemExpand' : posA.join(',').startsWith(item.pathA.join(',')) && item.childA.length > 0}"
      >
        <div v-if="!(mode === 'edit' && posA.join(',').startsWith(item.pathA.join(',')) && item.pathA.toString() === posA.toString())" class="itemTitle">
          {{item.title}}
        </div>
        <div v-if="mode === 'edit' && posA.join(',').startsWith(item.pathA.join(',')) && item.pathA.toString() === posA.toString()">
          <input v-select type="text" :value="item.title" @keyup.enter="updateItem" />
        </div>
      </div>
      <div v-if="mode === 'selected' && posA.join(',').startsWith(item.pathA.join(',')) && item.pathA.toString() === posA.toString()" class="modesubBar">
        <div class="modesub1" :class="{ 'modesub' : modesub === 'move' }">Move</div>
        <div class="modesub1" :class="{ 'modesub' : modesub === 'edit' }">Edit</div>
        <div class="modesub1" :class="{ 'modesub' : modesub === 'delete' }">Delete</div>
      </div>
    </div>

    <div class="item" v-if="mode === 'normal' && pos === itemA.length">
      <div
        class="addButton itemTitle"
        :class="{ 'itemBlue': pos === itemA.length }"
      >
        Add Item
      </div>
    </div>

    <div class="item" v-if="mode === 'add' && pos === itemA.length">
      <input v-select type="text" @keyup.enter="addItem" />
    </div>

  </div>
</div>
</template>

<script>
import { mapActions } from 'vuex';

function addItemHelper(title, itemA0, posA, step) {
  // recursively go through itemA and add new item
  if (step === posA.length - 1) {
    // posA[posA.length - 1] doesn't exist yet, so must push
    itemA0.push({
      pathA: posA,
      iid: 0,
      title: title,
      labelA: [],
      childA: [],
    });
    return itemA0;
  } else {
    itemA0[posA[step]].childA = addItemHelper(title, itemA0[posA[step]].childA, posA, step + 1);
  }
  return itemA0;
}

function updateItemHelper(title, itemA0, posA, step) {
  // recursively go through itemA and update title to new title
  if (step === posA.length - 1) {
    itemA0[posA[step]].title = title;
    return itemA0;
  } else {
    itemA0[posA[step]].childA = updateItemHelper(title, itemA0[posA[step]].childA, posA, step + 1);
  }
  return itemA0;
}

export default {
  name: 'app',
  props: ['itemA', 'pos'],
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
      let itemA0 = Array.from(this.$store.state.itemA);
      itemA0 = addItemHelper(e.target.value, itemA0, this.$store.state.posA, 0);
      this.$store.dispatch('updateMode', {mode: 'normal', modesub: ''});
      this.$store.dispatch('updateItemA', itemA0)
    },
    updateItem (e) {
      let itemA0 = Array.from(this.$store.state.itemA);
      itemA0 = updateItemHelper(e.target.value, itemA0, this.$store.state.posA, 0);
      this.$store.dispatch('updateMode', {mode: 'normal', modesub: ''});
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
    posA() {
      return this.$store.state.posA;
    }
  },
};
</script>

<style scoped>
.addButton {
  font-weight: 600;
  font-size: 0.8em;
  color: #999;
}

.item {
  height: 36px;
  line-height: 36px;
  vertical-align: middle;
  width: 95%;
  border: 1px solid #CCC;
  border-left: none;
}

.itemTitle {
  padding-left: 7px;
}

.itemBlue {
  background-color: #CCFFFF;
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
  z-index: 100;
  background-color: #CCC;
  width: 95%;
  height: 16px;
}

.modesub1 {
  text-align: center;
  width: 33.333%;
  z-index: 100;
  color: #654321;
  background-color: #CCC;
  line-height: 16px;
  font-size: 12px;
}

.modesub {
  background-color: #333;
}
</style>
