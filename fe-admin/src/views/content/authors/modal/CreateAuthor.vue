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
import { useAuthorStore } from "../../../../stores/authorStore";

const authorStore = useAuthorStore();
authorStore.getGeneralAuthorData();
const author = ref({
  name: "",
  biography: "",
  birth_day: "",
});

const resetAuthor = () => {
  author.value = {
    name: "",
    biography: "",
    birth_day: "",
  };
};
</script>
<template>
  <Dialog
    :open="authorStore.createDialogIsOpen"
    @update:open="(value:boolean)=>{authorStore.createDialogIsOpen = value;}"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Author</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Title</Label>
        <Input v-model="author.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="subject" class="text-center w-1/4">Subject</Label>
        <Input v-model="author.biography" placeholder="Biography" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Description</Label>
        <Input v-model="author.birth_day" placeholder="Birthday" />
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="authorStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await authorStore.createAuthor(author);
              authorStore.createDialogIsOpen = false;
              authorStore.getAuthorData();
              resetAuthor();
            }
          "
          >Add</Button
        >
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
