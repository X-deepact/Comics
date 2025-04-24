<template>
  <div class="border-b border-black/25 min-h-[40px] pb-2 mx-auto mt-[5px] mb-[15px]">
    <div class="flex flex-row justify-between ">
      <div class="block">
        <div class="inline-flex gap-4 mr-7">
          <Image src="src/assets/images/ic__lx.png" class="w-[30px] h-[30px] bg-no-repeat bg-center" />
          <div class="flex gap-2">
            <GenreChip
              v-for="btn in homePageGenres"
              :key="btn.name"
              :content="btn.name.length > 10 ? btn.name.slice(0, 10) + '...' : btn.name"
              is-navigating
              @click="onNavigateToCategoryPage(btn.id, 'genre')"
            />
          </div>
        </div>
        <div class="inline-flex gap-4">
          <Image src="src/assets/images/ic__jd.png" class="w-[30px] h-[30px] bg-no-repeat bg-center" />
          <div class="flex gap-3">
            <GenreChip
              v-for="btn in statusButtons"
              :key="btn.name"
              :content="$t(`category.${btn.name.toLowerCase()}`)"
              is-navigating
              @click="onNavigateToCategoryPage(btn.name, 'progress')"
            />
          </div>
        </div>
      </div>
      <div class="flex gap-2 items-center">
        <i class="pi pi-objects-column" style="color: #ff6600" />
        <Button
          size="small"
          class="font-bold"
          :label="$t('dashboard.allCategories')"
          @click="router.push('/category')"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

import { useRouter } from "vue-router";
import {storeToRefs} from "pinia";
import {useGenres} from "@/stores/genres/useGenres.ts";
import { getProgressesList} from "@/services/categoryData.ts";
import GenreChip from "@/components/common/GenreChip.vue";
import {onMounted} from "vue";

const router = useRouter();
const statusButtons = getProgressesList().slice(1);
const store = useGenres();
const { homePageGenres } = storeToRefs(store);

onMounted(() => {
  store.fetchGenres();
})

const onNavigateToCategoryPage = (genre: string, group: string) => {
  const query: Record<string, string> = {};
  query[`${group}`]= genre;
  return router.push({path: "/category", query })
};

</script>
