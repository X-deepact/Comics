<template>
  <div class="overflow-hidden relative w-full">
    <canvas
      ref="baseCanvas"
      class="relative w-full h-auto"
      :width="1024"
      :height="1537"
    ></canvas>
<!--    <div class="absolute w-full h-full top-0 left-0">-->
<!--      <div class="max-w-full xl:max-w-none h-full flex flex-col">-->
<!--        <canvas-->
<!--          :width="0"-->
<!--          :height="0"-->
<!--          class="w-full pointer-events-none"-->
<!--        >-->
<!--        </canvas>-->
<!--      </div>-->
<!--    </div>-->
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";

const props = defineProps<{
  imageSrc: string
}>();
const baseCanvas = ref(null)

onMounted(() => {
  const baseCtx = baseCanvas.value.getContext('2d')

  const img = new Image()
  img.src = props.imageSrc
  img.onload = () => {
    baseCtx.drawImage(img, 0, 0, 1024, 1537)
  }
})
</script>

<style scoped>

</style>