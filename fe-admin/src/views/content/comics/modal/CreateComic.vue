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
import { useComicStore } from "../../../../stores/comicStore";
import VueMultiselect from "vue-multiselect";
import { useAuthorStore } from "../../../../stores/authorStore";
import loadingImg from "@/assets/loading.svg";
import { useGenreStore } from "../../../../stores/genreStore";
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const isLoading = ref(false);
const comicStore = useComicStore();
const authorStore = useAuthorStore();
const genreStore = useGenreStore();
genreStore.getGeneralGenreData();
authorStore.getGeneralAuthorData();
const comic = ref({
  title: "",
  author: [],
  genre: [],
  code: "",
  cover: null as File | null,
  description: "",
  active: true,
  language: "all",
  audience: "all",
  status: "ongoing",
});

const previewUrl = ref<string | null>(null);
const resetComic = () => {
  comic.value = {
    title: "",
    author: [],
    genre: [],
    code: "",
    cover: null,
    description: "",
    active: true,
    language: "",
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
  if (comic.value.title == "") {
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
    :open="comicStore.createDialogIsOpen"
    @update:open="(value:boolean)=>{comicStore.createDialogIsOpen = value;}"
  >
    <DialogContent class="overflow-y-scroll max-h-screen">
      <DialogHeader>
        <DialogTitle>Create Comic</DialogTitle>
      </DialogHeader>
      <img
        v-if="previewUrl"
        :src="previewUrl"
        class="max-w-[200px] justify-self-center h-auto"
      />
      <Input
        type="file"
        placeholder="Cover"
        class="justify-self-center w-[50%]"
        @change="handleFileChange"
        accept="image/*"
      />
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Name</Label>
        <Input v-model="comic.title" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Description</Label>
        <Input v-model="comic.description" placeholder="description" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Author</Label>
        <VueMultiselect
          v-model="comic.author"
          :options="authorStore.generalauthorData"
          :multiple="true"
          :close-on-select="false"
          placeholder="Select authors"
          label="name"
          track-by="name"
        >
        </VueMultiselect>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Genre</Label>
        <VueMultiselect
          v-model="comic.genre"
          :options="genreStore.generalGenreData"
          :multiple="true"
          :close-on-select="false"
          placeholder="Select authors"
          label="name"
          track-by="name"
        >
        </VueMultiselect>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Code</Label>
        <Input v-model="comic.code" placeholder="code" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active</Label>
        <Select v-model="comic.active">
          <SelectTrigger>
            <SelectValue :value="comic.active.toString()" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem :value="true">On</SelectItem>
              <SelectItem :value="false">Off</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Language</Label>
        <Select v-model="comic.language">
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="all">All</SelectItem>
              <SelectItem value="en">English</SelectItem>
              <SelectItem value="ch">Chinese</SelectItem>
              <SelectItem value="vi">Vietnamese</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Audience</Label>
        <Select v-model="comic.audience">
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
        <Select v-model="comic.status">
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
        <Button
          variant="secondary"
          @click="comicStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              if (checkForm()) {
                isLoading = true;
                await comicStore.createComic(comic);
                isLoading = false;
                comicStore.createDialogIsOpen = false;
                comicStore.getComicData();
                resetComic();
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
