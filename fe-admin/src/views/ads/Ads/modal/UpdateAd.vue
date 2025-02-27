<template>
  <Dialog
    :open="adsStore.updateDialogIsOpen"
    @update:open="(value: boolean) => {
      adsStore.updateDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Advertisement</DialogTitle>
      </DialogHeader>
      
      <!-- Current image preview -->
      <div class="flex flex-col items-center gap-4">
        <img
          v-if="previewUrl"
          :src="previewUrl"
          class="max-w-[200px] h-auto"
          alt="Preview"
        />
        <img
          v-else-if="adsStore.selectedData.image"
          :src="adsStore.selectedData.image"
          class="max-w-[200px] h-auto"
          alt="Current"
        />
        <Input
          type="file"
          placeholder="Image"
          class="w-[50%]"
          @change="handleFileChange"
          accept="image/*"
        />
      </div>

      <div class="grid gap-4 py-4">
        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Title</Label>
          <Input v-model="adsStore.selectedData.title" class="col-span-3" />
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Type</Label>
          <Select v-model="adsStore.selectedData.type" class="col-span-3">
            <SelectTrigger>
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="internal">Internal</SelectItem>
              <SelectItem value="external">External</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">URL</Label>
          <Input v-model="adsStore.selectedData.direct_url" class="col-span-3" />
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Status</Label>
          <Select 
            v-model="adsStore.selectedData.status" 
            class="col-span-3"
          >
            <SelectTrigger>
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="active">Active</SelectItem>
              <SelectItem value="inactive">Inactive</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>
      <DialogFooter>
        <Button
          variant="secondary"
          @click="closeDialog"
        >
          Cancel
        </Button>
        <Button
          @click="handleUpdate"
        >
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAdsStore } from "@/stores/adsStore";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const adsStore = useAdsStore();
const previewUrl = ref<string | null>(null);

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    if (previewUrl.value) {
      URL.revokeObjectURL(previewUrl.value);
    }
    const file = target.files[0];
    adsStore.selectedData.image = file;
    previewUrl.value = URL.createObjectURL(file);
  }
};

const closeDialog = () => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
    previewUrl.value = null;
  }
  adsStore.updateDialogIsOpen = false;
};

const handleUpdate = async () => {
  const updateData = {
    id: adsStore.selectedData.id,
    title: adsStore.selectedData.title,
    type: adsStore.selectedData.type || 'internal',
    direct_url: adsStore.selectedData.direct_url,
    active_from: adsStore.selectedData.active_from,
    active_to: adsStore.selectedData.active_to,
    status: (adsStore.selectedData.status || 'active').toLowerCase(),
    image: adsStore.selectedData.image instanceof File ? adsStore.selectedData.image : undefined,
    comic_id: adsStore.selectedData.comic_id,
  };

  console.log('Sending update data:', updateData);

  await adsStore.updateAd(updateData);
  closeDialog();
  adsStore.getAdsData();
};
</script> 