import {defineStore, storeToRefs} from "pinia";
import {computed, onMounted, onUnmounted, ref, watch} from "vue";
import { API } from "@/services";
import type { ComicDetailResponse } from "@/services/comics/types.ts";
import {useLocale} from "@/stores/i18n/useLocale.ts";
import router from "@/router";
import {useRoute} from "vue-router";

export const useRecentlyUpdated = defineStore("recentlyUpdatedStore",() => {
	const isLoading = ref(false);
	const route = useRoute();
	const currentRecentlyUpdated = ref<ComicDetailResponse[]>([]);
	const { locale } = storeToRefs(useLocale());
	const totalCount = ref(0);
	const isShowingPagination = computed(() => {
		return totalCount.value > 10;
	});
	const totalPage = computed(() => {
		return Math.ceil(totalCount.value / 10);
	});
	const currentPage = ref(1);
	const disableNextButton = computed(() => {
		return currentPage.value === totalPage.value;
	});
	const disablePreviousButton = computed(() => {
		return currentPage.value === 1;
	});

	const fetchComics = async () => {
		isLoading.value = true;
		const response = await API.comic.getRecentlyUpdated(currentPage.value, locale.value);

		if (response.status === 200) {
			currentRecentlyUpdated.value = response.data;
			totalCount.value = response.total;
		}
		isLoading.value = false;
	};


	watch(currentPage, () => {
	  fetchComics();
	});

	watch(locale, (newValue) => {
		fetchComics();
	})

	onUnmounted(() => {
		currentPage.value = 1;
	})

	return { currentRecentlyUpdated, isShowingPagination, disableNextButton, disablePreviousButton,currentPage, fetchComics, isLoading }
})