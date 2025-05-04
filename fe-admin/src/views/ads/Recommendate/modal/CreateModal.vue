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
const selectedComic = ref<{ id: number, name: string } | null>(null);

// Add Comic interface
interface Comic {
  id: number;
  name: string;
}

const formData = ref({
  title: "",
  cover: null as File | null,
  position: String(RecommendPosition.TOPING),
  active_from: "",
  active_to: "",
});

// Add comic search functionality
const comicSearchKeyword = ref("");

// Function to select comic
const selectComic = (comic: Comic) => {
  selectedComic.value = comic;
  comicSearchKeyword.value = "";
  // Clear the search results in the store
  recommendStore.comicSearchResults = [];
};

// Function to remove selected comic
const removeComic = () => {
  selectedComic.value = null;
};

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
    cover: null,
    position: String(RecommendPosition.TOPING),
    active_from: "",
    active_to: "",
  };
  previewImage.value = "";
  selectedComic.value = null;
  if (fileInput.value) {
    fileInput.value.value = "";
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

    // Ensure cover is not null before submitting
    if (!formData.value.cover) {
      throw new Error("Cover image is required");
    }

    const submitData: {
      title: string;
      cover: File;
      position: number;
      active_from: number;
      active_to: number;
    } = {
      title: formData.value.title,
      cover: formData.value.cover!,
      position: Number(formData.value.position),
      active_from: Math.floor(activeFromDate.getTime() / 1000),
      active_to: Math.floor(activeToDate.getTime() / 1000),
    };

    const createdRecommend = await recommendStore.createRecommend(submitData);
    
    // Add the selected comic
    if (createdRecommend && createdRecommend.id && selectedComic.value) {
      await recommendStore.createRecommendComic(createdRecommend.id, selectedComic.value.id);
    }

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
      !formData.value.active_from || !formData.value.active_to || !selectedComic.value) {
    toast({
      description: "Please fill in all required fields and select a comic",
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

// Update debounce function to use Window.Timeout instead of NodeJS.Timeout
const debounce = <T extends (...args: any[]) => any>(
  func: T,
  wait: number
): (...args: Parameters<T>) => void => {
  let timeout: number | null = null;

  return function executedFunction(...args: Parameters<T>) {
    if (timeout) {
      window.clearTimeout(timeout);
    }

    timeout = window.setTimeout(() => {
      func(...args);
      timeout = null;
    }, wait);
  };
};

// Then use it as before
const handleComicSearch = debounce((value: string | null | undefined) => {
  recommendStore.searchComics(value);
}, 300);
</script>

<template>
  <Dialog
    :open="recommendStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
      recommendStore.createDialogIsOpen = value;
      if (!value) resetForm();
    }"
  >
    <DialogContent>
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
              <input
                ref="fileInput"
                type="file"
                accept="image/*"
                @change="handleFileChange"
                class="hidden"
              />
              <Button
                type="button"
                variant="outline"
                class="w-full"
                @click="fileInput?.click()"
              >
                Choose Image
              </Button>
              <div v-if="previewImage" class="mt-2">
                <img
                  :src="previewImage"
                  alt="Preview"
                  class="max-w-full max-h-[200px] object-contain rounded-md border"
                />
              </div>
            </div>
          </div>

          <div class="grid grid-cols-4 items-center gap-4">
            <Label class="text-right">Position *</Label>
            <div class="col-span-3">
              <Select v-model="formData.position">
                <SelectTrigger>
                  <SelectValue :placeholder="positionLabels[Number(formData.position)]" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem :value="String(RecommendPosition.TOPING)">
                    {{ positionLabels[RecommendPosition.TOPING] }}
                  </SelectItem>
                  <SelectItem :value="String(RecommendPosition.RECOMMEND_PRODUCTS)">
                    {{ positionLabels[RecommendPosition.RECOMMEND_PRODUCTS] }}
                  </SelectItem>
                  <SelectItem :value="String(RecommendPosition.RECOMMEND_MASTERPIECES)">
                    {{ positionLabels[RecommendPosition.RECOMMEND_MASTERPIECES] }}
                  </SelectItem>
                  <SelectItem :value="String(RecommendPosition.FASTEST_GROWING)">
                    {{ positionLabels[RecommendPosition.FASTEST_GROWING] }}
                  </SelectItem>
                  <SelectItem :value="String(RecommendPosition.TESTING_NEW_PRODUCTS)">
                    {{ positionLabels[RecommendPosition.TESTING_NEW_PRODUCTS] }}
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          <div class="grid grid-cols-4 items-center gap-4">
            <Label class="text-right">Active From *</Label>
            <Input
              type="date"
              v-model="formData.active_from"
              class="col-span-3"
            />
          </div>

          <div class="grid grid-cols-4 items-center gap-4">
            <Label class="text-right">Active To *</Label>
            <Input
              type="date"
              v-model="formData.active_to"
              class="col-span-3"
            />
          </div>
        </div>

        <div class="grid grid-cols-4 items-center gap-4">
          <Label class="text-right">Comic *</Label>
          <div class="col-span-3">
            <div class="flex gap-2">
              <Input
                v-model="comicSearchKeyword"
                placeholder="Search comics by name..."
                @input="(e: Event) => handleComicSearch((e.target as HTMLInputElement).value)"
                class="flex-1"
              />
            </div>

            <!-- Search Results -->
            <div v-if="recommendStore.comicSearchResults.length > 0" class="border rounded-lg shadow-sm mt-2">
              <div class="max-h-48 overflow-y-auto">
                <div
                  v-for="comic in recommendStore.comicSearchResults"
                  :key="comic.id"
                  class="flex items-center justify-between p-3 hover:bg-gray-50 cursor-pointer border-b last:border-b-0"
                  @click="selectComic(comic)"
                >
                  <span class="font-medium">{{ comic.name }}</span>
                  <Button size="sm" variant="ghost">
                    Select
                  </Button>
                </div>
              </div>
            </div>

            <!-- Loading State -->
            <div v-else-if="recommendStore.searchingComics" class="mt-2 text-sm text-gray-500">
              Searching...
            </div>

            <!-- No Results Message -->
            <div 
              v-else-if="comicSearchKeyword && recommendStore.comicSearchResults.length === 0" 
              class="mt-2 text-sm text-gray-500"
            >
              No comics found
            </div>

            <!-- Selected Comic -->
            <div v-if="selectedComic" class="mt-2 border rounded-lg shadow-sm">
              <div class="flex items-center justify-between p-3">
                <span class="font-medium">{{ selectedComic.name }}</span>
                <Button
                  size="sm"
                  variant="destructive"
                  @click="removeComic"
                >
                  Remove
                </Button>
              </div>
            </div>
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