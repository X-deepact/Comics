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
import { ref, watch } from "vue";
import { useAdStore } from "@/stores/adStore";
import loadingImg from "@/assets/loading.svg";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { toast } from "@/components/ui/toast/use-toast";
import axios from "axios";
import { authHeader } from "@/services/authHeader";

const adStore = useAdStore();
const isLoading = ref(false);
const previewImage = ref('');

// Add comic search functionality
const comicSearchKeyword = ref("");
const comicSearchResults = ref<{id: number, name: string}[]>([]);
const searchingComics = ref(false);
const selectedComic = ref<{id: number, name: string} | null>(null);

const ad = ref({
  title: "",
  image: "",
  active_from: "",
  active_to: "",
  type: "internal" as 'internal' | 'external',
  direct_url: "",
  comic_id: 0,
  status: "active" as 'active' | 'inactive',
});

// Watch for type changes to reset comic selection when switching to external
watch(() => ad.value.type, (newType) => {
  if (newType === 'external') {
    selectedComic.value = null;
    ad.value.comic_id = 0;
  }
});

const resetAd = () => {
  ad.value = {
    title: "",
    image: "",
    active_from: "",
    active_to: "",
    type: "internal" as 'internal' | 'external',
    direct_url: "",
    comic_id: 0,
    status: "active" as 'active' | 'inactive',
  };
  previewImage.value = '';
  comicSearchKeyword.value = '';
  comicSearchResults.value = [];
  selectedComic.value = null;
};

// Comic search functions
const handleComicSearch = debounce(async (value: string) => {
  if (!value || value.trim().length === 0) {
    comicSearchResults.value = [];
    return;
  }
  
  searchingComics.value = true;
  try {
    const API_URL = import.meta.env.VITE_API_URL;
    const response = await axios.get(`${API_URL}/general/comics`, {
      params: { name: value.trim() },
      headers: authHeader()
    });

    if (response.data.code === "SUCCESS") {
      comicSearchResults.value = response.data.data;
    } else {
      comicSearchResults.value = [];
      toast({
        description: response.data.msg,
        variant: "destructive",
      });
    }
  } catch (error: any) {
    comicSearchResults.value = [];
    toast({
      description: error.response?.data?.msg || "Failed to search comics",
      variant: "destructive",
    });
  } finally {
    searchingComics.value = false;
  }
}, 300);

const selectComic = (comic: {id: number, name: string}) => {
  selectedComic.value = comic;
  ad.value.comic_id = comic.id;
  comicSearchKeyword.value = "";
  comicSearchResults.value = [];
};

const removeComic = () => {
  selectedComic.value = null;
  ad.value.comic_id = 0;
};

// Debounce utility function
function debounce<T extends (...args: any[]) => any>(
  func: T,
  wait: number
): (...args: Parameters<T>) => void {
  let timeout: number | null = null;

  return function executedFunction(...args: Parameters<T>) {
    if (timeout) {
      clearTimeout(timeout);
    }

    timeout = setTimeout(() => {
      func(...args);
      timeout = null;
    }, wait);
  };
}

// Add type guard function
interface ResponseWithData {
  data: string;
}

function isResponseWithData(response: unknown): response is ResponseWithData {
  return typeof response === 'object' && response !== null && 'data' in response;
}

const handleFileUpload = async (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    const file = target.files[0];
    
    // Show preview
    const reader = new FileReader();
    reader.onload = (e) => {
      if (e.target?.result) {
        previewImage.value = e.target.result as string;
      }
    };
    reader.readAsDataURL(file);

    try {
      // Upload image and get filename
      const response = await adStore.uploadAdImage(file);
      ad.value.image = isResponseWithData(response) ? response.data : response as string;
    } catch (error) {
      previewImage.value = '';
      ad.value.image = '';
      toast({
        description: "Failed to upload image",
        variant: "destructive",
      });
    }
  }
};

