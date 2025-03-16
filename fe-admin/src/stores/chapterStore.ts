import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
import { useComicStore } from "./comicStore";
import { type DateValue } from "@internationalized/date";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();
const comicStore = useComicStore();

export interface Chapter {
  id: number;
  name: string;
  number: number;
  active: boolean;
  active_from: string;
  comic_id: number;
  cover: boolean;
  created_by_name: string;
  created_at: string;
  updated_by_name: string;
  updated_at: string;
}
export const useChapterStore = defineStore("chapterStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const chapteritemDialogIsOpen = ref(false);
  const chaptersData = ref<Chapter[]>([]);
  const selectedData = ref({
    id: 0,
    name: "",
    number: 0,
    active_from: null as DateValue | null,
    active: false,
    comic_id: 0,
    cover: false,
  });
  const searchKeyword = ref("");
  const isLoading = ref(true);
  const current_page = ref(1);
  const page_size = ref(10);
  const totalItems = ref(0);
  const sortBy = ref("");
  const sorting = ref("asc");
  async function getChapterData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/chapters?page=${current_page.value}&page_size=${page_size.value}&comic_id=${comicStore.selectedData.id}&sort_by=${sortBy.value}&sort=${sorting.value}`,
        {
          headers: authHeader(),
        }
      )
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        chaptersData.value = response.data.data;
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
  async function createChapter(data: any) {
    await axios
      .post(`${API_URL}/chapters`, data, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
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

  async function updateChapter(data: any) {
    await axios
      .put(`${API_URL}/chapters`, data, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
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
      name: "",
      number: 0,
      active: false,
      active_from: null as DateValue | null,
      comic_id: 0,
      cover: false,
    };
  }
  async function deleteChapter(id: any) {
    await axios
      .delete(`${API_URL}/chapters/${id}`, {
        headers: authHeader(),
      })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
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
      name: "",
      number: 0,
      active_from: null as DateValue | null,
      active: false,
      comic_id: 0,
      cover: false,
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
    chapteritemDialogIsOpen,
    chaptersData,
    selectedData,
    searchKeyword,
    isLoading,
    current_page,
    totalItems,
    page_size,
    sortBy,
    createChapter,
    updateChapter,
    getChapterData,
    deleteChapter,
    setSorting,
  };
});
