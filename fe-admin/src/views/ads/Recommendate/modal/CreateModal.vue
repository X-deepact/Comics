<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { ref, watch } from "vue";
import { useRecommendStore } from "@/stores/recommendStore";
import loadingImg from "@/assets/loading.svg";
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
import { cn } from "@/lib/utils";
import { useToast } from "@/components/ui/toast/use-toast";

const df = new DateFormatter("en-US", { dateStyle: "long" });
const activeFromValue = ref<DateValue>();
const activeToValue = ref<DateValue>();

const recommendStore = useRecommendStore();
const isLoading = ref(false);

const { toast } = useToast();

const formData = ref({
  title: "",  
  cover: "",
  position: 0,
  active_from: "",
  active_to: "",
});

const fileInput = ref<HTMLInputElement | null>(null);
const previewImage = ref<string>("");

const handleFileUpload = (event: Event) => {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    const file = input.files[0];
    
    // Check file type
    if (!file.type.startsWith('image/')) {
      toast({
        description: "Please upload an image file",
        variant: "destructive",
      });
      return;
    }

    // Check file size (e.g., max 5MB)
    if (file.size > 5 * 1024 * 1024) {
      toast({
        description: "File size should be less than 5MB",
        variant: "destructive",
      });
      return;
    }

    // Create preview
    const reader = new FileReader();
    reader.onload = (e) => {
      if (e.target?.result) {
        previewImage.value = e.target.result as string;
        formData.value.cover = e.target.result as string; // Store the base64 image data
      }
    };
    reader.readAsDataURL(file);
  }
};

const triggerFileInput = () => {
  fileInput.value?.click();
};

const validateForm = () => {
  if (!formData.value.title?.trim()) {
    toast({
      description: "Title is required",
      variant: "destructive",
    });
    return false;
  }
  if (!formData.value.cover) {
    toast({
      description: "Cover image is required",
      variant: "destructive",
    });
    return false;
  }
  if (!formData.value.active_from || !formData.value.active_to) {
    toast({
      description: "Active dates are required",
      variant: "destructive",
    });
    return false;
  }
  return true;
};

const handleSubmit = async () => {
  if (!validateForm()) return;
  
  isLoading.value = true;
  try {
    await recommendStore.createRecommend(formData.value);
    recommendStore.createDialogIsOpen = false;
    await recommendStore.getRecommendData();
    resetForm();
  } catch (error) {
    console.error('Error creating recommend:', error);
  } finally {
    isLoading.value = false;
  }
};

const resetForm = () => {
  formData.value = {
    title: "",
    cover: "",
    position: 0,
    active_from: "",
    active_to: "",
  };
  previewImage.value = "";
  if (fileInput.value) {
    fileInput.value.value = "";
  }
  activeFromValue.value = undefined;
  activeToValue.value = undefined;
};

const setActiveFrom = () => {
  if (activeFromValue.value) {
    const date = activeFromValue.value.toDate(getLocalTimeZone());
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    activeFromValue.value = undefined;
    formData.value.active_from = `${year}-${month}-${day}`;
  }
};

const setActiveTo = () => {
  if (activeToValue.value) {
    const date = activeToValue.value.toDate(getLocalTimeZone());
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    activeToValue.value = undefined;
    formData.value.active_to = `${year}-${month}-${day}`;
  }
};

const formatDate = (value: DateValue | undefined): string => {
  if (!value) return "";
  const date = value.toDate(getLocalTimeZone());
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  return `${year}-${month}-${day}`;
};

watch([activeFromValue, activeToValue], ([newFromValue, newToValue]) => {
  if (newFromValue) {
    formData.value.active_from = formatDate(newFromValue);
  }
  if (newToValue) {
    formData.value.active_to = formatDate(newToValue);
  }
});
</script>

<template>
  <Dialog
    :open="recommendStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
      recommendStore.createDialogIsOpen = value;
      if (!value) resetForm();
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Recommendation</DialogTitle>
        <DialogDescription>
          Add a new recommendation below.
        </DialogDescription>
      </DialogHeader>
      
      <div class="flex items-center gap-4">
        <Label for="title" class="text-center w-1/4">Title</Label>
        <Input
          id="title"
          v-model="formData.title"
          placeholder="Title"
        />
      </div>

      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Cover Image</Label>
        <div class="flex flex-col gap-2 w-3/4">
          <input
            ref="fileInput"
            type="file"
            accept="image/*"
            class="hidden"
            @change="handleFileUpload"
          />
          <Button
            type="button"
            variant="outline"
            @click="triggerFileInput"
          >
            Choose Image
          </Button>
          <div v-if="previewImage" class="mt-2">
            <img
              :src="previewImage"
              alt="Preview"
              class="max-w-[200px] max-h-[200px] object-contain"
            />
          </div>
        </div>
      </div>

      <div class="flex items-center gap-4">
        <Label for="position" class="text-center w-1/4">Position</Label>
        <Input
          id="position"
          type="number"
          v-model="formData.position"
          placeholder="Position"
        />
      </div>

      <div class="flex items-center gap-4">
        <Label for="active_from" class="text-center w-1/4">Active From</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              :class="
                cn(
                  'w-[280px] justify-start text-left font-normal',
                  !activeFromValue && 'text-muted-foreground'
                )
              "
            >
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{
                activeFromValue
                  ? df.format(activeFromValue.toDate(getLocalTimeZone()))
                  : "Pick a date"
              }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar v-model="activeFromValue" initial-focus />
          </PopoverContent>
        </Popover>
      </div>

      <div class="flex items-center gap-4">
        <Label for="active_to" class="text-center w-1/4">Active To</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              :class="
                cn(
                  'w-[280px] justify-start text-left font-normal',
                  !activeToValue && 'text-muted-foreground'
                )
              "
            >
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{
                activeToValue
                  ? df.format(activeToValue.toDate(getLocalTimeZone()))
                  : "Pick a date"
              }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar v-model="activeToValue" initial-focus />
          </PopoverContent>
        </Popover>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="recommendStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="handleSubmit"
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Add
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template> 