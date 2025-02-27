<template>
  <Dialog
    :open="adsStore.deleteDialogIsOpen"
    @update:open="(value: boolean) => {
      adsStore.deleteDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Delete Advertisement</DialogTitle>
        <DialogDescription>
          Are you sure you want to delete this advertisement?
        </DialogDescription>
      </DialogHeader>
      <DialogFooter>
        <Button
          variant="secondary"
          @click="adsStore.deleteDialogIsOpen = false"
        >
          Cancel
        </Button>
        <Button
          variant="destructive"
          @click="
            async () => {
              await adsStore.deleteAd(adsStore.selectedData.id);
              adsStore.deleteDialogIsOpen = false;
              adsStore.getAdsData();
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
import { useAdsStore } from "@/stores/adsStore";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

const adsStore = useAdsStore();
</script> 