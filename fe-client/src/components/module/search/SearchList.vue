<template>
  <template v-if="!props.data || props.data.length === 0">
    <ListComicSkeleton />
  </template>

  <div v-else class="grid grid-cols-6 gap-x-8 gap-y-4">
    <div v-for="item in props.data" :key="item.id">
      <ComicCardWithBanner
        :item="item"
        width="245px"
        height="21rem"
      >
        <template #comicDescription>
          <div class="mt-2 overflow-hidden">
            <div
                class="text-gray-800 font-bold text-sm line-clamp-2 hover:text-[#ff620e] hover:cursor-pointer"
                @click="onNavigateToComicDetails(item.id)"
            >
              {{ item.name }}
            </div>
            <div
                v-for="chapter in item.chapters"
                :key="chapter.id"
            >
              <span class="text-sm font-semibold hover:text-[#ff620e] hover:cursor-pointer">
                Chapter {{chapter.number}}
              </span>
              <span class="text-sm text-gray-500"> - {{ getTimeAgo(chapter.created_at) }}</span>
            </div>
          </div>
        </template>
      </ComicCardWithBanner>
    </div>
  </div>
</template>

<script setup lang="ts">
import ComicCardWithBanner from "@/components/module/recently-update/ComicCardWithBanner.vue";
import type { ComicDetailResponse } from "@/services/comics/types.ts";
import ListComicSkeleton from "@/components/skeleton/ListComicSkeleton.vue";
import {getTimeAgo} from "@/utils/date-formatted.ts";
import router from "@/router";

const props = defineProps<{
  data: ComicDetailResponse[];
}>();

const onNavigateToComicDetails = (id: number) => {
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString()
    }
  });
}
</script>
