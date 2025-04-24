<template>
  <Card class="w-full">
    <template #content>
      <div class="relative">
        <div
          class="h-[27rem] bg-layer"
          :style="{ '--bg-thumb': `url(${shortDetail.thumb})` }"
        />
        <div class="absolute inset-0 z-10 flex items-center">
          <div class="flex w-full px-6 sm:px-12 items-center gap-6">
            <div class="w-1/4 sm:w-1/6">
              <img :src="shortDetail.thumb" alt="short thumb" class="w-full object-cover rounded-lg" />
            </div>
            <div class="flex flex-col gap-3 w-full">
              <h1 class="text-2xl sm:text-4xl font-bold text-white">
                {{ shortDetail.translated_name }}
              </h1>
              <p class="text-white text-lg dark:text-orange-500 dark:text-dynamic">
                {{ formatToOnlyYear(shortDetail.release_date) }}
              </p>
              <div class="flex flex-wrap gap-2">
                <GenreChip
                    v-for="genre in shortDetail.genres"
                    :key="genre.name"
                    :content="genre.name"
                    is-navigating
                    @click="onNavigateToCategoryPage(genre.id)"
                />
              </div>
              <p class="text-white text-lg line-clamp-2 max-w-3xl mt-10">
                {{ shortDetail.description }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </template>
  </Card>
</template>

<script setup>
import { useShortDetails } from "@/stores/short-details/useShortDetails";
import { storeToRefs } from "pinia";
import { useRouter } from "vue-router";
import GenreChip from "@/components/common/GenreChip.vue";
import {formatToOnlyYear} from "@/utils/date-formatted.js";

const router = useRouter();

const shortDetailsStore = useShortDetails();
const { shortDetail } = storeToRefs(shortDetailsStore);

const onNavigateToCategoryPage = (genre) => {
  const query = {};
  query["genre"] = genre;
  return router.push({ path: "/short", query });
};
</script>

<style scoped>
.bg-layer::before {
  content: "";
  position: absolute;
  inset: 0;
  background-image: var(--bg-thumb);
  background-size: 100%;
  background-position: top;
  background-repeat: no-repeat;
  background-color: rgba(0, 0, 0, 0.6); /* dark overlay */
  background-blend-mode: darken;
  z-index: 0;
  pointer-events: none;
}

.bg-layer {
  position: relative;
  z-index: 0;
}
</style>
