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
import { useChapterStore } from "../../../../../stores/chapterStore";
const isLoading = ref(false);
const chapterStore = useChapterStore();
</script>
<template>
  <Dialog
    :open="chapterStore.deleteDialogIsOpen"
    @update:open="(value:boolean)=>{
      chapterStore.deleteDialogIsOpen = value;
      }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this chapter?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="chapterStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await chapterStore.deleteChapter(chapterStore.selectedData.id);
              isLoading = false;
              chapterStore.deleteDialogIsOpen = false;
              chapterStore.getChapterData();
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
