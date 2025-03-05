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
import { useToast } from "@/components/ui/toast/use-toast";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { RecommendPosition, positionLabels } from "@/stores/recommendStore";

const { toast } = useToast();
const isLoading = ref(false);
const recommendStore = useRecommendStore();
const previewImage = ref<string>("");
const fileInput = ref<HTMLInputElement | null>(null);

const formatDateForInput = (dateString: string) => {
  if (!dateString) return '';
  const date = new Date(dateString);
  return date.toISOString().split('T')[0]; // Format: YYYY-MM-DD
};

const formData = ref({
  id: 0,
  title: "",
  cover: "",
  position: RecommendPosition.COMPLETE_MASTERPIECE,
  active_from: "",
  active_to: "",
});

// Initialize form data when selectedData changes
watch(() => recommendStore.selectedData, (newData) => {
  if (newData) {
    formData.value = {
      id: newData.id,
      title: newData.title || "",
      cover: newData.cover || "",
      position: newData.position || RecommendPosition.COMPLETE_MASTERPIECE,
      active_from: formatDateForInput(newData.active_from),
      active_to: formatDateForInput(newData.active_to),
    };
  }
}, { immediate: true });

const handleFileUpload = (event: Event) => {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    const file = input.files[0];
    const reader = new FileReader();
    reader.onload = (e) => {
      if (e.target?.result) {
        previewImage.value = e.target.result as string;
        formData.value.cover = e.target.result as string;
      }
    };
    reader.readAsDataURL(file);
  }
};

const handleSubmit = async () => {
  // Validate form data

  if (!formData.value.title?.trim()) {
    toast({
      description: "Title is required",
      variant: "destructive",
    });
    return;
  }

  if (!formData.value.active_from || !formData.value.active_to) {
    toast({
      description: "Active dates are required",
      variant: "destructive",
    });
    return;
  }

  // Validate dates
  const fromDate = new Date(formData.value.active_from);
  const toDate = new Date(formData.value.active_to);
  
  if (isNaN(fromDate.getTime()) || isNaN(toDate.getTime())) {
    toast({
      description: "Invalid date format",
      variant: "destructive",
    });
    return;
  }

  if (fromDate > toDate) {
    toast({
      description: "Active from date must be before active to date",
      variant: "destructive",
    });
    return;
  }

  isLoading.value = true;
  try {
    const submitData = {
      ...formData.value,
      id: Number(formData.value.id),
      position: Number(formData.value.position),
      active_from: formData.value.active_from,
      active_to: formData.value.active_to,
    };

    await recommendStore.updateRecommend(submitData);
    console.log('Update successful');
    
    recommendStore.updateDialogIsOpen = false;
    await recommendStore.getRecommendData();
  } catch (error: any) {
    console.error('Update error:', {
      message: error.message,
      response: error.response?.data,
      status: error.response?.status,
      data: error.config?.data,
    });
    
    toast({
      description: error.response?.data?.message || "Failed to update recommendation",
      variant: "destructive",
    });
  } finally {
    isLoading.value = false;
  }
};
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
      
      <div class="grid gap-4">
        <div class=" gap-4">
          <div class="space-y-4">
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Title *</Label>
              <Input v-model="formData.title" class="col-span-3" placeholder="Enter title" />
            </div>
            <div class="space-y-4">
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Image</Label>
              <div class="col-span-3">
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
                  class="w-full"
                  @click="fileInput?.click()"
                >
                  Choose Image
                </Button>
                <div v-if="previewImage || formData.cover" class="mt-2">
                  <img
                    :src="previewImage || formData.cover"
                    alt="Preview"
                    class="max-w-full max-h-[200px] object-contain rounded-md border"
                  />
                </div>
              </div>
            </div>
          </div>
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Position *</Label>
              <div class="col-span-3">
              <Select v-model="formData.position" >
                <SelectTrigger>
                  <SelectValue :placeholder="positionLabels[formData.position]" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem :value="RecommendPosition.COMPLETE_MASTERPIECE">
                    {{ positionLabels[RecommendPosition.COMPLETE_MASTERPIECE] }}
                  </SelectItem>
                  <SelectItem :value="RecommendPosition.FASTEST_RISING">
                    {{ positionLabels[RecommendPosition.FASTEST_RISING] }}
                  </SelectItem>
                  <SelectItem :value="RecommendPosition.NEW_PUBLISHING">
                    {{ positionLabels[RecommendPosition.NEW_PUBLISHING] }}
                  </SelectItem>
                  <SelectItem :value="RecommendPosition.RECENTLY_UPDATE">
                    {{ positionLabels[RecommendPosition.RECENTLY_UPDATE] }}
                  </SelectItem>
                </SelectContent>
              </Select>
              </div>
            </div>

            
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Active From *</Label>
              <Input 
                type="date"
                v-model="formData.active_from"
                class="col-span-3"
                required
              />
            </div>

            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Active To *</Label>
              <Input 
                type="date"
                v-model="formData.active_to"
                class="col-span-3" 
                required
              />
            </div>
          </div>

          
        </div>
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
          <img v-if="isLoading" :src="loadingImg" alt="loading" class="mr-2 h-4 w-4" />
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template> 