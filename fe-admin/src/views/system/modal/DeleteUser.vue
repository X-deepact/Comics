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
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const isLoading = ref(false);
const userStore = useUserStore();
</script>
<template>
  <Dialog
    :open="userStore.deleteDialogIsOpen"
    @update:open="(value:boolean)=>{
      userStore.deleteDialogIsOpen = value;
      }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this user?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button
          variant="secondary"
          @click="userStore.deleteDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="
            async () => {
              try {
                isLoading = true;
                await userStore.deleteUser(userStore.selectedData.id);
                isLoading = false;
                userStore.deleteDialogIsOpen = false;
                userStore.getUserData();
              } catch (error) {
                toast.error('Failed to delete user. Please try again later.');
                isLoading = false;
              }
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
