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
            :label="button.translated_name == 'all' ? $t(`category.all`) : button.translated_name"
            :class="getButtonClass(button.id.toString(), 'genre')"
            @click="handleClick(button.id.toString(), 'genre')"
          />
        </div>

        <div
          class="flex gap-3 sm:gap-4 items-center overflow-x-auto whitespace-nowrap lg:overflow-visible lg:whitespace-normal lg:flex-wrap"
        >
          <p class="font-semibold inline-block">
            {{ $t("category.releaseYear") }}:
          </p>
          <Button
            size="large"
            v-for="button in selectionList.releaseYearList"
            :key="button.name"
            :label="button.name == 'all' ? $t(`category.all`) : button.name"
            :class="getButtonClass(button.name, 'releaseYear')"
            @click="handleClick(button.name, 'releaseYear')"
          />
        </div>
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import Card from "primevue/card";
import Button from "primevue/button";
import { useShort } from "@/stores/short/useShort.ts";
import { useRouter } from "vue-router";
import { watch } from "vue";

const store = useShort();
const { selectionTags, selectionList } = store;
const router = useRouter();

watch(
  () => selectionTags,
  (newTags) => {
    const { genre, releaseYear } = newTags;
    const query: Record<string, string> = {};
    if (genre !== "all") {
      query.genre = genre;
    }
    if (releaseYear !== "all") {
      query.releaseYear = releaseYear;
    }

    router.push({ path: "/short", query });
  },
  { deep: true }
);

const handleClick = (value: string, group: string) => {
  switch (group) {
    case "genre":
      selectionTags.genre = value;
      break;
    case "releaseYear":
      selectionTags.releaseYear = value;
      break;
  }
};
const getButtonClass = (value: string, group: string) => {
  let isActive = false;
  switch (group) {
    case "genre":
      isActive = selectionTags.genre === value;
      break;
    case "releaseYear":
      isActive = selectionTags.releaseYear === value;
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
