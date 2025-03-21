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
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { ref } from "vue";
import { useGenreStore } from "../../../../stores/genreStore";
import loadingImg from "@/assets/loading.svg";
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const genreStore = useGenreStore();
genreStore.getGenreData();
const isLoading = ref(false);
const genre = ref({
  name: "",
  position: 1,
  language: "en",
});
const resetGenre = () => {
  genre.value = {
    name: "",
    position: 1,
    language: "en",
  };
};
const checkForm = () => {
  if (genre.value.name) return true;
  toast({
    description: "Name cannot be empty!",
    variant: "destructive",
  });
};
</script>
<template>
  <Dialog :open="genreStore.createDialogIsOpen"
    @update:open="(value: boolean) => { genreStore.createDialogIsOpen = value; }">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Genre</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Name</Label>
        <Input v-model="genre.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="lang" class="text-center w-1/4">Language</Label>
        <Select v-model="genre.language">
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="en">English</SelectItem>
              <SelectItem value="zh">Chinese</SelectItem>
              <SelectItem value="vi">Vietnamese</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label for="position" class="text-center w-1/4">Position</Label>
        <Input v-model="genre.position" placeholder="Position" type="number" />
      </div>
      <DialogFooter class="sm:justify-end">
        <Button variant="secondary" @click="genreStore.createDialogIsOpen = false">
          Close
        </Button>
        <Button :disabled="isLoading" @click="
          async () => {
            if (checkForm()) {
              isLoading = true;
              await genreStore.createGenre(genre);
              isLoading = false;
              genreStore.createDialogIsOpen = false;
              genreStore.getGenreData();
              resetGenre();
            }
          }
        "><img v-if="isLoading" :src="loadingImg" size="icon" />Add</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
