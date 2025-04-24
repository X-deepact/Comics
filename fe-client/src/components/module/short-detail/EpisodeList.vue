<template>
  <div>
    <div class="flex w-full">
      <div class="flex flex-wrap justify-center sm:justify-start gap-2 mx-4 ">
        <div
            v-for="item in shortEpisodes"
            :key="item.id"
            class="p-4 bg-gray-800 text-white text-center py-3 rounded-lg cursor-pointer hover:bg-orange-500"
            @click="onNavigateToShortDetail(item.id)"
        >
          Episode {{ item.number }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from "pinia";
import { useShortDetails } from "@/stores/short-details/useShortDetails.ts";
import { useRouter } from "vue-router";

const { shortEpisodes, shortDetail } = storeToRefs(useShortDetails());
const router = useRouter();
const onNavigateToShortDetail = (id: number) => {
  router.push({
    name: "watch",
    params: {
      shortId: shortDetail.value?.id,
      episodeId: id.toString(),
    },
  });
};
</script>
