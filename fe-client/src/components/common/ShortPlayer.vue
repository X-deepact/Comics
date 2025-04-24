<template>
  <div
    class="outline-none bg-black"
    :tabindex="0"
    autofocus
    @keydown.prevent.space="playing = !playing"
    @keydown.right="currentTime += 10"
    @keydown.left="currentTime -= 10"
    @keydown.esc="() => { if (isFullscreen) exit}"
  >
    <div class="relative aspect-[9/16] max-h-[70vh] max-w-[400px] mx-auto shadow-lg">
      <video
        ref="video"
        crossorigin="anonymous"
        class="absolute w-full h-full"
        :loop="loop"
        @ended="() => onAutoPlay()"
        playsinline
      />
      <div
          v-if="waiting"
          class="absolute inset-0 grid place-items-center pointer-events-none bg-black bg-opacity-20"
      >
        <ProgressSpinner style="width: 50px; height: 50px" strokeWidth="8" fill="transparent"
                         animationDuration=".5s" aria-label="Custom ProgressSpinner"
        />
      </div>
    </div>
    <div class="flex flex-row items-center h-12 px-0 md:px-4 bg-gray-900">
      <ButtonWithTooltip>
        <template #btn>
          <Button
            class="text-orange-500"
            size="large"
            @click="playing = !playing"
            :icon="playing ? 'pi pi-pause' : 'pi pi-play'"
          />
        </template>
        <template #tooltip>
          <span>{{ playing ? 'Pause' : 'Play' }}</span>
        </template>
      </ButtonWithTooltip>

      <ButtonWithTooltip>
        <template #btn>
          <Button
            class="text-orange-500"
            size="large"
            @click="onPlayPreviousEp"
            icon="pi pi-step-backward"
            :disabled="isFirstEpisode"
          />
        </template>
        <template #tooltip>
          <span>Previous Episode</span>
        </template>
      </ButtonWithTooltip>

      <ButtonWithTooltip>
        <template #btn>
          <Button
            class="text-orange-500"
            size="large"
            @click="onPlayNextEp"
            icon="pi pi-step-forward"
            :disabled="isLastEpisode"
          />
        </template>
        <template #tooltip>
          <span>Next Episode</span>
        </template>
      </ButtonWithTooltip>

      <ButtonWithTooltip>
        <template #btn>
          <Button
            class="text-orange-500"
            size="large"
            @click="muted = !muted"
            :icon="muted ? 'pi pi-volume-off' : 'pi pi-volume-up'"
          />
        </template>
        <template #tooltip>
          <span>{{ muted ? 'Unmute' : 'Mute' }}</span>
        </template>
      </ButtonWithTooltip>

      <Scrubber
        v-model="currentTime"
        :max="duration"
        :secondary="endBuffer"
        class="flex-1 min-w-0"
      >
        <template #default="{ position, pendingValue }">
          <div
              class="absolute transform -translate-x-1/2 bg-black rounded px-2 bottom-0 mb-4 py-1 text-xs text-white"
              :style="{ left: position }"
          >
            {{ formatDuration(pendingValue) }}
          </div>
        </template>
      </Scrubber>

      <div class="hidden md:flex flex-col text-sm mx-0 sm:mx-2 text-orange-500">
        {{ formatDuration(currentTime) }} / {{ formatDuration(duration) }}
      </div>

      <MenuPlayer>
        <template #default="{ open }">
          <Button @click="open" icon="pi pi-language" size="large" class="text-orange-500"/>
        </template>
        <template #menu>
          <div class="absolute bottom-0 right-0 bg-black transform -translate-y-[2.5rem] rounded py-2 shadow">
            <MenuPlayerItem
                @click="() => { disableTrack(); }"
            >
              <span>Subtitle Off</span>
              <i v-if="selectedTrack === -1" class="pi pi-check" />
            </MenuPlayerItem>
            <MenuPlayerItem
                v-for="track in tracks"
                :key="track.id"
                @keydown.stop.prevent.enter.space="enableTrack(track)"
                @click="() => { enableTrack(track); }"
            >
              <span>{{ track.label }}</span>
              <i v-if="track.mode === 'showing'" class="pi pi-check" />
            </MenuPlayerItem>
          </div>
        </template>
      </MenuPlayer>

