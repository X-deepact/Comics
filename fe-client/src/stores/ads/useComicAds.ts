import {defineStore, storeToRefs} from "pinia";
import { watch, ref } from "vue";
import type { AdvertisementResponse } from "@/services/ads/types.ts";
import { API } from "@/services";
import {useLocale} from "@/stores/i18n/useLocale.ts";

export const useComicAds = defineStore('ComicAdsStore', () => {
	const { locale } = storeToRefs(useLocale());

	const comicAds = ref<AdvertisementResponse[]>([]);
	const fetchAds = async () => {
		const response = await API.ads.getAds(locale.value);
		if(response.data){
			comicAds.value = response.data;
		}else{
			comicAds.value = [];
		}
	}

	watch(locale, (newValue) => {
		fetchAds();
	  })

	return {
		comicAds,
		fetchAds
	}
})