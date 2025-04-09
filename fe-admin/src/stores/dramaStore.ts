import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface Drama {
  id: number;
  release_date: string;
  thumb: string;
  active: boolean;
  created_at: string;
  created_by_name: string;
  updated_at: string;
  updated_by_name: string;
  translations: [];
  genres: [];
}

export const useDramaStore = defineStore("dramaStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const updateActiveDialogIsOpen = ref(false);
  const episodeDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref({
    id: 0,
    thumb: undefined as string | undefined,
    release_date: "",
    translations: [] as {
      language: string;
      title: string;
      description: string;
    }[],
    genres: [],
  });
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const dramaData = ref<Drama[]>([]);
  const generalDramaGenreData = ref([]);
  const sortBy = ref("");
  const sorting = ref("asc");
  async function getDramaData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/dramas?page=${current_page.value}&page_size=${page_size.value}&sort=${sorting.value}&sort_by=${sortBy.value}`,
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

  async function createDrama(data: any) {
    const sendData = {
      thumb: data.thumb,
      translations: data.translations,
      release_date: data.release_date.split("T")[0],
      genres: data.genres.map((genre: any) => genre.id),
    };
    await axios
      .post(`${API_URL}/dramas`, sendData, { headers: authHeader() })
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

  async function updateDrama(data: any) {
    const send_data = {
      id: data.id,
      release_date: data.release_date,
      thumb: data.thumb,
      translations: data.translations,
      genres: data.genres.map((genre: any) => genre.id),
    };
    await axios
      .put(`${API_URL}/dramas`, send_data, { headers: authHeader() })
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
      thumb: undefined,
      release_date: "",
      translations: [],
      genres: [],
    };
  }
  async function deleteDrama(id: any) {
    await axios
      .delete(`${API_URL}/dramas/${id}`, {
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
      thumb: undefined,
      release_date: "",
      translations: [],
      genres: [],
    };
  }
  async function setActiveDrama(id: number) {
    await axios
      .put(`${API_URL}/dramas/active/${id}`, null, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        toast({
          description: "Updated active status successfully",
        });
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }
  async function uploadImage(data: any): Promise<string | undefined> {
    const formData = new FormData();
    formData.append("file", data);
    try {
      const response = await axios.post(
        `${API_URL}/uploads/drama-thumb`,
        formData,
        {
          headers: { ...authHeader(), "Content-Type": "multipart/form-data" },
        }
      );
      if (response.data.code == "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return undefined;
      }

      return response.data.data;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      return undefined;
    }
  }
  function setSorting() {
    if (sorting.value === "asc") {
      sorting.value = "desc";
    } else {
      sorting.value = "asc";
    }
  }
  async function getGeneralGenreData() {
    await axios
      .get(`${API_URL}/general/drama-genres`, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        generalDramaGenreData.value = response.data.data;
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
    updateActiveDialogIsOpen,
    episodeDialogIsOpen,
    selectedData,
    dramaData,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    isLoading,
    sortBy,
    generalDramaGenreData,
    getGeneralGenreData,
    uploadImage,
    getDramaData,
    createDrama,
    updateDrama,
    deleteDrama,
    setActiveDrama,
    setSorting,
  };
});
