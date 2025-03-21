import { defineStore } from "pinia";
import axios from "axios";
import { ref } from "vue";
import { authHeader } from "../services/authHeader";
import { useToast } from "@/components/ui/toast/use-toast";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();

export interface User {
  id: number;
  username: string;
  password: string;
  phone: string;
  email: string;
  FirstName: string;
  LastName: string;
  DisplayName: string;
  TierId: number;
  Birthday: string;
  created_at: string;
  updated_at: string;
  created_by_name: string;
  updated_by_name: string;
}
export const useUserStore = defineStore("userStore", () => {
  const createDialogIsOpen = ref(false);
  const updateDialogIsOpen = ref(false);
  const deleteDialogIsOpen = ref(false);
  const updateStatusDialogIsOpen = ref(false);
  const isLoading = ref(true);
  const selectedData = ref<User>({
    id: 0,
    username: "",
    password: "",
    phone: "",
    email: "",
    FirstName: "",
    LastName: "",
    DisplayName: "",
    TierId: 0,
    Birthday: "",
    created_at: "",
    updated_at: "",
    created_by_name: "",
    updated_by_name: "",
  });
  const current_page = ref(1);
  const page_size = ref(10);
  const searchKeyword = ref("");
  const totalItems = ref(0);
  const userData = ref<User[]>([]);
  const generalGenreData = ref<User[]>([]);

  async function getUserData() {
    isLoading.value = true;
    try {
      const params = {
        page: current_page.value,
        page_size: page_size.value,
        username: searchKeyword.value || undefined,
      };

      const response = await axios.get(`${API_URL}/user`, {
        headers: authHeader(),
        params,
      });
      if (response.data.code == "ERROR") {
        toast({
          description: response.data.msg,
          variant: "destructive",
        });
        return false;
      }
      userData.value = response.data.data;
      current_page.value = response.data.pagination.page;
      totalItems.value = response.data.pagination.total
        ? response.data.pagination.total
        : 0;
      page_size.value = response.data.pagination.page_size;
    } catch (error: any) {
      toast({
        description: `Error fetching user data: ${
          error.response?.data?.message || error.message
        }`,
        variant: "destructive",
      });
    } finally {
      isLoading.value = false;
    }
  }
  async function createUser(data: any) {
    await axios
      .post(`${API_URL}/user`, data, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        toast({
          description: "Created successfully",
        });
      })
      .catch((error) => {
        toast({
          description: `Error creating user: ${
            error.response?.data?.message || error.message
          }`,
          variant: "destructive",
        });
      });
  }
  async function updateUser(data: any) {
    await axios
      .put(`${API_URL}/user`, data, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        toast({
          description: "Updated successfully",
        });
      })
      .catch((error) => {
        toast({
          description: `Error updating user: ${
            error.response?.data?.message || error.message
          }`,
          variant: "destructive",
        });
      });
    selectedData.value = <User>{
      id: 0,
      username: "",
      password: "",
      phone: "",
      email: "",
      FirstName: "",
      LastName: "",
      DisplayName: "",
      TierId: 0,
      Birthday: "",
      created_at: "",
      updated_at: "",
      created_by_name: "",
      updated_by_name: "",
    };
  }
  async function deleteUser(id: any) {
    await axios
      .delete(`${API_URL}/user/${id}`, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        toast({
          description: "Deleted successfully",
        });
      })
      .catch((error) => {
        toast({
          description: `Error deleting user: ${
            error.response?.data?.message || error.message
          }`,
          variant: "destructive",
        });
      });
    selectedData.value = <User>{
      id: 0,
      username: "",
      password: "",
      phone: "",
      email: "",
      FirstName: "",
      LastName: "",
      DisplayName: "",
      TierId: 0,
      created_at: "",
      updated_at: "",
      created_by_name: "",
      updated_by_name: "",
    };
  }

  async function updateUserStatus(id: any) {
    await axios
      .put(`${API_URL}/user/${id}/active`, null, { headers: authHeader() })
      .then((response) => {
        if (response.data.code == "ERROR") {
          toast({
            description: response.data.msg,
            variant: "destructive",
          });
          return false;
        }
        toast({
          description: "Update Status successfully",
        });
      })
      .catch((error) => {
        toast({
          description: `Error updating user status: ${
            error.response?.data?.message || error.message
          }`,
          variant: "destructive",
        });
      });
    selectedData.value = <User>{
      id: 0,
      username: "",
      password: "",
      phone: "",
      email: "",
      FirstName: "",
      LastName: "",
      DisplayName: "",
      TierId: 0,
      created_at: "",
      updated_at: "",
      created_by_name: "",
      updated_by_name: "",
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
    userData,
    generalGenreData,
    getUserData,
    createUser,
    updateUser,
    deleteUser,
    updateUserStatus,
    updateStatusDialogIsOpen,
  };
});
