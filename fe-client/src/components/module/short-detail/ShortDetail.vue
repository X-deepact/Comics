<template>
  <div>
    <div v-if="shortDetailResponseErrorCode">
      <img class="w-full" src="@/assets/images/not-found.png" alt="404" />
    </div>
    <div v-else>
      <div v-if="isLoading" class="text-center text-gray-500 pt-10">
        Loading short details...
      </div>
      <div v-else>
        <ShortInfoDetail />
        <ListLayout>
          <template #main>
            <EpisodeList />
          </template>
        </ListLayout>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import ShortInfoDetail from "@/components/module/short-detail/ShortInfoDetail.vue";
import EpisodeList from "@/components/module/short-detail/EpisodeList.vue";
import ListLayout from "@/layouts/ListLayout.vue";
import { useShortDetails } from "@/stores/short-details/useShortDetails.ts";
import { storeToRefs } from "pinia";
const shortDetailsStore = useShortDetails();
const { shortDetailResponseErrorCode, isLoading } = storeToRefs(shortDetailsStore);
</script>
