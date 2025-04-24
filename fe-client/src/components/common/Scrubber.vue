<template>
  <div ref="scrubber"
       class="relative h-2 rounded cursor-pointer select-none bg-black dark:bg-white dark:bg-opacity-10 bg-opacity-20"
       @mousedown="scrubbing = true">
    <div class="relative overflow-hidden h-full w-full rounded">
      <div
          :style="{ transform: `translateX(${secondary / max * 100 - 100}%)` }"
          class="h-full absolute opacity-30 left-0 top-0 bg-orange-500 w-full rounded"
      />
      <div
          :style="{ transform: `translateX(${value / max * 100 - 100}%)` }"
          class="relative h-full w-full bg-orange-500 rounded"
      />
    </div>
    <div :class="{ '!opacity-100': scrubbing }" class="absolute inset-0 hover:opacity-100 opacity-0">
      <slot
          :pending-value="pendingValue"
          :position="`${Math.max(0, Math.min(elementX, elementWidth))}px`"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import {useEventListener, useMouseInElement, useVModel} from '@vueuse/core'
import {ref as deepRef, shallowRef, watch} from 'vue'

const props = defineProps({
  min: {type: Number, default: 0},
  max: {type: Number, default: 100},
  secondary: {type: Number, default: 0},
  modelValue: {type: Number, required: true},
})

const emit = defineEmits(['update:modelValue'])

const scrubber = deepRef()
const scrubbing = shallowRef(false)
const pendingValue = shallowRef(0)

useEventListener('mouseup', () => scrubbing.value = false, {passive: true})

const value = useVModel(props, 'modelValue', emit);
const {elementX, elementWidth} = useMouseInElement(scrubber);

watch([scrubbing, elementX], () => {
  const progress = Math.max(0, Math.min(1, (elementX.value) / elementWidth.value))
  pendingValue.value = progress * props.max
  if (scrubbing.value)
    value.value = pendingValue.value
})
</script>
