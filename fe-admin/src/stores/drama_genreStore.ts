import { defineStore } from "pinia";
import axios from "axios";
import { ref, computed } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";

const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface GenreTranslation {
  id?: number;
  genre_id?: number;
  language: string;
  translated_name: string;
  created_at?: string;
  updated_at?: string;
  created_by?: number;
  updated_by?: number;
}

export interface Genre {
  id: number;
  name: string;
  position: number;
  translations: GenreTranslation[];
  created_at: string;
  updated_at: string;
  created_by: number;
  updated_by: number;
}

export interface GenreCreateRequest {
  name: string;
  position: number;
  translations: {
    language: string;
    translated_name: string;
  }[];
}

export interface GenreFilters {
  name?: string;
  language?: string;
}

interface SortState {
  sort_by: string;
  sort_order: 'asc' | 'desc';
}

export const useDramaGenreStore = defineStore("dramaGenreStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<Genre>({
    id: 0,
    name: "",
    position: 0,
    translations: [],
    created_at: "",
    updated_at: "",
    created_by: 0,
    updated_by: 0,
  });
  
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const genreData = ref<Genre[]>([]);
  const filters = ref<GenreFilters>({});

  const sort = ref<SortState>({
    sort_by: 'id',
    sort_order: 'desc'
  });

  // Computed property for sorted data
  const sortedGenreData = computed(() => {
    if (!sort.value.sort_by) return genreData.value;

    return [...genreData.value].sort((a, b) => {
      const aValue = a[sort.value.sort_by as keyof typeof a] || '';
      const bValue = b[sort.value.sort_by as keyof typeof b] || '';

      if (sort.value.sort_order === 'asc') {
        return aValue < bValue ? -1 : aValue > bValue ? 1 : 0;
      } else {
        return bValue < aValue ? -1 : bValue > aValue ? 1 : 0;
      }
    });
  });

  async function getGenreData() {
    isLoading.value = true;
    try {
      const params = new URLSearchParams({
        page: current_page.value.toString(),
        page_size: page_size.value.toString(),
      });

      if (searchKeyword.value) params.append('name', searchKeyword.value);
      if (filters.value.language) params.append('language', filters.value.language);
      
      if (sort.value.sort_by) {
        params.append('sort_by', sort.value.sort_by);
        params.append('sort', sort.value.sort_order);
      }

      const response = await axios.get(
        `${API_URL}/genres-drama?${params.toString()}`,
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }

      genreData.value = response.data.data;
      current_page.value = response.data.pagination.page;
      totalItems.value = response.data.pagination.total || 0;
      page_size.value = response.data.pagination.page_size;

    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      genreData.value = [];
      totalItems.value = 0;
    } finally {
      isLoading.value = false;
    }
  }

  async function createGenre(data: GenreCreateRequest) {
    try {
      const response = await axios.post(
        `${API_URL}/genres-drama`,
        data,
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

  async function updateGenre(data: Partial<Genre>) {
    try {
      const response = await axios.put(
        `${API_URL}/genres-drama`,
        data,
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

  async function deleteGenre(id: number) {
    try {
      const response = await axios.delete(
        `${API_URL}/genres-drama/${id}`, 
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
        name: "",
        position: 0,
        translations: [],
        created_at: "",
        updated_at: "",
        created_by: 0,
        updated_by: 0,
      };
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
    }
  }

  async function getGenreById(id: number) {
    try {
      const response = await axios.get(
        `${API_URL}/genres-drama/${id}`,
        { headers: authHeader() }
      );

      if (response.data.code === "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return null;
      }

      return response.data.data;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      return null;
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

  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    selectedData,
    genreData: sortedGenreData,
    current_page,
    page_size,
    searchKeyword,
    totalItems,
    isLoading,
    getGenreData,
    createGenre,
    updateGenre,
    deleteGenre,
    getGenreById,
    filters,
    sort,
    updateSort,
  };
});
