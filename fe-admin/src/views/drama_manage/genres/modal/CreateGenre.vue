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
import { useDramaGenreStore } from "@/stores/drama_genreStore";
import loadingImg from "@/assets/loading.svg";
import { toast } from "@/components/ui/toast/use-toast";

const genreStore = useDramaGenreStore();
const isLoading = ref(false);

const languageOptions = [
  { value: 'en', label: 'English' },
  { value: 'zh', label: 'Chinese' },
  { value: 'vi', label: 'Vietnamese' }
];

const genre = ref({
  name: "",
  position: 0,
  translations: [
    { language: "en", translated_name: "" }
  ]
});

const resetGenre = () => {
  genre.value = {
    name: "",
    position: 0,
    translations: [
      { language: "en", translated_name: "" }
    ]
  };
};

const addTranslation = () => {
  genre.value.translations.push({ language: "", translated_name: "" });
};

const removeTranslation = (index: number) => {
  genre.value.translations.splice(index, 1);
};

const validateForm = () => {
  if (!genre.value.name.trim()) {
    toast({
      description: "Name is required",
      variant: "destructive",
    });
    return false;
  }

  // Validate all translations
  for (const translation of genre.value.translations) {
    if (!translation.language.trim()) {
      toast({
        description: "Language is required for all translations",
        variant: "destructive",
      });
      return false;
    }
    if (!translation.translated_name.trim()) {
      toast({
        description: "Translated name is required for all translations",
        variant: "destructive",
      });
      return false;
    }
  }

  return true;
};

const handleSubmit = async () => {
  try {
    if (!validateForm()) return;
    
    isLoading.value = true;
    
    await genreStore.createGenre(genre.value);
    genreStore.createDialogIsOpen = false;
    await genreStore.getGenreData();
    resetGenre();
    
    toast({
      description: "Genre created successfully",
    });
  } catch (error: any) {
    toast({
      description: error.response?.data?.message || "Failed to create genre",
      variant: "destructive",
    });
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <Dialog
    :open="genreStore.createDialogIsOpen"
    @update:open="(value: boolean) => {
      genreStore.createDialogIsOpen = value;
      if (!value) resetGenre();
    }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Genre</DialogTitle>
      </DialogHeader>

      <div class="grid gap-4">
        <div class="flex items-center gap-4">
          <Label for="name" class="text-right w-1/4">Name *</Label>
          <Input v-model="genre.name" placeholder="Name" />
        </div>
        <div class="flex items-center gap-4">
          <Label for="position" class="text-right w-1/4">Position</Label>
          <Input type="number" v-model.number="genre.position" placeholder="Position" />
        </div>
        
        <div class="border p-4 rounded-md mt-4">
          <div class="flex justify-between items-center mb-4">
            <h3 class="font-medium">Translations</h3>
            <Button type="button" variant="outline" size="sm" @click="addTranslation">
              Add Translation
            </Button>
          </div>
          
          <div v-for="(translation, index) in genre.translations" :key="index" class="mb-4 border-b pb-4 last:border-b-0">
            <div class="flex items-center gap-4 mb-2">
              <Label :for="`language-${index}`" class="text-right w-1/4">Language *</Label>
              <select 
                v-model="translation.language"
                :id="`language-${index}`"
                class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              >
                <option value="">Select Language</option>
                <option v-for="option in languageOptions" :key="option.value" :value="option.value">
                  {{ option.label }}
                </option>
              </select>
            </div>
            <div class="flex items-center gap-4">
              <Label :for="`translated-name-${index}`" class="text-right w-1/4">Translated Name *</Label>
              <div class="flex gap-2 w-3/4">
                <Input v-model="translation.translated_name" :id="`translated-name-${index}`" placeholder="Translated name" class="flex-1" />
                <Button 
                  v-if="genre.translations.length > 1" 
                  type="button" 
                  variant="destructive" 
                  size="icon" 
                  @click="removeTranslation(index)"
                >
                  <span class="sr-only">Remove</span>
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4"><path d="M18 6 6 18"></path><path d="m6 6 12 12"></path></svg>
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="genreStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          :disabled="isLoading"
          @click="handleSubmit"
        >
          <img v-if="isLoading" :src="loadingImg" size="icon" alt="Loading" />
          Add
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>