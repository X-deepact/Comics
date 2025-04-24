<template>
  <template v-if="!hasItems">
    {{ "Loading Topping Comic" }}
  </template>
  <div
    v-else
    class="grid gap-1 items-stretch"
    :class="[{ 'grid-cols-3': isLargeScreen }]"
  >
    <div>
      <Galleria
        unstyled
        ref="galleriaRef"
        :value="comicAds"
        :auto-play="true"
        :circular="true"
        :show-item-navigators="false"
        :show-thumbnails="false"
        :transitionInterval="4000"
      >
        <template #item="{ item }">
          <div
            class="relative overflow-hidden hover:cursor-pointer h-[20.5rem]"
            :class="[{ 'px-2': isMobile }]"
          >
            <img
              class="w-full h-full object-cover rounded-md"
              :src="item.ads_image"
              :alt="item.comic_id"
              @click="
                onNavigateToComicDetail(
                  item.comic_id,
                  item.ads_type,
                  item.ads_direct_url
                )
              "
            />
          </div>
        </template>
      </Galleria>
    </div>

    <div v-if="isLargeScreen" class="flex flex-col gap-1">
      <ComicCardWithBanner
        v-if="items?.[0]"
        class="mb-1"
        height="10rem"
        :item="items[0]"
        isHavingOverlayText
        :overlaying-text="getChapterText(items[0])"
        @click="onNavigateToComic(items[0].id)"
      />
      <div class="grid grid-cols-2 gap-1">
        <ComicCardWithBanner
          v-if="items?.[1]"
          height="10rem"
          :item="items[1]"
          isHavingOverlayText
          :overlaying-text="getChapterText(items[1])"
          @click="onNavigateToComic(items[1].id)"
        />
        <ComicCardWithBanner
          v-if="items?.[2]"
          height="10rem"
          :item="items[2]"
          isHavingOverlayText
          :overlaying-text="getChapterText(items[2])"
          @click="onNavigateToComic(items[2].id)"
        />
      </div>
    </div>

    <div v-if="isLargeScreen" class="flex flex-col gap-1">
      <div class="grid grid-cols-2 mb-1 gap-1">
        <ComicCardWithBanner
          v-if="items?.[3]"
          height="10rem"
          :item="items[3]"
          isHavingOverlayText
          :overlaying-text="getChapterText(items[3])"
          @click="onNavigateToComic(items[3].id)"
        />
        <ComicCardWithBanner
          v-if="items?.[4]"
          height="10rem"
          :item="items[4]"
          isHavingOverlayText
          :overlaying-text="getChapterText(items[4])"
          @click="onNavigateToComic(items[4].id)"
        />
      </div>
      <ComicCardWithBanner
        v-if="items?.[5]"
        height="10rem"
        :item="items[5]"
        isHavingOverlayText
        :overlaying-text="getChapterText(items[5])"
        @click="onNavigateToComic(items[5].id)"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import ComicCardWithBanner from "@/components/module/recently-update/ComicCardWithBanner.vue";
import { storeToRefs } from "pinia";
import { useComicAds } from "@/stores/ads/useComicAds.ts";
import type { ComicDetailResponse } from "@/services/comics/types.ts";
import { computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useResponsive } from "@/composables/useResponsive";

const { isDoubleExtraSmallAndExtraSmallScreen, isSmallAndMiddleScreen, isLargeScreen } = useResponsive();
const isMobile = isDoubleExtraSmallAndExtraSmallScreen || isSmallAndMiddleScreen;
const router = useRouter();
const store = useComicAds();
const { comicAds } = storeToRefs(store);
const props = defineProps<{
  items: ComicDetailResponse[] | undefined;
}>();
const hasItems = computed(
  () => Array.isArray(props.items) && props.items.length > 0
);

onMounted(() => {
  store.fetchAds();
});
function getChapterText(item?: ComicDetailResponse) {
  if (item?.chapters?.[0]) {
    return `Chapter ${item.chapters[0].number}${item.chapters[0].name ? `: ${item.chapters[0].name}` : ""}`;
  }
  return "";
}

const onNavigateToComic = (id: number) => {
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString(),
    },
  });
};

const onNavigateToComicDetail = (
  id: number,
  ads_type: String,
  ads_direct_url: String
) => {
  if (ads_type == "internal") {
    router.push({
      name: "comic-detail",
      params: {
        comicId: id.toString(),
      },
    });
  } else {
    window.open(ads_direct_url as string, "_blank");
  }
};
</script>
