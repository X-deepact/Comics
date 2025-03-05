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
const fileInput = ref<HTMLInputElement | null>(null);

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
  }
});

const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    const file = target.files[0];
    const reader = new FileReader();
    reader.onload = (e) => {
      if (e.target?.result) {
        adStore.selectedData.image = e.target.result as string;
      }
    };
    reader.readAsDataURL(file);
  }
};

const validateForm = () => {
  if (!adStore.selectedData.title || 
      !adStore.selectedData.image ||
      !adStore.selectedData.type ||
      !adStore.selectedData.direct_url ||
      !adStore.selectedData.status ||
      !adStore.selectedData.active_from ||
      !adStore.selectedData.active_to) {
    toast({
      description: "Enter correctly",
      variant: "destructive",
    });
    return false;
  }
  return true;
};

const handleSubmit = async () => {
  try {
    if (!validateForm()) return;
    
    isLoading.value = true;
    const formattedData = {
      ...adStore.selectedData,
      comic_id: parseInt(adStore.selectedData.comic_id.toString()),
      active_from: new Date(adStore.selectedData.active_from).toISOString(),
      active_to: new Date(adStore.selectedData.active_to).toISOString(),
    };
    await adStore.updateAd(formattedData);
    adStore.updateDialogIsOpen = false;
    adStore.getAdData();
  } catch (error: any) {
    toast({
      description: error.response?.data?.message || error.message,
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
        <div class=" gap-4">
          <div class="space-y-4">
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Title *</Label>
              <Input v-model="adStore.selectedData.title" class="col-span-3" placeholder="Enter title" />
            </div>
            <div class="space-y-4">
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Image *</Label>
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
                  <div v-if="adStore.selectedData.image" class="mt-2">
                    <img
                      :src="adStore.selectedData.image"
                      alt="Preview"
                      class="max-w-full max-h-[200px] object-contain rounded-md border"
                    />
                  </div>
                </div>
              </div>
            </div>
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Type *</Label>
              <div class="col-span-3">
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
            </div>
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">URL *</Label>
              <Input v-model="adStore.selectedData.direct_url" class="col-span-3" placeholder="Enter URL" />
            </div>
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Status *</Label>
              <div class="col-span-3">
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
            
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Start Date *</Label>
              <Input 
                type="datetime-local" 
                v-model="adStore.selectedData.active_from"
                class="col-span-3"
              />
            </div>
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">End Date *</Label>
              <Input 
                type="datetime-local" 
                v-model="adStore.selectedData.active_to"
                class="col-span-3"
              />
            </div>
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Comic ID</Label>
              <Input type="number" v-model="adStore.selectedData.comic_id" class="col-span-3" placeholder="Enter Comic ID" />
            </div>
          </div>
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