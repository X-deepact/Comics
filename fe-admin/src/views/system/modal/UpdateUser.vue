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
import { useUserStore } from "../../../stores/userStore";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
import { cn } from "@/lib/utils";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  DateFormatter,
  type DateValue,
  getLocalTimeZone,
} from "@internationalized/date";
import { CalendarIcon } from "lucide-vue-next";
import { useToast } from "@/components/ui/toast/use-toast";
const { toast } = useToast();
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});
const value = ref<DateValue | null>(null);

const isLoading = ref(false);
const userStore = useUserStore();

const getBirthday = () => {
  if (value.value) {
    const date = value.value.toDate(getLocalTimeZone());
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    value.value = undefined;
    return `${year}-${month}-${day}`;
  }
};
const checkForm = () => {
  if (!userStore.selectedData.username || 
      !userStore.selectedData.DisplayName ||
      !userStore.selectedData.FirstName ||
      !userStore.selectedData.LastName ||
      !userStore.selectedData.phone ||
      !userStore.selectedData.email ||
      !userStore.selectedData.password ||
      !userStore.selectedData.TierId) {
    toast({
      description: "Enter correctly",
      variant: "destructive",
    });
    return false;
  }
  if (userStore.selectedData.password.length < 8) {
    toast({
      description: "Password must be at least 8 characters long",
      variant: "destructive",
    });
    return false;
  }
  return true;
};
</script>
<template>
  <Dialog :open="userStore.updateDialogIsOpen" @update:open="
    (value: boolean) => {
      userStore.updateDialogIsOpen = value;
    }
  ">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update User</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="username" class="text-center w-1/4">User Name *</Label>
        <Input v-model="userStore.selectedData.username" placeholder="User Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="DisplayName" class="text-center w-1/4">Display Name *</Label>
        <Input v-model="userStore.selectedData.DisplayName" placeholder="Display Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="FirstName" class="text-center w-1/4">First Name *</Label>
        <Input v-model="userStore.selectedData.FirstName" placeholder="First Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="LastName" class="text-center w-1/4">Last Name *</Label>
        <Input v-model="userStore.selectedData.LastName" placeholder="Last Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="phone" class="text-center w-1/4">Phone *</Label>
        <Input v-model="userStore.selectedData.phone" placeholder="Phone" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="email" class="text-center w-1/4">Email *</Label>
        <Input v-model="userStore.selectedData.email" placeholder="Email" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="password" class="text-center w-1/4">Password *</Label>
        <Input type="password" v-model="userStore.selectedData.password" placeholder="Password" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="TierId" class="text-center w-1/4">Tier *</Label>
        <Input v-model="userStore.selectedData.TierId" placeholder="Tier" type="number" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="birth_day" class="text-center w-1/4">Birthday</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button variant="outline" :class="cn('font-normal w-3/4', !value && 'text-muted-foreground')
              ">
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{
                value
                  ? df.format(value.toDate(getLocalTimeZone()))
                  : "Pick a date"
              }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar v-model="value" initial-focus />
          </PopoverContent>
        </Popover>
      </div>
      <DialogFooter class="sm:justify-end">
        <Button variant="secondary" @click="userStore.updateDialogIsOpen = false">
          Close
        </Button>
        <Button :disabled="isLoading" @click="
          async () => {
            try {
              if (checkForm()) {
                isLoading = true;
                await userStore.updateUser({
                  id: userStore.selectedData.id,
                  username: userStore.selectedData.username,
                  DisplayName: userStore.selectedData.DisplayName,
                  FirstName: userStore.selectedData.FirstName,
                  LastName: userStore.selectedData.LastName,
                  phone: userStore.selectedData.phone,
                  email: userStore.selectedData.email,
                  password: userStore.selectedData.password,
                  TierId: userStore.selectedData.TierId,
                  Birthday: getBirthday(),
                });
                isLoading = false;
                userStore.updateDialogIsOpen = false;
                userStore.getUserData();
              }
            } catch (error) {
              toast({
                description: error instanceof Error ? error.message : 'An unexpected error occurred while updating user',
                variant: 'destructive'
              });
              isLoading = false;
            }
          }
        ">
          <img v-if="isLoading" :src="loadingImg" size="icon" />Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