const validateForm = () => {
  const requiredFields = {
    title: "Title",
    image: "Image",
    type: "Type",
    direct_url: "URL",
    status: "Status",
    active_from: "Start Date",
    active_to: "End Date"
  };

  for (const [field, label] of Object.entries(requiredFields)) {
    if (!ad.value[field as keyof typeof ad.value]) {
      toast({
        description: `${label} is required`,
        variant: "destructive",
      });
      return false;
    }
  }
  
  // Validate comic selection for internal ads
  if (ad.value.type === 'internal' && !ad.value.comic_id) {
    toast({
      description: "Please select a comic for internal ads",
      variant: "destructive",
    });
    return false;
  }
  
  return true;
};

const handleSubmit = async () => {
  try {
    if (!validateForm()) return;
    
    isLoading.value = true;
    
    const formattedData = {
      ...ad.value,
      comic_id: ad.value.type === 'internal' ? parseInt(ad.value.comic_id.toString()) : null,
      active_from: new Date(ad.value.active_from).toISOString(),
      active_to: new Date(ad.value.active_to).toISOString(),
    };
    
    await adStore.createAd(formattedData);
    adStore.createDialogIsOpen = false;
    await adStore.getAdData();
    resetAd();
    
    toast({
      description: "Ad created successfully",
    });
  } catch (error: any) {
    toast({
      description: error.response?.data?.message || "Failed to create ad",
      variant: "destructive",
    });
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <Dialog
    :open="adStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
      adStore.createDialogIsOpen = value;
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Ad</DialogTitle>
      </DialogHeader>

      <div class="grid gap-4">
        <div class="flex items-center gap-4">
          <Label for="title" class="text-right w-1/4">Title *</Label>
          <Input v-model="ad.title" placeholder="Title" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="image" class="text-right w-1/4">Image *</Label>
          <div class="w-full">
            <Input type="file" @change="handleFileUpload" accept="image/*" />
            <div v-if="previewImage" class="mt-2">
              <img
                :src="previewImage"
                alt="Preview"
                class="max-w-full h-32 object-contain rounded-md border"
              />
            </div>
          </div>
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_from" class="text-right w-1/4">Start Date *</Label>
          <Input type="datetime-local" v-model="ad.active_from" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="active_to" class="text-right w-1/4">End Date *</Label>
          <Input type="datetime-local" v-model="ad.active_to" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="type" class="text-right w-1/4">Type *</Label>
          <Select v-model="ad.type">
            <SelectTrigger>
              <SelectValue placeholder="Select type" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="internal">Internal</SelectItem>
              <SelectItem value="external">External</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div class="flex items-center gap-4">
          <Label for="direct_url" class="text-right w-1/4">URL</Label>
          <Input v-model="ad.direct_url" placeholder="URL" />
        </div>
        <div class="flex items-center gap-4" v-if="ad.type === 'internal'">
          <Label for="comic_id" class="text-right w-1/4">Comic *</Label>
          <div class="w-full">
            <div class="flex gap-2">
              <Input
                v-model="comicSearchKeyword"
                placeholder="Search comics by name..."
                @input="(e: any) => handleComicSearch((e.target as HTMLInputElement).value)"
                class="flex-1"
              />
            </div>

            <!-- Search Results -->
            <div v-if="comicSearchResults.length > 0" class="border rounded-lg shadow-sm mt-2">
              <div class="max-h-48 overflow-y-auto">
                <div
                  v-for="comic in comicSearchResults"
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
            <div v-else-if="searchingComics" class="mt-2 text-sm text-gray-500">
              Searching...
            </div>

            <!-- No Result Message -->
            <div 
              v-else-if="comicSearchKeyword && comicSearchResults.length === 0" 
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
        <div class="flex items-center gap-4">
          <Label for="status" class="text-right w-1/4">Status *</Label>
          <Select v-model="ad.status">
            <SelectTrigger>
              <SelectValue placeholder="Select status" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="active">Active</SelectItem>
              <SelectItem value="inactive">Inactive</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="adStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="handleSubmit"
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" />
          Add
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template> 