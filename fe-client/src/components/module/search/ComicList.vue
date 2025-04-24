<template>
  <div class="card">
    <DataView layout="grid" :value="data">
      <template #grid="slotProps">
        <div class="grid grid-cols-12 gap-2 mx-2 sm:mx-0">
          <div
              v-for="(item, index) in slotProps.items"
              :key="index"
              class="col-span-12 sm:col-span-6 md:col-span-4 xl:col-span-2"
          >
            <div class="bg-gray-100 sm:bg-transparent dark:bg-surface-900 rounded flex flex-col">
              <div
                  class="bg-surface-50 dark:bg-black flex flex-row sm:flex-col justify-center rounded gap-2 sm:gap-0"
              >
              <div
                  class="w-1/3 sm:w-auto mx-auto hover:cursor-pointer hover:shadow-[0_6px_6px_rgba(0,3,4,0.25)] relative"
                  @click="item.type === 'drama' ? onNavigateToShortDetails(item.id) : onNavigateToComicDetails(item.id)"
                >
                  <div 
                    v-if="item.type === 'drama'"
                    class="absolute bg-orange-500 text-white text-xs font-bold px-2 py-1 rounded"
                  >
                    {{ $t('short.pageTitle') }}
                  </div>
                  <img
                      class="rounded w-full max-w-[300px]"
                      :src="item.cover ? item.cover : item.thumb"
                      :alt="item.name"
                  />
                </div>
                <div
                    class="w-2/3 sm:w-full flex flex-col xl:flex-row justify-between xl:items-center flex-1 gap-6 mt-2"
                >
                  <div class="flex flex-col justify-between items-start gap-2">
                    <div>
                      <div v-if="item.authors" class="font-medium text-surface-500 text-surface-400 text-sm mt-1">
                        {{ $t('search.author') }}: {{ item.authors.map(a => a.name).join(', ') }}
                      </div>
                      <div class="text-xl font-bold cursor-pointer hover:text-orange-500"
                      @click="item.type === 'drama' ? onNavigateToShortDetails(item.id) : onNavigateToComicDetails(item.id)"
                      >
                        {{ item.name }}
                      </div>
                      <div class="text-sm font-medium mt-1">
                        {{ item.description.length > 30 ? item.description.slice(0, 30) + '...' : item.description }}
                      </div>
                      <div v-if="item.latest_episode" class="text-sm font-medium mt-1">
                        {{ $t('search.latestEpisode') }}: {{ item.latest_episode }}
                      </div>
                    </div>
                    <div
                        v-if="!isDoubleExtraSmallAndExtraSmallScreen"
                        class="flex flex-col justify-between gap-1 mt-2"
                    >
                      <div
                          v-for="(chap, index) in item.chapters"
                          :key="index"
                          class="text-sm font-medium cursor-pointer hover:text-orange-500"
                          @click="onNavigateToChapterDetails(item.id, chap.id)"
                      >
                        <span>{{ `Chap ${chap.number}: ${chap.name}` }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>
    </DataView>
  </div>

</template>

<script setup lang="ts">
import {useResponsive} from "@/composables/useResponsive.ts";
import type {ComicDetailResponse} from "@/services/comics/types.ts";
import router from "@/router";

const { data } = defineProps<{
  data: ComicDetailResponse[];
}>()
const { isDoubleExtraSmallAndExtraSmallScreen } = useResponsive();

const onNavigateToComicDetails = (id: number) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString()
    }
  });
};

const onNavigateToShortDetails = (id: number) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: "short-detail",
    params: {
        shortId: id.toString()
    }
  });
};

const onNavigateToChapterDetails = (comicId: number, chapterId: number) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: "chapter",
    params: {
      comicId: comicId.toString(),
      chapterId: chapterId.toString()
    }
  });
};
</script>
