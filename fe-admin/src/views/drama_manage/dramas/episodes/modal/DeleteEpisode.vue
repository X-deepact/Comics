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
const episodeStore = useEpisodeStore();
</script>
<template>
    <Dialog :open="episodeStore.deleteDialogIsOpen" @update:open="(value: boolean) => {
        episodeStore.deleteDialogIsOpen = value;
    }">
        <DialogContent>
            <DialogHeader>Do you really want to delete this episode?</DialogHeader>
            <DialogFooter class="justify-end gap-4">
                <Button variant="secondary" @click="episodeStore.deleteDialogIsOpen = false">
                    Close
                </Button>
                <Button :disabled="isLoading" @click="
                    async () => {
                        isLoading = true;
                        await episodeStore.deleteEpisode(episodeStore.selectedData.id);
                        isLoading = false;
                        episodeStore.deleteDialogIsOpen = false;
                        episodeStore.getEpisodeData();
                    }
                ">
                    <img v-if="isLoading" :src="loadingImg" size="icon" />
                    Delete
                </Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>
