<template>
  <div>
    <div class="flex justify-between">
      <div class="inline-flex gap-2">
        <Button
          v-if="firstChapter"
          type="button"
          class="ml-2 sm:ml-0 px-6 py-2 uppercase rounded bg-gray-200 hover:bg-gray-400 dark:bg-gray-400 dark:hover:bg-gray-500 flex justify-center items-center font-bold"
          label="Read from first chapter"
          @click="onNavigateToChapterDetail(firstChapter.id)"
        />

        <Button
          v-if="latestChapter"
          type="button"
          class="px-6 py-2 uppercase rounded bg-gray-200 hover:bg-gray-400 dark:bg-gray-400 dark:hover:bg-gray-500 justify-center items-center font-bold hidden sm:flex"
          label="Read from latest chapter"
          @click="onNavigateToChapterDetail(latestChapter.id)"
        />
      </div>
      <Button
        size="large"
        type="button"
        :icon="
          sortOrder === 'asc'
            ? 'pi pi-sort-amount-up'
            : 'pi pi-sort-amount-down'
        "
        class="cursor-pointer"
        @click="chapterToggleSort"
        v-tooltip.top="{
          value:
            sortOrder === 'asc'
              ? $t('message.sortByDesc')
              : $t('message.sortByAsc'),
          pt: {
            text: {
              style: {
                padding: '4px 10px',
              },
            },
          },
        }"
      />
    </div>

    <VirtualScroller
      :items="comicChapters"
      :itemSize="50"
      class="max-h-100 my-4 rounded-md"
      style="height: 500px"
    >
      <template #item="{ item }">
        <div
          class="p-2 cursor-pointer"
          :key="item.number"
          @click="onNavigateToChapterDetail(item.id)"
        >
          <div
            class="group flex items-center space-x-2 bg-gray-50 dark:bg-transparent p-2 hover:text-orange-500 cursor-pointer border-l-6 border-gray-400"
          >
            <div class="font-bold text-xl mr-5 group-hover:text-orange-500">
              {{ `Chapter ${item.number}` }}
            </div>
            <div>
              <div
                class="text-gray-400 text-xl italic group-hover:text-orange-500"
              >
                {{ item.name ?? "No title" }}
              </div>
              <div class="text-gray-500 text-sm group-hover:text-orange-500">
                {{ getTimeAgo(item.created_at) }}
              </div>
            </div>
          </div>
        </div>
      </template>
    </VirtualScroller>
  </div>
</template>

<script setup lang="ts">
import { getTimeAgo } from "@/utils/date-formatted.ts";
import { storeToRefs } from "pinia";
import { useComicDetails } from "@/stores/comic-details/useComicDetails.ts";
import { useRouter } from "vue-router";

const { comicChapters, sortOrder, comicDetail, firstChapter, latestChapter } =
  storeToRefs(useComicDetails());
const { chapterToggleSort } = useComicDetails();
const router = useRouter();
const onNavigateToChapterDetail = (id: number) => {
  router.push({
    name: "chapter",
    params: {
      comicId: comicDetail.value?.id,
      chapterId: id.toString(),
    },
  });
};
</script>
