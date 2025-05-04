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
import { DateFormatter, getLocalTimeZone } from "@internationalized/date";
import type { CalendarDate, DateValue } from "@internationalized/date";
import { CalendarIcon } from "lucide-vue-next";
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});
const isLoading = ref(false);
const authorStore = useAuthorStore();

const checkForm = () => {
  if (
    authorStore.selectedData.name === "" ||
    authorStore.selectedData.biography === "" ||
    authorStore.selectedData.birth_day?.toString() === ""
  ) {
    toast({
      description: "Fill all fields.",
      variant: "destructive",
    });
    return false;
  }
  return true;
};

const getBirthDay = (): DateValue | undefined => {
  return authorStore.selectedData.birth_day as DateValue || undefined;
};

const setBirthDay = (date: DateValue | undefined) => {
  authorStore.selectedData.birth_day = date as CalendarDate;
};
</script>
<template>
  <Dialog
    :open="authorStore.updateDialogIsOpen"
    @update:open="
    (value: boolean) => {
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
                cn(
                  'font-normal w-3/4',
                  !authorStore.selectedData.birth_day && 'text-muted-foreground'
                )
              "
            >
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{
                authorStore.selectedData.birth_day
                  ? df.format(
                      authorStore.selectedData.birth_day.toDate(
                        getLocalTimeZone()
                      )
                    )
                  : "Pick a date"
              }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar
              :model-value="getBirthDay()"
              @update:model-value="setBirthDay"
              initial-focus
            />
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
              if (checkForm()) {
                isLoading = true;
                await authorStore.updateAuthor({
                  id: authorStore.selectedData.id,
                  name: authorStore.selectedData.name,
                  birth_day: authorStore.selectedData.birth_day?.toString(),
                  biography: authorStore.selectedData.biography,
                });
                isLoading = false;
                authorStore.updateDialogIsOpen = false;
                authorStore.getAuthorData();
              }
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
