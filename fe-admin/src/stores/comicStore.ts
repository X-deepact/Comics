import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
import type { Author } from "./authorStore";
import type { Genre } from "./genreStore";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();
export interface Comic {
  id: number;
  cover: string;
  name: string;
  description: string;
  active: boolean;
  code: string;
  lang: string;
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
  const chapterDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref({
    id: 0,
    cover: "",
    name: "",
    description: "",
    active: false,
    code: "",
    lang: "",
    audience: "",
    status: "",
    authors: [] as Author[],
    genres: [] as Genre[],
  });
  const selectedAuthorForCreate = ref({});
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const comicData = ref<Comic[]>([]);
  const sortBy = ref("");
  const sorting = ref("asc");

  async function getComicData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/comics?page=${current_page.value}&page_size=${page_size.value}&q=${searchKeyword.value}&sort_by=${sortBy.value}&sort=${sorting.value}&language=&audience=`,
        { headers: authHeader() }
      )
      .then((response) => {
        comicData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total;
        page_size.value = response.data.pagination.page_size;
        isLoading.value = false;
      })
      .catch((error: any) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }
  function setSorting() {
    if (sorting.value === "asc") {
      sorting.value = "desc";
    } else {
      sorting.value = "asc";
    }
  }
  async function createComic(data: any) {
    const formData = new FormData();
    formData.append("cover", data.cover);
    formData.append("name", data.title);
    formData.append("description", data.description);
    formData.append("active", data.active);
    formData.append("code", data.code);
    formData.append("lang", data.language);
    formData.append("audience", data.audience);
    formData.append("status", data.status);
    data.author.forEach((author: any) => formData.append("authors", author.id));
    data.genre.forEach((genre: any) => formData.append("genres", genre.id));

    try {
      await axios.post(`${API_URL}/comics`, formData, {
        headers: { ...authHeader(), "Content-Type": "multipart/form-data" },
      });
      toast({
        description: "Created successfully",
      });
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
    }
  }
  async function updateComic(data: any) {
    const formData = new FormData();
    formData.append("id", data.id);
    formData.append("cover", data.cover);
    formData.append("name", data.name);
    formData.append("description", data.description);
    formData.append("active", data.active);
    formData.append("code", data.code);
    formData.append("lang", data.lang);
    formData.append("audience", data.audience);
    formData.append("status", data.status);
    data.author.forEach((author: any) => formData.append("authors", author.id));
    data.genre.forEach((genre: any) => formData.append("genres", genre.id));

    try {
      await axios.put(`${API_URL}/comics`, formData, {
        headers: { ...authHeader(), "Content-Type": "multipart/form-data" },
      });
      toast({
        description: "Updated successfully",
      });
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
    }
    selectedData.value = {
      id: 0,
      cover: "",
      name: "",
      description: "",
      active: false,
      code: "",
      lang: "",
      audience: "",
      status: "",
      authors: [],
      genres: [],
    };
  }
  async function deleteComic(id: any) {
    await axios
      .delete(`${API_URL}/comics/${id}`, {
        headers: authHeader(),
      })
      .then((response) => {
        toast({
          description: "Deleted Successfully",
        });
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
    selectedData.value = {
      id: 0,
      cover: "",
      name: "",
      description: "",
      active: false,
      code: "",
      lang: "",
      audience: "",
      status: "",
      authors: [],
      genres: [],
    };
  }
  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    chapterDialogIsOpen,
    isLoading,
    selectedData,
    selectedAuthorForCreate,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    comicData,
    sortBy,
    setSorting,
    getComicData,
    createComic,
    updateComic,
    deleteComic,
  };
});
