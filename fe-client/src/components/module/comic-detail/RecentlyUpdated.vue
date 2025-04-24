<template>
  <div class="">
    <div class="bg-white p-4 rounded-md">
      <h3 class="pl-4 text-2xl font-semibold">Recently Updated</h3>
      <div class="" @mouseleave="hoveredItem = null">
        <ul class="p-2 relative">
          <li
            v-for="(item, index) in recentlyUpdatedList"
            :key="item.id"
            class="p-2 cursor-pointer flex items-center relative w-full"
            @mouseenter="hoveredItem = item"
            @mouseleave="hoveredItem = null"
            @mousedown="selectItem(item)"
          >
            <!-- <span
              :class="{
                'text-orange-600 font-bold': index === 0,
                'text-orange-500 font-bold': index === 1,
                'text-orange-400 font-bold': index === 2,
                'text-gray-500': index !== 0,
              }"
            >
              {{ index + 1 }}.
            </span> -->

            <div class="ml-2 flex-grow">
              <!-- <div
                v-if="hoveredItem !== item"
                class="flex justify-between w-full"
              >
                <span>{{ item.title }}</span>
                <span>Reward: 1</span>
              </div> -->

              <div class="w-full bg-white p-2 rounded-md flex items-start">
                <img
                  :src="item.bannerImage"
                  alt="Cover"
                  class="w-16 h-20 object-cover rounded-md mr-3"
                />

                <div class="flex flex-col">
                  <h3 class="text-sm font-semibold">{{ item.title }}</h3>
                  <p class="text-xs text-gray-500 mt-1">
                    {{ $t('search.author') }}: {{ item.author }}
                  </p>
                  <p class="text-xs text-gray-600 mt-1">
                    {{ item.description }}
                  </p>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import type { ComicDetails } from "@/models/comic.model.ts";

const props = defineProps({
  recentlyUpdatedList: Array<ComicDetails>,
});

const hoveredItem = ref(null);
const emit = defineEmits(["selectItem"]);

const selectItem = (item) => {
  emit("selectItem", item);
};
</script>
