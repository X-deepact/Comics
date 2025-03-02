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
import { ref } from "vue";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const adStore = useAdStore();
const isLoading = ref(false);

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

const handleSubmit = async () => {
  try {
    isLoading.value = true;
    const formattedData = {
      ...adStore.selectedData,
      status: adStore.selectedData.status === 'active',
      comic_id: parseInt(adStore.selectedData.comic_id.toString()),
      active_from: new Date(adStore.selectedData.active_from).toISOString(),
      active_to: new Date(adStore.selectedData.active_to).toISOString(),
    };
    await adStore.updateAd(formattedData);
    adStore.updateDialogIsOpen = false;
    adStore.getAdData();
  } catch (error) {
    console.error('Error updating ad:', error);
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
          <Label for="title" class="text-right w-1/4">Title</Label>
          <Input v-model="adStore.selectedData.title" placeholder="Title" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="image" class="text-right w-1/4">Image</Label>
          <Input type="file" @change="handleFileUpload" accept="image/*" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_from" class="text-right w-1/4">Start Date</Label>
          <Input type="datetime-local" v-model="adStore.selectedData.active_from" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_to" class="text-right w-1/4">End Date</Label>
          <Input type="datetime-local" v-model="adStore.selectedData.active_to" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="type" class="text-right w-1/4">Type</Label>
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
          <Label for="status" class="text-right w-1/4">Status</Label>
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