<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogHeader,
} from "@/components/ui/dialog";
import { useDramaStore } from "../../../../stores/dramaStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
const isLoading = ref(false);
const dramaStore = useDramaStore();
</script>
<template>
    <Dialog :open="dramaStore.updateActiveDialogIsOpen" @update:open="(value: boolean) => {
        dramaStore.updateActiveDialogIsOpen = value;
    }">
        <DialogContent>
            <DialogHeader>Do you really want to update active status of this drama?</DialogHeader>
            <DialogFooter class="justify-end gap-4">
                <Button variant="secondary" @click="dramaStore.updateActiveDialogIsOpen = false">
                    Close
                </Button>
                <Button :disabled="isLoading" @click="
                    async () => {
                        isLoading = true;
                        await dramaStore.setActiveDrama(dramaStore.selectedData.id);
                        isLoading = false;
                        dramaStore.updateActiveDialogIsOpen = false;
                        dramaStore.getDramaData();
                    }
                ">
                    <img v-if="isLoading" :src="loadingImg" size="icon" />
                    Update
                </Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>
