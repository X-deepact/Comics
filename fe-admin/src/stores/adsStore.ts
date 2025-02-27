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
  comic_id?: number;
  status: string;
  created_at: string;
  updated_at: string;
  created_by_name: string;
  updated_by_name: string;
}

export interface AdUpdateRequest {
  id: number;
  title: string;
  active_from: string;
  active_to: string;
  type: string;
  direct_url: string;
  status: string;
  comic_id?: number;
  image?: File;
}

export interface AdCreateRequest {
  title: string;
  image: File;
  type: string;
  direct_url: string;
  active_from: string;
  active_to: string;
  status: string;
  comic_id?: number;
}

export const useAdsStore = defineStore("adsStore", () => {
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
  const adsData = ref<Ad[]>([]);

  async function getAdsData() {
    isLoading.value = true;
    try {
      const response = await axios.get(
        `${API_URL}/ads?page=${current_page.value}&page_size=${page_size.value}&title=${searchKeyword.value}`,
        { headers: authHeader() }
      );
      adsData.value = response.data.data;
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

  async function createAd(data: AdCreateRequest) {
    try {
      // Basic validation
      if (!data.title || !data.image || !data.active_from || !data.active_to) {
        throw new Error("Missing required fields");
      }

      // Convert image to base64
      const base64Image = await new Promise<string>((resolve) => {
        const reader = new FileReader();
        reader.onloadend = () => {
          const base64String = reader.result as string;
          resolve(base64String.split(',')[1]); // Remove data URL prefix
        };
        reader.readAsDataURL(data.image);
      });

      // Format dates as RFC3339 strings
      const formatDate = (dateStr: string) => {
        const date = new Date(dateStr);
        return date.toISOString();
      };

      // Prepare request data
      const requestData = {
        title: data.title,
        type: data.type.toLowerCase(),
        direct_url: data.direct_url,
        active_from: formatDate(data.active_from),
        active_to: formatDate(data.active_to),
        status: data.status.toLowerCase(),
        image: base64Image,
        comic_id: data.comic_id || null
      };

      // Log the request data (excluding image for brevity)
      console.log('Request Data:', {
        ...requestData,
        image: '[base64 string]'
      });

      const response = await axios.post(`${API_URL}/ads`, requestData, {
        headers: {
          ...authHeader(),
          'Content-Type': 'application/json'
        },
      });

      console.log('Response:', response.data);

      toast({
        description: "Created Successfully",
      });
    } catch (error: any) {
      console.error('Create Ad Error:', error.response?.data || error);
      toast({
        description: error.response?.data?.message || "Failed to create ad",
        variant: "destructive",
      });
    }
  }

  async function updateAd(data: AdUpdateRequest) {
    try {
      // Format the request data
      const requestData = {
        id: data.id,
        title: data.title,
        type: data.type,
        direct_url: data.direct_url,
        active_from: new Date(data.active_from).toISOString(),
        active_to: new Date(data.active_to).toISOString(),
        status: data.status.toLowerCase(), // Ensure status is lowercase
        comic_id: data.comic_id || null,
      };

      if (data.image instanceof File) {
        const formData = new FormData();
        Object.entries(requestData).forEach(([key, value]) => {
          if (value !== null) {
            formData.append(key, value.toString());
          }
        });
        formData.append("image", data.image);

        await axios.put(`${API_URL}/ads`, formData, {
          headers: { 
            ...authHeader(), 
            "Content-Type": "multipart/form-data" 
          },
        });
      } else {
        await axios.put(`${API_URL}/ads`, requestData, { 
          headers: { 
            ...authHeader(),
            'Content-Type': 'application/json'
          } 
        });
      }

      toast({
        description: "Updated Successfully",
      });
      
      // Reset selected data
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
      console.error('Update Ad Error:', error.response?.data || error);
      toast({
        description: error.response?.data?.message || error.message,
        variant: "destructive",
      });
    }
  }

  async function deleteAd(id: number) {
    try {
      await axios.delete(`${API_URL}/ads/${id}`, { headers: authHeader() });
      toast({
        description: "Deleted Successfully",
      });
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

  async function updateAdStatus(id: number, status: string) {
    try {
      const response = await axios.patch(
        `${API_URL}/ads/${id}/status`,
        { status: status.toLowerCase() },
        {
          headers: {
            ...authHeader(),
            'Content-Type': 'application/json'
          }
        }
      );

      if (response.data) {
        toast({
          description: "Status updated successfully",
        });
      }
    } catch (error: any) {
      console.error('Update Status Error:', error.response?.data || error);
      toast({
        description: error.response?.data?.message || "Failed to update status",
        variant: "destructive",
      });
      throw error;
    }
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
    adsData,
    getAdsData,
    createAd,
    updateAd,
    deleteAd,
    updateAdStatus,
  };
}); 