<!--      Configuration Menu-->
      <MenuPlayer>
        <template #default="{ open }">
          <Button @click="open" icon="pi pi-cog" size="large" class="text-orange-500" />
        </template>
        <template #menu="{ close }">
          <div class="absolute bottom-0 right-0 bg-black transform -translate-y-[2.5rem] rounded py-2 shadow">
            <MenuPlayerItem
                @click="() => { togglePictureInPicture(); close(); }"
            >
              <Button icon="pi pi-window-minimize" label="PiP" size="small" />
            </MenuPlayerItem>

            <MenuPlayerItem>
              <Button icon="pi pi-play-circle" label="Speeds" disabled size="small" />
              <select
                  id="speedSelect"
                  v-model.number="rate"
                  class="block w-[3.75rem] bg-gray-800 border border-gray-600 rounded text-sm text-white focus:outline-none focus:ring-2 focus:ring-orange-500"
              >
                <option
                    v-for="speed in playbackSpeeds"
                    :key="speed"
                    :value="speed"
                >
                  {{ speed }}x
                </option>
              </select>
            </MenuPlayerItem>

            <MenuPlayerItem
                @click="() => { loop = !loop }"
            >
              <Button icon="pi pi-replay" label="Loop" size="small" />
              <i v-if="loop" class="pi pi-check" />
            </MenuPlayerItem>

            <MenuPlayerItem>
              <div class="flex flex-row justify-between w-full">
                <Button @click="autoplay = !autoplay" label="Autoplay" size="small" />

                <div class="relative inline-block w-10 align-middle select-none transition duration-200 ease-in">
                  <input
                      type="checkbox"
                      id="autoplayToggle"
                      class="toggle-checkbox absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer"
                      :checked="autoplay"
                      @change="autoplay = !autoplay"
                  />
                  <label
                    for="autoplayToggle"
                    class="toggle-label block overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer"
                  ></label>
                </div>
              </div>
            </MenuPlayerItem>

            <MenuPlayerItem>
              <Button icon="pi pi-expand" label="Full Screen" @click="enter" size="small" />
            </MenuPlayerItem>
          </div>
        </template>
      </MenuPlayer>

      <ButtonWithTooltip>
        <template #btn>

        </template>
        <template #tooltip>
          <span>Fullscreen</span>
        </template>
      </ButtonWithTooltip>
    </div>
  </div>
</template>
<script setup lang="ts">
import {useMediaControls, useFullscreen} from '@vueuse/core'
import type { UseMediaTextTrackSource } from '@vueuse/core';
import {computed, shallowRef, useTemplateRef, watch, onMounted, ref} from 'vue';
import Scrubber from "@/components/common/Scrubber.vue";
import MenuPlayer from "@/components/common/MenuPlayer.vue";
import MenuPlayerItem from "@/components/common/MenuPlayerItem.vue";
import ButtonWithTooltip from "@/components/common/ButtonWithTooltip.vue";
import Hls from 'hls.js';
import {useShortDetails} from "@/stores/short-details/useShortDetails.ts";
import {useRouter} from "vue-router";
import {storeToRefs} from "pinia";

const props = defineProps<{
  source: string;
  tracks?: UseMediaTextTrackSource[]
}>();
const video = useTemplateRef<HTMLVideoElement>('video');
const loop = shallowRef(false);
const autoplay = shallowRef(true);
const startWhenInit = shallowRef(true);
const key = ref(0); // key to force page reload;
const playbackSpeeds = [0.25, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2];
const router = useRouter();
const { isFirstEpisode, isLastEpisode } = storeToRefs(useShortDetails());

const { isFullscreen, enter, exit } = useFullscreen(video);

const controls = useMediaControls(video, {
  src: {
    src: props.source,
    type: 'application/x-mpegURL',
  },
  tracks: props.tracks
});

const {
  playing,
  buffered,
  currentTime,
  duration,
  tracks,
  waiting,
  muted,
  rate,
  selectedTrack,
  enableTrack,
  disableTrack,
  togglePictureInPicture,
} = controls;

const endBuffer = computed(() => buffered.value.length > 0 ? buffered.value[buffered.value.length - 1][1] : 0);
function formatDuration(seconds: number) {
  return new Date(1000 * seconds).toISOString().slice(14, 19);
}

function setupHls() {
  if (!video.value || !props.source) return;

  if (Hls.isSupported()) {
    const hls = new Hls();

    hls.loadSource(props.source);
    hls.attachMedia(video.value);

    hls.on(Hls.Events.MANIFEST_PARSED, () => {
      if (startWhenInit.value) {
        playing.value = true;
      }
    })
  } else if (video.value.canPlayType('application/vnd.apple.mpegurl')) {
    video.value.src = props.source;
    video.value.load();
    if (startWhenInit.value) {
      playing.value = true;
    }
  }
}

function onAutoPlay() {
  if (autoplay.value) {
    onPlayNextEp();
    key.value++;
  }
}

function onPlayNextEp() {
  const { nextEpisode } = useShortDetails();
  router.push({
    name: "watch",
    params: {
      shortId: nextEpisode?.shortId,
      episodeId: nextEpisode?.episodeId,
    },
  });
}

function onPlayPreviousEp() {
  const { previousEpisode } = useShortDetails();
  router.push({
    name: "watch",
    params: {
      shortId: previousEpisode?.shortId,
      episodeId: previousEpisode?.episodeId,
    },
  });
}

onMounted(() => {
  setupHls();
})

watch(
    () => props.source,
    () => {
      setupHls();
    }
)

watch(
  () => props.tracks,
  (newTracks) => {
    if (!video.value) return;

    const oldTracks = video.value.querySelectorAll('track');
    oldTracks.forEach(track => track.remove());

    newTracks?.forEach((track) => {
      const trackEl = document.createElement('track');
      trackEl.kind = track.kind;
      trackEl.label = track.label;
      trackEl.srclang = track.srcLang;
      trackEl.src = track.src;
      video.value?.appendChild(trackEl);
    });
  },
  { immediate: true, deep: true }
);

</script>

<style>
.toggle-checkbox:checked {
  right: 0;
  border-color: #ff620e;
}
.toggle-checkbox:checked + .toggle-label {
  background-color: #ff620e;
}
</style>
