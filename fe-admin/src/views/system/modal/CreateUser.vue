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
import { useUserStore } from "../../../stores/userStore";
import loadingImg from "@/assets/loading.svg";
import { cn } from "@/lib/utils";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  DateFormatter,
  getLocalTimeZone,
  type DateValue
} from "@internationalized/date";
import { CalendarIcon } from "lucide-vue-next";
import { useToast } from "@/components/ui/toast/use-toast";

const { toast } = useToast();
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});
const value = ref<DateValue>();

const userStore = useUserStore();
userStore.getUserData();
const isLoading = ref(false);
const user = ref({
  username: "",
  password: "",
  phone: "",
  email: "",
  FirstName: "",
  LastName: "",
  DisplayName: "",
  TierId: 0,
  Birthday: "",
});

const resetUser = () => {
  user.value = {
    username: "",
    password: "",
    phone: "",
    email: "",
    FirstName: "",
    LastName: "",
    DisplayName: "",
    TierId: 0,
    Birthday: "",
  };
};

const setBirthday = () => {
  if (value.value) {
    const date = value.value.toDate(getLocalTimeZone());
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    user.value.Birthday = `${year}-${month}-${day}`;
  }
};

const validateEmail = (email: string) => {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
};

const checkForm = () => {
  if (!user.value.username || !user.value.email || !user.value.password || !user.value.phone || 
      !user.value.FirstName || !user.value.LastName || !user.value.DisplayName || 
      !user.value.TierId || !user.value.Birthday) {
    toast({
      description: "Enter correctly",
      variant: "destructive",
    });
    return false;
  }

  if (!validateEmail(user.value.email)) {
    toast({
      description: "Email format is incorrect",
      variant: "destructive",
    });
    return false;
  }

  if (user.value.password.length < 8) {
    toast({
      description: "Password must be at least 8 characters",
      variant: "destructive",
    });
    return false;
  }

  return true;
};
</script>

<template>
  <Dialog :open="userStore.createDialogIsOpen"
    @update:open="(value: boolean) => { userStore.createDialogIsOpen = value; }">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Subject</DialogTitle>
      </DialogHeader>
      <div class="flex items-center gap-4">
        <Label for="username" class="text-center w-1/4">User Name *</Label>
        <Input v-model="user.username" placeholder="User Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="DisplayName" class="text-center w-1/4">Display Name *</Label>
        <Input v-model="user.DisplayName" placeholder="Display Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="FirstName" class="text-center w-1/4">First Name *</Label>
        <Input v-model="user.FirstName" placeholder="First Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="LastName" class="text-center w-1/4">Last Name *</Label>
        <Input v-model="user.LastName" placeholder="Last Name" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="phone" class="text-center w-1/4">Phone *</Label>
        <Input v-model="user.phone" placeholder="Phone" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="email" class="text-center w-1/4">Email *</Label>
        <Input v-model="user.email" placeholder="Email" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="password" class="text-center w-1/4">Password *</Label>
        <Input type="password" v-model="user.password" placeholder="Password" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="TierId" class="text-center w-1/4">Tier *</Label>
        <Input v-model="user.TierId" placeholder="Tier" type="number" />
      </div>
      <div class="flex items-center gap-4">
        <Label for="birthday" class="text-center w-1/4">Birthday *</Label>
        <Popover>
          <PopoverTrigger as-child>
            <Button variant="outline" :class="cn(
              'w-[280px] justify-start text-left font-normal',
              !value && 'text-muted-foreground'
            )">
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{ value ? df.format(value.toDate(getLocalTimeZone())) : "Pick a date" }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar 
              v-model="value" 
              mode="single"
              :multiple="false"
              initial-focus 
              @update:model-value="(date: DateValue | undefined) => { 
                if (date) {
                  value = date;
                  setBirthday();
                }
              }" 
            />
          </PopoverContent>
        </Popover>
      </div>
      <DialogFooter class="sm:justify-end">
        <Button variant="secondary" @click="userStore.createDialogIsOpen = false">
          Close
        </Button>
        <Button :disabled="isLoading" @click="
          async () => {
            if (checkForm()) {
              isLoading = true;
              setBirthday();
              await userStore.createUser(user);
              isLoading = false;
              userStore.createDialogIsOpen = false;
              userStore.getUserData();
              resetUser();
            }
          }
        "><img v-if="isLoading" :src="loadingImg" size="icon" />Add</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
