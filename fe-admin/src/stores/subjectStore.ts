import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
const API_URL = import.meta.env.VITE_API_URL;

export interface Subject {
  id: number;
  name: string;
  createTime: string;
  operator: string;
  created_by: number;
  created_at: string;
  updated_by: number;
  updated_at: string;
  created_by_user: {};
  updated_by_user: {};
}

export const useSubjectStore = defineStore("subjectStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<Subject>({
    id: 0,
    name: "",
    createTime: "",
    operator: "",
    created_by: 0,
    created_at: "",
    updated_by: 0,
    updated_at: "",
    created_by_user: {},
    updated_by_user: {},
  });
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const subjectData = ref<Subject[]>([]);

  async function getSubjectData() {
    isLoading.value = true;
    await axios
      .get(
        `${API_URL}/subjects?page=${current_page.value}&page_size=${page_size.value}&q=${searchKeyword.value}`,
        { headers: authHeader() }
      )
      .then((response) => {
        subjectData.value = response.data.data;
        current_page.value = response.data.pagination.page;
        totalItems.value = response.data.pagination.total;
        page_size.value = response.data.pagination.page_size;
        isLoading.value = false;
      });
  }

  async function createSubject(data: any) {
    await axios
      .post(`${API_URL}/subjects`, data, { headers: authHeader() })
      .then((response) => {});
  }

  async function updateSubject(data: any) {
    await axios
      .put(`${API_URL}/subjects`, data, { headers: authHeader() })
      .then((response) => {});
    selectedData.value = <Subject>{
      id: 0,
      name: "",
      createTime: "",
      operator: "",
      created_by: 0,
      created_at: "",
      updated_by: 0,
      updated_at: "",
      created_by_user: {},
      updated_by_user: {},
    };
  }

  async function deleteSubject(id: any) {
    await axios
      .delete(`${API_URL}/subjects/${id}`, { headers: authHeader() })
      .then((response) => {});
    selectedData.value = <Subject>{
      id: 0,
      name: "",
      createTime: "",
      operator: "",
      created_by: 0,
      created_at: "",
      updated_by: 0,
      updated_at: "",
      created_by_user: {},
      updated_by_user: {},
    };
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
    subjectData,
    getSubjectData,
    createSubject,
    updateSubject,
    deleteSubject,
  };
}); 