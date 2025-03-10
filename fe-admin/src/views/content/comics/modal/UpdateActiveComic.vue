<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogHeader,
} from "@/components/ui/dialog";
import { useComicStore } from "../../../../stores/comicStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
const isLoading = ref(false);
const comicStore = useComicStore();
</script>
<template>
    <Dialog :open="comicStore.updateActiveDialogIsOpen" @update:open="(value: boolean) => {
        comicStore.updateActiveDialogIsOpen = value;
    }">
        <DialogContent>
            <DialogHeader>Do you really want to update active status of this comic?</DialogHeader>
            <DialogFooter class="justify-end gap-4">
                <Button variant="secondary" @click="comicStore.updateActiveDialogIsOpen = false">
                    Close
                </Button>
                <Button :disabled="isLoading" @click="
                    async () => {
                        isLoading = true;
                        await comicStore.setActiveComic(comicStore.selectedData.id);
                        isLoading = false;
                        comicStore.updateActiveDialogIsOpen = false;
                        comicStore.getComicData();
                    }
                ">
                    <img v-if="isLoading" :src="loadingImg" size="icon" />
                    Update
                </Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>
