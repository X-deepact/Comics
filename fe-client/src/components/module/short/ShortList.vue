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
                    class="w-1/3 sm:w-auto mx-auto hover:cursor-pointer hover:shadow-[0_6px_6px_rgba(0,3,4,0.25)]"
                    @click="onNavigateToShortDetails(item.id)"
                >
                  <img
                      class="rounded w-full max-w-[300px]"
                      :src="item.thumb"
                      :alt="item.title"
                  />
                </div>
                <div
                    class="w-2/3 sm:w-full flex flex-col xl:flex-row justify-between xl:items-center flex-1 gap-6 mt-2"
                >
                  <div class="flex flex-col justify-between items-start gap-1">
                    <div>
                      <div class="text-xl font-bold cursor-pointer hover:text-orange-500"
                           @click="onNavigateToShortDetails(item.id)"
                      >
                        {{ item.translated_name }}
                      </div>
                    </div>
                    <div
                        class="flex flex-col justify-between gap-1"
                    >
                      <div
                          class="text-sm font-medium cursor-pointer hover:text-orange-500"
                      >
                        <span>{{ $t('search.latestEpisode') }}: {{ `${item?.latest_episode === 0 ? '-' : item?.latest_episode}` }}
                        </span>
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
import type {ShortDetailResponse} from "@/services/short/types.ts";
import router from "@/router";

const { data } = defineProps<{
  data: ShortDetailResponse[];
}>()
const { isDoubleExtraSmallAndExtraSmallScreen } = useResponsive();

const onNavigateToShortDetails = (id: number) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: "short-detail",
    params: {
        shortId: id.toString()
    }
  });
};
</script>
