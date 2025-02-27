<template>
  <Dialog
    :open="recommendStore.updateDialogIsOpen"
    @update:open="recommendStore.updateDialogIsOpen = $event"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Recommendation</DialogTitle>
      </DialogHeader>
      <img
        v-if="previewUrl"
        :src="previewUrl"
        class="max-w-[200px] justify-self-center h-auto"
      />
      <Input
        type="file"
        placeholder="cover"
        class="justify-self-center w-[50%]"
        @change="handleFileChange"
        accept="image/*"
      />
      <div class="flex items-center gap-4">
        <Label class="text-right w-1/4">Title</Label>
        <Input v-model="recommendStore.selectedData.title" placeholder="Title" />
      </div>
      <div class="flex items-center gap-4">
        <Label class="text-right w-1/4">Position</Label>
        <Input
          type="number"
          v-model="recommendStore.selectedData.position"
          placeholder="Position"
        />
      </div>
      
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="recommendStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="
            async () => {
              await recommendStore.updateRecommend(recommendStore.selectedData);
              recommendStore.updateDialogIsOpen = false;
              recommendStore.getRecommendData();
              resetPreview();
            }
          "
        >
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
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
import { useRecommendStore } from "@/stores/recommendStore";

const recommendStore = useRecommendStore();
const previewUrl = ref<string | null>(null);

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    if (previewUrl.value) {
      URL.revokeObjectURL(previewUrl.value);
    }
    recommendStore.selectedData.cover = target.files[0];
    previewUrl.value = URL.createObjectURL(target.files[0]);
  }
};

const resetPreview = () => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
    previewUrl.value = null;
  }
};
</script> 