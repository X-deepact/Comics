<template>
  <div class="mx-2 sm:mx-0 mb-10">
    <div class="w-full inline-flex items-center justify-between mb-5">
      <div class="inline-flex gap-2">
        <Image :src="props.icon" alt="icon" />
        <span class="text-2xl">{{ props.text }}</span>
      </div>
      <Button
        v-if="!props.hideBtn"
        class="hover:text-[#ff620e]"
        size="small"
        icon="pi pi-angle-right"
        icon-pos="right"
        :label="$t(props.navBtnText)"
        @click="emits('navigate')"
      />
    </div>
    <slot name="main" />
  </div>
</template>

<script setup lang="ts">

interface ComponentProps {
  text: string;
  icon: string;
  navBtnText?: string;
  hideBtn?: boolean;
}

const props = withDefaults(defineProps<ComponentProps>(), {
  text: "独家作品",
  icon: "src/assets/images/ic__dj.png",
  navBtnText: "dashboard.more",
  hideBtn: false,
});
const emits = defineEmits(["navigate"]);
</script>
