import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
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
  const isLoading = ref(true);
  const selectedData = ref<Comic>({
    id: 0,
    cover: "",
    name: "",
    description: "",
    active: false,
    code: "",
    lang: "",
    audience: "",
    created_by: 0,
    created_at: "",
    updated_by: 0,
    updated_at: "",
    created_by_user: {},
    updated_by_user: {},
  });
  const selectedAuthorForCreate = ref({});
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const comicData = ref<Comic[]>([]);

  async function getComicData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/comics?page=${current_page.value}&page_size=${page_size.value}&q=${searchKeyword.value}&sort_by=&sort=asc&active=true&language=&audience=`,
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
  async function createComic(data: any) {
    const formData = new FormData();
    formData.append("cover", data.cover);
    formData.append("name", data.title);
    formData.append("description", data.description);
    formData.append("active", data.active);
    formData.append("code", data.code);
    formData.append("lang", data.language);
    formData.append("audience", data.audience);
    [1, 2, 3].forEach((genre) => formData.append("genres", genre.toString()));
    [1, 2, 3].forEach((author) =>
      formData.append("authors", author.toString())
    );

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
    await axios
      .put(`${API_URL}/comics`, data, { headers: authHeader() })
      .then((response) => {
        toast({
          description: "Updated Successfully",
        });
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
    selectedData.value = <Comic>{
      id: 0,
      cover: "",
      name: "",
      description: "",
      active: false,
      code: "",
      lang: "",
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
    selectedData.value = <Comic>{
      id: 0,
      cover: "",
      name: "",
      description: "",
      active: false,
      code: "",
      lang: "",
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
    selectedAuthorForCreate,
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
