<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogTitle,
  DialogHeader,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";

import { Input } from "@/components/ui/input";
import { ref } from "vue";
import { useGenreStore } from "../../../../stores/genreStore";

const genreStore = useGenreStore();
genreStore.getGenreData();
const genre = ref({
  name: "",
  position: "",
  lang: "",
});
const resetGenre = () => {
  genre.value = {
    name: "",
    position: "",
    lang: "",
  };
};
</script>
<template>
  <Dialog
    :open="genreStore.createDialogIsOpen"
    @update:open="(value:boolean)=>{genreStore.createDialogIsOpen = value;}"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Subject</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Name</Label>
        <Input v-model="genre.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="lang" class="text-center w-1/4">Language</Label>
        <Input v-model="genre.lang" placeholder="Language" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="position" class="text-center w-1/4">Position</Label>
        <Input v-model="genre.position" placeholder="Position" />
      </div>
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="genreStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await genreStore.createGenre(genre);
              genreStore.createDialogIsOpen = false;
              genreStore.getGenreData();
              resetGenre();
            }
          "
          >Add</Button
        >
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
