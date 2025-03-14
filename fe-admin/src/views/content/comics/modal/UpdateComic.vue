<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useComicStore } from "../../../../stores/comicStore";
import VueMultiselect from "vue-multiselect";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
import { useAuthorStore, type Author } from "../../../../stores/authorStore";
import { useGenreStore, type Genre } from "../../../../stores/genreStore";
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const isLoading = ref(false);
const comicStore = useComicStore();
const authorStore = useAuthorStore();
const genreStore = useGenreStore();
genreStore.getGeneralGenreData();
authorStore.getGeneralAuthorData();
const comic = ref({
  id: 0,
  name: "",
  author: [] as Author[],
  genre: [] as Genre[],
  code: "",
  cover: null as File | null,
  description: "",
  lang: "all",
  original_language: "en",
  audience: "all",
  status: "ongoing",
});
const previewUrl = ref<string | null>(null);
const resetComic = () => {
  comic.value = {
    id: 0,
    name: "",
    author: [],
    genre: [],
    code: "",
    cover: null,
    description: "",
    lang: "en",
    original_language: "en",
    audience: "",
    status: "ongoing",
  };
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
    previewUrl.value = null;
  }
};
const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    // Clear previous preview URL if it exists
    if (previewUrl.value) {
      URL.revokeObjectURL(previewUrl.value);
    }
    comic.value.cover = target.files[0];
    // Create new preview URL
    previewUrl.value = URL.createObjectURL(target.files[0]);
  }
};
const checkForm = () => {
  if (comicStore.selectedData.name == "") {
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
  <Dialog :open="comicStore.updateDialogIsOpen" @update:open="
    (value: boolean) => {
      comicStore.updateDialogIsOpen = value;
      resetComic();
    }
  ">
    <DialogContent class="overflow-y-scroll max-h-screen">
      <DialogHeader>
        <DialogTitle>Update Comic</DialogTitle>
      </DialogHeader>
      <img v-if="!previewUrl" :src="comicStore.selectedData.cover" class="max-w-[200px] justify-self-center h-auto" />
      <img v-if="previewUrl" :src="previewUrl" class="max-w-[200px] justify-self-center h-auto" />
      <Input type="file" placeholder="cover" class="justify-self-center w-[50%]" @change="handleFileChange"
        accept="image/*" />
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Name</Label>
        <Input v-model="comicStore.selectedData.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Description</Label>
        <Input v-model="comicStore.selectedData.description" placeholder="Description" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Author</Label>
        <VueMultiselect v-model="comicStore.selectedData.authors" :options="authorStore.generalauthorData"
          :multiple="true" :close-on-select="false" placeholder="Select authors" label="name" track-by="name">
        </VueMultiselect>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Genre</Label>
        <VueMultiselect v-model="comicStore.selectedData.genres" :options="genreStore.generalGenreData" :multiple="true"
          :close-on-select="false" placeholder="Select genres" label="name" track-by="name">
        </VueMultiselect>
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Code</Label>
        <Input v-model="comicStore.selectedData.code" placeholder="Code" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Original Language</Label>
        <Select v-model="comicStore.selectedData.original_language">
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="jp">Japanese</SelectItem>
              <SelectItem value="ko">Korean</SelectItem>
              <SelectItem value="zh-cn">Chinese(Mainland)</SelectItem>
              <SelectItem value="zh-tw">Chinese(Hong Kong & Taiwan)</SelectItem>
              <SelectItem value="en">English</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Language</Label>
        <Select v-model="comicStore.selectedData.lang">
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="all">All</SelectItem>
              <SelectItem value="en">English</SelectItem>
              <SelectItem value="zh">Chinese</SelectItem>
              <SelectItem value="vi">Vietnamese</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Audience</Label>
        <Select v-model="comicStore.selectedData.audience">
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="all">All</SelectItem>
              <SelectItem value="boys">Boys</SelectItem>
              <SelectItem value="girls">Girls</SelectItem>
              <SelectItem value="children">Children</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Status</Label>
        <Select v-model="comicStore.selectedData.status">
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="ongoing">Ongoing</SelectItem>
              <SelectItem value="completed">Completed</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <DialogFooter class="sm:justify-end">
        <Button variant="secondary" @click="comicStore.updateDialogIsOpen = false">
          Close
        </Button>
        <Button :disabled="isLoading" @click="
          async () => {
            if (checkForm()) {
              isLoading = true;
              comic.id = comicStore.selectedData.id;
              comic.name = comicStore.selectedData.name;
              comic.description = comicStore.selectedData.description;
              comic.code = comicStore.selectedData.code;
              comic.lang = comicStore.selectedData.lang;
              comic.original_language = comicStore.selectedData.original_language;
              comic.audience = comicStore.selectedData.audience;
              comic.status = comicStore.selectedData.status;
              comic.genre = comicStore.selectedData.genres;
              comic.author = comicStore.selectedData.authors;
              await comicStore.updateComic(comic);
              isLoading = false;
              comicStore.updateDialogIsOpen = false;
              comicStore.getComicData();
              resetComic();
            }
          }
        ">
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
<style src="vue-multiselect/dist/vue-multiselect.css"></style>
