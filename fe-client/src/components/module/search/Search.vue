<template>
  <ListLayout>
    <template #top>
      <div class="mx-2">
        “<span class="text-orange-600">{{ searchQuery }}</span>”（{{ searchResultTotal }}）
      </div>
    </template>
    <template #main>
      <template v-if="!searchResultList || searchResultList.length === 0">
        <ListComicSkeleton />
      </template>
      <ComicList v-else :data="searchResultList" />
    </template>
    <template #bottom>
      <div
        v-if="isShowingPagination"
        class="flex items-center justify-center flex-wrap gap-4"
      >
        <Button
          :disabled="disablePreviousButton"
          size="large"
          icon-pos="left"
          label="Previous"
          icon="pi pi-angle-left"
          :class="[
            'w-[7rem] border border-r-1 p-2 rounded flex',
            disablePreviousButton
              ? 'bg-gray-500 text-gray-300'
              : 'bg-[#ff6600] text-white',
          ]"
          @click="currentPage--"
        />
        <Button
          :disabled="disableNextButton"
          size="large"
          icon-pos="right"
          label="Next"
          icon="pi pi-angle-right"
          :class="[
            'w-[7rem] border border-r-1 p-2 rounded flex',
            disableNextButton
              ? 'bg-gray-500 text-gray-300'
              : 'bg-[#ff6600] text-white',
          ]"
          @click="currentPage++"
        />
      </div>
    </template>
  </ListLayout>
</template>

<script setup lang="ts">
import ListLayout from "@/layouts/ListLayout.vue";
import { storeToRefs } from "pinia";
import { useSearch } from "@/stores/search/useSearch.ts";
import { computed, watchEffect, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import ListComicSkeleton from "@/components/skeleton/ListComicSkeleton.vue";
import ComicList from "@/components/module/search/ComicList.vue";

const route = useRoute();
const router = useRouter();

const searchQuery = computed(() => route.query.query || "");

const searchStore = useSearch();
const {
  searchResultList,
  searchResultTotal,
  isShowingPagination,
  disableNextButton,
  disablePreviousButton,
  currentPage,
} = storeToRefs(searchStore);

watch(currentPage, (newPage) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    query: {
      ...route.query,
      page: newPage.toString(),
    },
  });
});

watchEffect(() => {
  searchStore.fetchSearchResultsList(searchQuery.value, currentPage.value);
});
</script>
