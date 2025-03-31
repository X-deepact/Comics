import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
import { type DateValue } from "@internationalized/date";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface Episode {}

export const useEpisodeStore = defineStore("episodeStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref({});
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const episodeData = ref<Episode[]>([]);
  const sortBy = ref("");
  const sorting = ref("asc");
  async function getDramaData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/episodes?page=${current_page.value}&page_size=${page_size.value}&sort=${sorting.value}&sort_by=${sortBy.value}`,
        {
          headers: authHeader(),
        }
      )
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        dramaData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total
          ? response.data.pagination.total
          : 0;
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
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
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
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
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
    selectedData.value = {
      id: 0,
      name: "",
      biography: "",
      birth_day: null as DateValue | null,
    };
  }
  async function deleteAuthor(id: any) {
    await axios
      .delete(`${API_URL}/author/${id}`, {
        headers: authHeader(),
      })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
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
      name: "",
      biography: "",
      birth_day: null as DateValue | null,
    };
  }

  function setSorting() {
    if (sorting.value === "asc") {
      sorting.value = "desc";
    } else {
      sorting.value = "asc";
    }
  }
  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    selectedData,
    dramaData,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    isLoading,
    sortBy,
    getDramaData,
    createAuthor,
    updateAuthor,
    deleteAuthor,
    setSorting,
  };
});
