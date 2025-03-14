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
import { useAdStore } from "@/stores/adStore";
import loadingImg from "@/assets/loading.svg";
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

const ad = ref({
  title: "",
  image: "",
  active_from: "",
  active_to: "",
  type: "internal" as 'internal' | 'external',
  direct_url: "",
  comic_id: 0,
  status: "active" as 'active' | 'inactive',
});

const resetAd = () => {
  ad.value = {
    title: "",
    image: "",
    active_from: "",
    active_to: "",
    type: "internal" as 'internal' | 'external',
    direct_url: "",
    comic_id: 0,
    status: "active" as 'active' | 'inactive',
  };
  previewImage.value = '';
};

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
      ad.value.image = response.data || response;
    } catch (error) {
      previewImage.value = '';
      ad.value.image = '';
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
    if (!ad.value[field as keyof typeof ad.value]) {
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
      ...ad.value,
      comic_id: ad.value.type === 'internal' ? parseInt(ad.value.comic_id.toString()) : null,
      active_from: new Date(ad.value.active_from).toISOString(),
      active_to: new Date(ad.value.active_to).toISOString(),
    };
    
    await adStore.createAd(formattedData);
    adStore.createDialogIsOpen = false;
    await adStore.getAdData();
    resetAd();
    
    toast({
      description: "Ad created successfully",
    });
  } catch (error: any) {
    toast({
      description: error.response?.data?.message || "Failed to create ad",
      variant: "destructive",
    });
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <Dialog
    :open="adStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
      adStore.createDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Ad</DialogTitle>
      </DialogHeader>

      <div class="grid gap-4">
        <div class="flex items-center gap-4">
          <Label for="title" class="text-right w-1/4">Title *</Label>
          <Input v-model="ad.title" placeholder="Title" />
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
          <Input type="datetime-local" v-model="ad.active_from" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_to" class="text-right w-1/4">End Date *</Label>
          <Input type="datetime-local" v-model="ad.active_to" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="type" class="text-right w-1/4">Type *</Label>
          <Select v-model="ad.type">
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
          <Input v-model="ad.direct_url" placeholder="URL" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="comic_id" class="text-right w-1/4">Comic ID</Label>
          <Input type="number" v-model="ad.comic_id" placeholder="Comic ID" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="status" class="text-right w-1/4">Status *</Label>
          <Select v-model="ad.status">
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
          @click="adStore.createDialogIsOpen = false"
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