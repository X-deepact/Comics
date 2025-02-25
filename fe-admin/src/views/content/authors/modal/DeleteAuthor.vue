<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
} from "@/components/ui/dialog";
import { useAuthorStore } from "../../../../stores/authorStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
const isLoading = ref(false);
const authorStore = useAuthorStore();
</script>
<template>
  <Dialog
    :open="authorStore.deleteDialogIsOpen"
    @update:open="(value:boolean)=>{
      authorStore.deleteDialogIsOpen = value;
      }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this author?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="authorStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await authorStore.deleteAuthor(authorStore.selectedData.id);
              isLoading = false;
              authorStore.deleteDialogIsOpen = false;
              authorStore.getAuthorData();
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
