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
  cover: string;
  position: number;
  active_from: string;
  active_to: string;
  created_at: string;
  updated_at: string;
  created_by: number;
  updated_by: number;
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
        `${API_URL}/recommend?page=${current_page.value}&page_size=${page_size.value}&title=${searchKeyword.value}`,
        {
          headers: authHeader(),
        }
      );
      recommendData.value = response.data.data;
      current_page.value = response.data.pagination.page;
      totalItems.value = response.data.pagination.total;
      page_size.value = response.data.pagination.page_size;
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
    console.log('Creating recommend with data:', data);
    const formattedData = {
      title: data.title,
      cover: data.cover,
      position: Number(data.position),
      active_from: Math.floor(new Date(data.active_from).getTime() / 1000), // Convert to Unix timestamp
      active_to: Math.floor(new Date(data.active_to).getTime() / 1000), // Convert to Unix timestamp
    };
    console.log('Formatted data:', formattedData);

    try {
      const response = await axios.post(
        `${API_URL}/recommend`, 
        formattedData, 
        { headers: authHeader() }
      );
      console.log('Create response:', response);
      toast({
        description: "Created Successfully",
      });
      return response.data;
    } catch (error: any) {
      console.error('Create error:', error.response?.data || error);
      toast({
        description: error.response?.data?.message || "Failed to create recommendation",
        variant: "destructive",
      });
      throw error;
    }
  }

  async function updateRecommend(data: any) {
    console.log('Updating recommend with data:', data);
    const formattedData = {
      id: data.id,
      title: data.title,
      cover: data.cover,
      position: Number(data.position),
      active_from: new Date(data.active_from * 1000).toISOString().split('T')[0],
      active_to: new Date(data.active_to * 1000).toISOString().split('T')[0],
    };
    console.log('Formatted data:', formattedData);

    try {
      const response = await axios.put(
        `${API_URL}/recommend`, 
        formattedData, 
        { headers: authHeader() }
      );
      console.log('Update response:', response);
      toast({
        description: "Updated Successfully",
      });
      return response.data;
    } catch (error: any) {
      console.error('Update error:', error.response?.data || error);
      toast({
        description: error.response?.data?.message || "Failed to update recommendation",
        variant: "destructive",
      });
      throw error;
    }
  }

  async function deleteRecommend(id: number) {
    await axios
      .delete(`${API_URL}/recommend/${id}`, {
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
  }

  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    selectedData,
    recommendData,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    isLoading,
    getRecommendData,
    createRecommend,
    updateRecommend,
    deleteRecommend,
  };
}); 