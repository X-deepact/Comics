import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
import { useChapterStore } from "./chapterStore";
import { type DateValue } from "@internationalized/date";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();
const chapterStore = useChapterStore();

export interface ChapterItem {
  id: number;
  image: string;
  page: number;
  chapter_id: number;
  active: boolean;
  active_from: string;
  created_by_name: string;
  created_at: string;
  updated_by_name: string;
  updated_at: string;
}

export const useChapterItemStore = defineStore("chapteritemStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const chapteritemsData = ref<ChapterItem[]>([]);
  const selectedData = ref({
    id: 0,
    image: "",
    page: 0,
    chapter_id: 0,
    active: true,
    active_from: null as DateValue | null,
  });
  const searchKeyword = ref("");
  const isLoading = ref(true);
  const current_page = ref(1);
  const page_size = ref(10);
  const totalItems = ref(0);
  const sortBy = ref("");
  const sorting = ref("asc");
  async function getChapterItemsData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/chapter-items?page=${current_page.value}&page_size=${page_size.value}&chapter_id=${chapterStore.selectedData.id}&sort_by=${sortBy.value}&sort=${sorting.value}`,
        {
          headers: authHeader(),
        }
      )
      .then((response) => {
        chapteritemsData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total
          ? response.data.pagination.total
          : 0;
        page_size.value = response.data.pagination.page_size;
        isLoading.value = false;
      })
      .catch((error: any) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }
  async function uploadImage(data: any): Promise<string | null> {
    const formData = new FormData();
    formData.append("file", data);
    try {
      const response = await axios.post(
        `${API_URL}/chapter-items/upload-image`,
        formData,
        {
          headers: { ...authHeader(), "Content-Type": "multipart/form-data" },
        }
      );
      return response.data.data.data;
    } catch (error: any) {
      toast({
        description: error.message,
        variant: "destructive",
      });
      return null;
    }
  }

  async function createChapterItem(data: any) {
    await axios
      .post(`${API_URL}/chapter-items`, data, {
        headers: authHeader(),
      })
      .then((response) => {
        toast({
          description: "Created Successfully",
        });
      })
      .catch((error: any) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
  }
  async function updateChapterItem(data: any) {
    await axios
      .put(`${API_URL}/chapter-items`, data, {
        headers: authHeader(),
      })
      .then((response) => {
        toast({
          description: "Updated Successfully",
        });
      })
      .catch((error) => {
        toast({
          description: error.message,
          variant: "destructive",
        });
      });
    selectedData.value = {
      id: 0,
      image: "",
      page: 0,
      chapter_id: 0,
      active: true,
      active_from: null as DateValue | null,
    };
  }
  async function deleteChapterItem(id: any) {
    await axios
      .delete(`${API_URL}/chapter-items/${id}`, {
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
    selectedData.value = {
      id: 0,
      image: "",
      page: 0,
      chapter_id: 0,
      active: true,
      active_from: null as DateValue | null,
    };
  }
  function setSorting() {
    if (sorting.value === "asc") {
      sorting.value = "desc";
    } else {
      sorting.value = "asc";
    }
  }
  return {
    createDialogIsOpen,
    updateDialogIsOpen,
    deleteDialogIsOpen,
    chapteritemsData,
    selectedData,
    searchKeyword,
    isLoading,
    current_page,
    totalItems,
    page_size,
    sortBy,
    createChapterItem,
    updateChapterItem,
    getChapterItemsData,
    deleteChapterItem,
    uploadImage,
    setSorting,
  };
});
