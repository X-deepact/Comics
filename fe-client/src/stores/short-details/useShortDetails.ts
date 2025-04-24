import {defineStore, storeToRefs} from "pinia";
import {computed, ref, watch} from "vue";
import { API } from "@/services";
import type {ShortDramaDetailResponse} from "@/services/short/types.ts";
import { useRoute } from "vue-router";
import {useLocale} from "@/stores/i18n/useLocale.ts";

export const useShortDetails = defineStore("ShortDetailStore", () => {
  const shortDetail = ref<ShortDramaDetailResponse | undefined>(undefined);
  const route = useRoute();
  const { locale } = storeToRefs(useLocale());
  const shortDetailResponseErrorCode = ref(false);
  const isLoading = ref(false);

  const shortEpisodes = computed(() => {
    if (!shortDetail.value) {
      return [];
    }
    return shortDetail.value?.episodes?.sort((a, b) => {
      return a.number! - b?.number!;
    });
  });

  const currentEpisodeIndex = computed(() => {
    const episodeId = Number(route.params.episodeId);
    return shortEpisodes.value.findIndex((ep) => ep.id === episodeId);
  });

  const isFirstEpisode = computed(() => currentEpisodeIndex.value === 0);
  const isLastEpisode = computed(() => currentEpisodeIndex.value === shortEpisodes.value.length - 1);

  // using this variable for autoplay and play next ep functionality
  const nextEpisode = computed(() => {
    const index = currentEpisodeIndex.value;
    if (index === -1 || index >= shortEpisodes.value.length - 1) return null;

    return {
      shortId: route.params.shortId,
      episodeId: shortEpisodes.value[index + 1]?.id,
    };
  });

  const previousEpisode = computed(() => {
    const index = currentEpisodeIndex.value;
    if (index <= 0) return null;

    return {
      shortId: route.params.shortId,
      episodeId: shortEpisodes.value[index - 1]?.id,
    };
  });
  const fetchShortDetails = async () => {
    shortDetailResponseErrorCode.value = false;
    shortDetail.value = undefined;
    isLoading.value = true;

    const { shortId } = route.params;
    if (shortId === undefined) {
      shortDetailResponseErrorCode.value = true;
      return;
    }

    const response = await API.service.get<ShortDramaDetailResponse>
    (`short-drama/${shortId}`, {
        params: {
          language: locale.value,
        }
      }
    );

    if (response.code == 'ERROR') {
      shortDetailResponseErrorCode.value = true;
    } else {
      shortDetail.value = response.data;
    }
    isLoading.value = false;
  };

  watch(
    () => route.params,
    async () => {
      await fetchShortDetails();
    },
    {
      deep: true,
      immediate: true
    }
  );

  watch(
    () => locale.value,
    async () => {
      await fetchShortDetails();
    }
  )

  return {
    shortDetail,
    fetchShortDetails,
    shortEpisodes,
    shortDetailResponseErrorCode,
    isLoading,
    nextEpisode,
    previousEpisode,
    isFirstEpisode,
    isLastEpisode,
  };
});
