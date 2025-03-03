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
import { useChapterItemStore } from "../../../../../../stores/chapteritemStore";
const isLoading = ref(false);
const chapteritemStore = useChapterItemStore();
</script>
<template>
    <Dialog :open="chapteritemStore.deleteDialogIsOpen" @update:open="(value: boolean) => {
        chapteritemStore.deleteDialogIsOpen = value;
    }">
        <DialogContent>
            <DialogHeader>Do you really want to delete this chapter item?</DialogHeader>
            <DialogFooter class="justify-end gap-4">
                <Button variant="secondary" @click="chapteritemStore.deleteDialogIsOpen = false">
                    Close
                </Button>
                <Button :disabled="isLoading" @click="
                    async () => {
                        isLoading = true;
                        await chapteritemStore.deleteChapterItem(chapteritemStore.selectedData.id);
                        isLoading = false;
                        chapteritemStore.deleteDialogIsOpen = false;
                        chapteritemStore.getChapterItemsData();
                    }
                ">
                    <img v-if="isLoading" :src="loadingImg" size="icon" />
                    Delete
                </Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>
