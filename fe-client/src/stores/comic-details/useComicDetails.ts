import {defineStore, storeToRefs} from "pinia";
import {computed, ref, watch} from "vue";
import { API } from "@/services";
import type {ComicDetailResponse} from "@/services/comics/types.ts";
import { useRoute } from "vue-router";
import {useLocale} from "@/stores/i18n/useLocale.ts";

export const useComicDetails = defineStore("ComicDetailStore", () => {
  const comicDetail = ref<ComicDetailResponse | undefined>(undefined);
  const route = useRoute();
  const sortOrder = ref<'asc' | 'desc'>('desc');
  const { locale } = storeToRefs(useLocale());
  const comicDetailResponseErrorCode = ref(false);
  const isLoading = ref(false);

  const comicChapters = computed(() => {
    if (!comicDetail.value) {
      return [];
    }
    return comicDetail.value?.chapters?.sort((a, b) => {
      if (sortOrder.value === 'asc') {
        return a.number! - b.number!;
      } else {
        return b.number! - a.number!;
      }
    });
  });

  const latestChapter = computed(() => {
    if (!comicDetail.value || comicDetail.value.chapters?.length === 0) return null;

    return comicDetail.value.chapters?.reduce((first, chapter) =>
      chapter.number! > first.number! ? chapter : first
    );
  });

  const firstChapter = computed(() => {
    if (!comicDetail.value || comicDetail.value.chapters?.length === 0) return null;

    return comicDetail.value.chapters?.reduce((first, chapter) =>
      chapter.number! < first.number! ? chapter : first
    );
  });

  const fetchComicDetails = async () => {
    comicDetailResponseErrorCode.value = false;
    comicDetail.value = undefined;
    isLoading.value = true;

    const { comicId } = route.params;
    if (comicId === undefined) {
      comicDetailResponseErrorCode.value = true;
      return;
    }

    const response = await API.comic.getComicDetails(comicId as string, locale.value);
    if (response.code == 'ERROR') {
      comicDetailResponseErrorCode.value = true;
    } else {
      comicDetail.value = response.data;
    }
    isLoading.value = false;
  };

  watch(
    () => route.params,
    async () => {
      await fetchComicDetails();
    },
    {
      deep: true,
      immediate: true
    }
  );

  watch(
    () => locale.value,
    async () => {
      await fetchComicDetails();
    }
  )


  const chapterToggleSort = () => {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc';
  };

  return {
    comicDetail,
    fetchComicDetails,
    comicChapters,
    chapterToggleSort,
    sortOrder,
    firstChapter,
    latestChapter,
    comicDetailResponseErrorCode,
    isLoading
  };
});
