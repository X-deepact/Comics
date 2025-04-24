<template>
  <div
    :class="['relative overflow-hidden cursor-pointer rounded-md', containerClass]"
    :style="customStyle"
  >
    <Image
      :src="props.src"
      :alt="alt"
      class="w-full h-auto object-contain"
    />
    <div class="absolute left-2 bottom-2 px-2 max-w-[130px] text-[12px]
    leading-[22px] text-white bg-black/65 z-[999] overflow-hidden whitespace-nowrap
    text-ellipsis"
    >
      {{ props.overlayText }}
    </div>
  </div>
</template>
<script setup lang="ts">
import { useRouter } from "vue-router";
import { computed } from "vue";

const props = defineProps({
  src: { type: String, required: true },
  alt: { type: String, required: true },
  path: { type: String, required: false },
  overlayText: { type: String, default: '' },
  containerClass: { type: String, default: '' },
  width: { type: String, default: '' },
  height: { type: String, default: '' },
});

const router = useRouter();

const customStyle = computed(() => {
  const style: Record<string, string> = {};
  if (props.width) {
    style.width = props.width;
  }
  if (props.height) {
    style.height = props.height;
  }
  return style;
});

</script>
