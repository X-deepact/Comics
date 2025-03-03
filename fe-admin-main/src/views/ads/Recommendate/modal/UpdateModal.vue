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

const { toast } = useToast();
const isLoading = ref(false);
const recommendStore = useRecommendStore();
const previewImage = ref<string>("");
const fileInput = ref<HTMLInputElement | null>(null);

const formData = ref({
  id: 0,
  title: "",
  cover: "",
  position: 0,
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
      position: newData.position || 0,
      active_from: "",  // We'll handle dates separately
      active_to: "",    // We'll handle dates separately
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
  if (!formData.value.title?.trim()) {
    toast({
      description: "Title is required",
      variant: "destructive",
    });
    return;
  }

  const activeFromDate = document.getElementById('active_from') as HTMLInputElement;
  const activeToDate = document.getElementById('active_to') as HTMLInputElement;

  if (!activeFromDate.value || !activeToDate.value) {
    toast({
      description: "Active dates are required",
      variant: "destructive",
    });
    return;
  }

  isLoading.value = true;
  try {
    const submitData = {
      ...formData.value,
      position: Number(formData.value.position),
      active_from: activeFromDate.value,
      active_to: activeToDate.value,
    };

    await recommendStore.updateRecommend(submitData);
    recommendStore.updateDialogIsOpen = false;
    await recommendStore.getRecommendData();
  } catch (error: any) {
    console.error('Update error:', error);
    toast({
      description: error.message || "Failed to update recommendation",
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
              <Label class="text-right">Title</Label>
              <Input v-model="formData.title" class="col-span-3" placeholder="Enter title" />
            </div>
            <div class="space-y-4">
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Cover Image</Label>
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
              <Label class="text-right">Position</Label>
              <Input 
                type="number" 
                v-model="formData.position" 
                class="col-span-3"
                placeholder="Enter position number"
              />
            </div>

            
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Active From</Label>
              <Input 
                id="active_from"
                type="date" 
                class="col-span-3"
                required
              />
            </div>

            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Active To</Label>
              <Input 
                id="active_to"
                type="date"
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