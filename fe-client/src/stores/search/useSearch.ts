import {defineStore, storeToRefs} from "pinia";
import { ref, watch, computed } from "vue";
import { API } from "@/services";
import type {ComicDetailResponse} from "@/services/comics/types.ts";
import {useLocale} from "@/stores/i18n/useLocale.ts";

export const useSearch = defineStore("searchStore", () => {
    const searchResult = ref<ComicDetailResponse[]>([]);
    const searchResultList = ref<ComicDetailResponse[]>([]);
    const searchResultTotal = ref(0);
    const searchQuery = ref("");
    const { locale } = storeToRefs(useLocale());

    const isShowingPagination = computed(() => {
      return searchResultTotal.value > 10;
    });
    const totalPage = computed(() => {
      return Math.ceil(searchResultTotal.value / 10);
    });
    const currentPage = ref(1);
    const disableNextButton = computed(() => {
      return currentPage.value === totalPage.value;
    });
    const disablePreviousButton = computed(() => {
      return currentPage.value === 1;
    });

    let abortController: AbortController | null = null;

    const fetchSearchResults = async (query: string = "") => {
      if (abortController) {
        abortController.abort();
      }

      abortController = new AbortController();

      const response = await API.comic.getComicsSearch(query, 1 , locale.value, abortController.signal);

      if (response.success) {
        searchResult.value = response.data as ComicDetailResponse[];
      }
    };

    const fetchSearchResultsList = async (query: string = "", page: number = 1) => {
      if (abortController) {
        abortController.abort();
      }

      abortController = new AbortController();
      const response = await API.comic.getComicsSearch(query, page, locale.value, abortController.signal);

      if (response.success) {
        searchResultList.value = response.data;
        searchResultTotal.value = response.total;
      }
    };
  
    watch(searchQuery, async (newQuery) => {
      await fetchSearchResults(newQuery);
    });

    watch(currentPage, async (newPage) => {
      await fetchSearchResultsList(searchQuery.value, newPage);
    });
  
    return { searchResult, searchResultTotal, searchQuery, fetchSearchResults, searchResultList, fetchSearchResultsList, isShowingPagination, disableNextButton, disablePreviousButton,currentPage,
   };
  });
