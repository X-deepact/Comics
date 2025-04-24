<template>
  <div class="mx-auto w-full overflow-x-hidden py-2">
    <ShortPlayer
        v-if="tracks.length > 0 && streamUrl?.video"
        :key="route.params.episodeId"
        ref="player"
        :tracks="tracks"
        :source="streamUrl.video"
    />
  </div>
  <section class="px-4 sm:px-8 lg:px-20 py-6 sm:py-10">
    <div class="mx-auto max-w-full sm:max-w-md md:max-w-lg lg:max-w-3xl xl:max-w-5xl">
      <div>
        <h1 class="text-xl lg:text-2xl 2xl:text-3xl font-bold">
          {{ shortDetail?.translated_name }}
        </h1>
        <span class="block text-sm sm:text-base lg:text-lg mt-1 sm:mt-2">
          {{ `${shortDetail?.name} (${getYear(shortDetail?.release_date as string)})` }}
        </span>
      </div>
      <div class="mt-6 flex flex-col gap-6 sm:gap-8">
        <div class="flex justify-between">
          <div class="flex gap-2 sm:gap-4">
            <Button class="px-4 h-8 border rounded" icon="pi pi-plus" label="Favourite" icon-pos="left" />
            <Button class="px-4 h-8 border rounded" icon="pi pi-bell" label="Following" icon-pos="left" />
          </div>
          <Button icon="pi pi-ellipsis-v" />
        </div>

        <div>
          <span class="block text-base sm:text-lg font-bold">Episodes:</span>
          <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-3 mt-3">
            <Button
              class="w-full truncate px-3 py-2 bg-gray-800 text-white text-center rounded-lg cursor-pointer hover:bg-orange-500"
              v-for="episode in shortEpisodes"
              :label="`Episode ${episode.number}`"
              :key="episode.id"
              @click="onPlayEpisode(episode.id as number)"
            />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import ShortPlayer from "@/components/common/ShortPlayer.vue";
import {useShortDetails} from "@/stores/short-details/useShortDetails.ts";
import {storeToRefs} from "pinia";
import {useRoute, useRouter} from "vue-router";
import {useSubtitle} from "@/stores/subtitles/useSubtitle.ts";
import {onMounted, watch} from "vue";
import {useStream} from "@/stores/short/useStream.ts";
import {getYear} from "date-fns";

const router = useRouter();
const route = useRoute();
const store = useShortDetails();
const { shortEpisodes, shortDetail } = storeToRefs(store);

const subtitleStore = useSubtitle();
const { fetchSubtitles } = subtitleStore;
const { tracks } = storeToRefs(subtitleStore);

const streamStore = useStream();
const { fetchStreamUrl } = streamStore;
const { streamUrl } = storeToRefs(streamStore);

function onPlayEpisode(id: number) {
  const { shortId } = route.params;
  router.push({
    name: "watch",
    params: {
      shortId,
      episodeId: id.toString(),
    },
  });
}

onMounted(async () => {
  await fetchAllData();
});

watch(
    () => route.params,
    async () => {
      await fetchAllData();
    },
    { deep: true }
);

async function fetchAllData() {
  await Promise.allSettled([fetchSubtitles(), fetchStreamUrl()]);
}
</script>
