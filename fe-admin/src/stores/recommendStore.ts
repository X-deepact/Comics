import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export enum RecommendPosition {
  COMPLETE_MASTERPIECE = 1,
  FASTEST_RISING = 2,
  NEW_PUBLISHING = 3,
  RECENTLY_UPDATE = 4
}

export const positionLabels = {
  [RecommendPosition.COMPLETE_MASTERPIECE]: "Complete Masterpiece",
  [RecommendPosition.FASTEST_RISING]: "Fastest Rising",
  [RecommendPosition.NEW_PUBLISHING]: "New Publishing",
  [RecommendPosition.RECENTLY_UPDATE]: "Recently Update"
};

export interface Recommend {
  id: number;
  title: string;
  cover: string;
  position: RecommendPosition;
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
  const sortBy = ref("updated_at"); // Change default sort to updated_at
  const sortDirection = ref("desc"); // Default to newest first

  async function getRecommendData() {
    isLoading.value = true;
    try {
      const params = new URLSearchParams({
        page: current_page.value.toString(),
        page_size: page_size.value.toString(),
        sort_by: sortBy.value,
        sort: sortDirection.value,
      });

      console.log('Request URL:', `${API_URL}/recommend?${params.toString()}`);

      const response = await axios.get(
        `${API_URL}/recommend?${params.toString()}`,
        {
          headers: authHeader(),
        }
      );

      // If searching, filter results client-side
      if (searchKeyword.value.trim()) {
        recommendData.value = response.data.data.filter((item: any) => 
          item.title.toLowerCase().includes(searchKeyword.value.toLowerCase())
        );
      } else {
        recommendData.value = response.data.data;
      }

      current_page.value = response.data.pagination.page;
      totalItems.value = response.data.pagination.total;
      page_size.value = response.data.pagination.page_size;
    } catch (error: any) {
      console.error('API Error:', error.response || error);
      toast({
        description: error.response?.data?.message || error.message,
        variant: "destructive",
      });
    } finally {
      isLoading.value = false;
    }
  }

  // Debounced search function
  let searchTimeout: NodeJS.Timeout;
  async function handleSearch(keyword: string) {
    if (searchTimeout) {
      clearTimeout(searchTimeout);
    }

    searchTimeout = setTimeout(() => {
      searchKeyword.value = keyword;
      current_page.value = 1;
      getRecommendData();
    }, 300);
  }

  async function createRecommend(data: any) {
    console.log('Creating recommend with data:', data);
    const formattedData = {
      title: data.title,
      cover: data.cover,
      position: Number(data.position),
      active_from: Math.floor(new Date(data.active_from).getTime() / 1000),
      active_to: Math.floor(new Date(data.active_to).getTime() / 1000),
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
      console.log('Raw update data received:', data);

      if (!data.id) {
        throw new Error('Missing ID for update');
      }

      // Format the dates as ISO strings instead of Unix timestamps
      const requestData = {
        id: Number(data.id),
        title: data.title?.trim(),
        cover: data.cover || '',
        position: Number(data.position),
        // Format dates as ISO strings (YYYY-MM-DD)
        active_from: data.active_from,
        active_to: data.active_to
      };

      console.log('Formatted request data:', requestData);
      console.log('Request URL:', `${API_URL}/recommend`);

      const response = await axios.put(
        `${API_URL}/recommend`,
        requestData,
        { 
          headers: {
            ...authHeader(),
            'Content-Type': 'application/json'
          }
        }
      );

      console.log('Update response:', response);
      toast({
        description: "Updated Successfully",
      });
      return response.data;
    } catch (error: any) {
      console.error('Update error details:', {
        message: error.message,
        response: error.response?.data,
        status: error.response?.status,
        requestData: error.config?.data,
      });

      // More specific error message
      const errorMessage = error.response?.data?.message 
        || error.response?.data?.error 
        || "Failed to update recommendation";
      
      toast({
        description: errorMessage,
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

  // Add this new function to handle sorting
  function handleSort(sortKey: string) {
    if (sortBy.value === sortKey) {
      // Toggle direction if clicking same column
      sortDirection.value = sortDirection.value === "asc" ? "desc" : "asc";
    } else {
      // New column, set as default desc
      sortBy.value = sortKey;
      sortDirection.value = "desc";
    }
    current_page.value = 1; // Reset to first page
    getRecommendData();
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
    sortBy,
    sortDirection,
    getRecommendData,
    handleSearch,
    createRecommend,
    updateRecommend,
    deleteRecommend,
    handleSort,
  };
}); 