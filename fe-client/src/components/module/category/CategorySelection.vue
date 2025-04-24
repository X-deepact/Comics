<template>
  <Card
    class="justify-self-center border-t border-[#ff620e] p-10 shadow-xl my-10 mx-2 sm:mx-0"
  >
    <template #content>
      <div class="flex flex-col gap-4">
        <div
        class="flex gap-3 sm:gap-4 items-center overflow-x-auto whitespace-nowrap lg:overflow-visible lg:whitespace-normal lg:flex-wrap"
        >
          <p class="font-semibold inline-block">{{ $t("category.genres") }}:</p>
          <Button
            size="small"
            v-for="button in selectionList.genresList"
            :key="button.id"
            :label="button.name == 'all' ? $t(`category.all`) : button.name"
            :class="getButtonClass(button.id.toString(), 'genre')"
            @click="handleClick(button.id.toString(), 'genre')"
          />
        </div>

        <div
          class="flex gap-3 sm:gap-4 items-center overflow-x-auto whitespace-nowrap lg:overflow-visible lg:whitespace-normal lg:flex-wrap"
        >
          <p class="font-semibold inline-block">
            {{ $t("category.progress") }}:
          </p>
          <Button
            size="large"
            v-for="button in selectionList.progressList"
            :key="button.name"
            :label="$t(`category.${button.name}`)"
            :class="getButtonClass(button.name, 'progress')"
            @click="handleClick(button.name, 'progress')"
          />
        </div>

        <div
          class="flex gap-3 sm:gap-4 items-center overflow-x-auto whitespace-nowrap lg:overflow-visible lg:whitespace-normal lg:flex-wrap"
        >
          <p class="font-semibold inline-block">
            {{ $t("category.audience") }}:
          </p>
          <Button
            size="large"
            v-for="button in selectionList.audiencesList"
            :key="button.name"
            :label="$t(`category.${button.name}`)"
            :class="getButtonClass(button.name, 'audience')"
            @click="handleClick(button.name, 'audience')"
          />
        </div>

        <div
          class="flex gap-3 sm:gap-4 items-center overflow-x-auto whitespace-nowrap lg:overflow-visible lg:whitespace-normal lg:flex-wrap"
        >
          <p class="font-semibold inline-block">
            {{ $t("category.originalLanguage") }}:
          </p>
          <Button
            size="large"
            v-for="button in selectionList.originalLanguagesList"
            :key="button.name"
            :label="$t(`category.${button.name}`)"
            :class="getButtonClass(button.name, 'originalLanguage')"
            @click="handleClick(button.name, 'originalLanguage')"
          />
        </div>
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import Card from "primevue/card";
import Button from "primevue/button";
import { useCategory } from "@/stores/category/useCategory.ts";
import { useRouter } from "vue-router";
import { watch } from "vue";

import { useResponsive } from "@/composables/useResponsive";

const { isLargeScreen } = useResponsive();

const store = useCategory();
const { selectionTags, selectionList } = store;
const router = useRouter();

watch(
  () => selectionTags,
  (newTags) => {
    const { genre, progress, audience, originalLanguage } = newTags;
    const query: Record<string, string> = {};
    if (genre !== "all") {
      query.genre = genre;
    }
    if (progress !== "all") {
      query.progress = progress;
    }
    if (audience !== "all") {
      query.audience = audience;
    }
    if (originalLanguage !== "all") {
      query.originalLanguage = originalLanguage;
    }

    router.push({ path: "/category", query });
  },
  { deep: true }
);

const handleClick = (value: string, group: string) => {
  switch (group) {
    case "genre":
      selectionTags.genre = value;
      break;
    case "progress":
      selectionTags.progress = value;
      break;
    case "audience":
      selectionTags.audience = value;
      break;
    case "originalLanguage":
      selectionTags.originalLanguage = value;
      break;
  }
};
const getButtonClass = (value: string, group: string) => {
  let isActive = false;
  switch (group) {
    case "genre":
      isActive = selectionTags.genre === value;
      break;
    case "progress":
      isActive = selectionTags.progress === value;
      break;
    case "audience":
      isActive = selectionTags.audience === value;
      break;
    case "originalLanguage":
      isActive = selectionTags.originalLanguage === value;
      break;
  }
  return isActive
    ? "bg-orange-500 text-sm text-white px-4 h-8 rounded-full capitalize drop-shadow-md"
    : "bg-black/25 text-white/80 text-sm text-black px-4 h-8 rounded-full capitalize drop-shadow-md";
};
</script>
<style scoped>
.overflow-x-auto {
  overflow-x: auto;
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.overflow-x-auto::-webkit-scrollbar {
  display: none;
}
</style>
