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
import { useAdStore } from "@/stores/adStore";
import loadingImg from "@/assets/loading.svg";
import { ref, watch } from "vue";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { toast } from "@/components/ui/toast/use-toast";

const adStore = useAdStore();
const isLoading = ref(false);
const previewImage = ref('');
const imageChanged = ref(false);

const formatDateForInput = (dateString: string) => {
  if (!dateString) return '';
  const date = new Date(dateString);
  // Format: YYYY-MM-DDThh:mm
  return date.toISOString().slice(0, 16);
};

// Watch for changes to selectedData to format dates when modal opens
watch(() => adStore.updateDialogIsOpen, (newValue) => {
  if (newValue && adStore.selectedData) {
    // Format dates when modal opens
    adStore.selectedData.active_from = formatDateForInput(adStore.selectedData.active_from);
    adStore.selectedData.active_to = formatDateForInput(adStore.selectedData.active_to);
    // Set preview image
    previewImage.value = adStore.selectedData.image;
    // Reset the imageChanged flag when opening the modal
    imageChanged.value = false;
  }
});

const handleFileUpload = async (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    const file = target.files[0];
    
    // Show preview
    const reader = new FileReader();
    reader.onload = (e) => {
      if (e.target?.result) {
        previewImage.value = e.target.result as string;
      }
    };
    reader.readAsDataURL(file);

    try {
      // Upload image and get filename
      const response = await adStore.uploadAdImage(file);
      // Store just the filename, not the entire response object
      adStore.selectedData.image = response.data || response;
      // Set the flag to indicate that the image has been changed
      imageChanged.value = true;
    } catch (error) {
      previewImage.value = '';
      adStore.selectedData.image = '';
      toast({
        description: "Failed to upload image",
        variant: "destructive",
      });
    }
  }
};

const validateForm = () => {
  const requiredFields = {
    title: "Title",
    image: "Image",
    type: "Type", 
    direct_url: "URL",
    status: "Status",
    active_from: "Start Date",
    active_to: "End Date"
  };

  for (const [field, label] of Object.entries(requiredFields)) {
    if (!adStore.selectedData[field as keyof typeof adStore.selectedData]) {
      toast({
        description: `${label} is required`,
        variant: "destructive",
      });
      return false;
    }
  }
  return true;
};

const handleSubmit = async () => {
  try {
    if (!validateForm()) return;
    
    isLoading.value = true;
    
    const formattedData = {
      ...adStore.selectedData,
      comic_id: adStore.selectedData.type === 'internal' ? parseInt(adStore.selectedData.comic_id.toString()) : null,
      active_from: new Date(adStore.selectedData.active_from).toISOString(),
      active_to: new Date(adStore.selectedData.active_to).toISOString(),
      // Only include image in the update if it was changed
      image: imageChanged.value ? adStore.selectedData.image : undefined
    };
    
    await adStore.updateAd(formattedData);
    adStore.updateDialogIsOpen = false;
    await adStore.getAdData();
    
    toast({
      description: "Ad updated successfully",
    });
  } catch (error: any) {
    toast({
      description: error.response?.data?.message || "Failed to update ad",
      variant: "destructive",
    });
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <Dialog
    :open="adStore.updateDialogIsOpen"
    @update:open="(value: boolean) => {
      adStore.updateDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Ad</DialogTitle>
      </DialogHeader>
      
      <div class="grid gap-4">
        <div class="flex items-center gap-4">
          <Label for="title" class="text-right w-1/4">Title *</Label>
          <Input v-model="adStore.selectedData.title" placeholder="Title" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="image" class="text-right w-1/4">Image *</Label>
          <div class="w-full">
            <Input type="file" @change="handleFileUpload" accept="image/*" />
            <div v-if="previewImage" class="mt-2">
              <img
                :src="previewImage"
                alt="Preview"
                class="max-w-full h-32 object-contain rounded-md border"
              />
            </div>
          </div>
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_from" class="text-right w-1/4">Start Date *</Label>
          <Input type="datetime-local" v-model="adStore.selectedData.active_from" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_to" class="text-right w-1/4">End Date *</Label>
          <Input type="datetime-local" v-model="adStore.selectedData.active_to" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="type" class="text-right w-1/4">Type *</Label>
          <Select v-model="adStore.selectedData.type">
            <SelectTrigger>
              <SelectValue placeholder="Select type" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="internal">Internal</SelectItem>
              <SelectItem value="external">External</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div class="flex items-center gap-4">
          <Label for="direct_url" class="text-right w-1/4">URL</Label>
          <Input v-model="adStore.selectedData.direct_url" placeholder="URL" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="comic_id" class="text-right w-1/4">Comic ID</Label>
          <Input type="number" v-model="adStore.selectedData.comic_id" placeholder="Comic ID" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="status" class="text-right w-1/4">Status *</Label>
          <Select v-model="adStore.selectedData.status">
            <SelectTrigger>
              <SelectValue placeholder="Select status" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="active">Active</SelectItem>
              <SelectItem value="inactive">Inactive</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="adStore.updateDialogIsOpen = false"
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