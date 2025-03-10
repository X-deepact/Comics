import { defineStore } from "pinia";
import axios from "axios";
import router from "../router/router";
import { useToast } from "@/components/ui/toast/use-toast";
const API_URL = import.meta.env.VITE_API_URL;
const { toast } = useToast();
interface User {
  username: string;
  email: string;
  full_name: string;
  role: string;
  active: boolean;
  password_changed_at: string;
  created_at: string;
}
export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: localStorage.getItem("user") || (null as unknown as User),
    access_token: localStorage.getItem("access_token") || "",
  }),
  actions: {
    async login(username: string, password: string) {
      try {
        const { data } = await axios.post(`${API_URL}/user/login`, {
          username,
          password,
        });
        if (data.data.code == "ERROR") {
          toast({
            description: data.data.msg,
            variant: "destructive",
          });
          return false;
        }
        this.access_token = data.data.access_token;
        this.user = data.data.user;
        localStorage.setItem("user", data.data.user);
        localStorage.setItem("access_token", data.data.access_token);
        axios.defaults.headers.common[
          "Authorization"
        ] = `Bearer ${data.access_token}`;
        return true;
      } catch (error) {
        console.error(error);
        return false;
      }
    },
    logout() {
      this.user = null as unknown as User;
      this.access_token = "";
      localStorage.removeItem("access_token");
      localStorage.removeItem("user");
      delete axios.defaults.headers.common["Authorization"];
      router.push({ name: "login" });
    },
  },
});
