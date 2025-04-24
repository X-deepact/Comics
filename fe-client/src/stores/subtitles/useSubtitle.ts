import {defineStore} from "pinia";
import {computed, ref} from "vue";
import type {ComputedRef} from "vue";
import {useRoute} from "vue-router";
import type {ShortSubtitle} from "@/services/subtitles/types.ts";
import {API} from "@/services";
import type {UseMediaTextTrackSource} from "@vueuse/core";

export const useSubtitle = defineStore("subtitle store", () => {
  const subtitles = ref<ShortSubtitle[]>([]);
  const route = useRoute();
  const tracks: ComputedRef<UseMediaTextTrackSource[]> = computed(() => {
    return subtitles.value.map((subtitle) => {
      return {
        src: subtitle.subtitle_url,
        kind: 'subtitles',
        label: getLabel(subtitle.language) as string,
        srcLang: subtitle.language
      }
    })
  })

  const fetchSubtitles = async () => {
    const { episodeId } = route.params;

    const response = await API.service.getList<ShortSubtitle>(
      `subtitles/${episodeId}`
    )

    if (response.status === 200) subtitles.value = response.data;
  };
  function getLabel(label: string) {
    switch (label) {
      case 'en':
        return 'English'
      case 'vi':
        return 'Vietnamese'
      case 'zh':
        return 'Chinese'
    }
  }

  return {
    subtitles,
    fetchSubtitles,
    tracks
  }
})