import { defineStore } from "pinia";
import axios from "axios";
import { ref, computed } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export enum RecommendPosition {
  COMPLETE_MASTERPIECE = 10,
  FASTEST_RISING = 20,
  NEW_PUBLISHING = 30,
  RECENTLY_UPDATE = 40
}

export const positionLabels: { [key: number]: string } = {
  [RecommendPosition.COMPLETE_MASTERPIECE]: "Complete Masterpiece",
  [RecommendPosition.FASTEST_RISING]: "Fastest Rising",
  [RecommendPosition.NEW_PUBLISHING]: "New Publishing",
  [RecommendPosition.RECENTLY_UPDATE]: "Recently Update"
};

export interface Comic {
  id: number;
  name: string;
}

export interface Recommend {
  id: number;
  title: string;
  cover: string;
  position: RecommendPosition;
  active_from: string;
  active_to: string;
  created_at: string;
  updated_at: string;
  created_by: string;
  updated_by: string;
  comics?: Comic[];
}

export interface ApiResponse<T> {
  data: T;
  message: string;
  status: string;
}

export interface PaginationResponse<T> {
  data: T[];
  pagination: {
    total: number;
    page: number;
    page_size: number;
  };
}

export interface RecommendCreateRequest {
  title: string;
  cover: string | File;
  position: number;
  active_from: number;  // Unix timestamp in seconds
  active_to: number;    // Unix timestamp in seconds
}

export interface RecommendUpdateRequest {
  id: number;
  title: string;
  cover: string;
  position: number;
  active_from: number;  // Unix timestamp in seconds
  active_to: number;    // Unix timestamp in seconds
}

export interface RecommendComicRequest {
  recommend_id: number;
  comic_id: number;
}

export const useRecommendStore = defineStore("recommendStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(false);
  const selectedData = ref<Recommend | null>(null);
  const recommendData = ref<Recommend[]>([]);
  const current_page = ref(1);
  const page_size = ref(10);
  const totalItems = ref(0);
  const searchKeyword = ref("");
  const sortBy = ref("updated_at");
  const sortDirection = ref<"asc" | "desc">("desc");

  async function getRecommendData() {
    isLoading.value = true;
    try {
      const params: Record<string, any> = {
        page: current_page.value,
        page_size: page_size.value,
        sort: sortDirection.value,
        sort_by: sortBy.value
      };

      if (searchKeyword.value.trim()) {
        params.title = searchKeyword.value.trim();
      }

      const response = await axios.get(`${API_URL}/recommend`, {
        params: params,
        headers: authHeader(),
      });

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      recommendData.value = response.data.data;
      if (response.data.pagination) {
        totalItems.value = response.data.pagination.total || 0;
        current_page.value = response.data.pagination.page;
        page_size.value = response.data.pagination.page_size;
      }
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      recommendData.value = [];
      totalItems.value = 0;
    } finally {
      isLoading.value = false;
    }
  }

  function handleSort(sortKey: string) {
    if (sortBy.value === sortKey) {
      sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
    } else {
      sortBy.value = sortKey;
      sortDirection.value = 'desc';
    }
    getRecommendData();
  }

  async function getRecommendById(id: number) {
    try {
      const response = await axios.get<ApiResponse<Recommend>>(
        `${API_URL}/recommend/${id}`,
        { headers: authHeader() }
      );
      return response.data.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || error.message,
        variant: "destructive",
      });
      throw error;
    }
  }

  async function createRecommend(data: RecommendCreateRequest) {
    try {
      const formData = new FormData();
      
      formData.append('title', data.title);
      formData.append('position', data.position.toString());
      formData.append('active_from', (data.active_from * 1000).toString());
      formData.append('active_to', (data.active_to * 1000).toString());

      if (data.cover instanceof File) {
        formData.append('cover', data.cover);
      } else {
        formData.append('cover', data.cover.toString());
      }

      const response = await axios.post(
        `${API_URL}/recommend`,
        formData,
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      toast({
        description: "Created Successfully",
      });
      return response.data.data;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      throw error;
    }
  }

  async function updateRecommend(data: RecommendUpdateRequest) {
    try {
      const formData = new FormData();
      
      formData.append('id', data.id.toString());
      formData.append('title', data.title);
      formData.append('position', data.position.toString());
      formData.append('active_from', (data.active_from * 1000).toString());
      formData.append('active_to', (data.active_to * 1000).toString());

      // Only append cover if it's a new file
      if (data.cover instanceof File) {
        formData.append('cover', data.cover);
      }
      // Don't append anything for cover if using existing image

      
      const response = await axios.put<ApiResponse<Recommend>>(
        `${API_URL}/recommend`,
        formData,
        { 
          headers: {
            ...authHeader(),
            'accept': 'application/json',
          }
        }
      );
      
      toast({
        description: response.data.message || "Updated Successfully",
      });
      return response.data.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || "Failed to update recommendation",
        variant: "destructive",
      });
      throw error;
    }
  }

  async function deleteRecommend(id: number) {
    try {
      const response = await axios.delete<ApiResponse<null>>(
        `${API_URL}/recommend/${id}`,
        { headers: authHeader() }
      );
      toast({
        description: response.data.message || "Deleted Successfully",
      });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || "Failed to delete recommendation",
        variant: "destructive",
      });
      throw error;
    }
  }

  async function createRecommendComic(recommendId: number, comicId: number) {
    try {
      const requestData: RecommendComicRequest = {
        recommend_id: recommendId,
        comic_id: comicId
      };
      const response = await axios.post<ApiResponse<{ id: number }>>(
        `${API_URL}/recommend/comic`,
        requestData,
        { headers: authHeader() }
      );
      toast({
        description: response.data.message || "Comic added successfully",
      });
      return response.data.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || "Failed to add comic",
        variant: "destructive",
      });
      throw error;
    }
  }

  async function deleteRecommendComic(recommendId: number, comicId: number) {
    try {
      const requestData: RecommendComicRequest = {
        recommend_id: recommendId,
        comic_id: comicId
      };
      const response = await axios.delete<ApiResponse<{ id: number }>>(
        `${API_URL}/recommend/comic`,
        { 
          headers: authHeader(),
          data: requestData
        }
      );
      toast({
        description: response.data.message || "Comic removed successfully",
      });
      return response.data.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || "Failed to remove comic",
        variant: "destructive",
      });
      throw error;
    }
  }

  function handleSearch(keyword: string) {
    searchKeyword.value = keyword;
    current_page.value = 1; // Reset to first page when searching
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
    getRecommendById,
    createRecommend,
    updateRecommend,
    deleteRecommend,
    createRecommendComic,
    deleteRecommendComic,
    handleSort,
    handleSearch,
  };
}); 
