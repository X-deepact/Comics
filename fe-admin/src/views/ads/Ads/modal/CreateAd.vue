<template>
  <Dialog
    :open="adsStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
      adsStore.createDialogIsOpen = value;
      if (!value) resetForm();
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Advertisement</DialogTitle>
      </DialogHeader>

      <div class="flex flex-col items-center gap-4">
        <img
          v-if="previewUrl"
          :src="previewUrl"
          class="max-w-[200px] h-auto"
          alt="Preview"
        />
        <Input
          type="file"
          placeholder="Image"
          class="w-[50%]"
          @change="handleFileChange"
          accept="image/*"
          required
        />
      </div>

      <div class="grid gap-4 py-4">
        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Title</Label>
          <Input v-model="newAd.title" class="col-span-3" required />
        </div>

        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Type</Label>
          <Select v-model="newAd.type" class="col-span-3">
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
          <Input v-model="newAd.direct_url" class="col-span-3" required />
        </div>

        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Start Date</Label>
          <Input
            type="datetime-local"
            v-model="newAd.active_from"
            class="col-span-3"
            required
          />
        </div>

        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">End Date</Label>
          <Input
            type="datetime-local"
            v-model="newAd.active_to"
            class="col-span-3"
            required
          />
        </div>

        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Status</Label>
          <Select v-model="newAd.status" class="col-span-3">
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
        <Button variant="secondary" @click="closeDialog">Cancel</Button>
        <Button
          @click="handleCreate"
          :disabled="!isFormValid"
        >
          Create
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
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

const newAd = ref({
  title: "",
  image: null as File | null,
  type: "internal",
  direct_url: "",
  active_from: "",
  active_to: "",
  status: "active",
});

const isFormValid = computed(() => {
  return (
    newAd.value.title &&
    newAd.value.image &&
    newAd.value.direct_url &&
    newAd.value.active_from &&
    newAd.value.active_to
  );
});

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    if (previewUrl.value) {
      URL.revokeObjectURL(previewUrl.value);
    }
    newAd.value.image = target.files[0];
    previewUrl.value = URL.createObjectURL(target.files[0]);
  }
};

const resetForm = () => {
  newAd.value = {
    title: "",
    image: null,
    type: "internal",
    direct_url: "",
    active_from: "",
    active_to: "",
    status: "active",
  };
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
    previewUrl.value = null;
  }
};

const closeDialog = () => {
  adsStore.createDialogIsOpen = false;
  resetForm();
};

const handleCreate = async () => {
  if (!newAd.value.image) {
    return;
  }

  await adsStore.createAd({
    title: newAd.value.title,
    image: newAd.value.image,
    type: newAd.value.type,
    direct_url: newAd.value.direct_url,
    active_from: newAd.value.active_from,
    active_to: newAd.value.active_to,
    status: newAd.value.status,
  });

  closeDialog();
  adsStore.getAdsData();
};
</script> 