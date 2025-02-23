<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useGenreStore } from "../../../../stores/genreStore";

const genreStore = useGenreStore();
</script>
<template>
  <Dialog
    :open="genreStore.updateDialogIsOpen"
    @update:open="
      (value : boolean) => {
        genreStore.updateDialogIsOpen = value;
      }
    "
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Genre</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Name</Label>
        <Input v-model="genreStore.selectedData.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="lang" class="text-center w-1/4">Language</Label>
        <Input v-model="genreStore.selectedData.lang" placeholder="Language" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="position" class="text-center w-1/4">Position</Label>
        <Input
          v-model="genreStore.selectedData.position"
          placeholder="Position"
        />
      </div>
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="genreStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await genreStore.updateGenre({
                id: genreStore.selectedData.id,
                name: genreStore.selectedData.name,
                language: genreStore.selectedData.lang,
                position: genreStore.selectedData.position,
              });
              genreStore.updateDialogIsOpen = false;
              genreStore.getGenreData();
            }
          "
        >
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
