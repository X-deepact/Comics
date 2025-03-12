<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { ref } from "vue";
import { useRecommendStore, RecommendPosition, positionLabels } from "@/stores/recommendStore";
import loadingImg from "@/assets/loading.svg";
import { useToast } from "@/components/ui/toast/use-toast";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const { toast } = useToast();
const isLoading = ref(false);
const recommendStore = useRecommendStore();
const previewImage = ref<string>("");
const fileInput = ref<HTMLInputElement | null>(null);

const formData = ref({
  title: "",
  cover: "",
  position: RecommendPosition.COMPLETE_MASTERPIECE,
  active_from: "",
  active_to: "",
});

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    const file = target.files[0];
    formData.value.cover = file;
    previewImage.value = URL.createObjectURL(file);
  }
};

const resetForm = () => {
  formData.value = {
    title: "",
    cover: "",
    position: RecommendPosition.COMPLETE_MASTERPIECE,
    active_from: "",
    active_to: "",
  };
  previewImage.value = "";
  if (fileInput.value) {
    fileInput.value.value = "";
  }
};

const handleSubmit = async () => {
  if (!validateForm()) return;
  
  isLoading.value = true;
  try {
    // Convert date strings to Unix timestamps (start of day)
    const activeFromDate = new Date(formData.value.active_from);
    const activeToDate = new Date(formData.value.active_to);

    // Set time to start of day (00:00:00)
    activeFromDate.setHours(0, 0, 0, 0);
    activeToDate.setHours(0, 0, 0, 0);

    const submitData = {
      title: formData.value.title,
      cover: formData.value.cover,
      position: Number(formData.value.position),
      active_from: Math.floor(activeFromDate.getTime() / 1000),
      active_to: Math.floor(activeToDate.getTime() / 1000),
    };

    await recommendStore.createRecommend(submitData);
    recommendStore.createDialogIsOpen = false;
    await recommendStore.getRecommendData();
    resetForm();
  } catch (error: any) {
    toast({
      description: error.response?.data?.message || error.message,
      variant: "destructive",
    });
  } finally {
    isLoading.value = false;
  }
};

const validateForm = () => {
  if (!formData.value.title || !formData.value.cover || 
      !formData.value.active_from || !formData.value.active_to) {
    toast({
      description: "Please fill in all required fields",
      variant: "destructive",
    });
    return false;
  }

  const fromDate = new Date(formData.value.active_from);
  const toDate = new Date(formData.value.active_to);

  if (isNaN(fromDate.getTime()) || isNaN(toDate.getTime())) {
    toast({
      description: "Invalid date format",
      variant: "destructive",
    });
    return false;
  }

  if (toDate <= fromDate) {
    toast({
      description: "Active To date must be after Active From date",
      variant: "destructive",
    });
    return false;
  }

  return true;
};
</script>

<template>
  <Dialog
    :open="recommendStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
      recommendStore.createDialogIsOpen = value;
      if (!value) resetForm();
    }"
  >
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Create Recommendation</DialogTitle>
      </DialogHeader>
      <div class="grid gap-4 py-4">
        <div class="grid gap-4">
          <div class="grid grid-cols-4 items-center gap-4">
            <Label class="text-right">Title *</Label>
            <Input
              id="title"
              v-model="formData.title"
              class="col-span-3"
              required
            />
          </div>

          <div class="grid grid-cols-4 items-center gap-4">
            <Label class="text-right">Cover *</Label>
            <div class="col-span-3">
              <Input
                ref="fileInput"
                type="file"
                accept="image/*"
                @change="handleFileChange"
                class="mb-2"
                required
              />
              <img
                v-if="previewImage"
                :src="previewImage"
                alt="Preview"
                class="max-w-[200px] max-h-[200px] object-contain"
              />
            </div>
          </div>

          <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Position *</Label>
              <div class="col-span-3">
              <Select v-model="formData.position" >
                <SelectTrigger>
                  <SelectValue :placeholder="positionLabels[formData.position]" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem :value="RecommendPosition.TOPING">
                    {{ positionLabels[RecommendPosition.TOPING] }}
                  </SelectItem>
                  <SelectItem :value="RecommendPosition.RECOMMEND_PRODUCTS">
                    {{ positionLabels[RecommendPosition.RECOMMEND_PRODUCTS] }}
                  </SelectItem>
                  <SelectItem :value="RecommendPosition.RECOMMEND_MASTERPIECES">
                    {{ positionLabels[RecommendPosition.RECOMMEND_MASTERPIECES] }}
                  </SelectItem>
                  <SelectItem :value="RecommendPosition.FASTEST_GROWING">
                    {{ positionLabels[RecommendPosition.FASTEST_GROWING] }}
                  </SelectItem>
                  <SelectItem :value="RecommendPosition.TESTING_NEW_PRODUCTS">
                    {{ positionLabels[RecommendPosition.TESTING_NEW_PRODUCTS] }}
                  </SelectItem>
                </SelectContent>
              </Select>
              </div>
            </div>

          <div class="grid grid-cols-4 items-center gap-4">
            <Label class="text-right" for="active_from">
              Active From *
            </Label>
            <Input
              id="active_from"
              type="date"
              v-model="formData.active_from"
              class="col-span-3"
            />
          </div>

          <div class="grid grid-cols-4 items-center gap-4">
            <Label class="text-right" for="active_to">
              Active To *
            </Label>
            <Input
              id="active_to"
              type="date"
              v-model="formData.active_to"
              class="col-span-3"
            />
          </div>
        </div>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="recommendStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="handleSubmit"
        >
          <img v-if="isLoading" :src="loadingImg" alt="loading" class="mr-2 h-4 w-4" />
          Add
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template> 