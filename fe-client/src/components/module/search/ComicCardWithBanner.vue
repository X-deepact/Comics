<template>
  <Card unstyled class="card-container overflow-hidden">
    <template #title>
      <div
        class="relative w-full max-h-[17rem] overflow-hidden hover:shadow-[0_6px_6px_rgba(0,3,4,0.25)]"
      >
        <!-- <Image
          :src="props.item.bannerImage"
          alt="banner-image"
          class="w-full h-auto object-cover"
        /> -->
        <Image
          src="/src/assets/images/comic-banner.jpg"
          alt="banner-image"
          class="w-full h-auto object-cover"
        />

        <!-- Description Overlay -->
        <div
          class="absolute bottom-0 left-0 w-full bg-gradient-to-t from-black/70 to-transparent p-3 text-white text-sm"
        >
          <p class="line-clamp-2">
            {{ props.item.description }}
          </p>
        </div>
      </div>
    </template>
    <template #subtitle>
      <Button
        class="mt-2 font-bold hover:text-[#ff620e] w-full justify-start overflow-hidden whitespace-nowrap text-ellipsis cursor-pointer"
        variant="text"
        :label="props.item.title"
        @click="onNavigateToComicDetails(props.item.id)"
      />
      <!-- <div class="flex flex-col text-sm text-gray-600 mt-1">
        <div class="flex flex-row items-center">
          <span class="hover:cursor-pointer hover:text-[#ff620e]">
            {{ $t("message.popularity") }} :
          </span>
          <span class="ml-1">{{ props.item.popularity }}</span>
        </div>
      </div> -->
    </template>
  </Card>
</template>

<script setup lang="ts">
import type { ComicDetails } from "@/models/comic.model.ts";
import router from "@/router";

const props = defineProps<{
  item: ComicDetails;
}>();

const onNavigateToComicDetails = (id: number) => {
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString(),
    },
  });
};
</script>
