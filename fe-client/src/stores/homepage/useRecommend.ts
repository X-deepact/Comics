import {defineStore, storeToRefs} from "pinia";
import {computed, onMounted, ref, watch} from "vue";
import {API} from "@/services";
import type {ComicRecommendResponse} from "@/services/comics/types.ts";
import {POSITION_ENUM} from "@/models/common.enum.ts";
import {useLocale} from "@/stores/i18n/useLocale.ts";

export const useRecommend = defineStore("RecommendStore", () => {
  const recommendComics = ref<ComicRecommendResponse[]>([]);
  const { locale } = storeToRefs(useLocale());

  const fetchRecommendation = async () => {
     try {
       const response = await API.comic.getRecommendComics(locale.value);

       if (response.status) {
         recommendComics.value = response.data;
       }
     } catch (e: any) {
       console.log(e.message)
     }
  }

  const toppingList = computed(() => {
    if (!recommendComics.value) {
      return [];
    }
    return recommendComics
      .value
      .filter((value: ComicRecommendResponse) => value.position === POSITION_ENUM.TOPING.toString());
  });
  
  const recommendationsList = computed(() => {
    if (!recommendComics.value) {
      return [];
    }

    const list = recommendComics
      .value
      .filter((value: ComicRecommendResponse) => value.position === POSITION_ENUM.RECOMMEND_PRODUCTS.toString());

    if (!list || list.length === 0) {
      return [];
    }

    return list[0]?.comics;
  })

  const completeMasterpieceList = computed(() => {
    if (!recommendComics.value) {
      return [];
    }

    const list = recommendComics
      .value
      .filter((value: ComicRecommendResponse) => value.position === POSITION_ENUM.COMPLETED_MASTERPIECES.toString());

    if (!list || list.length === 0) {
      return [];
    }

    return list[0]?.comics;
  });

  const fastestRisingList = computed(() => {
    if (!recommendComics.value) {
      return [];
    }

    const list = recommendComics
      .value
      .filter((value: ComicRecommendResponse) => value.position === POSITION_ENUM.FASTEST_GROWING.toString());

    if (!list || list.length === 0) {
      return [];
    }

    return list[0]?.comics;
  });

  const newProductList = computed(() => {
    if (!recommendComics.value) {
      return [];
    }

    const list = recommendComics
      .value
      .filter((value: ComicRecommendResponse) => value.position === POSITION_ENUM.TESTING_NEW_PRODUCTS.toString());

    if (!list || list.length === 0) {
      return [];
    }

    return list[0]?.comics;
  });

  watch(locale, (newValue) => {
    fetchRecommendation();
  })

  return {
    recommendComics,
    recommendationsList,
    completeMasterpieceList,
    fastestRisingList,
    newProductList,
    toppingList,
    fetchRecommendation
  };
})