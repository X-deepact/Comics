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
  type: string;
  direct_url: string;
  comic_id: number;
  status: string | boolean | null;
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

  async function getAdData() {
    isLoading.value = true;
    try {
      const response = await axios.get(
        `${API_URL}/ads?page=${current_page.value}&page_size=${page_size.value}&title=${searchKeyword.value}`,
        { headers: authHeader() }
      );
      adData.value = response.data.data;
      totalItems.value = response.data.pagination.total;
      isLoading.value = false;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      isLoading.value = false;
    }
  }

  async function createAd(data: any) {
    try {
      const requestData: AdCreateRequest = {
        title: data.title,
        image: data.image,
        active_from: data.active_from,
        active_to: data.active_to || null,
        type: data.type as 'internal' | 'external',
        direct_url: data.direct_url || '',
        comic_id: data.type === 'internal' ? Number(data.comic_id) || null : null,
        status: data.status === 'active' ? 'active' : 'inactive'
      };

      console.log('Creating ad with data:', requestData);

      const response = await axios.post(
        `${API_URL}/ads`, 
        requestData, 
        { headers: authHeader() }
      );

      toast({ description: "Created Successfully" });
      return response.data;
    } catch (error: any) {
      console.error('Error creating ad:', error.response?.data || error.message);
      toast({
        description: error.response?.data?.message || error.message,
        variant: "destructive",
      });
      throw error;
    }
  }

  async function updateAd(data: any) {
    try {
      const requestData = {
        id: data.id,
        title: data.title,
        image: data.image,
        active_from: data.active_from,
        active_to: data.active_to || null,
        type: data.type as 'internal' | 'external',
        direct_url: data.direct_url || '',
        comic_id: data.type === 'internal' ? Number(data.comic_id) || null : null,
        status: data.status === 'active' ? 'active' : 'inactive'
      };

      console.log('Updating ad with data:', requestData);

      const response = await axios.put(
        `${API_URL}/ads`, 
        requestData, 
        { headers: authHeader() }
      );

      toast({ description: "Updated Successfully" });
      return response.data;
    } catch (error: any) {
      console.error('Error updating ad:', error.response?.data || error.message);
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
  };
}); 