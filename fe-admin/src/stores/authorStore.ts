import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface Author {
  id: number;
  name: string;
  biography: string;
  birth_day: string;
  created_at: string;
  updated_at: string;
  created_by: string;
  updated_by: string;
}

export const useAuthorStore = defineStore("authorStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<Author>({
    id: 0,
    name: "",
    biography: "",
    birth_day: "",
    created_by: "",
    created_at: "",
    updated_by: "",
    updated_at: "",
  });
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const authorData = ref<Author[]>([]);
  async function getAuthorData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/author?page=${current_page.value}&page_size=${page_size.value}`,
        {
          headers: authHeader(),
        }
      )
      .then((response) => {
        authorData.value = response.data.data;
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

  async function createAuthor(data: any) {
    await axios
      .post(`${API_URL}/author`, data, { headers: authHeader() })
      .then((response) => {
        toast({
          description: "Created Successfully",
        });
      })
      .catch((error: any) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }

  async function updateAuthor(data: any) {
    await axios
      .put(`${API_URL}/author`, data, { headers: authHeader() })
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
    selectedData.value = <Author>{
      id: 0,
      name: "",
      biography: "",
      birth_day: "",
      created_by: "",
      created_at: "",
      updated_by: "",
      updated_at: "",
    };
  }
  async function deleteAuthor(id: any) {
    await axios
      .delete(`${API_URL}/author/${id}`, {
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
    selectedData.value = <Author>{
      id: 0,
      name: "",
      biography: "",
      birth_day: "",
      created_by: "",
      created_at: "",
      updated_by: "",
      updated_at: "",
    };
  }
  async function getGeneralAuthorData() {
    await axios
      .get(`${API_URL}/general/authors`, { headers: authHeader() })
      .then((response) => {
        authorData.value = response.data;
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }
  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    selectedData,
    authorData,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    isLoading,
    getAuthorData,
    createAuthor,
    updateAuthor,
    deleteAuthor,
    getGeneralAuthorData,
  };
});
