<template>
  <div class="p-8">
    <Component
      title="Comic"
      :totalItems="comicStore.totalItems"
      :onAddClick="handleAdd"
      @update:search="
        (keyword) => {
          comicStore.searchKeyword = keyword;
          comicStore.getComicData();
        }
      "
    />

    <TableComponent :subjects="comicStore.comicData" :columns="ColumnNames" />

    <Pagination
      :currentPage="comicStore.current_page"
      :pageSize="comicStore.page_size"
      :totalPages="
        comicStore.totalItems % comicStore.page_size === 0
          ? comicStore.totalItems / comicStore.page_size
          : Math.ceil(comicStore.totalItems / comicStore.page_size)
      "
      @update:pageSize="
        (pageSize) => {
          comicStore.page_size = pageSize;
          comicStore.getComicData();
        }
      "
      @update:currentPage="
        (currentPage) => {
          comicStore.current_page = currentPage;
          comicStore.getComicData();
        }
      "
    />

    <Modal
      :isVisible="showAddModal"
      title="Add New Comic"
      description="Fill in the details for the new comic."
      :newSubject="newComic"
      @update:isVisible="showAddModal = $event"
      @confirm="saveNewComic"
    >
      <template #default="{ newSubject }">
        <div class="grid gap-4 py-4">
          <div class="grid grid-cols-2 gap-4">
            <div class="space-y-4">
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Title</Label>
                <Input v-model="newComic.title" class="col-span-3" />
              </div>
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Author</Label>
                <Input v-model="newComic.author" class="col-span-3" />
              </div>
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Subject</Label>
                <div class="col-span-3">
                  <Select v-model="newComic.subject" multiple>
                    <SelectTrigger>
                      <SelectValue
                        :placeholder="
                          newComic.subject.length ? '' : 'Select subjects'
                        "
                        :value="newComic.subject.join(', ')"
                      />
                    </SelectTrigger>
                    <SelectContent class="bg-white">
                      <SelectItem
                        v-for="subject in availableSubjects"
                        :key="subject"
                        :value="subject"
                      >
                        {{ subject }}
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Original Language</Label>
                <div class="col-span-3">
                  <Select v-model="newComic.language">
                    <SelectTrigger>
                      <SelectValue placeholder="Select language" />
                    </SelectTrigger>
                    <SelectContent class="bg-white">
                      <Input
                        v-model="languageSearchQuery"
                        @input="handleSearchInput"
                        placeholder="Search language..."
                        class="p-2 mb-2"
                      />
                      <SelectItem
                        v-for="language in filteredCountries"
                        :key="language.value"
                        :value="language.value"
                      >
                        {{ language.label }}
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Progress</Label>
                <div class="col-span-3">
                  <Select v-model="newComic.progress">
                    <SelectTrigger>
                      <SelectValue placeholder="Select progress" />
                    </SelectTrigger>
                    <SelectContent class="bg-white">
                      <SelectItem value="ongoing">Ongoing</SelectItem>
                      <SelectItem value="stopped">Completed</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Audience</Label>
                <div class="col-span-3">
                  <Select v-model="newComic.audience">
                    <SelectTrigger>
                      <SelectValue placeholder="Select audience" />
                    </SelectTrigger>
                    <SelectContent class="bg-white">
                      <SelectItem value="all">Boy's Comics</SelectItem>
                      <SelectItem value="teen">Girl's Comics</SelectItem>
                      <SelectItem value="teen">Children's Comics</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Update Time</Label>
                <div class="col-span-3">
                  <Input
                    type="datetime-local"
                    v-model="newComic.updateTime"
                    class="w-full"
                  />
                </div>
              </div>
            </div>
            <div class="space-y-4">
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Cover</Label>
                <div class="col-span-3">
                  <Input
                    type="file"
                    accept="image/*"
                    @change="handleCoverUpload"
                  />
                  <div v-if="newComic.cover" class="mt-2">
                    <img
                      :src="newComic.cover"
                      alt="Preview"
                      class="w-32 h-32 object-cover"
                    />
                  </div>
                </div>
              </div>
              <div class="grid grid-cols-4 items-center gap-4">
                <Label class="text-right">Introduction</Label>
                <Textarea
                  v-model="newComic.introduction"
                  class="col-span-3"
                  rows="6"
                />
              </div>
            </div>
          </div>
        </div>
      </template>
    </Modal>
  </div>
</template>

<script setup>
import { ref, computed, watch } from "vue";
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import Modal from "@/lib/Modal.vue";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { useComicStore } from "../../../stores/comicStore";
import { ColumnNames } from "./columnHeader";
// Define columns structure
const comicStore = useComicStore();
comicStore.getComicData();

const showAddModal = ref(false);

// New comic template
const newComic = ref({
  title: "",
  author: "",
  subject: [],
  language: "",
  progress: "",
  audience: "",
  updateTime: "",
  cover: "",
  introduction: "",
});

// Event handlers
const handleAdd = () => {
  newComic.value = {
    title: "",
    author: "",
    subject: [],
    language: "",
    progress: "",
    audience: "",
    updateTime: "",
    cover: "",
    introduction: "",
  };
  showAddModal.value = true;
};

const saveNewComic = () => {
  comics.value.push({ ...newComic.value });
  newComic.value = {
    title: "",
    author: "",
    subject: [],
    language: "",
    progress: "",
    audience: "",
    updateTime: "",
    cover: "",
    introduction: "",
  };
  showAddModal.value = false;
};

const handleEdit = (comic) => {
  console.log("Edit:", comic.title);
};

const handleDelete = (comic) => {
  comics.value = comics.value.filter((c) => c.title !== comic.title);
  console.log("Delete:", comic.title);
};

const handleCoverUpload = (event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newComic.value.cover = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};
</script>
