<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useRecommendStore } from "@/stores/recommendStore";
import loadingImg from "@/assets/loading.svg";
import { ref, watch } from "vue";
import { useToast } from "@/components/ui/toast/use-toast";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { RecommendPosition, positionLabels } from "@/stores/recommendStore";

const { toast } = useToast();
const isLoading = ref(false);
const recommendStore = useRecommendStore();
const previewImage = ref<string>("");
const fileInput = ref<HTMLInputElement | null>(null);

const formatDateForInput = (dateValue: string | number) => {
  if (!dateValue) return '';
  try {
    const date = new Date(dateValue);
    if (isNaN(date.getTime())) return '';
    
    return date.toISOString().split('T')[0];
  } catch (error) {
    toast({
      description: `Date formatting error: ${error}`,
      variant: "destructive"
    });
    return '';
  }
};

const formData = ref({
  id: 0,
  title: "",
  cover: "",
  position: RecommendPosition.COMPLETE_MASTERPIECE,
  active_from: "",
  active_to: "",
});

// Initialize form data when selectedData changes
watch(() => recommendStore.selectedData, (newData) => {
  if (newData) {
    const activeFrom = typeof newData.active_from === 'number' 
      ? formatDateForInput(newData.active_from * 1000)
      : formatDateForInput(newData.active_from);
    
    const activeTo = typeof newData.active_to === 'number'
      ? formatDateForInput(newData.active_to * 1000)
      : formatDateForInput(newData.active_to);

    formData.value = {
      id: newData.id,
      title: newData.title || "",
      cover: "", // Reset cover since we'll use previewImage for the existing image
      position: newData.position || RecommendPosition.COMPLETE_MASTERPIECE,
      active_from: activeFrom,
      active_to: activeTo,
    };
    // Store the existing image URL in previewImage
    previewImage.value = newData.cover || "";
  }
}, { immediate: true });

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    const file = target.files[0];
    formData.value.cover = file;
    previewImage.value = URL.createObjectURL(file);
  }
};

const handleSubmit = async () => {
  if (!validateForm()) return;

  isLoading.value = true;
  try {
    const activeFromDate = new Date(formData.value.active_from);
    const activeToDate = new Date(formData.value.active_to);

    activeFromDate.setHours(0, 0, 0, 0);
    activeToDate.setHours(0, 0, 0, 0);

    const submitData = {
      id: formData.value.id,
      title: formData.value.title,
      // Only include cover if it's a new file, otherwise don't send it at all
      ...(formData.value.cover instanceof File ? { cover: formData.value.cover } : {}),
      position: Number(formData.value.position),
      active_from: Math.floor(activeFromDate.getTime() / 1000),
      active_to: Math.floor(activeToDate.getTime() / 1000),
    };

    await recommendStore.updateRecommend(submitData);
    recommendStore.updateDialogIsOpen = false;
    await recommendStore.getRecommendData();
  } catch (error: any) {
    toast({
      description: `Update error: ${error.response?.data?.message || error.message}`,
      variant: "destructive",
    });
  } finally {
    isLoading.value = false;
  }
};

const validateForm = () => {
  if (!formData.value.title?.trim()) {
    toast({
      description: "Title is required",
      variant: "destructive",
    });
    return false;
  }
  if (!formData.value.position) {
    toast({
      description: "Position is required",
      variant: "destructive",
    });
    return false;
  }
  if (!formData.value.active_from || !formData.value.active_to) {
    toast({
      description: "Active dates are required",
      variant: "destructive",
    });
    return false;
  }

  // Validate dates
  const fromDate = new Date(formData.value.active_from);
  const toDate = new Date(formData.value.active_to);
  
  if (isNaN(fromDate.getTime()) || isNaN(toDate.getTime())) {
    toast({
      description: "Invalid date format",
      variant: "destructive",
    });
    return false;
  }

  if (fromDate > toDate) {
    toast({
      description: "Active from date must be before active to date",
      variant: "destructive",
    });
    return false;
  }

  return true;
};
</script>

<template>
  <Dialog
    :open="recommendStore.updateDialogIsOpen"
    @update:open="(value: boolean) => {
      recommendStore.updateDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Recommendation</DialogTitle>
        <DialogDescription>
          Update the recommendation details below.
        </DialogDescription>
      </DialogHeader>
      
      <div class="grid gap-4">
        <div class=" gap-4">
          <div class="space-y-4">
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Title *</Label>
              <Input v-model="formData.title" class="col-span-3" placeholder="Enter title" />
            </div>
            <div class="space-y-4">
            <div class="grid grid-cols-4 items-center gap-4">
              <Label class="text-right">Image</Label>
              <div class="col-span-3">
                <input
                  ref="fileInput"
                  type="file"
                  accept="image/*"
                  class="hidden"
                  @change="handleFileChange"
                />
                <Button
                  type="button"
                  variant="outline"
                  class="w-full"
                  @click="fileInput?.click()"
                >
                  Choose Image
                </Button>
                <div v-if="previewImage || formData.cover" class="mt-2">
                  <img
                    :src="previewImage || formData.cover"
                    alt="Preview"
                    class="max-w-full max-h-[200px] object-contain rounded-md border"
                  />
                </div>
              </div>
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
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="recommendStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="handleSubmit"
        >
          <img v-if="isLoading" :src="loadingImg" alt="loading" class="mr-2 h-4 w-4" />
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>