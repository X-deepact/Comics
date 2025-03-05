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
import loadingImg from "@/assets/loading.svg";
import { useChapterStore } from "../../../../../stores/chapterStore";
import { useComicStore } from "../../../../../stores/comicStore";
const comicStore = useComicStore();
const chapterStore = useChapterStore();
const isLoading = ref(false);
const chapter = ref({
  comic_id: 0,
  number: 1,
  name: "",
  cover: true,
  active: true,
});
const resetChapter = () => {
  chapter.value = {
    comic_id: 0,
    number: 1,
    name: "",
    cover: true,
    active: true,
  };
};
</script>
<template>
  <Dialog
    :open="chapterStore.createDialogIsOpen"
    @update:open="(value: boolean) => { chapterStore.createDialogIsOpen = value; }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Chapter</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Name</Label>
        <Input v-model="chapter.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Number</Label>
        <Input v-model="chapter.number" placeholder="Number" type="number" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active</Label>
        <Select v-model="chapter.active">
          <SelectTrigger>
            <SelectValue :value="chapter.active.toString()" />
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
        <Label class="text-center w-1/4">Cover</Label>
        <Select v-model="chapter.cover">
          <SelectTrigger>
            <SelectValue :value="chapter.cover.toString()" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem :value="true">On</SelectItem>
              <SelectItem :value="false">Off</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="chapterStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              chapter.comic_id = comicStore.selectedData.id;
              await chapterStore.createChapter(chapter);
              isLoading = false;
              chapterStore.createDialogIsOpen = false;
              chapterStore.getChapterData();
              resetChapter();
            }
          "
          ><img v-if="isLoading" :src="loadingImg" size="icon" />Add</Button
        >
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
