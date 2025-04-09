import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
import { useDramaStore } from "./dramaStore";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();
const dramaStore = useDramaStore();

export interface Episode {
  id: number;
  drama_id: number;
  number: number;
  video: string;
  active: boolean;
  created_at: string;
  created_by_name: string;
  updated_at: string;
  updated_by_name: string;
  subtitles: [];
}

export const useEpisodeStore = defineStore("episodeStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const updateActiveDialogIsOpen = ref(false);
  const playDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref({
    id: 0,
    drama_id: 0,
    number: 0,
    video: "",
    active: true,
    subtitles: [],
  });
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const episodeData = ref<Episode[]>([]);
  const sortBy = ref("");
  const sorting = ref("asc");

  async function getEpisodeData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/episodes?page=${current_page.value}&page_size=${page_size.value}&sort=${sorting.value}&sort_by=${sortBy.value}&drama_id=${dramaStore.selectedData.id}`,
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
        episodeData.value = response.data.data;
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

  async function createEpisode(data: any) {
    await axios
      .post(`${API_URL}/episodes`, data, { headers: authHeader() })
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

  async function updateEpisode(data: any) {
    await axios
      .put(`${API_URL}/episodes`, data, { headers: authHeader() })
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
      drama_id: 0,
      number: 0,
      video: "",
      active: true,
      subtitles: [],
    };
  }
  async function deleteEpisode(id: any) {
    await axios
      .delete(`${API_URL}/episodes/${id}`, {
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
      drama_id: 0,
      number: 0,
      video: "",
      active: true,
      subtitles: [],
    };
  }

  function setSorting() {
    if (sorting.value === "asc") {
      sorting.value = "desc";
    } else {
      sorting.value = "asc";
    }
  }
  async function uploadVideo(data: any, id: any): Promise<string | undefined> {
    const formData = new FormData();
    formData.append("file", data);
    formData.append("id", id);
    try {
      const response = await axios.post(
        `${API_URL}/uploads/episode-video`,
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
  async function uploadSubtitle(data: any): Promise<string | undefined> {
    const formData = new FormData();
    formData.append("file", data);
    try {
      const response = await axios.post(
        `${API_URL}/uploads/episode-sub`,
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
  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    updateActiveDialogIsOpen,
    selectedData,
    episodeData,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    isLoading,
    sortBy,
    playDialogIsOpen,
    getEpisodeData,
    createEpisode,
    updateEpisode,
    deleteEpisode,
    setSorting,
    uploadVideo,
    uploadSubtitle,
  };
});
