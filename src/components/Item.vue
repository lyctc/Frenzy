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
</style>
