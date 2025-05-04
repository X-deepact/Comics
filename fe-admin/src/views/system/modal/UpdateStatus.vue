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

const isLoading = ref(false);
const userStore = useUserStore();
const { toast } = useToast();
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
              try {
                isLoading = true;
                await userStore.updateUserStatus(userStore.selectedData.id);
                isLoading = false;
                userStore.updateStatusDialogIsOpen = false;
                userStore.getUserData();
              } catch (error) {
                toast({
                  description: 'Failed to update user status. Please try again.',
                  variant: 'destructive'
                });
                isLoading = false;
              }
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
