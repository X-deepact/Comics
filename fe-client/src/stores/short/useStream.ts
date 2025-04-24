import {defineStore} from "pinia";
import {ref} from "vue";
import {useRoute} from "vue-router";
import {API} from "@/services";
import type {EpisodeResponse} from "@/services/short/types.ts";

export const useStream = defineStore("stream", () => {
  const streamUrl = ref<EpisodeResponse | undefined>();
  const route = useRoute();

  const fetchStreamUrl = async () => {
    const { shortId, episodeId } = route.params;

    const response = await API.service.get<EpisodeResponse>(
      `episode/${episodeId}`
    )

    if (response.status === 200) streamUrl.value = response.data;
  }

  return {
    streamUrl,
    fetchStreamUrl
  }
})