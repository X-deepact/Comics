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
import { useToast } from "@/components/ui/toast/use-toast";
import { cn } from "@/lib/utils";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  DateFormatter,
  type DateValue,
  getLocalTimeZone,
} from "@internationalized/date";
import { CalendarIcon } from "lucide-vue-next";

const { toast } = useToast();
const comicStore = useComicStore();
const chapterStore = useChapterStore();
const isLoading = ref(false);
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});
const value = ref<DateValue>();

const chapter = ref({
  comic_id: 0,
  number: 1,
  name: "",
  cover: true,
  active: true,
  active_from: null as string | null,
});

const resetChapter = () => {
  chapter.value = {
    comic_id: 0,
    number: 1,
    name: "",
    cover: true,
    active: true,
    active_from: null,
  };
  value.value = undefined;
};

const setActiveFrom = () => {
  if (value.value) {
    const date = value.value.toDate(getLocalTimeZone());
    chapter.value.active_from = date.toISOString();
  }
};

const checkForm = () => {
  if (!chapter.value.name) {
    toast({
      description: "Name is required",
      variant: "destructive",
    });
    return false;
  }
  
  if (!value.value) {
    toast({
      description: "Active From date is required",
      variant: "destructive",
    });
    return false;
  }
  
  return true;
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
        <Input
          v-model="chapter.number"
          placeholder="Number"
          type="number"
          :min="1"
        />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active From</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button variant="outline" :class="cn(
              'w-[280px] justify-start text-left font-normal',
              !value && 'text-muted-foreground'
            )
              ">
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{
                value
                  ? df.format(value.toDate(getLocalTimeZone()))
                  : "Pick a date"
              }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar v-model="value" initial-focus />
          </PopoverContent>
        </Popover>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active</Label>
        <Select :v-model="chapter.active">
          <SelectTrigger>
            <SelectValue :value="chapter.active.toString()" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="true">On</SelectItem>
              <SelectItem value="false">Off</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Cover</Label>
        <Select :v-model="chapter.cover">
          <SelectTrigger>
            <SelectValue :value="chapter.cover.toString()" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value="true">On</SelectItem>
              <SelectItem value="false">Off</SelectItem>
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
              if (checkForm()) {
                isLoading = true;
                chapter.comic_id = comicStore.selectedData.id;
                setActiveFrom();
                await chapterStore.createChapter(chapter);
                isLoading = false;
                chapterStore.createDialogIsOpen = false;
                chapterStore.getChapterData();
                resetChapter();
              }
            }
          "
          ><img v-if="isLoading" :src="loadingImg" size="icon" />Add</Button
        >
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
