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
import { useChapterItemStore } from "../../../../../../stores/chapteritemStore";
import { useChapterStore } from "../../../../../../stores/chapterStore";
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
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});
const value = ref<DateValue>();
const chapterStore = useChapterStore();
const chapteritemStore = useChapterItemStore();
const isLoading = ref(false);
const previewUrl = ref<string | null>(null);
const chapteritem = ref({
  chapter_id: 0,
  page: 1,
  image: null as string | null,
  active_from: null as string | null,
  active: true,
});
const resetChapterItem = () => {
  chapteritem.value = {
    chapter_id: 0,
    page: 1,
    image: null as string | null,
    active_from: null as string | null,
    active: true,
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
const setActiveFrom = () => {
  if (value.value) {
    const date = value.value.toDate(getLocalTimeZone());
    chapteritem.value.active_from = date.toISOString();
  }
};
const checkForm = () => {
  if (!value.value || !image.value) {
    toast({
      description: "Fill in all data.",
      variant: "destructive",
    });
    return false;
  }
  return true;
}
</script>
<template>
  <Dialog :open="chapteritemStore.createDialogIsOpen"
    @update:open="(value: boolean) => { chapteritemStore.createDialogIsOpen = value; }">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Chapter Item</DialogTitle>
      </DialogHeader>

      <img v-if="previewUrl" :src="previewUrl" class="max-w-[200px] justify-self-center h-auto" />
      <Input type="file" placeholder="Image" class="justify-self-center w-[50%]" @change="handleFileChange"
        accept="image/*" />

      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Page</Label>
        <Input v-model="chapteritem.page" placeholder="Page" type="number" />
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
        <Select v-model="chapteritem.active">
          <SelectTrigger>
            <SelectValue :value="chapteritem.active" />
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
        <Button variant="secondary" @click="chapteritemStore.createDialogIsOpen = false">
          Close
        </Button>
        <Button :disabled="isLoading" @click="
          async () => {
            if (checkForm()) {
              isLoading = true;
              chapteritem.image = await chapteritemStore.uploadImage(image);
              chapteritem.chapter_id = chapterStore.selectedData.id;
              setActiveFrom();
              await chapteritemStore.createChapterItem(chapteritem);
              chapteritemStore.getChapterItemsData();
              isLoading = false;
              chapteritemStore.createDialogIsOpen = false;
              resetChapterItem();
            }
          }
        ">
          <img v-if="isLoading" :src="loadingImg" size="icon" />Add</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
