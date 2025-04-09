<script setup lang="ts">
import {
    Dialog,
    DialogContent,
    DialogHeader,
} from "@/components/ui/dialog";
import { useEpisodeStore } from "../../../../../stores/episodeStore";
import M3uPlayer from "../../../../../lib/Player.vue"
import { langConverter } from "../../../../../lib/utils";
const episodeStore = useEpisodeStore();
const formattedSubtitles = (datas: any) => {
    if (Array.isArray(datas)) {
        return datas.map((track: any) => ({
            language: langConverter(track.language),
            kind: 'subtitles',
            src: track.url
        }));
    } else {
        return [];
    }
}
</script>
<template>
    <Dialog :open="episodeStore.playDialogIsOpen" @update:open="(value: boolean) => {
        episodeStore.playDialogIsOpen = value;
    }">
        <DialogContent class="min-w-[80%]">
            <DialogHeader class="flex justify-center items-center py-1 font-medium">{{ episodeStore.selectedData.number
                }}
            </DialogHeader>
            <div class="w-full h-auto">
                <M3uPlayer :videoSource="episodeStore.selectedData.video"
                    :subtitles="formattedSubtitles(episodeStore.selectedData.subtitles)" />
            </div>
        </DialogContent>
    </Dialog>
</template>
