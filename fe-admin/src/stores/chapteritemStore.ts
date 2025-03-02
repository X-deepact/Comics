import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
import { useComicStore } from "./comicStore";
import { useChapterStore } from "./chapterStore";
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
  const selectedData = ref<ChapterItem>({
    id: 0,
    image: "",
    page: 0,
    chapter_id: 0,
    active: true,
    active_from: "",
    created_by_name: "",
    created_at: "",
    updated_by_name: "",
    updated_at: "",
  });
  const searchKeyword = ref("");
  const isLoading = ref(true);
  const current_page = ref(1);
  const page_size = ref(10);
  const totalItems = ref(0);
  async function getChapterItemsData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/chapter-items?page=${current_page.value}&page_size=${page_size.value}&chapter_id=${chapterStore.selectedData.id}`,
        {
          headers: authHeader(),
        }
      )
      .then((response) => {
        chapteritemsData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total;
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
  async function createChapterItem(data: any) {
    const formData = new FormData();
    formData.append("active", data.active);
    formData.append("active_from", data.active_from);
    formData.append("chapter_id", data.chapter_id);
    formData.append("image", data.image);
    formData.append("page", data.page);
    await axios
      .post(`${API_URL}/chapter-items`, formData, {
        headers: { ...authHeader(), "Content-Type": "multipart/form-data" },
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
    const formData = new FormData();
    formData.append("active", data.active);
    formData.append("active_from", data.active_from);
    formData.append("id", data.id);
    formData.append("image", data.image);
    formData.append("page", data.page);
    await axios
      .put(`${API_URL}/chapter-items`, formData, {
        headers: { ...authHeader(), "Content-Type": "multipart/form-data" },
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
    selectedData.value = <ChapterItem>{
      id: 0,
      image: "",
      page: 0,
      chapter_id: 0,
      active: true,
      active_from: "",
      created_by_name: "",
      created_at: "",
      updated_by_name: "",
      updated_at: "",
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
    selectedData.value = <ChapterItem>{
      id: 0,
      image: "",
      page: 0,
      chapter_id: 0,
      active: true,
      active_from: "",
      created_by_name: "",
      created_at: "",
      updated_by_name: "",
      updated_at: "",
    };
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
    createChapterItem,
    updateChapterItem,
    getChapterItemsData,
    deleteChapterItem,
  };
});
