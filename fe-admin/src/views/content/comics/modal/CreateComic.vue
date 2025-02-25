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
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { Input } from "@/components/ui/input";
import { ref, defineEmits } from "vue";
import { useComicStore } from "../../../../stores/comicStore";
import MultiSelect from "primevue/multiselect";
import { Author, useAuthorStore } from "../../../../stores/authorStore";

const comicStore = useComicStore();
const authorStore = useAuthorStore();
authorStore.getGeneralAuthorData();
const comic = ref({
  title: "",
  author: <Author>[],
  subject: [],
  code: "",
  cover: null as File | null,
  description: "",
  active: true,
  language: "all",
  audience: "all",
});

const previewUrl = ref<string | null>(null);
const resetComic = () => {
  comic.value = {
    title: "",
    author: [],
    subject: [],
    code: "",
    cover: null,
    description: "",
    active: true,
    language: "",
    audience: "",
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
</script>
<template>
  <Dialog
    :open="comicStore.createDialogIsOpen"
    @update:open="(value:boolean)=>{comicStore.createDialogIsOpen = value;}"
  >
    <DialogContent>
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
        placeholder="cover"
        class="justify-self-center w-[50%]"
        @change="handleFileChange"
        accept="image/*"
      />
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Title</Label>
        <Input v-model="comic.title" placeholder="Title" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="author" class="text-center w-1/4">Author</Label>
        <Input v-model="comic.author" placeholder="Author" />
        <MultiSelect
          v-model="comicStore.selectedAuthorForCreate"
          :options="authorStore.authorData"
          optionLabel="name"
          filter
          placeholder="Select Authors"
        />
      </div>
      <div class="flex items-center gap-4">
        <Label for="subject" class="text-center w-1/4">Subject</Label>
        <Input v-model="comic.subject" placeholder="Subject" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Description</Label>
        <Input v-model="comic.description" placeholder="description" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Code</Label>
        <Input v-model="comic.code" placeholder="code" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Active</Label>
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
        <Label for="name" class="text-center w-1/4">Language</Label>
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
        <Label for="name" class="text-center w-1/4">Audience</Label>
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
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="comicStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await comicStore.createComic(comic);
              comicStore.createDialogIsOpen = false;
              comicStore.getComicData();
              resetComic();
            }
          "
          >Add</Button
        >
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
