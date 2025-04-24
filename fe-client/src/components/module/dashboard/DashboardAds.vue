<template>
  <div class="w-full my-10 inline-block">
    <Image :src="props.adImage" alt="image-banner" />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  adImage: string,
}>();

</script>
