<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter } from "@/components/ui/dialog";
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
const drama = ref({
  release_date: "2025-04-01",
  thumb: null as string | null,
  translations: [{ language: "en", title: "", description: "" }],
  genres: [],
});

const previewUrl = ref<string | null>(null);
const resetDrama = () => {
  drama.value = {
    release_date: "2025-04-01",
    thumb: null as string | null,
    translations: [{ language: "en", title: "", description: "" }],
    genres: [],
  };
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
  if (drama.value.translations.length == 0) {
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
  <Dialog
    :open="dramaStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
    dramaStore.createDialogIsOpen = value;
    resetDrama();
  }"
  >
    <DialogContent
      class="flex flex-col min-w-[80%] overflow-y-auto max-h-screen h-[80%]"
    >
      <p class="text-3xl font-medium">Create Drama</p>
      <div class="flex flex-row gap-4 h-full">
        <div class="flex flex-col gap-2 w-[20%]">
          <div class="border border-gray-100 h-[150px] rounded-lg">
            <p
              v-if="!previewUrl"
              class="flex justify-center items-center h-full text-gray-500"
            >
              Thumbnail
            </p>
            <img
              v-if="previewUrl"
              :src="previewUrl"
              class="w-full h-full justify-self-center h-auto"
            />
          </div>
          <Input
            type="file"
            placeholder="Cover"
            class="justify-self-center"
            @change="handleFileChange"
            accept="image/*"
          />
          <div class="flex flex-col gap-4">
            <Label class="pt-3 whitespace-nowrap">Release Date</Label>
            <Input type="date" v-model="drama.release_date" />
          </div>
          <div class="flex flex-col gap-4">
            <Label>Genre</Label>
            <VueMultiselect
              v-model="drama.genres"
              :options="dramaStore.generalDramaGenreData"
              :multiple="true"
              :close-on-select="false"
              placeholder="Select genres"
              label="name"
              track-by="name"
            >
            </VueMultiselect>
          </div>
        </div>
        <div class="flex flex-col gap-4 border-l px-4 w-full">
          <div class="flex justify-between items-center w-full">
            <Button
              variant="outline"
              @click="
                () => {
                  drama.translations = [
                    ...drama.translations,
                    { language: 'en', title: '', description: '' },
                  ];
                }
              "
            >
              Add Translation
            </Button>
          </div>
          <div
            v-for="(translation, index) in drama.translations"
            :key="index"
            class="flex flex-row items-center gap-2 w-full"
          >
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
            <Button
              variant="destructive"
              @click="
                () => {
                  drama.translations.splice(index, 1);
                }
              "
              >Delete
            </Button>
          </div>
        </div>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="dramaStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              if (checkForm()) {
                isLoading = true;
                const response = await dramaStore.uploadImage(image);
                drama.thumb = response ? response : null;
                await dramaStore.createDrama(drama);
                isLoading = false;
                dramaStore.createDialogIsOpen = false;
                dramaStore.getDramaData();
                resetDrama();
              }
            }
          "
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Add</Button
        >
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
<style src="vue-multiselect/dist/vue-multiselect.css"></style>
