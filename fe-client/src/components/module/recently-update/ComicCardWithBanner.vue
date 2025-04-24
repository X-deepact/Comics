<template>
  <Card unstyled>
    <template #title>
      <div
        :class="[
          'overflow-hidden hover:shadow-[0_6px_6px_rgba(0,3,4,0.25)] cursor-pointer rounded-md relative px-2 md:px-0',
          !width ? 'w-full' : '',
        ]"
        :style="containerStyle"
      >
        <img
          :src="props.item.cover"
          alt="banner-image"
          class="w-full h-full object-cover rounded-md"
          @click="onNavigateToComicDetails(props.item.id)"
        />
        <div
          v-if="isHavingOverlayText"
          class="z-10 absolute left-2 bottom-2 px-2 max-w-1/2 text-[12px] leading-[22px] text-white bg-black/65 overflow-hidden whitespace-nowrap text-ellipsis"
        >
          {{ overlayingText }}
        </div>
        <div
          v-if="isShowRating"
          class="absolute right-2 bottom-2 px=2 text-[18px] leading-[26px] text-[#ff620e] z-[999]"
        >
          {{ ratingText }}
        </div>
      </div>
    </template>
    <template #subtitle>
      <slot name="comicDescription" />
    </template>
  </Card>
</template>

<script setup lang="ts">
import { computed } from "vue";
import type { ComicDetailResponse } from "@/services/comics/types.ts";
import router from "@/router";
import { useResponsive } from "@/composables/useResponsive";

const props = defineProps<{
  item: ComicDetailResponse;
  width?: string;
  height?: string;
  isHavingOverlayText?: boolean;
  overlayingText?: string;
  isShowRating?: boolean;
  ratingText?: string | number;
}>();

const containerStyle = computed(() => {
  const style: Record<string, string> = {};

  if (props.width) {
    style.width = props.width;
  }
  if (props.height) {
    style.height = props.height;
  }

  return style;
});

const onNavigateToComicDetails = (id: number) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString(),
    },
  });
};
</script>
