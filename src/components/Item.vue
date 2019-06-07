<template>
<div class="item">
  {{todo.title}}
</div>
</template>

<script>
export default {
  name: 'app',
  props: [
    'todo'
  ],
}
</script>

<style scoped>
.item {
  margin: 5px 0px;
  width: 200px;
  padding: 5px;
}
</style>
