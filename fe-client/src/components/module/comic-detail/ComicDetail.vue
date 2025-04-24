<template>
  <div>
    <div v-if="comicDetailResponseErrorCode">
      <img class="w-full" src="@/assets/images/not-found.png" alt="404" />
    </div>
    <div v-else>
      <div v-if="isLoading" class="text-center text-gray-500 pt-10">
        Loading comic details...
      </div>
      <div v-else>
        <BookDetail />
        <ListLayout>
          <template #main>
            <ChapterList />
          </template>
        </ListLayout>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import BookDetail from "@/components/module/comic-detail/BookDetail.vue";
import ChapterList from "@/components/module/comic-detail/ChapterList.vue";
import ListLayout from "@/layouts/ListLayout.vue";
import { useComicDetails } from "@/stores/comic-details/useComicDetails.ts";
import { storeToRefs } from "pinia";
const comicDetailsStore = useComicDetails();
const { comicDetailResponseErrorCode, isLoading } = storeToRefs(comicDetailsStore);
</script>
