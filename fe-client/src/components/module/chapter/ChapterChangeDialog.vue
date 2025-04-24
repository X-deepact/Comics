<template>
  <Dialog
    v-model:visible="visible"
    header="Chuyen Chuong"
  >

  </Dialog>
</template>
<script setup lang="ts">

import { ref } from "vue";

const visible = ref(false);

const toggleVisible = () => {
  visible.value = !visible.value;
}

defineExpose({
  toggleVisible
})

</script>

<style scoped>

</style>