<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
} from "@/components/ui/dialog";
import { useUserStore } from "../../../stores/userStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
const isLoading = ref(false);
const userStore = useUserStore();
</script>
<template>
  <Dialog
    :open="userStore.updateStatusDialogIsOpen"
    @update:open="(value:boolean)=>{
      userStore.updateStatusDialogIsOpen = value;
      }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to update this user's status?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="userStore.updateStatusDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              isLoading = true;
              await userStore.updateUserStatus(userStore.selectedData.id);
              isLoading = false;
              userStore.updateStatusDialogIsOpen = false;
              userStore.getUserData();
            }
          "
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Update Status
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
