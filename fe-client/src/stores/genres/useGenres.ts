import { defineStore, storeToRefs } from "pinia";
import { watch, ref } from "vue";
import { API } from "@/services";
import { useLocale } from "@/stores/i18n/useLocale.ts";

export const useGenres = defineStore("GenreStore", () => {
  const homePageGenres = ref();
  const { locale } = storeToRefs(useLocale());

  const fetchGenres = async () => {
    const response = await API.genre.getGenres({
			isHomePage: true,
			locale: locale.value
		});
    homePageGenres.value = [...response.data];
  }

  watch(locale, (newValue) => {
    fetchGenres();
  })

  return { homePageGenres, fetchGenres };
});
