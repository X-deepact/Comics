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
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useRecommendStore } from "@/stores/recommendStore";
import loadingImg from "@/assets/loading.svg";
import { ref, watch } from "vue";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { CalendarIcon } from "lucide-vue-next";
import { useToast } from "@/components/ui/toast/use-toast";

const { toast } = useToast();
const isLoading = ref(false);
const recommendStore = useRecommendStore();

const fromDate = ref<Date | null>(null);
const toDate = ref<Date | null>(null);

const previewImage = ref<string>("");
const fileInput = ref<HTMLInputElement | null>(null);

const validateForm = () => {
  if (!recommendStore.selectedData.title?.trim()) {
    toast({
      description: "Title is required",
      variant: "destructive",
    });
    return false;
  }
  if (!recommendStore.selectedData.position) {
    toast({
      description: "Position is required",
      variant: "destructive",
    });
    return false;
  }
  if (!fromDate.value || !toDate.value) {
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
    const submitData = {
      ...recommendStore.selectedData,
      position: Number(recommendStore.selectedData.position),
      active_from: fromDate.value?.toISOString().split('T')[0],
      active_to: toDate.value?.toISOString().split('T')[0],
    };
    
    console.log('Submitting update data:', submitData);
    await recommendStore.updateRecommend(submitData);
    recommendStore.updateDialogIsOpen = false;
    await recommendStore.getRecommendData();
  } catch (error) {
    console.error('Error updating recommend:', error);
  } finally {
    isLoading.value = false;
  }
};

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
        recommendStore.selectedData.cover = e.target.result as string;
      }
    };
    reader.readAsDataURL(file);
  }
};

const triggerFileInput = () => {
  fileInput.value?.click();
};

// Initialize dates when modal opens
watch(() => recommendStore.updateDialogIsOpen, (isOpen) => {
  if (isOpen && recommendStore.selectedData) {
    console.log('Selected data:', recommendStore.selectedData);
    // Convert Unix timestamp to Date object if exists
    fromDate.value = recommendStore.selectedData.active_from ? 
      new Date(recommendStore.selectedData.active_from * 1000) : null;
    toDate.value = recommendStore.selectedData.active_to ? 
      new Date(recommendStore.selectedData.active_to * 1000) : null;
  } else {
    fromDate.value = null;
    toDate.value = null;
  }
});
</script>

<template>
  <Dialog
    :open="recommendStore.updateDialogIsOpen"
    @update:open="(value: boolean) => {
      recommendStore.updateDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Recommendation</DialogTitle>
        <DialogDescription>
          Update the recommendation details below.
        </DialogDescription>
      </DialogHeader>
      
      <div class="flex items-center gap-4">
        <Label for="title" class="text-center w-1/4">Title</Label>
        <Input
          id="title"
          v-model="recommendStore.selectedData.title"
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
          <div v-if="previewImage || recommendStore.selectedData.cover" class="mt-2">
            <img
              :src="previewImage || recommendStore.selectedData.cover"
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
          v-model="recommendStore.selectedData.position"
          placeholder="Position"
        />
      </div>

      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active From</Label>
        <Popover>
          <PopoverTrigger asChild>
            <Button
              variant="outline"
              class="w-[280px] justify-start text-left font-normal"
            >
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{ fromDate ? fromDate.toLocaleDateString() : "Select date" }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0" align="start">
            <Calendar
              v-model="fromDate"
              mode="single"
            />
          </PopoverContent>
        </Popover>
      </div>

      <div class="flex items-center gap-4">
        <Label class="text-center w-1/4">Active To</Label>
        <Popover>
          <PopoverTrigger asChild>
            <Button
              variant="outline"
              class="w-[280px] justify-start text-left font-normal"
            >
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{ toDate ? toDate.toLocaleDateString() : "Select date" }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0" align="start">
            <Calendar
              v-model="toDate"
              mode="single"
            />
          </PopoverContent>
        </Popover>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="recommendStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="handleSubmit"
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template> 