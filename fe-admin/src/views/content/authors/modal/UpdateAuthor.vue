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
import { useAuthorStore } from "../../../../stores/authorStore";

const authorStore = useAuthorStore();
</script>
<template>
  <Dialog
    :open="authorStore.updateDialogIsOpen"
    @update:open="
      (value : boolean) => {
        authorStore.updateDialogIsOpen = value;
      }
    "
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Author</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Name</Label>
        <Input v-model="authorStore.selectedData.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="biography" class="text-center w-1/4">Biography</Label>
        <Input
          v-model="authorStore.selectedData.biography"
          placeholder="biography"
        />
      </div>
      <div class="flex items-center gap-4">
        <Label for="birth_day" class="text-center w-1/4">Birthday</Label>
        <Input
          v-model="authorStore.selectedData.birth_day"
          placeholder="Birthday"
        />
      </div>
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="authorStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await authorStore.updateAuthor({
                id: authorStore.selectedData.id,
                name: authorStore.selectedData.name,
                birth_day: authorStore.selectedData.birth_day,
                biography: authorStore.selectedData.biography,
              });
              authorStore.updateDialogIsOpen = false;
              authorStore.getAuthorData();
            }
          "
        >
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
