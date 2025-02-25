<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
} from "@/components/ui/dialog";
import { useGenreStore } from "../../../../stores/genreStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
const isLoading = ref(false);
const genreStore = useGenreStore();
</script>
<template>
  <Dialog
    :open="genreStore.deleteDialogIsOpen"
    @update:open="(value:boolean)=>{
      genreStore.deleteDialogIsOpen = value;
      }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this genre?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="genreStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await genreStore.deleteGenre(genreStore.selectedData.id);
              isLoading = false;
              genreStore.deleteDialogIsOpen = false;
              genreStore.getGenreData();
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
