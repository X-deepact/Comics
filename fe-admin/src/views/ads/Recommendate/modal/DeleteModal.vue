<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
} from "@/components/ui/dialog";
import { useRecommendStore } from "@/stores/recommendStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";

const isLoading = ref(false);
const recommendStore = useRecommendStore();
</script>

<template>
  <Dialog
    :open="recommendStore.deleteDialogIsOpen"
    @update:open="(value: boolean) => {
      recommendStore.deleteDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this recommendation?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="recommendStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await recommendStore.deleteRecommend(recommendStore.selectedData.id);
              isLoading = false;
              recommendStore.deleteDialogIsOpen = false;
              recommendStore.getRecommendData();
            }
          "
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Delete
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template> 