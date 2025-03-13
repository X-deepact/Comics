import { defineStore } from "pinia";
import axios from "axios";
import { ref, computed } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";

const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface Ad {
  id: number;
  title: string;
  image: string;
  active_from: string;
  active_to: string;
  type: 'internal' | 'external';
  direct_url: string;
  comic_id: number;
  status: 'active' | 'inactive';
  created_at: string;
  updated_at: string;
  created_by_name: string;
  updated_by_name: string;
}

export interface AdCreateRequest {
  title: string;
  image: string;
  active_from: string;
  active_to: string | null;
  type: 'internal' | 'external';
  direct_url: string;
  comic_id: number | null;
  status: 'active' | 'inactive';
}

export interface AdFilters {
  title?: string;
  type?: 'internal' | 'external';
  status?: 'active' | 'inactive';
}

interface SortState {
  sort_by: string;
  sort_order: 'asc' | 'desc';
}

export const useAdStore = defineStore("adStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<Ad>({
    id: 0,
    title: "",
    image: "",
    active_from: "",
    active_to: "",
    type: "internal",
    direct_url: "",
    comic_id: 0,
    status: "active" as 'active' | 'inactive',
    created_at: "",
    updated_at: "",
    created_by_name: "",
    updated_by_name: "",
  });
  
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const adData = ref<Ad[]>([]);
  const filters = ref<AdFilters>({});

  const sort = ref<SortState>({
    sort_by: 'id',
    sort_order: 'desc'
  });

  // Add computed property for sorted data
  const sortedAdData = computed(() => {
    if (!sort.value.sort_by) return adData.value;

    return [...adData.value].sort((a, b) => {
      const sortField = sort.value.sort_by === 'updated_by' ? 'updated_by_name' : sort.value.sort_by;
      const aValue = a[sortField as keyof typeof a] || '';
      const bValue = b[sortField as keyof typeof b] || '';

      if (sort.value.sort_order === 'asc') {
        return aValue < bValue ? -1 : aValue > bValue ? 1 : 0;
      } else {
        return bValue < aValue ? -1 : bValue > aValue ? 1 : 0;
      }
    });
  });

  async function getAdData() {
    isLoading.value = true;
    try {
      const params = new URLSearchParams({
        page: current_page.value.toString(),
        page_size: page_size.value.toString(),
      });

      if (searchKeyword.value) params.append('title', searchKeyword.value);
      if (filters.value.type) params.append('type', filters.value.type);
      if (filters.value.status) params.append('status', filters.value.status);
      
      if (sort.value.sort_by === 'updated_at') {
        params.append('sort_by', 'updated_at');
        params.append('sort_order', sort.value.sort_order.toUpperCase());
      }

      const response = await axios.get(
        `${API_URL}/ads?${params.toString()}`,
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      adData.value = response.data.data;
      current_page.value = response.data.pagination.page;
      totalItems.value = response.data.pagination.total || 0;
      page_size.value = response.data.pagination.page_size;

    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      adData.value = [];
      totalItems.value = 0;
    } finally {
      isLoading.value = false;
    }
  }

  async function createAd(data: AdCreateRequest) {
    try {
      const response = await axios.post(
        `${API_URL}/ads`,
        {
          ...data,
          comic_id: data.type === 'internal' ? Number(data.comic_id) : null,
        },
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      toast({ description: "Created Successfully" });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      throw error;
    }
  }

  async function updateAd(data: Ad) {
    try {
      const response = await axios.put(
        `${API_URL}/ads`,
        {
          ...data,
          comic_id: data.type === 'internal' ? Number(data.comic_id) : null,
        },
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      toast({ description: "Updated Successfully" });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      throw error;
    }
  }

  async function deleteAd(id: number) {
    try {
      const response = await axios.delete(
        `${API_URL}/ads/${id}`, 
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      toast({ description: "Deleted Successfully" });
      selectedData.value = {
        id: 0,
        title: "",
        image: "",
        active_from: "",
        active_to: "",
        type: "internal",
        direct_url: "",
        comic_id: 0,
        status: "active" as 'active' | 'inactive',
        created_at: "",
        updated_at: "",
        created_by_name: "",
        updated_by_name: "",
      };
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
    }
  }

  async function updateAdStatus(id: number, status: 'active' | 'inactive') {
    try {
      const response = await axios.patch(
        `${API_URL}/ads/${id}/status`,
        { status },
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      toast({ description: "Status Updated Successfully" });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      throw error;
    }
  }

  function updateSort(sortKey: string) {
    if (sort.value.sort_by === sortKey) {
      sort.value.sort_order = sort.value.sort_order === 'asc' ? 'desc' : 'asc';
    } else {
      sort.value.sort_by = sortKey;
      sort.value.sort_order = 'desc';
    }
  }

  const uploadImage = async (file: File) => {
    try {
      const formData = new FormData();
      formData.append('image', file);

      const response = await axios.post(
        `${API_URL}/upload-image`,
        formData,
        { 
          headers: { 
            ...authHeader(),
            'Content-Type': 'multipart/form-data'
          } 
        }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return null;
      }

      return response.data.filename; // Adjust based on your API response structure
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || error.message,
        variant: "destructive",
      });
      throw error;
    }
  };

  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    selectedData,
    adData: sortedAdData,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    isLoading,
    getAdData,
    createAd,
    updateAd,
    deleteAd,
    filters,
    updateAdStatus,
    sort,
    updateSort,
    uploadImage
  };
}); 
