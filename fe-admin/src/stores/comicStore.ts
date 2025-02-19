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
  language: string;
  audience: string;
  created_by: number;
  created_at: string;
  updated_by: number;
  updated_at: string;
  created_by_user: {};
  updated_by_user: {};
}

export const useComicStore = defineStore("comicStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<Comic>({
    id: 0,
    cover: "",
    title: "",
    description: "",
    active: false,
    code: "",
    language: "",
    audience: "",
    created_by: 0,
    created_at: "",
    updated_by: 0,
    updated_at: "",
    created_by_user: {},
    updated_by_user: {},
  });
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const comicData = ref<Comic[]>([]);

  async function getComicData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/comics?page=${current_page.value}&page_size=${page_size.value}&q=&sort_by=&sort=asc&active=true&language=&audience=`,
        { headers: authHeader() }
      )
      .then((response) => {
        comicData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total;
        page_size.value = response.data.pagination.page_size;
        isLoading.value = false;
      });
  }
  async function createComic(data: any) {
    await axios
      .post(`${API_URL}/comics`, data, { headers: authHeader() })
      .then((response) => {});
  }
  async function updateComic(data: any) {
    await axios
      .put(`${API_URL}/comics`, data, { headers: authHeader() })
      .then((response) => {});
    selectedData.value = <Comic>{
      id: 0,
      cover: "",
      title: "",
      description: "",
      active: false,
      code: "",
      language: "",
      audience: "",
      created_by: 0,
      created_at: "",
      updated_by: 0,
      updated_at: "",
      created_by_user: {},
      updated_by_user: {},
    };
  }
  async function deleteComic(id: any) {
    await axios
      .delete(`${API_URL}/comics/${id}`, {
        headers: authHeader(),
      })
      .then((response) => {});
    selectedData.value = <Comic>{
      id: 0,
      cover: "",
      title: "",
      description: "",
      active: false,
      code: "",
      language: "",
      audience: "",
      created_by: 0,
      created_at: "",
      updated_by: 0,
      updated_at: "",
      created_by_user: {},
      updated_by_user: {},
    };
  }
  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    isLoading,
    selectedData,
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
