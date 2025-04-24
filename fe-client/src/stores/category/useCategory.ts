import { defineStore, storeToRefs } from "pinia";
import type { CategoryButton } from "@/services/categoryData";
import {
	getAudiencesList,
	getOriginalLanguageList,
	getProgressesList
} from "@/services/categoryData";
import {computed, onUnmounted, reactive, ref, watch} from "vue";
import type { GenreResponse } from "@/services/genres/types.ts";
import { API } from "@/services";
import { useRoute } from "vue-router";
import type {ComicDetailResponse} from "@/services/comics/types.ts";
import { useLocale } from "@/stores/i18n/useLocale.ts";
export const useCategory = defineStore("categoryStore", () => {
	const { locale } = storeToRefs(useLocale());
	const isLoading = ref(false);
	const defaultCategory: GenreResponse = {
		id: 0,
		name: "all",
		position: 0
	}
	const route = useRoute();
  const comicData = ref<ComicDetailResponse[]>([]);
	const totalCount = ref(0);
	const isShowingPagination = computed(() => {
		return totalCount.value > 10;
	});
	const totalPage = computed(() => {
		return Math.ceil(totalCount.value / 10);
	});
	const currentPage = ref(1);
	const selectionList = reactive({
		genresList: [] as GenreResponse[],
		progressList: [] as CategoryButton[],
		audiencesList: [] as CategoryButton[],
		originalLanguagesList: [] as CategoryButton[],
	});
	const selectionTags = reactive({
		genre: "all",
		progress: "all",
		audience: "all",
		originalLanguage: "all",
	});
	const disableNextButton = computed(() => {
		return currentPage.value === totalPage.value;
	});
	const disablePreviousButton = computed(() => {
		return currentPage.value === 1;
	});

	async function fetchComicForCategory() {
		comicData.value = [];

		isLoading.value = true;
		const {
			genre,
			progress,
			audience,
			originalLanguage,
			page
		} = route.query;

		if (page === undefined) {
			currentPage.value = 1;
		}

		selectionTags.genre = typeof genre === "string" ? genre : "0";
		selectionTags.progress = typeof progress === "string" ? progress : "all";
		selectionTags.audience = typeof audience === "string" ? audience : "all";
		selectionTags.originalLanguage =
			typeof originalLanguage === "string" ? originalLanguage : "all";

		const response = await API.comic.getComics({
			genre: selectionTags.genre,
			progress: selectionTags.progress === "all" ? "" : selectionTags.progress,
			audience: selectionTags.audience === "all" ? "" : selectionTags.audience,
			originalLanguage: selectionTags.originalLanguage === "all" ? "" : selectionTags.originalLanguage,
			locale: locale.value,
			page: currentPage.value.toString()
		});

		if (response.status === 200) {
			comicData.value = response.data;
			totalCount.value = response.total;
		}
		isLoading.value = false;
	}
	
	watch(locale, async () => {
		await initialize();
	})

	const initialize = async () => {
		selectionList.audiencesList = getAudiencesList();
		selectionList.originalLanguagesList = getOriginalLanguageList();
		selectionList.progressList = getProgressesList();
		const response = await API.genre.getGenres({
			isHomePage: false,
			locale: locale.value
		});
		selectionList.genresList = [defaultCategory, ...response.data];

		fetchComicForCategory();
	}

	onUnmounted(() => {
		currentPage.value = 1;
	})

	return {
	  selectionList,
	  selectionTags,
	  comicData,
	  fetchComicForCategory,
	  isShowingPagination,
	  disableNextButton,
	  disablePreviousButton,
	  currentPage,
	  isLoading,
	  initialize
  };
});
