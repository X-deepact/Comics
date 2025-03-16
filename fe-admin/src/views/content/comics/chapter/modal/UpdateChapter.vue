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
  getLocalTimeZone,
} from "@internationalized/date";
import { CalendarIcon } from "lucide-vue-next";
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});

const { toast } = useToast();
const comicStore = useComicStore();
const isLoading = ref(false);
const chapterStore = useChapterStore();
const checkForm = () => {
  if (chapterStore.selectedData.name) return true;
  toast({
    description: "Name is required",
    variant: "destructive",
  });
  return false;
};
</script>
<template>
  <Dialog
    :open="chapterStore.updateDialogIsOpen"
    @update:open="
        (value: boolean) => {
            chapterStore.updateDialogIsOpen = value;
        }
    "
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Chapter</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Name</Label>
        <Input v-model="chapterStore.selectedData.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Number</Label>
        <Input
          v-model="chapterStore.selectedData.number"
          placeholder="Number"
          type="number"
        />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active From</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button variant="outline" :class="cn(
              'w-[280px] justify-start text-left font-normal',
              !chapterStore.selectedData.active_from && 'text-muted-foreground'
            )
              ">
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{
                chapterStore.selectedData.active_from
                  ? df.format(chapterStore.selectedData.active_from.toDate(getLocalTimeZone()))
                  : "Pick a date"
              }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar v-model="chapterStore.selectedData.active_from" initial-focus />
          </PopoverContent>
        </Popover>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active</Label>
        <Select v-model="chapterStore.selectedData.active">
          <SelectTrigger>
            <SelectValue :value="chapterStore.selectedData.active.toString()" />
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
        <Select v-model="chapterStore.selectedData.cover">
          <SelectTrigger>
            <SelectValue :value="chapterStore.selectedData.cover.toString()" />
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
          @click="chapterStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              if (checkForm()) {
                isLoading = true;
                await chapterStore.updateChapter({
                  id: chapterStore.selectedData.id,
                  name: chapterStore.selectedData.name,
                  number: chapterStore.selectedData.number,
                  active_from: chapterStore.selectedData.active_from?.toDate(getLocalTimeZone()).toISOString(),
                  active: chapterStore.selectedData.active,
                  cover: chapterStore.selectedData.cover,
                  comic_id: comicStore.selectedData.id,
                });
                isLoading = false;
                chapterStore.updateDialogIsOpen = false;
                chapterStore.getChapterData();
              }
            }
          "
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
