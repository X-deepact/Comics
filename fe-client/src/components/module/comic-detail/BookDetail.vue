<template>
  <div class="flex justify-center">
    <Card class="w-full relative">
      <template #content>
        <div class="py-1 sm:py-10 relative">
          <div class="absolute inset-0 bg-cover bg-center flex blur-[10px] sm:blur-[5px]"
            :style="{
              backgroundImage: `url(${comicDetail.cover})`,
              backgroundSize: '100%',
              backgroundPosition: 'top',
              backgroundRepeat: 'no-repeat',
            }"
          />
          <div class="grid grid-cols-6 gap-4 relative mx-2 lg:mt-0 sm:mx-5 lg:mx-2 2xl:mx-0">
            <div class="col-span-2 flex justify-end relative">
              <div class="relative px-2">
                <img
                  :src="comicDetail.cover"
                  alt="Book Cover"
                  class="w-full object-cover rounded-lg border-1"
                />
              </div>
            </div>
            <div class="col-span-4 flex flex-col justify-between">
              <div class="bg-opacity-90">
                <div>
                  <h1 class="!mt-0 text-2xl sm:text-4xl lg:text-5xl text-black font-bold sm:text-white">
                    {{ comicDetail.name }}
                  </h1>
                  <p class="text-black sm:text-white text-lg dark:text-orange-500 dark:text-dynamic">
                  <span
                      v-for="(author, index) in comicDetail.authors || []"
                      :key="index"
                  >
                    {{ author.name }}
                    <span v-if="index < (comicDetail.authors.length || 0) - 1">,</span>
                  </span>
                  </p>
                </div>
                <div>
                  <div class="flex flex-wrap gap-2 my-3 cursor-pointer pt-5">
                    <GenreChip
                        v-for="genre in comicDetail.genres"
                        :key="genre.name"
                        :content="genre.name"
                        is-navigating
                        @click="onNavigateToCategoryPage(genre.id)"
                    />
                  </div>
                  <div class="px-2">
                    <p class="text-sm text-black sm:text-white mt-5 dark:text-orange-500 dark:text-dynamic">
                      {{ comicDetail.description }}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>
    </Card>
  </div>
</template>

<script setup>
import { useComicDetails } from "@/stores/comic-details/useComicDetails";
import { storeToRefs } from "pinia";
import { useRouter } from "vue-router";
import GenreChip from "@/components/common/GenreChip.vue";

const router = useRouter();

const comicDetailsStore = useComicDetails();
const { comicDetail } = storeToRefs(comicDetailsStore);

const onNavigateToCategoryPage = (genre) => {
  const query = {};
  query["genre"] = genre;
  return router.push({ path: "/category", query });
};
</script>

<style scoped>
.p-card {
  box-shadow: none !important;
}
</style>
