<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
} from "@/components/ui/dialog";
import { useAdStore } from "@/stores/adStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";

const isLoading = ref(false);
const adStore = useAdStore();
</script>

<template>
  <Dialog
    :open="adStore.deleteDialogIsOpen"
    @update:open="(value: boolean) => {
      adStore.deleteDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this ad?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="adStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await adStore.deleteAd(adStore.selectedData.id);
              isLoading = false;
              adStore.deleteDialogIsOpen = false;
              adStore.getAdData();
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