<template>
  <div class="flex">
    <div class="relative rounded-full border border-[#ff620e] inline-flex items-center px-2">
      <i class="pi pi-search text-[#ff620e] ml-2" />
      <AutoComplete
        v-model="query"
        :suggestions="groupedSearchResult"
        option-label="name"
        option-group-label="label"
        option-group-children="items"
        @complete="onSearch"
        class="px-2 py-2 w-[15rem] md:w-[20rem] mr-[1rem] rounded-full"
        pt:overlay:class="mt-[1.5rem] border border-[#ff620e]"
        appendTo="self"
        :panelStyle="{ maxHeight: '20rem', overflowY: 'auto' }"
        :placeholder="$t('navigator.searchPlaceholder')"
        :loading="false"
      >
        <template #option="slotProps">
          <div
            class="flex px-4 py-3 hover:bg-gray-100 w-full cursor-pointer"
            :key="slotProps.option.id"
            @click="onNavigateToDetail(slotProps.option)"
          >
            <img
              alt="cover"
              :src="slotProps.option.cover || slotProps.option.thumb"
              class="w-16 h-20 object-cover rounded mr-4"
            />
            <div class="flex-1">
              <h3 class="font-bold break-words whitespace-normal">{{ slotProps.option.name }}</h3>
              <p
                v-if="slotProps.option.authors?.length"
                class="text-sm text-gray-600 break-words whitespace-normal"
              >
                {{ slotProps.option.authors[0].name }}
              </p>
              <p
                v-else-if="slotProps.option.description"
                class="text-sm text-gray-600 truncate break-words whitespace-normal"
              >
                {{ slotProps.option.description.length > 20 ? slotProps.option.description.slice(0, 20) + '...' : slotProps.option.description }}
              </p>
            </div>
          </div>
        </template>

        <template #header>
          <div class="font-medium px-3 py-2 border-b border-[#ff620e]">
            {{ $t('navigator.allResults') }}
          </div>
        </template>
        <template #optiongroup="slotProps">
          <div class="text-black px-3 py-2 bg-white sticky top-0 z-10 border-b border-[#ff620e]" :key="slotProps.option.label">
            {{ slotProps.option.label }}
          </div>
        </template>

        <template #footer>
          <div class="px-4 py-2 border-t border-[#ff620e] text-center bg-white sticky bottom-0 z-20 shadow-sm">
            <Button
              text
              size="small"
              severity="secondary"
              :label="$t('navigator.showAllResults')"
              @click="onShowAll"
            />
          </div>
        </template>
        <template #empty>
          <div class="w-full font-semibold justify-center flex py-5 italic">
            {{ $t('navigator.noResults') }}
          </div>
        </template>
      </AutoComplete>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue';
import { useRouter } from 'vue-router';
import { useSearch } from '@/stores/search/useSearch';
import { storeToRefs } from 'pinia';
import AutoComplete from 'primevue/autocomplete';
import Button from 'primevue/button';
import {i18n} from "@/main.ts";

const router = useRouter();
const query = ref('');
const { searchQuery, searchResult } = storeToRefs(useSearch());

const comics = computed(() => {
  if (!searchResult.value) {
    return [];
  }
  return searchResult.value.filter((result) => result.type === 'comic');
});
const dramas = computed(() => {
  if (!searchResult.value) {
    return [];
  }
  return searchResult.value.filter((result) => result.type === 'drama');
});
const groupedSearchResult = computed(() => {
  return [
    ...(comics.value.length
        ? [{ label: `${i18n.global.t("navigator.comicResults")}`, items: comics.value }]
        : []),
    ...(dramas.value.length
        ? [{ label: `${i18n.global.t("navigator.shortDramaResults")}`, items: dramas.value }]
        : [])
  ];
})

const onSearch = (event: any) => {
  searchQuery.value = event.query;
};

const onNavigateToDetail = (item: any) => {
  if (item.type === 'comic') {
    router.push({ name: 'comic-detail', params: { comicId: item.id } });
  } else if (item.type === 'drama') {
    router.push({ name: 'short-detail', params: { shortId: item.id } });
  }
};

const onShowAll = () => {
  router.push({
    path: '/search',
    query: { query: query.value }
  });
};

</script>
