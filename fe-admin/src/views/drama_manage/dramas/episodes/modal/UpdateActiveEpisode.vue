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
                        const formdata = {
                            id: episodeStore.selectedData.id,
                            drama_id: episodeStore.selectedData.drama_id,
                            number: episodeStore.selectedData.number,
                            video: episodeStore.selectedData.video,
                            active: !episodeStore.selectedData.active,
                            subtitles: episodeStore.selectedData.subtitles
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
