import {defineStore} from "pinia";
import { computed, ref, watch } from "vue";
import type { ChapterDetailResponse } from "@/services/chapters/type.ts";
import {API} from "@/services";
import {useRoute} from "vue-router";

export const useChapter = defineStore("Chapter Store", () => {
	const currentChapter = ref<ChapterDetailResponse>();
	const route = useRoute();

	const chapterItems = computed(() => {
		return currentChapter.value?.items;
	});
	
	const isLatestChapter = computed(() => {
		return currentChapter.value?.next_chapter === null;
	});
	
	const isFirstChapter = computed(() => {
		return currentChapter.value?.prev_chapter === null;
	});

	const fetchChapter = async () => {
		currentChapter.value = undefined;

		const { chapterId } = route.params;

    if (!chapterId) {
        return;
    }
    const response = await API.chapter.getChapterDetail(chapterId as string);
    currentChapter.value = response.data;
	}

	watch(
    () => route.params,
		async () => {
			await fetchChapter();
		},
{
		deep: true,
		immediate: true
    }
	)

	return {
		currentChapter,
		chapterItems,
		isLatestChapter,
		isFirstChapter
	}
})