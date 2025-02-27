import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";

const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface Recommend {
  id: number;
  title: string;
  cover: string | File;
  position: number;
  active_from: string;
  active_to: string;
  created_by: number;
  created_at: string;
  updated_by: number;
  updated_at: string;
}

export const useRecommendStore = defineStore("recommendStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<Recommend>({
    id: 0,
    title: "",
    cover: "",
    position: 0,
    active_from: "",
    active_to: "",
    created_by: 0,
    created_at: "",
    updated_by: 0,
    updated_at: "",
  });

  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const recommendData = ref<Recommend[]>([]);

  async function getRecommendData() {
    isLoading.value = true;
    try {
      const response = await axios.get(
        `${API_URL}/recommend?page=${current_page.value}&page_size=${page_size.value}&q=${searchKeyword.value}`,
        { headers: authHeader() }
      );
      
      // Assuming the API returns { data: [], pagination: { total, page, page_size } }
      if (response.data.data) {
        recommendData.value = response.data.data;
        totalItems.value = response.data.pagination?.total || response.data.data.length;
        current_page.value = response.data.pagination?.page || 1;
        page_size.value = response.data.pagination?.page_size || 10;
      } else {
        // If API returns array directly
        recommendData.value = response.data;
        totalItems.value = response.data.length;
      }
      
      isLoading.value = false;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      isLoading.value = false;
    }
  }

  async function createRecommend(data: any) {
    const formData = new FormData();
    if (data.cover instanceof File) {
      formData.append("cover", data.cover);
    }
    formData.append("title", data.title);
    formData.append("position", data.position.toString());
    formData.append("active_from", new Date(data.active_from).toISOString());
    formData.append("active_to", new Date(data.active_to).toISOString());

    try {
      await axios.post(`${API_URL}/recommend`, formData, {
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

  async function updateRecommend(data: any) {
    try {
      const updateData = {
        id: data.id,
        title: data.title,
        position: data.position,
        active_from: data.active_from ? new Date(data.active_from).toISOString().split('T')[0] : null,
        active_to: data.active_to ? new Date(data.active_to).toISOString().split('T')[0] : null,
      };

      if (data.cover instanceof File) {
        const formData = new FormData();
        formData.append("cover", data.cover);
        Object.entries(updateData).forEach(([key, value]) => {
          if (value !== null) {
            formData.append(key, value.toString());
          }
        });

        await axios.put(`${API_URL}/recommend`, formData, {
          headers: { ...authHeader(), "Content-Type": "multipart/form-data" },
        });
      } else {
        await axios.put(`${API_URL}/recommend`, updateData, { 
          headers: authHeader() 
        });
      }

      toast({
        description: "Updated Successfully",
      });
      resetSelectedData();
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
    }
  }

  async function deleteRecommend(id: number) {
    try {
      await axios.delete(`${API_URL}/recommend/${id}`, {
        headers: authHeader(),
      });
      toast({
        description: "Deleted Successfully",
      });
      selectedData.value = {
        id: 0,
        title: "",
        cover: "",
        position: 0,
        active_from: "",
        active_to: "",
        created_by: 0,
        created_at: "",
        updated_by: 0,
        updated_at: "",
      };
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
    }
  }

  function resetSelectedData() {
    selectedData.value = {
      id: 0,
      title: "",
      cover: "",
      position: 0,
      active_from: "",
      active_to: "",
      created_by: 0,
      created_at: "",
      updated_by: 0,
      updated_at: "",
    };
  }

  function setSelectedDataForUpdate(data: any) {
    selectedData.value = {
      ...data,
      active_from: data.active_from ? new Date(data.active_from).toISOString().split('T')[0] : '',
      active_to: data.active_to ? new Date(data.active_to).toISOString().split('T')[0] : '',
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
    recommendData,
    getRecommendData,
    createRecommend,
    updateRecommend,
    deleteRecommend,
    setSelectedDataForUpdate,
  };
}); 