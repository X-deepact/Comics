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

const comicStore = useComicStore();

const comic = ref({
  title: "",
  code: "",
  cover: "",
  description: "",
  active: true,
  language: "all",
  audience: "all",
});
const resetComic = () => {
  comic.value = {
    title: "",
    code: "",
    cover: "",
    description: "",
    active: true,
    language: "",
    audience: "",
  };
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
      <div class="min-h-[100px]"></div>
      <Input
        v-model="comic.cover"
        type="file"
        placeholder="cover"
        class="justify-self-center w-[50%]"
      />
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Title</Label>
        <Input v-model="comic.title" placeholder="title" />
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
