<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
} from "@/components/ui/dialog";
import { useAuthorStore } from "../../../../stores/authorStore";

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
          @click="
            async () => {
              await authorStore.deleteAuthor(authorStore.selectedData.id);
              authorStore.deleteDialogIsOpen = false;
              authorStore.getAuthorData();
            }
          "
          >Delete
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
