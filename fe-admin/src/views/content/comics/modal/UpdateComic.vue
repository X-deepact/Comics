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
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useComicStore } from "../../../../stores/comicStore";

const comicStore = useComicStore();
</script>
<template>
  <Dialog
    :open="comicStore.updateDialogIsOpen"
    @update:open="
      (value : boolean) => {
        comicStore.updateDialogIsOpen = value;
      }
    "
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Comic</DialogTitle>
      </DialogHeader>
      <div class="min-h-[100px]"></div>
      <Input
        type="file"
        placeholder="cover"
        class="justify-self-center w-[50%]"
        @change="
          (event) => (comicStore.selectedData.cover = event.target.files[0])
        "
      />
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Title</Label>
        <Input v-model="comicStore.selectedData.title" placeholder="title" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Description</Label>
        <Input
          v-model="comicStore.selectedData.description"
          placeholder="description"
        />
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Code</Label>
        <Input v-model="comicStore.selectedData.code" placeholder="code" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Active</Label>
        <Select
          :modelValue="comicStore.selectedData.active"
          @update:modelValue="(val) => (comicStore.selectedData.active = val)"
        >
          <SelectTrigger>
            <SelectValue :value="comicStore.selectedData.active.toString()" />
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
        <Select v-model="comicStore.selectedData.language">
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
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="comicStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await comicStore.updateComic({
                id: comicStore.selectedData.id,
                title: comicStore.selectedData.title,
                code: comicStore.selectedData.code,
                cover: comicStore.selectedData.cover,
                description: comicStore.selectedData.description,
                active: comicStore.selectedData.active,
                language: comicStore.selectedData.language,
                audience: comicStore.selectedData.audience,
              });
              comicStore.updateDialogIsOpen = false;
              comicStore.getComicData();
            }
          "
        >
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
