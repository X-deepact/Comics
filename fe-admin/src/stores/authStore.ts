import { defineStore } from "pinia";
import axios from "axios";
import router from "../router/router";
const API_URL = import.meta.env.VITE_API_URL;

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
        this.access_token = data.access_token;
        this.user = data.user;
        localStorage.setItem("user", data.user);
        localStorage.setItem("access_token", data.access_token);
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
