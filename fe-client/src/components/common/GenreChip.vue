<template>
  <Chip @click="handleClick">
    <span
      class="bg-white text-[#ff6600] rounded-full w-fit h-6 flex border border-[#ff6600]
      items-center justify-center px-2 hover:bg-[#ff6600] hover:text-white
      cursor-pointer"
    >
     {{ content }}
    </span>
  </Chip>
</template>
<script setup lang="ts">
const props = defineProps<{
  content: string
  isNavigating?: boolean
}>();
const emits = defineEmits(["click"]);
const handleClick = () => {
  if (props.isNavigating) {
    emits("click");
  }
}
</script>
