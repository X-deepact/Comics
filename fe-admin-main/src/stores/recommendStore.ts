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
  const isLoading = ref(false);
  const selectedData = ref<Recommend | null>(null);
  const recommendData = ref([]);
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);

  async function getRecommendData() {
    isLoading.value = true;
    try {
      const params = new URLSearchParams({
        page: current_page.value.toString(),
        page_size: page_size.value.toString(),
      });

      // Only add search parameter if there's a search keyword
      if (searchKeyword.value.trim()) {
        params.append('search', searchKeyword.value.trim());
      }

      console.log('Search URL:', `${API_URL}/recommend?${params.toString()}`); // Debug log

      const response = await axios.get(
        `${API_URL}/recommend?${params.toString()}`,
        {
          headers: authHeader(),
        }
      );

      console.log('Search response:', response.data); // Debug log

      recommendData.value = response.data.data;
      current_page.value = response.data.pagination.page;
      totalItems.value = response.data.pagination.total;
      page_size.value = response.data.pagination.page_size;
    } catch (error: any) {
      console.error('Search error:', error); // Debug log
      toast({
        description: error.message,
        variant: "destructive",
      });
    } finally {
      isLoading.value = false;
    }
  }

  // Reset page when searching
  async function handleSearch(keyword: string) {
    searchKeyword.value = keyword;
    current_page.value = 1; // Reset to first page when searching
    await getRecommendData();
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
    try {
      // Format the data according to API requirements
      const requestData = {
        id: data.id,
        title: data.title,
        cover: data.cover,
        position: Number(data.position),
        active_from: data.active_from,
        active_to: data.active_to
      };

      const response = await axios.put(
        `${API_URL}/recommend`, // Remove the ID from URL
        requestData,
        { headers: authHeader() }
      );

      toast({
        description: "Updated Successfully",
      });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || "Update failed",
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
    handleSearch,
    createRecommend,
    updateRecommend,
    deleteRecommend,
  };
}); 