import {defineStore, storeToRefs} from "pinia";
import type {CategoryButton} from "@/services/categoryData";
import {getReleaseYearList,} from "@/services/categoryData";
import {computed, onUnmounted, reactive, ref, watch} from "vue";
import {API} from "@/services";
import {useRoute} from "vue-router";
import type {ShortDramaDetailResponse, GenresForShortDramaResponse} from "@/services/short/types.ts";
import {useLocale} from "@/stores/i18n/useLocale.ts";

export const useShort = defineStore("shortStore", () => {
  const {locale} = storeToRefs(useLocale());
  const isLoading = ref(false);
  const defaultCategory: GenresForShortDramaResponse = {
    id: 0,
    translated_name: "all",
    position: 0
  }
  const route = useRoute();
  const shortData = ref<ShortDramaDetailResponse[]>([]);
  const totalCount = ref(0);
  const isShowingPagination = computed(() => {
    return totalCount.value > 10;
  });
  const totalPage = computed(() => {
    return Math.ceil(totalCount.value / 10);
  });
  const currentPage = ref(1);
  const selectionList = reactive({
    genresList: [] as GenresForShortDramaResponse[],
    releaseYearList: [] as CategoryButton[],
  });
  const selectionTags = reactive({
    genre: "all",
    releaseYear: "all",
  });
  const disableNextButton = computed(() => {
    return currentPage.value === totalPage.value;
  });
  const disablePreviousButton = computed(() => {
    return currentPage.value === 1;
  });

  async function fetchShortListForShort() {
    shortData.value = [];

    isLoading.value = true;
    const {
      genre,
      releaseYear,
      page
    } = route.query;

    if (page === undefined) {
      currentPage.value = 1;
    }

    selectionTags.genre = typeof genre === "string" ? genre : "0";
    selectionTags.releaseYear = typeof releaseYear === "string" ? releaseYear : "all";

    const response = await API.service.getList<ShortDramaDetailResponse>(`short-drama`, {
      params: {
        page_size: 12,
        page: currentPage.value.toString(),
        language: locale.value,
        genre_id: selectionTags.genre,
        release_year: selectionTags.releaseYear === "all" ? "" : selectionTags.releaseYear,
      }
    });

    if (response.status === 200) {
      shortData.value = response.data;
      totalCount.value = response.total;
    }
    isLoading.value = false;
  }

  watch(locale, async () => {
    await initialize();
  })

  const initialize = async () => {
    selectionList.releaseYearList = getReleaseYearList();
    const response = await API.service.get<GenresForShortDramaResponse[]>(
        `/short-drama-genre/all`,
        {
          params: {
            language: locale.value,
          }
        }
    )
    selectionList.genresList = [defaultCategory, ...(response.data || [])];

    fetchShortListForShort();
  }

  onUnmounted(() => {
    currentPage.value = 1;
  })

  return {
    selectionList,
    selectionTags,
    shortData,
    fetchShortListForShort,
    isShowingPagination,
    disableNextButton,
    disablePreviousButton,
    currentPage,
    isLoading,
    initialize
  };
});
