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
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const isLoading = ref(false);
const genreStore = useGenreStore();
const checkForm = () => {
  if (genreStore.selectedData.name) return true;
  toast({
    description: "Name cannot be empty!",
    variant: "destructive",
  });
};
</script>
<template>
  <Dialog :open="genreStore.updateDialogIsOpen" @update:open="
    (value: boolean) => {
      genreStore.updateDialogIsOpen = value;
    }
  ">
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
        <Select v-model="genreStore.selectedData.lang">
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="en">English</SelectItem>
              <SelectItem value="ch">Chinese</SelectItem>
              <SelectItem value="vi">Vietnamese</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label for="position" class="text-center w-1/4">Position</Label>
        <Input v-model="genreStore.selectedData.position" placeholder="Position" type="number" />
      </div>
      <DialogFooter class="sm:justify-end">
        <Button variant="secondary" @click="genreStore.updateDialogIsOpen = false">
          Close
        </Button>
        <Button :disabled="isLoading" @click="
          async () => {
            if (checkForm()) {
              isLoading = true;
              await genreStore.updateGenre({
                id: genreStore.selectedData.id,
                name: genreStore.selectedData.name,
                language: genreStore.selectedData.lang,
                position: genreStore.selectedData.position,
              });
              isLoading = false;
              genreStore.updateDialogIsOpen = false;
              genreStore.getGenreData();
            }
          }
        ">
          <img v-if="isLoading" :src="loadingImg" size="icon" />Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
