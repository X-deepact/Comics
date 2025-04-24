<template>
  <Section
      :icon="sliderIcon"
      :text="sliderHeader"
      :nav-btn-text="sliderNavText"
  >
    <template #main>
      <div class="relative">
        <Button
          v-if="!atStart"
          @click="scrollLeft"
          class="absolute left-2 top-1/2 -translate-y-1/2 w-10 h-10 text-white
          rounded-full bg-orange-500 hover:bg-orange-600 active:bg-orange-600 z-10"
          icon="pi pi-angle-left"
        />

        <Button
          v-if="!atEnd"
          @click="scrollRight"
          class="absolute right-2 top-1/2 -translate-y-1/2 w-10 h-10 text-white
          rounded-full bg-orange-500 hover:bg-orange-600 active:bg-orange-600 z-10"
          icon="pi pi-angle-right"
        />

        <div ref="sliderContainer" class="slider-container mx-2 sm:mx-0 overflow-x-auto">
          <div class="w-max flex items-center justify-center gap-2 mx-auto pb-[10px]">
            <div
              class="flex flex-col rounded-lg w-[12rem] sm:w-[14rem] md:w-[15rem]"
              v-for="item in items"
              :key="item.id"
            >
              <img
                :src="item.cover"
                :alt="item.name"
                class="w-full rounded-lg object-cover cursor-pointer hover:shadow-[0_6px_6px_rgba(0,3,4,0.25)]"
                @click="onNavigateToComicDetails(item.id)"
              />
              <Button
                class="mt-2 font-bold leading-[20px] hover:text-[#ff620e] justify-start
                overflow-hidden whitespace-nowrap text-ellipsis"
                variant="text"
                :label="item.name"
                @click="onNavigateToComicDetails(item.id)"
              />
              <div v-if="item.chapters && item.chapters.length > 0" class="justify-start">
              <span class="font-semibold hover:text-[#ff620e] cursor-pointer">
                Chapter {{ item.chapters[0].number }}
              </span>
              <span
                class="text-gray-500"
                v-if="item.chapters[0].name"
              >
                {{ `: ${item.chapters[0].name}` }}
              </span>
              </div>
              <div v-else class="text-gray-800/50 truncate max-w-[15ch] overflow-hidden whitespace-nowrap">
                {{ item.description }}
              </div>
              
            </div>
          </div>
        </div>
      </div>
    </template>
  </Section>
</template>

<script setup lang="ts">
import type {ComicDetailResponse} from "@/services/comics/types.ts";
import Section from "@/components/common/Section.vue";
import router from "@/router";
import {onMounted, onUnmounted, ref} from "vue";

const { items, sliderHeader, sliderIcon, sliderNavText } = defineProps<{
  items: ComicDetailResponse[];
  sliderHeader: string;
  sliderIcon: string;
  sliderNavText?: string;
}>();

const atStart = ref(true);
const atEnd = ref(false);
const sliderContainer = ref<HTMLElement | null>(null);

const checkScrollPosition = () => {
  if (!sliderContainer.value) return;

  const { scrollLeft, scrollWidth, clientWidth } = sliderContainer.value;
  atStart.value = scrollLeft <= 0;
  atEnd.value = scrollLeft + clientWidth >= scrollWidth;
};

const scrollLeft = () => {
  if (!sliderContainer.value) return;
  sliderContainer.value.scrollTo({
    left: sliderContainer.value.scrollLeft - 200,
    behavior: "smooth",
  });
  checkScrollPosition();
};

const scrollRight = () => {
  if (!sliderContainer.value) return;
  sliderContainer.value.scrollTo({
    left: sliderContainer.value.scrollLeft + 200,
    behavior: "smooth",
  });
  checkScrollPosition();
};

const onNavigateToComicDetails = (id: number) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString(),
    },
  });
};

onMounted(() => {
  if (sliderContainer.value) {
    sliderContainer.value.addEventListener("scroll", checkScrollPosition);
    checkScrollPosition();
  }
});

onUnmounted(() => {
  if (sliderContainer.value) {
    sliderContainer.value.removeEventListener("scroll", checkScrollPosition);
  }
});
</script>

<style scoped>
.slider-container::-webkit-scrollbar {
  height: 7px;
}

.slider-container::-webkit-scrollbar-thumb {
  background-color: #ff620e;
  border-radius: 10px;
}

.slider-container::-webkit-scrollbar-thumb:hover {
  background-color: #ff620e;
  border-radius: 10px;
}

.slider-container::-webkit-scrollbar-track {
  border-radius: 10px;
}
</style>