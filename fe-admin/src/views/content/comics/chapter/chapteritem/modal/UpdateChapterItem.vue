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
import { useChapterItemStore } from "../../../../../../stores/chapteritemStore";
import { cn } from "@/lib/utils";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { DateFormatter, getLocalTimeZone } from "@internationalized/date";
import { CalendarIcon } from "lucide-vue-next";
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});

const isLoading = ref(false);
const chapteritemStore = useChapterItemStore();
const previewUrl = ref<string | null>(null);
const chapteritem = ref({
  id: 0,
  page: 1,
  image: null as string | null,
  active_from: null as string | null,
  active: true,
});
const resetChapterItem = () => {
  chapteritem.value = {
    id: 0,
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
</script>
<template>
  <Dialog
    :open="chapteritemStore.updateDialogIsOpen"
    @update:open="
    (value: boolean) => {
      chapteritemStore.updateDialogIsOpen = value;
    }
  "
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Chapter Item</DialogTitle>
      </DialogHeader>
      <img
        v-if="!previewUrl"
        :src="chapteritemStore.selectedData.image"
        class="max-w-[200px] justify-self-center h-auto"
      />
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
        <Label class="text-center w-1/4">Page</Label>
        <Input v-model="chapteritemStore.selectedData.page" type="number" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active From</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              :class="
                cn(
                  'w-[280px] justify-start text-left font-normal',
                  !chapteritemStore.selectedData.active_from &&
                    'text-muted-foreground'
                )
              "
            >
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{
                chapteritemStore.selectedData.active_from
                  ? df.format(
                      chapteritemStore.selectedData.active_from.toDate(
                        getLocalTimeZone()
                      )
                    )
                  : "Pick a date"
              }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar
              :v-model="chapteritemStore.selectedData.active_from"
              initial-focus
            />
          </PopoverContent>
        </Popover>
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active</Label>
        <Select :v-model="chapteritem.active">
          <SelectTrigger>
            <SelectValue :value="chapteritem.active" />
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
          @click="chapteritemStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              chapteritem.image = await chapteritemStore.uploadImage(image);
              chapteritem.id = chapteritemStore.selectedData.id;
              chapteritem.page = chapteritemStore.selectedData.page;
              const data = chapteritemStore.selectedData.active_from
                ?.toDate(getLocalTimeZone())
                .toISOString();
              chapteritem.active_from = data ? data : null;
              await chapteritemStore.updateChapterItem(chapteritem);
              chapteritemStore.getChapterItemsData();
              isLoading = false;
              chapteritemStore.updateDialogIsOpen = false;
              resetChapterItem();
            }
          "
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
