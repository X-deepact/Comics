<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
} from "@/components/ui/dialog";
import { useComicStore } from "../../../../stores/comicStore";

const comicStore = useComicStore();
</script>
<template>
  <Dialog
    :open="comicStore.deleteDialogIsOpen"
    @update:open="(value:boolean)=>{
      comicStore.deleteDialogIsOpen = value;
      }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this comic?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="comicStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await comicStore.deleteComic(comicStore.selectedData.id);
              comicStore.deleteDialogIsOpen = false;
              comicStore.getComicData();
            }
          "
          >Delete
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
