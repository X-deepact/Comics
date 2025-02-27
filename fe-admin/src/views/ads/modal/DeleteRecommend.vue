<template>
  <Dialog
    :open="recommendStore.deleteDialogIsOpen"
    @update:open="recommendStore.deleteDialogIsOpen = $event"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Delete Recommendation</DialogTitle>
      </DialogHeader>
      <div class="py-4">Do you really want to delete this recommendation?</div>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="recommendStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          variant="destructive"
          @click="
            async () => {
              await recommendStore.deleteRecommend(recommendStore.selectedData.id);
              recommendStore.deleteDialogIsOpen = false;
              recommendStore.getRecommendData();
            }
          "
        >
          Delete
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { useRecommendStore } from "@/stores/recommendStore";

const recommendStore = useRecommendStore();
</script> 