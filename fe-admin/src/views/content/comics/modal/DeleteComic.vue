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
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await comicStore.deleteComic(comicStore.selectedData.id);
              isLoading = false;
              comicStore.deleteDialogIsOpen = false;
              comicStore.getComicData();
            }
          "
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Delete
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
