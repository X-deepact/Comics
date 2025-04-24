<template>
  <ListLayout>
    <template #top>
      <CategorySelection />
    </template>
    <template #main>
      <template v-if="isLoading">
        <ListComicSkeleton />
      </template>
      <div v-else>
        <div v-if="!categoryStore.comicData || categoryStore.comicData.length === 0">
          <p class="text-center text-lg">No comics found</p>
        </div>
        <ComicList v-else :data="categoryStore.comicData" />
      </div>
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
          :class="['w-[7rem] border border-r-1 p-2 rounded flex',
            disablePreviousButton
              ? 'bg-gray-500 text-gray-300'
              : 'bg-[#ff6600] text-white'
          ]"
          @click="currentPage--"
        />
        <Button
          :disabled="disableNextButton"
          size="large"
          icon-pos="right"
          label="Next"
          icon="pi pi-angle-right"
          :class="['w-[7rem] border border-r-1 p-2 rounded flex',
            disableNextButton
              ? 'bg-gray-500 text-gray-300'
              : 'bg-[#ff6600] text-white'
          ]"
          @click="currentPage++"
        />
      </div>
    </template>
  </ListLayout>
</template>

<script setup lang="ts">
import { useCategory } from "@/stores/category/useCategory.ts";
import ComicList from "@/components/module/recently-update/ComicList.vue";
import ListLayout from "@/layouts/ListLayout.vue";
import CategorySelection from "@/components/module/category/CategorySelection.vue";
import {onMounted, watch} from "vue";
import {useRoute, useRouter} from "vue-router";
import {storeToRefs} from "pinia";
import ListComicSkeleton from "@/components/skeleton/ListComicSkeleton.vue";
const route = useRoute();
const router = useRouter();
const categoryStore = useCategory();
const { currentPage,
  isShowingPagination,
  disablePreviousButton,
  disableNextButton,
  isLoading
} = storeToRefs(categoryStore);

onMounted(() => {
  categoryStore.initialize();
})


watch(
  () => route.query,
  () => {
    categoryStore.fetchComicForCategory();
  },
  { deep: true }
);

watch(currentPage, (newPage) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    query: {
      ...route.query,
      page: newPage.toString(),
    },
  });
 }
)
</script>
