<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
} from "@/components/ui/dialog";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { ref } from "vue";
import { useDramaStore } from "../../../../stores/dramaStore";
import VueMultiselect from "vue-multiselect";
import loadingImg from "@/assets/loading.svg";
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const isLoading = ref(false);
const dramaStore = useDramaStore();
dramaStore.getGeneralGenreData();

const previewUrl = ref<string | null>(null);
const resetDrama = () => {
  image.value = null;
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
    previewUrl.value = null;
  }
};
const image = ref<File | null>(null);
const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    // Clear previous preview URL if it exists
    if (previewUrl.value) {
      URL.revokeObjectURL(previewUrl.value);
    }
    image.value = target.files[0];
    // Create new preview URL
    previewUrl.value = URL.createObjectURL(target.files[0]);
  }
};
const checkForm = () => {
  if (dramaStore.selectedData.translations.length == 0) {
    toast({
      description: "Name is required",
      variant: "destructive",
    });
    return false;
  } else {
    return true;
  }
};
</script>
<template>
  <Dialog :open="dramaStore.updateDialogIsOpen" @update:open="(value: boolean) => {
    dramaStore.updateDialogIsOpen = value;
    resetDrama();
  }">
    <DialogContent class="flex flex-col min-w-[80%] overflow-y-auto max-h-screen h-[80%]">
      <p class="text-3xl font-medium">Update Drama</p>
      <div class="flex flex-row gap-4 h-full">
        <div class="flex flex-col gap-2 w-[20%]">
          <div class="border border-gray-100 h-[150px] rounded-lg">
            <img v-if="previewUrl" :src="previewUrl" class="w-full h-full justify-self-center h-auto" />
            <img v-if="!previewUrl" :src="dramaStore.selectedData.thumb"
              class="w-full h-full justify-self-center h-auto" />
          </div>
          <Input type="file" class="justify-self-center" @change="handleFileChange" accept="image/*" />
          <div class="flex flex-col gap-4">
            <Label class="pt-3 whitespace-nowrap">Release Date</Label>
            <Input type="date" v-model="dramaStore.selectedData.release_date" />
          </div>
          <div class="flex flex-col gap-4">
            <Label>Genre</Label>
            <VueMultiselect v-model="dramaStore.selectedData.genres" :options="dramaStore.generalDramaGenreData"
              :multiple="true" :close-on-select="false" placeholder="Select genres" label="name" track-by="name">
            </VueMultiselect>
          </div>
        </div>
        <div class="flex flex-col gap-4 border-l px-4 w-full">
          <div class="flex justify-between items-center w-full">
            <Button variant="outline" @click="
              () => {
                dramaStore.selectedData.translations.push({
                  language: 'en',
                  title: '',
                  description: ''
                })
              }">
              Add
            </Button>
          </div>
          <div v-for="(translation, index) in dramaStore.selectedData.translations" :key="index"
            class="flex flex-row items-center gap-2 w-full">
            <Select v-model="translation.language">
              <SelectTrigger class="w-[20%]">
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
            <Label class="col-span-1">Title</Label>
            <Input v-model="translation.title" class="w-[50%]" />
            <Label>Description</Label>
            <Input v-model="translation.description" class="min-w-[40%]" />
            <Button variant="destructive" @click="
              () => {
                dramaStore.selectedData.translations.splice(index, 1);
              }
            ">Delete
            </Button>
          </div>
        </div>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button variant="secondary" @click="dramaStore.updateDialogIsOpen = false">
          Close
        </Button>
        <Button :disabled="isLoading" @click="
          async () => {
            if (checkForm()) {
              isLoading = true;
              if (image) {
                dramaStore.selectedData.thumb = await dramaStore.uploadImage(image);
              } else { dramaStore.selectedData.thumb = undefined }
              await dramaStore.updateDrama(dramaStore.selectedData);
              isLoading = false;
              dramaStore.updateDialogIsOpen = false;
              dramaStore.getDramaData();
              resetDrama();
            }
          }
        ">
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Update</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
<style src="vue-multiselect/dist/vue-multiselect.css"></style>
