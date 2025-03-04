import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
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
  status: string;
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
  status: string;
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
    status: "active",
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

  async function getAdData() {
    isLoading.value = true;
    try {
      const params = new URLSearchParams({
        page: current_page.value.toString(),
        page_size: page_size.value.toString(),
      });

      if (sort.value.sort_by) {
        params.append('sort_by', sort.value.sort_by);
        params.append('sort_order', sort.value.sort_order);
      }

      if (searchKeyword.value) params.append('title', searchKeyword.value);
      if (filters.value.type) params.append('type', filters.value.type);
      if (filters.value.status) params.append('status', filters.value.status);

      console.log('Request params:', params.toString());

      const response = await axios.get(
        `${API_URL}/ads?${params.toString()}`,
        { headers: authHeader() }
      );

      console.log('API Response:', response.data);

      if (response.data && response.data.data && response.data.data.items) {
        adData.value = response.data.data.items;
        totalItems.value = response.data.data.pagination.total;
      } else {
        adData.value = [];
        totalItems.value = 0;
      }

      console.log('Processed Data:', adData.value);
    } catch (error: any) {
      console.error('Error fetching ads:', error);
      toast({
        description: error.response?.data?.message || error.message,
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
          status: data.status === 'active' ? 'active' : 'inactive',
          comic_id: data.type === 'internal' ? Number(data.comic_id) : null,
        },
        { headers: authHeader() }
      );
      toast({ description: "Created Successfully" });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || error.message,
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
          status: data.status === 'active' ? 'active' : 'inactive',
          comic_id: data.type === 'internal' ? Number(data.comic_id) : null,
        },
        { headers: authHeader() }
      );
      toast({ description: "Updated Successfully" });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || error.message,
        variant: "destructive",
      });
      throw error;
    }
  }

  async function deleteAd(id: number) {
    try {
      await axios.delete(`${API_URL}/ads/${id}`, { headers: authHeader() });
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
        status: "active",
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
      toast({ description: "Status Updated Successfully" });
      return response.data;
    } catch (error: any) {
      toast({
        description: error.response?.data?.message || error.message,
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
    console.log('Sort updated:', sort.value);
    getAdData();
  }

  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    selectedData,
    adData,
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
    updateSort
  };
}); 