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
import { useAuthorStore } from "../../../../stores/authorStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
import { Textarea } from "@/components/ui/textarea";
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
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});
const value = ref<DateValue>();
const isLoading = ref(false);
const authorStore = useAuthorStore();
const getBirthday = () => {
  if (value.value) {
    const date = value.value.toDate(getLocalTimeZone());
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    value.value = undefined;
    return `${year}-${month}-${day}`;
  }
};
</script>
<template>
  <Dialog
    :open="authorStore.updateDialogIsOpen"
    @update:open="
      (value : boolean) => {
        authorStore.updateDialogIsOpen = value;
      }
    "
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Author</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Name</Label>
        <Input v-model="authorStore.selectedData.name" placeholder="Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="biography" class="text-center w-1/4">Biography</Label>
        <Textarea
          v-model="authorStore.selectedData.biography"
          placeholder="Biography"
        />
      </div>
      <div class="flex items-center gap-4">
        <Label for="birth_day" class="text-center w-1/4">Birthday</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              :class="
                cn('font-normal w-3/4', !value && 'text-muted-foreground')
              "
            >
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
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="authorStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await authorStore.updateAuthor({
                id: authorStore.selectedData.id,
                name: authorStore.selectedData.name,
                birth_day: getBirthday(),
                biography: authorStore.selectedData.biography,
              });
              isLoading = false;
              authorStore.updateDialogIsOpen = false;
              authorStore.getAuthorData();
            }
          "
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
