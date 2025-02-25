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
import { Input } from "@/components/ui/input";
import { ref } from "vue";
import { useAuthorStore } from "../../../../stores/authorStore";
import loadingImg from "@/assets/loading.svg";
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

const authorStore = useAuthorStore();
const isLoading = ref(false);
authorStore.getGeneralAuthorData();

interface Author {
  name: string;
  biography: string;
  birth_day: string;
}

const author = ref<Author>({
  name: "",
  biography: "",
  birth_day: "",
});
const resetAuthor = () => {
  author.value = {
    name: "",
    biography: "",
    birth_day: "",
  };
};
const setBirthday = () => {
  if (value.value) {
    const date = value.value.toDate(getLocalTimeZone());
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    value.value = undefined;
    author.value.birth_day = `${year}-${month}-${day}`;
  }
};
</script>
<template>
  <Dialog
    :open="authorStore.createDialogIsOpen"
    @update:open="(value:boolean)=>{authorStore.createDialogIsOpen = value;}"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Author</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="name" class="text-center w-1/4">Name</Label>
        <Input v-model="author.name" placeholder="Name" />
      </div>
      <div class="flex items-start gap-4">
        <Label for="subject" class="text-center w-1/4 py-3">Biography</Label>
        <Textarea v-model="author.biography" placeholder="Biography" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="birthday" class="text-center w-1/4">Birthday</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              :class="
                cn(
                  'w-[280px] justify-start text-left font-normal',
                  !value && 'text-muted-foreground'
                )
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
          @click="authorStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              setBirthday();
              await authorStore.createAuthor(author);
              isLoading = false;
              authorStore.createDialogIsOpen = false;
              authorStore.getAuthorData();
              resetAuthor();
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
