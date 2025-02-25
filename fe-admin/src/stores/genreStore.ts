import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface Genre {
  id: number;
  name: string;
  position: number;
  lang: string;
  created_at: string;
  updated_at: string;
  created_by_name: string;
  updated_by_name: string;
}
export const useGenreStore = defineStore("genreStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<Genre>({
    id: 0,
    name: "",
    position: 0,
    lang: "",
    created_at: "",
    updated_at: "",
    created_by_name: "",
    updated_by_name: "",
  });
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const genreData = ref<Genre[]>([]);

  async function getGenreData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/genre?page=${current_page.value}&page_size=${page_size.value}&name=&language=`,
        { headers: authHeader() }
      )
      .then((response) => {
        genreData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total;
        page_size.value = response.data.pagination.page_size;
        isLoading.value = false;
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }
  async function createGenre(data: any) {
    await axios
      .post(`${API_URL}/genre`, data, { headers: authHeader() })
      .then((response) => {
        toast({
          description: "Created successfully",
        });
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }
  async function updateGenre(data: any) {
    await axios
      .put(`${API_URL}/genre`, data, { headers: authHeader() })
      .then((response) => {
        toast({
          description: "Updated successfully",
        });
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
    selectedData.value = <Genre>{
      id: 0,
      name: "",
      position: 0,
      lang: "",
      created_at: "",
      updated_at: "",
      created_by_name: "",
      updated_by_name: "",
    };
  }
  async function deleteGenre(id: any) {
    await axios
      .delete(`${API_URL}/genre/${id}`, { headers: authHeader() })
      .then(() => {
        toast({
          description: "Deleted successfully",
        });
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
    selectedData.value = <Genre>{
      id: 0,
      name: "",
      position: 0,
      lang: "",
      created_at: "",
      updated_at: "",
      created_by_name: "",
      updated_by_name: "",
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
    genreData,
    getGenreData,
    createGenre,
    updateGenre,
    deleteGenre,
  };
});
