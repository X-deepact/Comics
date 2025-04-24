<template>
  <main class="my-12 flex-grow transition origin-top duration-500">
    <div class="w-full max-w-[64rem] mx-auto px-2 mb-12">
      <Button
        class="mb-2 hover:text-orange-500 py-1 flex items-center"
        icon-pos="left"
        icon="pi pi-arrow-left"
        :label="currentChapter?.comic_name"
        @click="onNavigateBackToComic(currentChapter?.comic_id as number)"
      />
    </div>
    <div class="w-full max-w-[40rem] mx-auto px-2 mb-12">
      <h1 class="text-xl text-gray-600 mb-6">
        Chapter {{currentChapter?.number}}: {{ currentChapter?.name }}
      </h1>
      <div class="flex items-stretch flex-col gap-2">
        <Button
          :disabled="isLatestChapter"
          class="flex justify-center
          items-center bg-orange-500
          hover:bg-orange-600 px-3 py-3
          text-sm text-white uppercase
          font-bold rounded-xl"
          label="Next Chapter"
          @click="onNavigateToNextChap"
        />
        <Button
          :disabled="isFirstChapter"
          class="flex justify-center
          items-center bg-gray-800
          hover:bg-gray-700 px-3 py-3
          text-sm text-white uppercase
          font-bold rounded-xl"
          label="Previous Chapter"
          @click="onNavigateToPreviousChap"
        />
      </div>
    </div>
    <div class="select-none flex flex-col items-center">
      <div
        v-for="item in chapterItems"
        :key="item.page"
        class="transition transform max-w-full mb-6 last:mb-0 lg:max-w-[64rem]"
      >
        <div class="overflow-hidden relative w-full">
          <ChapterImage
            :image-src="item.image"
          />
        </div>
      </div>
    </div>
    <div class="mt-12 w-full flex justify-center">
      <Button
        :disabled="isLatestChapter"
        :label="onNextChapBtnLabel"
        class="w-full max-w-[40rem] py-5 px-3 text-center uppercase
          font-bold rounded-xl bg-orange-500 bg-opacity-20 text-white hover:bg-orange-600"
        @click="onNavigateToNextChap"
      />
    </div>
    <div class="mt-2 w-full flex justify-center">
      <div class="max-w-[40rem] w-full flex gap-2">
        <Button
          :disabled="isFirstChapter"
          class="w-full max-w-[40rem] py-3 px-3 text-center uppercase
          font-bold rounded-xl bg-gray-900  text-gray-300"
          label="Previous Chap"
          @click="onNavigateToPreviousChap"
        />
        <Button
          disabled
          class="w-full max-w-[40rem] py-3 px-3 text-center uppercase
          font-bold rounded-xl bg-gray-900 text-gray-300"
          label="To chapter"
          @click="chapterChangeDialog?.toggleVisible"
        />
      </div>
    </div>
  </main>
  <ChapterChangeDialog ref="chapterChangeDialog" />
</template>

<script setup lang="ts">
import {storeToRefs} from "pinia";
import {useChapter} from "@/stores/chapters/useChapter.ts";
import ChapterImage from "@/components/module/chapter/ChapterImage.vue";
import ChapterChangeDialog from "@/components/module/chapter/ChapterChangeDialog.vue";
import {computed, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

const chapterChangeDialog = ref(null);
const { currentChapter, chapterItems, isLatestChapter, isFirstChapter } = storeToRefs(useChapter());

const onNextChapBtnLabel = computed(() => {
  if (isLatestChapter.value) return "This is the latest chapter";
  return `Continue to watch Chapter ${currentChapter.value?.number! + 1}`;
});

const onNavigateBackToComic = (id: number) => {
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString(),
    },
  })
}

const onNavigateToPreviousChap = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: 'chapter',
    params: {
      comicId: currentChapter.value?.comic_id,
      chapterId: currentChapter.value?.prev_chapter,
    },
  })
}

const onNavigateToNextChap = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: 'chapter',
    params: {
      comicId: currentChapter.value?.comic_id,
      chapterId: currentChapter.value?.next_chapter,
    },
  })
}
</script>
