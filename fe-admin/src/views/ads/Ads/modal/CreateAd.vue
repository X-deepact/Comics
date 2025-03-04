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

const adStore = useAdStore();
const isLoading = ref(false);

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
};

const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    const file = target.files[0];
    const reader = new FileReader();
    reader.onload = (e) => {
      if (e.target?.result) {
        ad.value.image = e.target.result as string;
      }
    };
    reader.readAsDataURL(file);
  }
};

const handleSubmit = async () => {
  try {
    isLoading.value = true;
    const formattedData = {
      ...ad.value,
      comic_id: parseInt(ad.value.comic_id.toString()),
      active_from: new Date(ad.value.active_from).toISOString(),
      active_to: new Date(ad.value.active_to).toISOString(),
    };
    await adStore.createAd(formattedData);
    adStore.createDialogIsOpen = false;
    adStore.getAdData();
    resetAd();
  } catch (error) {
    console.error('Error creating ad:', error);
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
          <Label for="title" class="text-right w-1/4">Title</Label>
          <Input v-model="ad.title" placeholder="Title" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="image" class="text-right w-1/4">Image</Label>
          <Input type="file" @change="handleFileUpload" accept="image/*" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_from" class="text-right w-1/4">Start Date</Label>
          <Input type="datetime-local" v-model="ad.active_from" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_to" class="text-right w-1/4">End Date</Label>
          <Input type="datetime-local" v-model="ad.active_to" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="type" class="text-right w-1/4">Type</Label>
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
          <Label for="status" class="text-right w-1/4">Status</Label>
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