<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogHeader,
} from "@/components/ui/dialog";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
import { useEpisodeStore } from "../../../../../stores/episodeStore";

// Add interface for subtitle type
interface Subtitle {
    language: string;
    url: string;
}

const isLoading = ref(false);
const episodeStore = useEpisodeStore()
</script>
<template>
    <Dialog :open="episodeStore.updateActiveDialogIsOpen" @update:open="(value: boolean) => {
        episodeStore.updateActiveDialogIsOpen = value;
    }">
        <DialogContent>
            <DialogHeader>Do you really want to update active status of this episode?</DialogHeader>
            <DialogFooter class="justify-end gap-4">
                <Button variant="secondary" @click="episodeStore.updateActiveDialogIsOpen = false">
                    Close
                </Button>
                <Button :disabled="isLoading" @click="
                    async () => {
                        isLoading = true;
                        let video = episodeStore.selectedData.video;

                        if (video.includes(`http`)) {
                            const parts = video.split(`/`);
                            video = parts.slice(-2).join(`/`);
                        }

                        const subtitles = episodeStore.selectedData.subtitles.map((item: Subtitle) => {
                            if (item.url.includes('http')) {
                                const url = item.url.split('/').pop();
                                return { language: item.language, url: url };
                            } else {
                                return { ...item };
                            }
                        });

                        console.log(subtitles)

                        const formdata = {
                            id: episodeStore.selectedData.id,
                            drama_id: episodeStore.selectedData.drama_id,
                            number: episodeStore.selectedData.number,
                            video: video,
                            active: !episodeStore.selectedData.active,
                            subtitles: subtitles
                        }
                        await episodeStore.updateEpisode(formdata);
                        isLoading = false;
                        episodeStore.updateActiveDialogIsOpen = false;
                        episodeStore.getEpisodeData();
                    }
                ">
                    <img v-if="isLoading" :src="loadingImg" size="icon" />
                    Update
                </Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>
