import { defineStore } from "pinia";
import { ref } from "vue";

export const useLocale = defineStore("LocaleStore", () => {
	const locale = ref("en");
	
	const setLocale = (selectedLocale: string) => {
		locale.value = selectedLocale;
	}
	
	return { locale, setLocale }
})