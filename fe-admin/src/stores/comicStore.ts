import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
const API_URL = import.meta.env.VITE_API_URL;

export interface Comic {
  id: number;
  cover: string;
  title: string;
  description: string;
  active: boolean;
  code: string;
  created_by: number;
  created_at: string;
  updated_by: number;
  updated_at: string;
}

export const useComicStore = defineStore("comicStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const comicData = ref<Comic[]>([]);

  async function getComicData() {
    await axios
      .get(
        `${API_URL}/comic/comics?page=${current_page.value}&page_size=${page_size.value}`,
        { headers: authHeader() }
      )
      .then((response) => {
        comicData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total;
        page_size.value = response.data.pagination.page_size;
      });
  }
  async function createComic(data: any) {
    await axios.post(`${API_URL}/comic/create`, data).then((response) => {});
  }
  async function updateComic(data: any) {
    await axios.put(`${API_URL}/comic/update`, data).then((response) => {});
  }
  async function deleteComic(id: number) {
    await axios.delete(`${API_URL}/comic/delete/${id}`).then((response) => {});
  }
  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    comicData,
    getComicData,
    createComic,
    updateComic,
    deleteComic,
  };
});
