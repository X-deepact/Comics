<template>
  <div class="relative inline-block">
    <Select
      @change="onChangeLocale"
      v-model="selectedLocale"
      :options="localeOptions"
      optionLabel="label"
      placeholder="Select Locale"
      class="min-w-30 max-w-50 h-10 p-2 border border-gray-300 rounded-lg shadow-md focus:outline-none focus:ring-2 focus:ring-orange-500"
    />
  </div>
</template>

<script setup lang="ts">
import Trans from "@/utils/translation.ts";
import { ref } from "vue";
import { useLocale } from "@/stores/i18n/useLocale.ts";

interface LocaleOption {
  label: string;
  value: string;
}

const { setLocale } = useLocale();

const LOCALE_LABELS: Record<string, string> = {
  en: "English",
  vi: "Tiếng Việt",
  zh: "中文",
};
const locales = Trans.supportedLocales;
const localeOptions = locales.map((locale: string) => ({
  label: LOCALE_LABELS[locale] ?? locale,
  value: locale,
}));
const selectedLocale = ref<LocaleOption>(
  localeOptions.find((option: LocaleOption) => option.value === "en") ||
    localeOptions[0]
);

const onChangeLocale = async (): Promise<void> => {
  const newLocale = selectedLocale.value.value;
  setLocale(newLocale);
  await Trans.switchLanguage(newLocale);
};
</script>
